package main

type Switch struct {
	command Command
}

func (s *Switch) press() {
	s.command.execute()
}
