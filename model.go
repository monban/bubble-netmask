package netmask

import (
	"fmt"
	"net"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type Model struct {
	Ip          net.IP
	Size        int
	ActiveStyle lipgloss.Style
	Width       int
}

func New(networkAddress string) Model {
	var m Model
	m.Ip = net.ParseIP(networkAddress)
	i := net.IPMask(m.Ip.DefaultMask())
	_, sz := i.Size()
	m.Size = sz
	m.Width = 24
	m.ActiveStyle = lipgloss.NewStyle().
		Margin(0).
		Padding(0).
		Height(2).
		Width(24).
		Align(lipgloss.Center).
		BorderBackground(lipgloss.Color("63")).
		BorderForeground(lipgloss.Color("223")).
		BorderTop(false).
		BorderRight(true).
		BorderBottom(false).
		BorderLeft(true)

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
	return m, nil
}

func (m Model) View() string {
	var left string
	var right string
	cidr := m.AsCidr()
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatal(err)
	}
	mask := NetMaskString(ipnet.Mask)

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
	data := fmt.Sprintf(" %15s \n %15s ", cidr, mask)
	str := lipgloss.JoinHorizontal(lipgloss.Center, left, data, right)

	return str
}

func (m Model) AsCidr() string {
	return fmt.Sprintf("%s/%d", m.Ip.String(), m.Size)
}

func (m Model) AsMask() string {
	_, ipnet, err := net.ParseCIDR(m.AsCidr())
	if err != nil {
		log.Fatal(err)
	}
	return net.IP(ipnet.Mask).String()
}

func mask(ip net.IP, i int) {
	net.ParseCIDR(fmt.Sprintf("%s/%d", ip.String(), i))
}

func NetMaskString(mask net.IPMask) string {
	return net.IP(mask).String()
}
