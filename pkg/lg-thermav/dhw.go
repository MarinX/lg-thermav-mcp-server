package lgthermav

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func getDHWTemperature(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_dhw_temperature",
			mcp.WithDescription("Get domestic hot water (DHW) current temperature"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.DHWTemperature()
			if err != nil {
				return nil, fmt.Errorf("error getting current DHW temperature: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}

func getDHWTargetTemperature(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_target_dhw_temperature",
			mcp.WithDescription("Get target domestic hot water (DHW) temperature"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.DHWTargetTemperature()
			if err != nil {
				return nil, fmt.Errorf("error getting target DHW temperature: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}

func isDHWRunning(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("is_dhw_running",
			mcp.WithDescription("Check if heat pump is heating domestic hot water (DHW)"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.DHWHeatRunning()
			if err != nil {
				return nil, fmt.Errorf("error getting DHW status: %w", err)
			}
			return mcp.NewToolResultText(strconv.FormatBool(r)), nil
		}
}
