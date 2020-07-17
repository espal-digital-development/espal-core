package validators

import (
	"strings"
)

func (f *Form) renderTextAreaField(field FormField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)
	out.WriteString(`<textarea name="`)
	out.WriteString(field.Name())
	out.WriteString(`"`)
	if field.Placeholder() != "" {
		out.WriteString(` placeholder="`)
		out.WriteString(field.Placeholder())
		out.WriteString(`"`)
	}
	if !field.Optional() {
		out.WriteString(` required`)
	}
	out.WriteString(`>`)
	if field.Value() != "" {
		out.WriteString(field.Value())
	}
	out.WriteString(`</textarea>`)
	return out.String()
}
