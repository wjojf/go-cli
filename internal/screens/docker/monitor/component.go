package monitor

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		AlignHorizontal(lipgloss.Center)
)

func (m *Model) refreshTable() table.Model {
	columns := getColumns()
	rows := getRows(m.data)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

func getColumns() []table.Column {
	return []table.Column{
		{Title: "Name", Width: 25},
		{Title: "Status", Width: 10},
		{Title: "CPU", Width: 5},
		{Title: "Memory", Width: 7},
		{Title: "Network", Width: 7},
	}
}

func getRows(data []MonitorStatRow) []table.Row {
	var rows []table.Row
	for _, d := range data {
		rows = append(rows, d.Row())
	}
	return rows
}
