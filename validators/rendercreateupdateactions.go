package validators

import (
	"strings"
)

// RenderCreateUpdateActions will render all admin create/update
// actions of an admin module overview page.
func (form *Form) RenderCreateUpdateActions(fieldName string, url string) string {
	out := strings.Builder{}

	form.perror(out.WriteString(`<br>`))

	form.perror(out.WriteString(form.renderInputTypeField(form.field(fieldName))))

	form.perror(out.WriteString(`<input type="submit" value="`))
	form.perror(out.WriteString(form.translationsRepository.Singular(form.language.ID(), "save")))
	form.perror(out.WriteString(`" class="save"><br>`))

	form.perror(out.WriteString(`<input type="submit" value="`))
	form.perror(out.WriteString(form.translationsRepository.Singular(form.language.ID(), "saveAndReturn")))
	form.perror(out.WriteString(`" class="save" data-a="1"><br>`))

	form.perror(out.WriteString(`<input type="submit" value="`))
	form.perror(out.WriteString(form.translationsRepository.Singular(form.language.ID(), "saveAndCreateAnother")))
	form.perror(out.WriteString(`" class="save" data-a="2"><br>`))

	form.perror(out.WriteString(`<input type="submit" value="`))
	form.perror(out.WriteString(form.translationsRepository.Singular(form.language.ID(), "saveAndClone")))
	form.perror(out.WriteString(`" class="save" data-a="3"><br>`))

	form.perror(out.WriteString(`<button class="wide"><a href="`))
	form.perror(out.WriteString(url))
	form.perror(out.WriteString(`" class="delete">`))
	form.perror(out.WriteString(form.translationsRepository.Singular(form.language.ID(), "cancel")))
	form.perror(out.WriteString(`</a></button>`))

	return out.String()
}
