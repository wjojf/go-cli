package types

type Action interface {
	Execute() error
	Name() string
}

type ActionList []Action

type MockAction struct{}

func (m MockAction) Execute() error {
	return nil
}

func (m MockAction) Name() string {
	return "Mock Action"
}
