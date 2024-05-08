package optionslist

import (
	"fmt"

	"github.com/wjojf/go-ssh-tui/pkg/component"
	"github.com/wjojf/go-ssh-tui/pkg/utils"
)

type OptionsList[T utils.Stringable] struct {
	component.Base

	options  []T
	selected map[int]struct{}

	cursor int
}

func NewOptionsList[T utils.Stringable](opts Opts[T]) *OptionsList[T] {
	l := &OptionsList[T]{
		options:  opts.Options,
		selected: make(map[int]struct{}),
	}

	if opts.Style != nil {
		l.SetStyle(*opts.Style)
	}

	return l
}

// Render implements the component interface
func (l *OptionsList[T]) Render() string {

	var s string

	for index, option := range l.options {

		_, selected := l.selected[index]

		s += l.renderSingleItem(selected, option)
		if index < len(l.options)-1 {
			s += "\n"
		}
	}

	return s
}

func (l *OptionsList[T]) renderSingleItem(selected bool, option T) string {

	var selectedStr = " "
	if selected {
		selectedStr = "X"
	}

	s := fmt.Sprintf("[ %v ] %v", selectedStr, option.String())

	return l.Style.Render(s)
}

func (l *OptionsList[T]) Toggle(index int) {
	if _, ok := l.selected[index]; ok {
		delete(l.selected, index)
		return
	}

	l.selected[index] = struct{}{}
}

func (l *OptionsList[T]) Up() int {
	if l.cursor > 0 {
		l.cursor--
	}
	return l.cursor
}

func (l *OptionsList[T]) Down() int {
	if l.cursor < len(l.options)-1 {
		l.cursor++
	}
	return l.cursor
}

func (l *OptionsList[T]) Selected() []T {
	var selected []T
	for index := range l.selected {
		selected = append(selected, l.options[index])
	}
	return selected
}
