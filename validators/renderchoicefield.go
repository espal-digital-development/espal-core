package validators

import (
	"strings"
)

// nolint:funlen
func (f *Form) renderChoiceField(field *formField) string {
	out := strings.Builder{}
	f.addLabel(field, &out)

	if field.Searchable() {
		// This parent box is used for the searchResults to be found from the input
		out.WriteString(`<div>`)

		out.WriteString(`<input type="text" search-for="`)
		out.WriteString(field.Name())
		out.WriteString(`"`)
		if field.SearchableDataPath() != "" {
			out.WriteString(` url="`)
			out.WriteString(field.SearchableDataPath())
			out.WriteString(`"`)
		}
		out.WriteString(` placeholder="`)
		out.WriteString(field.Placeholder())
		out.WriteString(`"`)
		if field.Multiple() && field.Value() != "" {
			out.WriteString(` style="display: none;"`)
		}
		out.WriteString(`>`)

		out.WriteString(`<div>`)
		out.WriteString(`<span class="selectedBadge"><span class="selectedValue"></span>`)
		out.WriteString(`<span class="clearSelectedValue">x</span></span>`)
		for _, choice := range field.Choices() {
			if !field.ChoiceIsSelected(choice) {
				continue
			}
			out.WriteString(`<span class="selectedBadge"`)
			if choice.Value() != "" {
				out.WriteString(` style="display: inline-block;" data-id="`)
				out.WriteString(choice.Value())
				out.WriteString(`"`)
			}
			out.WriteString(`>`)

			out.WriteString(`<span class="selectedValue">`)
			out.WriteString(choice.Display())
			out.WriteString(`</span>`)

			out.WriteString(`<span class="clearSelectedValue">x</span>`)

			out.WriteString(`</span>`)
		}
		out.WriteString(`</div>`)

		out.WriteString(`<div class="searchResults"></div>`)
		out.WriteString(`<div class="noSearchResults">`)
		out.WriteString(f.translationsRepository.Singular(f.language.ID(), "noResultsFound"))
		out.WriteString(`</div>`)
	}

	out.WriteString(`<select name="`)
	out.WriteString(field.Name())
	out.WriteString(`"`)
	if field.Searchable() {
		out.WriteString(` class="selectSearch hidden"`)
	}
	if !field.Optional() {
		out.WriteString(` multiple`)
	}
	out.WriteString(`>`)

	if !field.Multiple() && field.Optional() {
		out.WriteString(`<option value="">`)
		if field.NoSelectionText() != "" {
			out.WriteString(field.NoSelectionText())
		}
		out.WriteString(`</option>`)
	}
	for _, choice := range field.Choices() {
		out.WriteString(`<option value="`)
		out.WriteString(choice.Value())
		out.WriteString(`"`)
		if field.ChoiceIsSelected(choice) {
			out.WriteString(` selected`)
		}
		out.WriteString(`>`)
		out.WriteString(choice.Display())
		out.WriteString(`</option>`)
	}

	out.WriteString(`</select>`)

	if field.Searchable() {
		out.WriteString(`</div>`)
	}

	return out.String()
}
