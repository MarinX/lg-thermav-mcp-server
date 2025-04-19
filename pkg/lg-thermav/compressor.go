package lgthermav

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func getCompressorSpeed(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_compressor_speed",
			mcp.WithDescription("Get the outdoor unit compressor speed in o/min"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.CompressorSpeed()
			if err != nil {
				return nil, fmt.Errorf("error getting current compressor speed: %w", err)
			}

			txt := fmt.Sprintf("%d o/min", r)
			return mcp.NewToolResultText(txt), nil
		}
}

func isCompressorRunning(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("is_compressor_running",
			mcp.WithDescription("Check if the outdoor unit compressor is running"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.CompressorSpeed()
			if err != nil {
				return nil, fmt.Errorf("error getting current compressor speed: %w", err)
			}

			txt := fmt.Sprintf("%d o/min", r)
			return mcp.NewToolResultText(txt), nil
		}
}
