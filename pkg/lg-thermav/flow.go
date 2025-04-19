package lgthermav

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func getFlowRate(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_flow_rate",
			mcp.WithDescription("Get current flow rate in L/min"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.FlowRate()
			if err != nil {
				return nil, fmt.Errorf("error getting flow rate: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}
