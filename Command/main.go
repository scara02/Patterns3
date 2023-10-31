package main

func main() {
	lamp := &Lamp{}

	turnOn := &TurnOn{
		device: lamp,
	}

	turnOff := &TurnOff{
		device: lamp,
	}

	on := &Switch{
		command: turnOn,
	}
	on.press()

	off := &Switch{
		command: turnOff,
	}
	off.press()
}
