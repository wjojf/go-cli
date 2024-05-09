package component

// Tree is a collection of static components.
type Tree map[string]Component

func NewTree() Tree {
	return make(map[string]Component)
}

func (t Tree) Add(name string, c Component) {
	t[name] = c
}

func (t Tree) Render() string {
	var s string
	for _, c := range t {
		s += c.Render()
	}
	return s
}
