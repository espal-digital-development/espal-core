package validators

import (
	"strings"
)

// RenderCreateUpdateActions will render all admin create/update actions of an admin module overview page.
func (f *Form) RenderCreateUpdateActions(fieldName string, url string) string {
	out := strings.Builder{}

	out.WriteString(`<br>`)

	out.WriteString(f.renderInputTypeField(f.field(fieldName)))

	out.WriteString(`<input type="submit" value="`)
	out.WriteString(f.translationsRepository.Singular(f.language.ID(), "save"))
	out.WriteString(`" class="save"><br>`)

	out.WriteString(`<input type="submit" value="`)
	out.WriteString(f.translationsRepository.Singular(f.language.ID(), "saveAndReturn"))
	out.WriteString(`" class="save" data-a="1"><br>`)

	out.WriteString(`<input type="submit" value="`)
	out.WriteString(f.translationsRepository.Singular(f.language.ID(), "saveAndCreateAnother"))
	out.WriteString(`" class="save" data-a="2"><br>`)

	out.WriteString(`<input type="submit" value="`)
	out.WriteString(f.translationsRepository.Singular(f.language.ID(), "saveAndClone"))
	out.WriteString(`" class="save" data-a="3"><br>`)

	out.WriteString(`<button class="wide"><a href="`)
	out.WriteString(url)
	out.WriteString(`" class="delete">`)
	out.WriteString(f.translationsRepository.Singular(f.language.ID(), "cancel"))
	out.WriteString(`</a></button>`)

	return out.String()
}
