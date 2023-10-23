package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	netmask "github.com/monban/bubble-netmask"
)

func main() {
	p := tea.NewProgram(netmask.New("192.168.1.0"))

	var m tea.Model
	var err error

	if m, err = p.Run(); err != nil {
		log.Fatal(err)
	}

	model := m.(netmask.Model)
	log.Info("netmask selected", "AsMask", model.AsMask(), "AsCIDR", model.AsCidr())
}
