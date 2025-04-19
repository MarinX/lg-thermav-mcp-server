package thermavmodbus

import "github.com/simonvetter/modbus"

// helper function
func readRegister[T any](cli *modbus.ModbusClient, addr uint16, regtype modbus.RegType, converter func(uint16) T) (T, error) {
	var zero T
	val, err := cli.ReadRegister(addr, regtype)
	if err != nil {
		return zero, err
	}

	return converter(val), nil
}

// InletTemperature reads the inlet temperature
func (c *Client) InletTemperature() (float32, error) {
	return readRegister(c.cli, 2, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// OutletTemperature reads the outlet temperature
func (c *Client) OutletTemperature() (float32, error) {
	return readRegister(c.cli, 3, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// FlowRate reads flow rate (L/min)
func (c *Client) FlowRate() (float32, error) {
	return readRegister(c.cli, 8, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// OutsideTemperature reads outdoor temperature
func (c *Client) OutsideTemperature() (float32, error) {
	return readRegister(c.cli, 12, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// CompressorSpeed reads speed of compressor (o/min)
func (c *Client) CompressorSpeed() (int, error) {
	return readRegister(c.cli, 24, modbus.INPUT_REGISTER, func(u uint16) int {
		return int(u * 60)
	})
}

// LiquidGasInTemperature reads the liquid gas temperature
func (c *Client) LiquidGasInTemperature() (float32, error) {
	return readRegister(c.cli, 16, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// LiquidGasOutTemperature reads the liquid gas temperature after passing through the system
func (c *Client) LiquidGasOutTemperature() (float32, error) {
	return readRegister(c.cli, 17, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// HotGasTemperature reads the hot gas temperature
func (c *Client) HotGasTemperature() (float32, error) {
	return readRegister(c.cli, 16, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// SteamBeforeEvaporatorTemperature reads the steam temperature before the evaporator
func (c *Client) SteamBeforeEvaporatorTemperature() (float32, error) {
	return readRegister(c.cli, 20, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// SteamAfterEvaporatorTemperature reads the steam temperature after the evaporator
func (c *Client) SteamAfterEvaporatorTemperature() (float32, error) {
	return readRegister(c.cli, 21, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// PressureCondenser reads the pressure of the condenser in Bar
func (c *Client) PressureCondenser() (float32, error) {
	return readRegister(c.cli, 22, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// VaporPressureEvaporator reads the vapor pressure of the evaporator
func (c *Client) VaporPressureEvaporator() (float32, error) {
	return readRegister(c.cli, 23, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// ExtractionTemperature reads the extraction temperature
func (c *Client) ExtractionTemperature() (float32, error) {
	return readRegister(c.cli, 18, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// UnderflowTargetTemperature reads the underfloor heating target temperature
func (c *Client) UnderflowTargetTemperature() (float32, error) {
	return readRegister(c.cli, 2, modbus.HOLDING_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// DHWTargetTemperature reads the domestic hot water target temperature
func (c *Client) DHWTargetTemperature() (float32, error) {
	return readRegister(c.cli, 8, modbus.HOLDING_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// DHWTemperature reads the domestic hot water temperature
func (c *Client) DHWTemperature() (float32, error) {
	return readRegister(c.cli, 5, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}

// IndoorTemperature reads the indoor temperature (thermostat)
func (c *Client) IndoorTemperature() (float32, error) {
	return readRegister(c.cli, 7, modbus.INPUT_REGISTER, func(u uint16) float32 {
		return float32(u) / 10
	})
}
