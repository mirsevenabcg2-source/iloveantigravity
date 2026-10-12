package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Entity struct {
	Type         string   `json:"type"`
	Name         string   `json:"name"`
	EntityType   string   `json:"entityType"`
	Observations []string `json:"observations"`
}

type Relation struct {
	Type         string `json:"type"`
	From         string `json:"from"`
	To           string `json:"to"`
	RelationType string `json:"relationType"`
}

type KnowledgeGraph struct {
	mu        sync.RWMutex
	filePath  string
	Entities  map[string]Entity
	Relations []Relation
}

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
	return ""
}

func resolveMemoryFilePath() string {
	if envPath := os.Getenv("MEMORY_FILE_PATH"); envPath != "" {
		return envPath
	}
	cwd, _ := os.Getwd()
	root := findProjectRoot(cwd)
	if root != "" {
		agentsDir := filepath.Join(root, ".agents")
		if _, err := os.Stat(agentsDir); err == nil {
			return filepath.Join(agentsDir, "memory.jsonl")
		}
		return filepath.Join(root, ".memory.jsonl")
	}
	home, _ := os.UserHomeDir()
	globalDir := filepath.Join(home, ".gemini", "antigravity")
	_ = os.MkdirAll(globalDir, 0755)
	return filepath.Join(globalDir, "global_memory.jsonl")
}

func NewKnowledgeGraph(filePath string) *KnowledgeGraph {
	kg := &KnowledgeGraph{
		filePath: filePath,
		Entities: make(map[string]Entity),
	}
	kg.load()
	return kg
}

func (kg *KnowledgeGraph) load() {
	kg.mu.Lock()
	defer kg.mu.Unlock()

	kg.Entities = make(map[string]Entity)
	kg.Relations = nil

	file, err := os.Open(kg.filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var raw map[string]interface{}
		if err := json.Unmarshal([]byte(line), &raw); err != nil {
			continue
		}
		itemType, _ := raw["type"].(string)
		if itemType == "entity" {
			var ent Entity
			if err := json.Unmarshal([]byte(line), &ent); err == nil {
				kg.Entities[ent.Name] = ent
			}
		} else if itemType == "relation" {
			var rel Relation
			if err := json.Unmarshal([]byte(line), &rel); err == nil {
				kg.Relations = append(kg.Relations, rel)
			}
		}
	}
}

