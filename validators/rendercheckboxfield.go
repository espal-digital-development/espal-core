package validators

import (
	"strings"
)

func (form *Form) renderCheckboxField(field FormField) string {
	out := strings.Builder{}
	form.addLabel(field, &out)
	form.perror(out.WriteString(`<p><input type="checkbox" name="`))
	form.perror(out.WriteString(field.Name()))
	form.perror(out.WriteString(`"`))
	if field.Value() == "1" {
		form.perror(out.WriteString(` checked`))
	}
	form.perror(out.WriteString(`>`))
	if field.Placeholder() != "" {
		form.perror(out.WriteString(`<span> `))
		form.perror(out.WriteString(field.Placeholder()))
		form.perror(out.WriteString(`</span>`))
	}
	form.perror(out.WriteString(`</p>`))
	return out.String()
}
