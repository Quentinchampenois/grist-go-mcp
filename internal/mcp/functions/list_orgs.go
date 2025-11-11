package functions_shell_history

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/quentinchampenois/go-grist-api"
)

type InputListOrgs struct{}

type OutputListOrgs struct {
	Orgs []grist.Org `json:"orgs" jsonschema:"List of organizations"`
}

func ListOrgs(ctx context.Context, req *mcp.CallToolRequest, input InputListOrgs) (
	*mcp.CallToolResult,
	OutputListOrgs,
	error,
) {
	endpoint := os.Getenv("GRIST_ENDPOINT")
	apiKey := os.Getenv("GRIST_API_KEY")
	c, err := grist.NewGristClient(ctx, endpoint, apiKey)
	if err != nil {
		return nil, OutputListOrgs{}, fmt.Errorf("failed to create Grist client: %w", err)
	}

	orgs, err := grist.ListOrgs(c)
	if err != nil {
		return nil, OutputListOrgs{}, fmt.Errorf("failed to list orgs: %w", err)
	}

	return nil, OutputListOrgs{
		Orgs: orgs,
	}, nil
}
