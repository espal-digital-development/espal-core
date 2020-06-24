package validators

import (
	"strings"
)

func (form *Form) renderTextAreaField(field FormField) string {
	out := strings.Builder{}
	form.addLabel(field, &out)
	form.perror(out.WriteString(`<textarea name="`))
	form.perror(out.WriteString(field.Name()))
	form.perror(out.WriteString(`"`))
	if field.Placeholder() != "" {
		form.perror(out.WriteString(` placeholder="`))
		form.perror(out.WriteString(field.Placeholder()))
		form.perror(out.WriteString(`"`))
	}
	if !field.Optional() {
		form.perror(out.WriteString(` required`))
	}
	form.perror(out.WriteString(`>`))
	if field.Value() != "" {
		form.perror(out.WriteString(field.Value()))
	}
	form.perror(out.WriteString(`</textarea>`))
	return out.String()
}
