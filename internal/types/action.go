package types

var (
	DefaultActionList = ActionList{
		MockAction{},
		MockAction{},
		MockAction{},
	}
)

type Action interface {
	Name() string
	Description() string
}

type MockAction struct{}

func (m MockAction) Name() string {
	return "MockAction"
}

func (m MockAction) Description() string {
	return "MockAction description"
}

type ActionList []Action
