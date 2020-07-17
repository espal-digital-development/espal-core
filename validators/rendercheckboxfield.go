package validators

import (
	"strings"
)

func (f *Form) renderCheckboxField(field FormField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)
	out.WriteString(`<p><input type="checkbox" name="`)
	out.WriteString(field.Name())
	out.WriteString(`"`)
	if field.Value() == "1" {
		out.WriteString(` checked`)
	}
	out.WriteString(`>`)
	if field.Placeholder() != "" {
		out.WriteString(`<span> `)
		out.WriteString(field.Placeholder())
		out.WriteString(`</span>`)
	}
	out.WriteString(`</p>`)
	return out.String()
}
