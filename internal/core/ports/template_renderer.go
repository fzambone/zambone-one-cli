// Package ports defines interfaces for external dependencies.
package ports

// TemplateRenderer defines operations for rendering templates into files.
type TemplateRenderer interface {
	// Render processes a template and returns a map of file paths to content.
	// templateName specifies which template to render.
	// data contains the variables to populate the template.
	Render(templateName string, data interface{}) (map[string][]byte, error)
}
