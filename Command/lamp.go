package main

import "fmt"

type Lamp struct {
	isOn bool
}

func (l *Lamp) on() {
	l.isOn = true
	fmt.Println("Lamp is turned on")
}

func (l *Lamp) off() {
	l.isOn = false
	fmt.Println("Lamp is turned off")
}
