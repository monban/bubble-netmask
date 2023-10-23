package netmask

import (
	"fmt"
	"net"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Size  int
	IpNet net.IPNet
}

func New(networkAddress string) Model {
	var m Model
	m.IpNet.IP = net.ParseIP(networkAddress)
	m.IpNet.Mask = m.IpNet.IP.DefaultMask()
	sz, _ := m.IpNet.Mask.Size()
	m.Size = sz
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			m.Size--
		case "right", "l":
			m.Size++
		case "enter":
			return m, nil
		}
	}
	if m.Size < 0 {
		m.Size = 0
	} else if m.Size > 32 {
		m.Size = 32
	}
	m.IpNet.Mask = net.CIDRMask(m.Size, 32)
	return m, nil
}

func (m Model) View() string {
	var left string
	var right string

	if m.Size == 0 {
		left = "|\n|"
	} else {
		left = "<\n<"
	}

	if m.Size == 32 {
		right = "|\n|"
	} else {
		right = ">\n>"
	}
	data := fmt.Sprintf(" %15s \n %15s ", &m.IpNet, NetMaskString(m.IpNet.Mask))
	str := lipgloss.JoinHorizontal(lipgloss.Center, left, data, right)

	return str
}

func (m Model) AsMask() string {
	return NetMaskString(m.IpNet.Mask)
}

func (m Model) Run() (net.IPMask, error) {
	var n tea.Model
	var err error

	p := tea.NewProgram(m)
	if n, err = p.Run(); err != nil {
		return nil, fmt.Errorf("getting netmask: %w", err)
	}
	m = n.(Model)

	return m.IpNet.Mask, nil
}

func NetMaskString(mask net.IPMask) string {
	return net.IP(mask).String()
}
