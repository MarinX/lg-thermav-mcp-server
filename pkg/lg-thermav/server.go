package lgthermav

import (
	"github.com/mark3labs/mcp-go/server"
)

type ModbusClient interface {
	DHWTargetTemperature() (float32, error)
	DHWTemperature() (float32, error)
	DHWHeatRunning() (bool, error)
	CompressorSpeed() (int, error)
	CompressorRunning() (bool, error)
	IndoorTemperature() (float32, error)
	OutsideTemperature() (float32, error)
	InletTemperature() (float32, error)
	OutletTemperature() (float32, error)
	FlowRate() (float32, error)
	UnderflowTargetTemperature() (float32, error)
}

// NewServer creates a new Upcloud MCP server with the specified Upcloud client and logger.
func NewServer(client ModbusClient, version string, readOnly bool) *server.MCPServer {
	// Create a new MCP server
	s := server.NewMCPServer(
		"lg-thermav-server",
		version,
		server.WithResourceCapabilities(true, true),
		server.WithLogging())

	// DHW tools
	s.AddTool(getDHWTargetTemperature(client))
	s.AddTool(getDHWTemperature(client))
	s.AddTool(isDHWRunning(client))

	// Compressor
	s.AddTool(getCompressorSpeed(client))
	s.AddTool(isCompressorRunning(client))

	// Temps
	s.AddTool(getIndoorTemp(client))
	s.AddTool(getOutdoorTemp(client))
	s.AddTool(getInletTemp(client))
	s.AddTool(getOutletTemp(client))
	s.AddTool(getUnderflowTargetTemp(client))

	// Flow rate
	s.AddTool(getFlowRate(client))

	return s
}
