package fakes

// StaticRenderer implements ports.TemplateRenderer by returning pre-configured output.
// Used in tests to provide known template content without actually rendering.
type StaticRenderer struct {
	// Output is the map that Render() will return
	Output map[string][]byte

	// Err is returned by Rende() if set (for testing error paths)
	Err error

	// Captured inputs for test assertions
	CalledWithTemplate string
	CalledWithData     interface{}
}

// NewStaticRenderer creates a renderer that returns a given output.
func NewStaticRenderer(output map[string][]byte) *StaticRenderer {
	return &StaticRenderer{
		Output: output,
	}
}

// Render returns the pre-configured output and records what it was called with.
func (r *StaticRenderer) Render(templateName string, data interface{}) (map[string][]byte, error) {
	r.CalledWithTemplate = templateName
	r.CalledWithData = data

	if r.Err != nil {
		return nil, r.Err
	}

	return r.Output, nil
}
