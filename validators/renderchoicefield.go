package validators

import (
	"strings"
)

func (f *Form) renderChoiceField(field *formField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)

	if field.Searchable() {
		// This parent box is used for the searchResults to be found from the input
		f.perror(out.WriteString(`<div>`))

		f.perror(out.WriteString(`<input type="text" search-for="`))
		f.perror(out.WriteString(field.Name()))
		f.perror(out.WriteString(`"`))
		if field.SearchableDataPath() != "" {
			f.perror(out.WriteString(` url="`))
			f.perror(out.WriteString(field.SearchableDataPath()))
			f.perror(out.WriteString(`"`))
		}
		f.perror(out.WriteString(` placeholder="`))
		f.perror(out.WriteString(field.Placeholder()))
		f.perror(out.WriteString(`"`))
		if field.Multiple() && field.Value() != "" {
			f.perror(out.WriteString(` style="display: none;"`))
		}
		f.perror(out.WriteString(`>`))

		f.perror(out.WriteString(`<div>`))
		f.perror(out.WriteString(`<span class="selectedBadge"><span class="selectedValue"></span><span class="clearSelectedValue">x</span></span>`))
		for _, choice := range field.Choices() {
			if !field.ChoiceIsSelected(choice) {
				continue
			}
			f.perror(out.WriteString(`<span class="selectedBadge"`))
			if choice.Value() != "" {
				f.perror(out.WriteString(` style="display: inline-block;" data-id="`))
				f.perror(out.WriteString(choice.Value()))
				f.perror(out.WriteString(`"`))
			}
			f.perror(out.WriteString(`>`))

			f.perror(out.WriteString(`<span class="selectedValue">`))
			f.perror(out.WriteString(choice.Display()))
			f.perror(out.WriteString(`</span>`))

			f.perror(out.WriteString(`<span class="clearSelectedValue">x</span>`))

			f.perror(out.WriteString(`</span>`))
		}
		f.perror(out.WriteString(`</div>`))

		f.perror(out.WriteString(`<div class="searchResults"></div>`))
		f.perror(out.WriteString(`<div class="noSearchResults">`))
		f.perror(out.WriteString(f.translationsRepository.Singular(f.language.ID(), "noResultsFound")))
		f.perror(out.WriteString(`</div>`))
	}

	f.perror(out.WriteString(`<select name="`))
	f.perror(out.WriteString(field.Name()))
	f.perror(out.WriteString(`"`))
	if field.Searchable() {
		f.perror(out.WriteString(` class="selectSearch hidden"`))
	}
	if !field.Optional() {
		f.perror(out.WriteString(` multiple`))
	}
	f.perror(out.WriteString(`>`))

	if !field.Multiple() && field.Optional() {
		f.perror(out.WriteString(`<option value="">`))
		if field.NoSelectionText() != "" {
			f.perror(out.WriteString(field.NoSelectionText()))
		}
		f.perror(out.WriteString(`</option>`))
	}
	for _, choice := range field.Choices() {
		f.perror(out.WriteString(`<option value="`))
		f.perror(out.WriteString(choice.Value()))
		f.perror(out.WriteString(`"`))
		if field.ChoiceIsSelected(choice) {
			f.perror(out.WriteString(` selected`))
		}
		f.perror(out.WriteString(`>`))
		f.perror(out.WriteString(choice.Display()))
		f.perror(out.WriteString(`</option>`))
	}

	f.perror(out.WriteString(`</select>`))

	if field.Searchable() {
		f.perror(out.WriteString(`</div>`))
	}

	return out.String()
}
