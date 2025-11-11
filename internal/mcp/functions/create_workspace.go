package functions_shell_history

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/quentinchampenois/go-grist-api"
)

type InputCreateWorkspace struct {
	OrgID         int64  `json:"orgID" jsonschema:"Org ID where the workspace will be created"`
	WorkspaceName string `json:"workspaceName" jsonschema:"Name of the workspace to create"`
}

type OutputCreateWorkspace struct {
	WorkspaceID  int64  `json:"workspaceID" jsonschema:"ID of the workspace created"`
	WorkspaceURL string `json:"workspaceURL" jsonschema:"URL of the workspace created"`
}

func CreateWorkspace(ctx context.Context, req *mcp.CallToolRequest, input InputCreateWorkspace) (
	*mcp.CallToolResult,
	OutputCreateWorkspace,
	error,
) {
	endpoint := os.Getenv("GRIST_ENDPOINT")
	apiKey := os.Getenv("GRIST_API_KEY")
	c, err := grist.NewGristClient(ctx, endpoint, apiKey)
	if err != nil {
		return nil, OutputCreateWorkspace{}, fmt.Errorf("failed to create Grist client: %w", err)
	}

	workspaceID, err := grist.CreateWorkspace(c, input.OrgID, input.WorkspaceName)
	if err != nil {
		return nil, OutputCreateWorkspace{}, fmt.Errorf("failed to create workspace: %w", err)
	}

	return nil, OutputCreateWorkspace{
		WorkspaceID:  *workspaceID,
		WorkspaceURL: fmt.Sprintf("%s/o/docs/ws/%d", endpoint, workspaceID),
	}, nil
}
