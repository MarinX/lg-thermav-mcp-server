package lgthermav

import (
	"context"
	"fmt"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDHW(t *testing.T) {
	tests := []struct {
		name           string
		runFor         func() server.ToolHandlerFunc
		expectError    bool
		expectedErrMsg string
		expectedOutput string
	}{
		{
			name: "get current DHW temperature",
			runFor: func() server.ToolHandlerFunc {
				_, h := getDHWTemperature(&MockModbus{DHWCurrentTemp: 25.2})
				return h
			},
			expectedOutput: "25.2°C",
		},
		{
			name: "get target DHW temperature",
			runFor: func() server.ToolHandlerFunc {
				_, h := getDHWTargetTemperature(&MockModbus{DHWTargetTemp: 51.1})
				return h
			},
			expectedOutput: "51.1°C",
		},
		{
			name: "get DHW status",
			runFor: func() server.ToolHandlerFunc {
				_, h := isDHWRunning(&MockModbus{DHWRunning: true})
				return h
			},
			expectedOutput: "true",
		},
		{
			name: "handle error",
			runFor: func() server.ToolHandlerFunc {
				_, h := getDHWTemperature(&MockModbus{Error: fmt.Errorf("Mock error")})
				return h
			},
			expectError:    true,
			expectedErrMsg: "Mock error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := tc.runFor()
			request := createMCPRequest(nil)

			result, err := handler(context.Background(), request)
			if tc.expectError {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErrMsg)
				return
			}
			require.NoError(t, err)
			if !assert.Equal(t, 1, len(result.Content)) {
				return
			}

			textContent := result.Content[0].(mcp.TextContent)
			assert.Equal(t, tc.expectedOutput, textContent.Text)
		})
	}
}
