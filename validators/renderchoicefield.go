package validators

import (
	"strings"
)

func (form *Form) renderChoiceField(field *formField) string {
	out := strings.Builder{}
	form.addLabel(field, &out)

	if field.Searchable() {
		// This parent box is used for the searchResults to be found from the input
		form.perror(out.WriteString(`<div>`))

		form.perror(out.WriteString(`<input type="text" search-for="`))
		form.perror(out.WriteString(field.Name()))
		form.perror(out.WriteString(`"`))
		if field.SearchableDataPath() != "" {
			form.perror(out.WriteString(` url="`))
			form.perror(out.WriteString(field.SearchableDataPath()))
			form.perror(out.WriteString(`"`))
		}
		form.perror(out.WriteString(` placeholder="`))
		form.perror(out.WriteString(field.Placeholder()))
		form.perror(out.WriteString(`"`))
		if field.Multiple() && field.Value() != "" {
			form.perror(out.WriteString(` style="display: none;"`))
		}
		form.perror(out.WriteString(`>`))

		form.perror(out.WriteString(`<div>`))
		form.perror(out.WriteString(`<span class="selectedBadge"><span class="selectedValue"></span><span class="clearSelectedValue">x</span></span>`))
		for _, choice := range field.Choices() {
			if !field.ChoiceIsSelected(choice) {
				continue
			}
			form.perror(out.WriteString(`<span class="selectedBadge"`))
			if choice.Value() != "" {
				form.perror(out.WriteString(` style="display: inline-block;" data-id="`))
				form.perror(out.WriteString(choice.Value()))
				form.perror(out.WriteString(`"`))
			}
			form.perror(out.WriteString(`>`))

			form.perror(out.WriteString(`<span class="selectedValue">`))
			form.perror(out.WriteString(choice.Display()))
			form.perror(out.WriteString(`</span>`))

			form.perror(out.WriteString(`<span class="clearSelectedValue">x</span>`))

			form.perror(out.WriteString(`</span>`))
		}
		form.perror(out.WriteString(`</div>`))

		form.perror(out.WriteString(`<div class="searchResults"></div>`))
		form.perror(out.WriteString(`<div class="noSearchResults">`))
		form.perror(out.WriteString(form.translationsRepository.Singular(form.language.ID(), "noResultsFound")))
		form.perror(out.WriteString(`</div>`))
	}

	form.perror(out.WriteString(`<select name="`))
	form.perror(out.WriteString(field.Name()))
	form.perror(out.WriteString(`"`))
	if field.Searchable() {
		form.perror(out.WriteString(` class="selectSearch hidden"`))
	}
	if !field.Optional() {
		form.perror(out.WriteString(` multiple`))
	}
	form.perror(out.WriteString(`>`))

	if !field.Multiple() && field.Optional() {
		form.perror(out.WriteString(`<option value="">`))
		if field.NoSelectionText() != "" {
			form.perror(out.WriteString(field.NoSelectionText()))
		}
		form.perror(out.WriteString(`</option>`))
	}
	for _, choice := range field.Choices() {
		form.perror(out.WriteString(`<option value="`))
		form.perror(out.WriteString(choice.Value()))
		form.perror(out.WriteString(`"`))
		if field.ChoiceIsSelected(choice) {
			form.perror(out.WriteString(` selected`))
		}
		form.perror(out.WriteString(`>`))
		form.perror(out.WriteString(choice.Display()))
		form.perror(out.WriteString(`</option>`))
	}

	form.perror(out.WriteString(`</select>`))

	if field.Searchable() {
		form.perror(out.WriteString(`</div>`))
	}

	return out.String()
}
