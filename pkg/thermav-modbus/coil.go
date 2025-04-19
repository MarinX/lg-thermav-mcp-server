package thermavmodbus

import "github.com/simonvetter/modbus"

// helper function
func readCoil(cli *modbus.ModbusClient, addr uint16) (bool, error) {
	val, err := cli.ReadDiscreteInput(addr)
	return val, err
}

// BackupHeater1Running checks if the backup heater 1 is running
func (c *Client) BackupHeater1Running() (bool, error) {
	return readCoil(c.cli, 10)
}

// BackupHeating2Running checks if the backup heater 2 is running
func (c *Client) BackupHeating2Running() (bool, error) {
	return readCoil(c.cli, 11)
}

// BoosterHeaterRunning checks if the booster heater is running
func (c *Client) BoosterHeaterRunning() (bool, error) {
	return readCoil(c.cli, 12)
}

// PumpRunning checks if the pump is running
func (c *Client) PumpRunning() (bool, error) {
	return readCoil(c.cli, 1)
}

// CompressorRunning checks if the compressor is running
func (c *Client) CompressorRunning() (bool, error) {
	return readCoil(c.cli, 3)
}

// DefrostRunning checks if the defrost cycle is active
func (c *Client) DefrostRunning() (bool, error) {
	return readCoil(c.cli, 4)
}

// DHWHeatRunning checks if the domestic hot water heating is active
func (c *Client) DHWHeatRunning() (bool, error) {
	return readCoil(c.cli, 5)
}

// HasErrorStatus checks if there is an error status
func (c *Client) HasErrorStatus() (bool, error) {
	return readCoil(c.cli, 13)
}
