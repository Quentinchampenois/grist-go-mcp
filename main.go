package main

import (
	"context"
	functionsgristmcp "grist-mcp-server/internal/mcp/functions"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "grist-mcp", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "list_orgs", Description: "List all organizations"}, functionsgristmcp.ListOrgs)
	mcp.AddTool(server, &mcp.Tool{Name: "create_workspace", Description: "Create a workspace in organization"}, functionsgristmcp.CreateWorkspace)
	mcp.AddTool(server, &mcp.Tool{Name: "create_new_docs", Description: "Create new documents in a workspace"}, functionsgristmcp.CreateNewDocs)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
