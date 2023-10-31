package main

type TurnOff struct {
	device Device
}

func (c *TurnOff) execute() {
	c.device.off()
}
