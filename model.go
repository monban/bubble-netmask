package netmask

import (
	"fmt"
	"net"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type Model struct {
	Ip   net.IP
	Size int
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
	cidr := m.AsCidr()
	_, mask, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf(
		"%s\n%s\n",
		cidr,
		NetMaskString(mask.Mask))
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
