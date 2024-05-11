package actions

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/wjojf/go-ssh-tui/internal/types/action"
	"testing"
)

func TestAdapterItem(t *testing.T) {

	var item list.Item = ItemAction{
		Action: action.MockAction{},
	}

	if _, ok := item.(list.DefaultItem); !ok {
		t.Errorf("ItemAction should implement DefaultItem")
	}

}
