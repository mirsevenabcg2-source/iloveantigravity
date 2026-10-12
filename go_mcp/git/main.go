package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func runGit(repoPath string, args ...string) (string, error) {
	cmdArgs := append([]string{"-C", repoPath}, args...)
	cmd := exec.Command("git", cmdArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	outStr := stdout.String()
	errStr := stderr.String()

	if err != nil {
		if strings.TrimSpace(errStr) != "" {
			return "", fmt.Errorf("%s", strings.TrimSpace(errStr))
		}
		if strings.TrimSpace(outStr) != "" {
			return "", fmt.Errorf("%s", strings.TrimSpace(outStr))
		}
		return "", err
	}

	res := strings.TrimSpace(outStr)
	if res == "" && strings.TrimSpace(errStr) != "" {
		res = strings.TrimSpace(errStr)
	}
	return res, nil
}

func main() {
	s := server.NewMCPServer("git", "1.0.0")

	// 1. git_status
	s.AddTool(mcp.NewTool("git_status",
		mcp.WithDescription("Shows the working tree status"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		out, err := runGit(repoPath, "status")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 2. git_diff_unstaged
	s.AddTool(mcp.NewTool("git_diff_unstaged",
		mcp.WithDescription("Shows changes in the working directory that are not yet staged"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithNumber("context_lines", mcp.Description("Number of context lines to show")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		ctxLines := request.GetInt("context_lines", 3)
		out, err := runGit(repoPath, "diff", fmt.Sprintf("-U%d", ctxLines))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 3. git_diff_staged
	s.AddTool(mcp.NewTool("git_diff_staged",
		mcp.WithDescription("Shows changes that are staged for commit"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithNumber("context_lines", mcp.Description("Number of context lines to show")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		ctxLines := request.GetInt("context_lines", 3)
		out, err := runGit(repoPath, "diff", "--staged", fmt.Sprintf("-U%d", ctxLines))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 4. git_diff
	s.AddTool(mcp.NewTool("git_diff",
		mcp.WithDescription("Shows differences between branches or commits"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithString("target", mcp.Required(), mcp.Description("Target branch or commit to diff against")),
		mcp.WithNumber("context_lines", mcp.Description("Number of context lines to show")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		target, err := request.RequireString("target")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		ctxLines := request.GetInt("context_lines", 3)
		out, err := runGit(repoPath, "diff", target, fmt.Sprintf("-U%d", ctxLines))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 5. git_commit
	s.AddTool(mcp.NewTool("git_commit",
		mcp.WithDescription("Records changes to the repository"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithString("message", mcp.Required(), mcp.Description("Commit message")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		msg, err := request.RequireString("message")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		out, err := runGit(repoPath, "commit", "-m", msg)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 6. git_add
	s.AddTool(mcp.NewTool("git_add",
		mcp.WithDescription("Adds file contents to the staging area"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithArray("files", mcp.Required(), mcp.Description("List of files to stage")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		files := request.GetStringSlice("files", nil)
		if len(files) == 0 {
			return mcp.NewToolResultError("files list cannot be empty"), nil
		}
		args := append([]string{"add"}, files...)
		out, err := runGit(repoPath, args...)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if out == "" {
			out = fmt.Sprintf("Successfully staged %d file(s)", len(files))
		}
		return mcp.NewToolResultText(out), nil
	})

	// 7. git_reset
	s.AddTool(mcp.NewTool("git_reset",
		mcp.WithDescription("Unstages all staged changes"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		out, err := runGit(repoPath, "reset")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if out == "" {
			out = "Successfully unstaged all changes"
		}
		return mcp.NewToolResultText(out), nil
	})

	// 8. git_log
	s.AddTool(mcp.NewTool("git_log",
		mcp.WithDescription("Shows the commit logs"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithNumber("max_count", mcp.Description("Maximum number of commits to show")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		maxCount := request.GetInt("max_count", 10)
		out, err := runGit(repoPath, "log", "-n", strconv.Itoa(maxCount), "--oneline", "--decorate")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 9. git_create_branch
	s.AddTool(mcp.NewTool("git_create_branch",
		mcp.WithDescription("Creates a new branch from an optional base branch"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithString("branch_name", mcp.Required(), mcp.Description("Name of the branch to create")),
		mcp.WithString("base_branch", mcp.Description("Optional base branch")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		branchName, err := request.RequireString("branch_name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		baseBranch := request.GetString("base_branch", "")
		args := []string{"checkout", "-b", branchName}
		if baseBranch != "" {
			args = append(args, baseBranch)
		}
		out, err := runGit(repoPath, args...)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 10. git_checkout
	s.AddTool(mcp.NewTool("git_checkout",
		mcp.WithDescription("Switches branches"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithString("branch_name", mcp.Required(), mcp.Description("Branch name to switch to")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		branchName, err := request.RequireString("branch_name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		out, err := runGit(repoPath, "checkout", branchName)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 11. git_show
	s.AddTool(mcp.NewTool("git_show",
		mcp.WithDescription("Shows the contents of a commit"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithString("revision", mcp.Required(), mcp.Description("Commit hash or reference")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		rev, err := request.RequireString("revision")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		out, err := runGit(repoPath, "show", rev)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	// 12. git_branch
	s.AddTool(mcp.NewTool("git_branch",
		mcp.WithDescription("List Git branches"),
		mcp.WithString("repo_path", mcp.Required(), mcp.Description("Path to the git repository")),
		mcp.WithString("branch_type", mcp.Description("Branch type: local, remote, all")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		repoPath, err := request.RequireString("repo_path")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		branchType := request.GetString("branch_type", "local")
		args := []string{"branch"}
		if branchType == "all" {
			args = append(args, "-a")
		} else if branchType == "remote" {
			args = append(args, "-r")
		}
		out, err := runGit(repoPath, args...)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(out), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
