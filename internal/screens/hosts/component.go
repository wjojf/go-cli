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

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Choose a host (press Enter to select)"

	l.SetFilteringEnabled(false)
	l.SetShowFilter(false)

	return l
}
