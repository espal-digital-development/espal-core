package validators

import (
	"strings"
)

// RenderCreateUpdateActions will render all admin create/update
// actions of an admin module overview page.
func (f *Form) RenderCreateUpdateActions(fieldName string, url string) string {
	out := strings.Builder{}

	f.perror(out.WriteString(`<br>`))

	f.perror(out.WriteString(f.renderInputTypeField(f.field(fieldName))))

	f.perror(out.WriteString(`<input type="submit" value="`))
	f.perror(out.WriteString(f.translationsRepository.Singular(f.language.ID(), "save")))
	f.perror(out.WriteString(`" class="save"><br>`))

	f.perror(out.WriteString(`<input type="submit" value="`))
	f.perror(out.WriteString(f.translationsRepository.Singular(f.language.ID(), "saveAndReturn")))
	f.perror(out.WriteString(`" class="save" data-a="1"><br>`))

	f.perror(out.WriteString(`<input type="submit" value="`))
	f.perror(out.WriteString(f.translationsRepository.Singular(f.language.ID(), "saveAndCreateAnother")))
	f.perror(out.WriteString(`" class="save" data-a="2"><br>`))

	f.perror(out.WriteString(`<input type="submit" value="`))
	f.perror(out.WriteString(f.translationsRepository.Singular(f.language.ID(), "saveAndClone")))
	f.perror(out.WriteString(`" class="save" data-a="3"><br>`))

	f.perror(out.WriteString(`<button class="wide"><a href="`))
	f.perror(out.WriteString(url))
	f.perror(out.WriteString(`" class="delete">`))
	f.perror(out.WriteString(f.translationsRepository.Singular(f.language.ID(), "cancel")))
	f.perror(out.WriteString(`</a></button>`))

	return out.String()
}
