package hosts

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/wjojf/go-ssh-tui/internal/types/host"
)

func GetList(hosts ...host.Host) list.Model {

	var items []list.Item
	for _, h := range hosts {
		items = append(items, ItemHost{Host: h})
	}

	return list.New(items, list.NewDefaultDelegate(), 0, 0)
}
