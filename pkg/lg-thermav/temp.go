package lgthermav

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func getIndoorTemp(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_indoor_temperature",
			mcp.WithDescription("Get current indoor temperature which is reported by thermostat"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.IndoorTemperature()
			if err != nil {
				return nil, fmt.Errorf("error getting current temperature: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}

func getOutdoorTemp(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_outdoor_temperature",
			mcp.WithDescription("Get current outdoor (outside) temperature which is reported by outdoor unit"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.OutsideTemperature()
			if err != nil {
				return nil, fmt.Errorf("error getting temperature: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}

func getInletTemp(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_inlet_temperature",
			mcp.WithDescription("Get current inlet temperature"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.InletTemperature()
			if err != nil {
				return nil, fmt.Errorf("error getting temperature: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}

func getOutletTemp(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_outlet_temperature",
			mcp.WithDescription("Get current outlet temperature"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.OutletTemperature()
			if err != nil {
				return nil, fmt.Errorf("error getting temperature: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}

func getUnderflowTargetTemp(svc ModbusClient) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_target_underflow_temp",
			mcp.WithDescription("Get target underflow temperature"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := svc.UnderflowTargetTemperature()
			if err != nil {
				return nil, fmt.Errorf("error getting temperature: %w", err)
			}

			txt := fmt.Sprintf("%.1f°C", r)
			return mcp.NewToolResultText(txt), nil
		}
}
