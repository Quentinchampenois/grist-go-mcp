package functions_shell_history

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/quentinchampenois/go-grist-api"
)

type InputCreateNewDocs struct {
	DocNames    []string `json:"names" jsonschema:"Name of the docs to create"`
	WorkspaceID int64    `json:"workspaceID" jsonschema:"ID of the workspace where the docs will be created"`
}

type OutputCreateNewDocs struct {
	WorkspaceID int64    `json:"workspaceID" jsonschema:"ID of the workspace used"`
	DocsIDs     []string `json:"docIDs" jsonschema:"IDs of the docs created"`
}

func CreateNewDocs(ctx context.Context, req *mcp.CallToolRequest, input InputCreateNewDocs) (
	*mcp.CallToolResult,
	OutputCreateNewDocs,
	error,
) {
	endpoint := os.Getenv("GRIST_ENDPOINT")
	apiKey := os.Getenv("GRIST_API_KEY")
	c, err := grist.NewGristClient(ctx, endpoint, apiKey)
	if err != nil {
		return nil, OutputCreateNewDocs{}, fmt.Errorf("failed to create Grist client: %w", err)
	}

	workspace, err := grist.DescribeWorkspace(c, input.WorkspaceID)
	if err != nil {
		return nil, OutputCreateNewDocs{}, fmt.Errorf("failed to describe workspace : %w", err)
	}
	var docIds []string
	if len(input.DocNames) == 0 {
		return nil, OutputCreateNewDocs{}, fmt.Errorf("no doc names provided")
	}
	for _, docName := range input.DocNames {
		docId, err := workspace.CreateDoc(c, docName, false)
		if err != nil {
			return nil, OutputCreateNewDocs{}, fmt.Errorf("failed to create doc: %w", err)
		}
		docIds = append(docIds, *docId)
	}
	return nil, OutputCreateNewDocs{
		WorkspaceID: workspace.ID,
		DocsIDs:     docIds,
	}, nil
}
