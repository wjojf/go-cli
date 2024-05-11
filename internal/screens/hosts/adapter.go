package hosts

import "github.com/wjojf/go-ssh-tui/internal/types/host"

type ItemHost struct {
	Host host.Host
}

func (h ItemHost) FilterValue() string {
	return string(h.Host)
}

func (h ItemHost) Title() string {
	return string(h.Host)
}

func (h ItemHost) Description() string {
	ip, ok := h.Host.IP()
	if !ok {
		return "No IP found"
	}
	return ip
}
