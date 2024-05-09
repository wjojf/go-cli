package component

// Component is a static part of the screen that can be rendered.
// It can be a logo, a piece of text, etc.
type Component interface {
	Render() string
}
