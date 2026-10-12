package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	_ "modernc.org/sqlite"
)

func findProjectRoot(startDir string) string {
	curr, err := filepath.Abs(startDir)
	if err != nil {
		return startDir
	}
	markers := []string{".git", ".agents", ".agent", "package.json", "pyproject.toml", "Cargo.toml", "go.mod"}

	for {
		for _, m := range markers {
			if _, err := os.Stat(filepath.Join(curr, m)); err == nil {
				return curr
			}
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}
	return startDir
}

func resolveDBPath(customPath string) string {
	if customPath != "" {
		return customPath
	}
	if envPath := os.Getenv("SQLITE_DB_PATH"); envPath != "" {
		return envPath
	}

	cwd, _ := os.Getwd()
	root := findProjectRoot(cwd)

	commonNames := []string{"dev.db", "app.db", "db.sqlite3", "database.sqlite", "data.db", "sqlite.db"}
	for _, name := range commonNames {
		cand := filepath.Join(root, name)
		if _, err := os.Stat(cand); err == nil {
			return cand
		}
	}

	entries, _ := os.ReadDir(root)
	for _, e := range entries {
		if !e.IsDir() {
			ext := strings.ToLower(filepath.Ext(e.Name()))
			if ext == ".sqlite" || ext == ".db" || ext == ".sqlite3" {
				return filepath.Join(root, e.Name())
			}
		}
	}

	return filepath.Join(root, ".sqlite_mcp.db")
}

func main() {
	dbFlag := flag.String("db-path", "", "Path to SQLite database file")
	flag.Parse()

	dbPath := resolveDBPath(*dbFlag)
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open SQLite database %s: %v\n", dbPath, err)
		os.Exit(1)
	}
	defer db.Close()

	s := server.NewMCPServer("sqlite", "1.0.0")

	// 1. read_query
	s.AddTool(mcp.NewTool("read_query",
		mcp.WithDescription("Execute a SELECT query and return rows"),
		mcp.WithString("query", mcp.Required(), mcp.Description("The SELECT query to execute")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, err := request.RequireString("query")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		rows, err := db.QueryContext(ctx, query)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		defer rows.Close()

		cols, err := rows.Columns()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		var results []map[string]interface{}
		for rows.Next() {
			colVals := make([]interface{}, len(cols))
			colPointers := make([]interface{}, len(cols))
			for i := range colVals {
				colPointers[i] = &colVals[i]
			}
			if err := rows.Scan(colPointers...); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			rowMap := make(map[string]interface{})
			for i, col := range cols {
				val := colVals[i]
				if b, ok := val.([]byte); ok {
					rowMap[col] = string(b)
				} else {
					rowMap[col] = val
				}
			}
			results = append(results, rowMap)
		}

		outBytes, _ := json.MarshalIndent(results, "", "  ")
		return mcp.NewToolResultText(string(outBytes)), nil
	})

	// 2. write_query
	s.AddTool(mcp.NewTool("write_query",
		mcp.WithDescription("Execute an INSERT, UPDATE, or DELETE query"),
		mcp.WithString("query", mcp.Required(), mcp.Description("The SQL query to execute")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, err := request.RequireString("query")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		res, err := db.ExecContext(ctx, query)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		affected, _ := res.RowsAffected()
		lastID, _ := res.LastInsertId()
		return mcp.NewToolResultText(fmt.Sprintf("Query executed successfully. Rows affected: %d, LastInsertId: %d", affected, lastID)), nil
	})

	// 3. create_table
	s.AddTool(mcp.NewTool("create_table",
		mcp.WithDescription("Execute a CREATE TABLE statement"),
		mcp.WithString("query", mcp.Required(), mcp.Description("The CREATE TABLE SQL query")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, err := request.RequireString("query")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if _, err := db.ExecContext(ctx, query); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText("Table created successfully."), nil
	})

	// 4. list_tables
	s.AddTool(mcp.NewTool("list_tables",
		mcp.WithDescription("List all tables in the SQLite database"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		rows, err := db.QueryContext(ctx, "SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		defer rows.Close()

		var tables []string
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err == nil {
				tables = append(tables, name)
			}
		}
		outBytes, _ := json.MarshalIndent(tables, "", "  ")
		return mcp.NewToolResultText(string(outBytes)), nil
	})

	// 5. describe_table
	s.AddTool(mcp.NewTool("describe_table",
		mcp.WithDescription("Describe the schema of a specific table"),
		mcp.WithString("table_name", mcp.Required(), mcp.Description("Name of the table to inspect")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tableName, err := request.RequireString("table_name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		rows, err := db.QueryContext(ctx, fmt.Sprintf("PRAGMA table_info(%s)", tableName))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		defer rows.Close()

		var columns []map[string]interface{}
		for rows.Next() {
			var cid int
			var name, colType string
			var notNull, pk int
			var dfltValue sql.NullString
			if err := rows.Scan(&cid, &name, &colType, &notNull, &dfltValue, &pk); err == nil {
				columns = append(columns, map[string]interface{}{
					"cid":        cid,
					"name":       name,
					"type":       colType,
					"notnull":    notNull,
					"dflt_value": dfltValue.String,
					"pk":         pk,
				})
			}
		}
		outBytes, _ := json.MarshalIndent(columns, "", "  ")
		return mcp.NewToolResultText(string(outBytes)), nil
	})

	// 6. append_insight
	s.AddTool(mcp.NewTool("append_insight",
		mcp.WithDescription("Append a business or technical insight to the insights table"),
		mcp.WithString("insight", mcp.Required(), mcp.Description("The insight text to record")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		insight, err := request.RequireString("insight")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		_, _ = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS insights (id INTEGER PRIMARY KEY AUTOINCREMENT, insight TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)")
		_, err = db.ExecContext(ctx, "INSERT INTO insights (insight) VALUES (?)", insight)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText("Insight successfully recorded."), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}
