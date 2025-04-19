package lgthermav

import "github.com/mark3labs/mcp-go/mcp"

type MockModbus struct {
	Error          error
	DHWTargetTemp  float32
	DHWCurrentTemp float32
	DHWRunning     bool
	CpSpeed        int  // compressor speed
	CpRunning      bool // compressor running
	IndoorTemp     float32
	OutTemp        float32
	InletTemp      float32
	OutletTemp     float32
	FR             float32 // flow rate
	UFTemp         float32 // underflow temp
}

func (m *MockModbus) DHWTargetTemperature() (float32, error) {
	return m.DHWTargetTemp, m.Error
}

func (m *MockModbus) DHWTemperature() (float32, error) {
	return m.DHWCurrentTemp, m.Error
}

func (m *MockModbus) DHWHeatRunning() (bool, error) {
	return m.DHWRunning, m.Error
}

func (m *MockModbus) CompressorSpeed() (int, error) {
	return m.CpSpeed, m.Error
}

func (m *MockModbus) CompressorRunning() (bool, error) {
	return m.CpRunning, m.Error
}

func (m *MockModbus) IndoorTemperature() (float32, error) {
	return m.IndoorTemp, m.Error
}

func (m *MockModbus) OutsideTemperature() (float32, error) {
	return m.OutTemp, m.Error
}

func (m *MockModbus) InletTemperature() (float32, error) {
	return m.InletTemp, m.Error
}

func (m *MockModbus) OutletTemperature() (float32, error) {
	return m.OutletTemp, m.Error
}

func (m *MockModbus) FlowRate() (float32, error) {
	return m.FR, m.Error
}

func (m *MockModbus) UnderflowTargetTemperature() (float32, error) {
	return m.UFTemp, m.Error
}

func createMCPRequest(args map[string]any) mcp.CallToolRequest {
	return mcp.CallToolRequest{
		Params: struct {
			Name      string         `json:"name"`
			Arguments map[string]any `json:"arguments,omitempty"`
			Meta      *struct {
				ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
			} `json:"_meta,omitempty"`
		}{
			Arguments: args,
		},
	}
}
