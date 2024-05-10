package actions

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/wjojf/go-ssh-tui/internal/types"
	"testing"
)

func TestAdapterItem(t *testing.T) {

	var item list.Item = FilterableAction{
		Action: types.MockAction{},
	}

	if _, ok := item.(list.DefaultItem); !ok {
		t.Errorf("FilterableAction should implement DefaultItem")
	}

}
