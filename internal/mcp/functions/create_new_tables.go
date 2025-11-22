package functions_shell_history

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/quentinchampenois/go-grist-api"
)

type InputCreateNewTables struct {
	Tables     grist.Tables `json:"tables" jsonschema:"tables to create in document"`
	DocumentID string       `json:"documentID" jsonschema:"ID of the document where the tables will be created"`
}

type OutputCreateNewTables struct {
	DocumentID string   `json:"documentID" jsonschema:"ID of the document used"`
	TablesIDs  []string `json:"tablesIDs" jsonschema:"IDs of the tables created"`
}

func CreateNewTables(ctx context.Context, req *mcp.CallToolRequest, input InputCreateNewTables) (
	*mcp.CallToolResult,
	OutputCreateNewTables,
	error,
) {
	endpoint := os.Getenv("GRIST_ENDPOINT")
	apiKey := os.Getenv("GRIST_API_KEY")
	c, err := grist.NewGristClient(ctx, endpoint, apiKey)
	if err != nil {
		return nil, OutputCreateNewTables{}, fmt.Errorf("failed to create Grist client: %w", err)
	}

	doc, err := grist.DescribeDoc(c, input.DocumentID)
	if err != nil {
		return nil, OutputCreateNewTables{}, fmt.Errorf("failed to describe document (documentID/'%s') : %w", input.DocumentID, err)
	}

	newTables, err := doc.CreateTables(c, input.Tables)
	if err != nil {
		return nil, OutputCreateNewTables{}, fmt.Errorf("failed to create tables in documnet (documentID/'%s') : %w", input.DocumentID, err)
	}

	var tableIds []string
	for _, table := range newTables.Tables {
		tableIds = append(tableIds, table.ID)
	}
	return nil, OutputCreateNewTables{
		DocumentID: doc.ID,
		TablesIDs:  tableIds,
	}, nil
}
