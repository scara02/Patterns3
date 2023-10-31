package main

type TurnOn struct {
	device Device
}

func (c *TurnOn) execute() {
	c.device.on()
}
