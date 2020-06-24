package validators

import (
	"sort"
	"strconv"
)

type options []ChoiceOption

// Len returns the current options length.
func (options options) Len() int {
	return len(options)
}

// Less returns if the first value at the index is less than the second.
func (options options) Less(i, j int) bool {
	return options[i].Display() < options[j].Display()
}

// Swap swaps the values at the given indexes.
func (options options) Swap(i, j int) {
	options[i], options[j] = options[j], options[i]
}

// GetCountryOptionsForLanguage returns all countries translated in the given language.
func (validators *Validators) GetCountryOptionsForLanguage(language language) []ChoiceOption {
	options, ok := validators.countryChoicesCache.Load(language.Code())
	if !ok {
		entries := validators.countriesRepository.All()
		options = make([]ChoiceOption, 0, len(entries))
		for k := range entries {
			options = append(options, &choiceOption{
				value:   strconv.FormatUint(uint64(entries[k].ID()), 10),
				display: entries[k].Translate(language.ID()),
			})
		}
		sort.Sort(options)
		validators.countryChoicesCache.Store(language.Code(), []ChoiceOption(options))
	}
	return []ChoiceOption(options)
}

// GetLanguageOptionsForLanguage returns all languages translated in the given language.
func (validators *Validators) GetLanguageOptionsForLanguage(language language) []ChoiceOption {
	options, ok := validators.languageChoicesCache.Load(language.Code())
	if !ok {
		entries := validators.languagesRepository.All()
		options = make([]ChoiceOption, 0, len(entries))
		for k := range entries {
			options = append(options, &choiceOption{
				value:   strconv.FormatUint(uint64(entries[k].ID()), 10),
				display: entries[k].Translate(language.ID()),
			})
		}
		sort.Sort(options)
		validators.languageChoicesCache.Store(language.Code(), []ChoiceOption(options))
	}
	return []ChoiceOption(options)
}

// GetCurrencyOptionsForLanguage returns all currencies translated in the given language.
func (validators *Validators) GetCurrencyOptionsForLanguage(language language) []ChoiceOption {
	options, ok := validators.currencyChoicesCache.Load(language.Code())
	if !ok {
		entries := validators.currenciesRepository.All()
		options = make([]ChoiceOption, 0, len(entries))
		for k := range entries {
			options = append(options, &choiceOption{
				value:   strconv.FormatUint(uint64(entries[k].ID()), 10),
				display: entries[k].Translate(language.ID()) + " (" + entries[k].Symbol() + ")",
			})
		}
		sort.Sort(options)
		validators.currencyChoicesCache.Store(language.Code(), []ChoiceOption(options))
	}
	return []ChoiceOption(options)
}
