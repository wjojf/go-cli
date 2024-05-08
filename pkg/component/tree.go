package component

type Tree []Component

func NewTree(components ...Component) Tree {

	tree := make(Tree, 0, len(components))
	for _, c := range components {
		tree = append(tree, c)
	}

	return tree
}

func (t Tree) Render() string {
	var s string
	for _, c := range t {
		s += c.Render()
	}
	return s
}