func (kg *KnowledgeGraph) save() error {
	kg.mu.RLock()
	defer kg.mu.RUnlock()

	dir := filepath.Dir(kg.filePath)
	_ = os.MkdirAll(dir, 0755)

	file, err := os.Create(kg.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, ent := range kg.Entities {
		ent.Type = "entity"
		b, _ := json.Marshal(ent)
		_, _ = writer.WriteString(string(b) + "\n")
	}
	for _, rel := range kg.Relations {
		rel.Type = "relation"
		b, _ := json.Marshal(rel)
		_, _ = writer.WriteString(string(b) + "\n")
	}
	return writer.Flush()
}

func main() {
	filePath := resolveMemoryFilePath()
	kg := NewKnowledgeGraph(filePath)

	s := server.NewMCPServer("memory", "1.0.0")

	// 1. read_graph
	s.AddTool(mcp.NewTool("read_graph",
		mcp.WithDescription("Read the entire knowledge graph"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		kg.load()
		kg.mu.RLock()
		defer kg.mu.RUnlock()

		var ents []Entity
		for _, e := range kg.Entities {
			ents = append(ents, e)
		}
		res := map[string]interface{}{
			"entities":  ents,
			"relations": kg.Relations,
		}
		b, _ := json.MarshalIndent(res, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	// 2. create_entities
	s.AddTool(mcp.NewTool("create_entities",
		mcp.WithDescription("Create multiple new entities in the knowledge graph"),
		mcp.WithArray("entities", mcp.Required(), mcp.Description("Array of entity objects")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Entities []Entity `json:"entities"`
		}
		if err := request.BindArguments(&args); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		kg.mu.Lock()
		for _, e := range args.Entities {
			if e.Name != "" {
				e.Type = "entity"
				kg.Entities[e.Name] = e
			}
		}
		kg.mu.Unlock()
		_ = kg.save()
		return mcp.NewToolResultText(fmt.Sprintf("Created %d entities", len(args.Entities))), nil
	})

	// 3. create_relations
	s.AddTool(mcp.NewTool("create_relations",
		mcp.WithDescription("Create multiple new relations between entities in the knowledge graph"),
		mcp.WithArray("relations", mcp.Required(), mcp.Description("Array of relation objects")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Relations []Relation `json:"relations"`
		}
		if err := request.BindArguments(&args); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		kg.mu.Lock()
		for _, r := range args.Relations {
			if r.From != "" && r.To != "" {
				r.Type = "relation"
				kg.Relations = append(kg.Relations, r)
			}
		}
		kg.mu.Unlock()
		_ = kg.save()
		return mcp.NewToolResultText(fmt.Sprintf("Created %d relations", len(args.Relations))), nil
	})

	// 4. add_observations
	s.AddTool(mcp.NewTool("add_observations",
		mcp.WithDescription("Add new observations to existing entities in the knowledge graph"),
		mcp.WithArray("observations", mcp.Required(), mcp.Description("Array of {entityName, contents}")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Observations []struct {
				EntityName string   `json:"entityName"`
				Contents   []string `json:"contents"`
			} `json:"observations"`
		}
		if err := request.BindArguments(&args); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		kg.mu.Lock()
		count := 0
		for _, obs := range args.Observations {
			if ent, ok := kg.Entities[obs.EntityName]; ok {
				ent.Observations = append(ent.Observations, obs.Contents...)
				kg.Entities[obs.EntityName] = ent
				count += len(obs.Contents)
			}
		}
		kg.mu.Unlock()
		_ = kg.save()
		return mcp.NewToolResultText(fmt.Sprintf("Added %d observations", count)), nil
	})

	// 5. search_nodes
	s.AddTool(mcp.NewTool("search_nodes",
		mcp.WithDescription("Search for nodes in the knowledge graph based on a query"),
		mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, _ := request.RequireString("query")
		q := strings.ToLower(query)

		kg.load()
		kg.mu.RLock()
		defer kg.mu.RUnlock()

		var matched []Entity
		for _, e := range kg.Entities {
			if strings.Contains(strings.ToLower(e.Name), q) || strings.Contains(strings.ToLower(e.EntityType), q) {
				matched = append(matched, e)
				continue
			}
			for _, o := range e.Observations {
				if strings.Contains(strings.ToLower(o), q) {
					matched = append(matched, e)
					break
				}
			}
		}
		b, _ := json.MarshalIndent(matched, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	// 6. open_nodes
	s.AddTool(mcp.NewTool("open_nodes",
		mcp.WithDescription("Open specific nodes in the knowledge graph by their names"),
		mcp.WithArray("names", mcp.Required(), mcp.Description("Names of entities to open")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Names []string `json:"names"`
		}
		if err := request.BindArguments(&args); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		kg.load()
		kg.mu.RLock()
		defer kg.mu.RUnlock()

		nameMap := make(map[string]bool)
		for _, n := range args.Names {
			nameMap[n] = true
		}

		var ents []Entity
		for name := range nameMap {
			if e, ok := kg.Entities[name]; ok {
				ents = append(ents, e)
			}
		}
		var rels []Relation
		for _, r := range kg.Relations {
			if nameMap[r.From] || nameMap[r.To] {
				rels = append(rels, r)
			}
		}
		res := map[string]interface{}{
			"entities":  ents,
			"relations": rels,
		}
		b, _ := json.MarshalIndent(res, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	// 7. delete_entities
	s.AddTool(mcp.NewTool("delete_entities",
		mcp.WithDescription("Delete multiple entities and their associated relations"),
		mcp.WithArray("entityNames", mcp.Required(), mcp.Description("Names of entities to delete")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			EntityNames []string `json:"entityNames"`
		}
		if err := request.BindArguments(&args); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		nameMap := make(map[string]bool)
		for _, n := range args.EntityNames {
			nameMap[n] = true
		}

		kg.mu.Lock()
		for n := range nameMap {
			delete(kg.Entities, n)
		}
		var keptRels []Relation
		for _, r := range kg.Relations {
			if !nameMap[r.From] && !nameMap[r.To] {
				keptRels = append(keptRels, r)
			}
		}
		kg.Relations = keptRels
		kg.mu.Unlock()
		_ = kg.save()
		return mcp.NewToolResultText(fmt.Sprintf("Deleted %d entities", len(args.EntityNames))), nil
	})

	// 8. delete_observations
	s.AddTool(mcp.NewTool("delete_observations",
		mcp.WithDescription("Delete specific observations from entities in the knowledge graph"),
		mcp.WithArray("deletions", mcp.Required(), mcp.Description("Array of {entityName, observations}")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Deletions []struct {
				EntityName   string   `json:"entityName"`
				Observations []string `json:"observations"`
			} `json:"deletions"`
		}
		if err := request.BindArguments(&args); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		kg.mu.Lock()
		for _, del := range args.Deletions {
			if ent, ok := kg.Entities[del.EntityName]; ok {
				obsMap := make(map[string]bool)
				for _, o := range del.Observations {
					obsMap[o] = true
				}
				var kept []string
				for _, o := range ent.Observations {
					if !obsMap[o] {
						kept = append(kept, o)
					}
				}
				ent.Observations = kept
				kg.Entities[del.EntityName] = ent
			}
		}
		kg.mu.Unlock()
		_ = kg.save()
		return mcp.NewToolResultText("Observations deleted successfully"), nil
	})

	// 9. delete_relations
	s.AddTool(mcp.NewTool("delete_relations",
		mcp.WithDescription("Delete multiple relations from the knowledge graph"),
		mcp.WithArray("relations", mcp.Required(), mcp.Description("Array of relations to delete")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Relations []Relation `json:"relations"`
		}
		if err := request.BindArguments(&args); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		kg.mu.Lock()
		for _, toDel := range args.Relations {
			var kept []Relation
			for _, r := range kg.Relations {
				if r.From == toDel.From && r.To == toDel.To && r.RelationType == toDel.RelationType {
					continue
				}
				kept = append(kept, r)
			}
			kg.Relations = kept
		}
		kg.mu.Unlock()
		_ = kg.save()
		return mcp.NewToolResultText("Relations deleted successfully"), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}
