package validators

// NewDateTimeField returns a new instance of FormField with the type DateTime.
func (form *Form) NewDateTimeField(name string) FormField {
	return form.defaultChecks(&formField{
		name:  name,
		_type: DateTimeFormField,
	})
}

// SetMinYear sets the form field min year.
func (formField *formField) SetMinYear(minYear uint) {
	formField.minYear = minYear
}

// MinYear gets the form field min year.
func (formField *formField) MinYear() uint {
	return formField.minYear
}

// SetMaxYear sets the form field max year.
func (formField *formField) SetMaxYear(maxYear uint) {
	formField.maxYear = maxYear
}

// MaxYear gets the form field max year.
func (formField *formField) MaxYear() uint {
	return formField.maxYear
}

// SetMinMonth sets the form field min month.
func (formField *formField) SetMinMonth(minMonth uint) {
	formField.minMonth = minMonth
}

// MinMonth gets the form field min month.
func (formField *formField) MinMonth() uint {
	return formField.minMonth
}

// SetMaxMonth sets the form field max month.
func (formField *formField) SetMaxMonth(maxMonth uint) {
	formField.maxMonth = maxMonth
}

// MaxMonth gets the form field max month.
func (formField *formField) MaxMonth() uint {
	return formField.maxMonth
}

// SetMinDay sets the form field min day.
func (formField *formField) SetMinDay(minDay uint) {
	formField.minDay = minDay
}

// MinDay gets the form field min day.
func (formField *formField) MinDay() uint {
	return formField.minDay
}

// SetMaxDay sets the form field max day.
func (formField *formField) SetMaxDay(maxDay uint) {
	formField.maxDay = maxDay
}

// MaxDay gets the form field max day.
func (formField *formField) MaxDay() uint {
	return formField.maxDay
}

// SetMaxHour sets the form field max hour.
func (formField *formField) SetMaxHour(maxHour uint) {
	formField.maxHour = maxHour
}

// MaxDay gets the form field max hour.
func (formField *formField) MaxHour() uint {
	return formField.maxHour
}

// SetMinHour sets the form field min hour.
func (formField *formField) SetMinHour(minHour uint) {
	formField.minDay = minHour
}

// MinHour gets the form field min hour.
func (formField *formField) MinHour() uint {
	return formField.minHour
}

// SetMaxMinute sets the form field max minute.
func (formField *formField) SetMaxMinute(maxMinute uint) {
	formField.maxMinute = maxMinute
}

// MaxMinute gets the form field max minute.
func (formField *formField) MaxMinute() uint {
	return formField.maxMinute
}

// SetMinMinute sets the form field min minute.
func (formField *formField) SetMinMinute(minMinute uint) {
	formField.minMinute = minMinute
}

// MinMinute gets the form field min minute.
func (formField *formField) MinMinute() uint {
	return formField.minMinute
}

// SetMaxSecond sets the form field max second.
func (formField *formField) SetMaxSecond(maxSecond uint) {
	formField.maxSecond = maxSecond
}

// MaxSecond gets the form field max second.
func (formField *formField) MaxSecond() uint {
	return formField.maxSecond
}

// SetMinSecond sets the form field min second.
func (formField *formField) SetMinSecond(minSecond uint) {
	formField.minSecond = minSecond
}

// MinSecond gets the form field min second.
func (formField *formField) MinSecond() uint {
	return formField.minSecond
}

// SetExlucdeYear marks the field's date to not show and use the year part.
func (formField *formField) SetExcludeYear() {
	formField.excludeYear = true
}

// ExcludeYear returns if the field's date should not show and use the year part.
func (formField *formField) ExcludeYear() bool {
	return formField.excludeYear
}

// SetExcludeMonth marks the field's date to not show and use the month part.
func (formField *formField) SetExcludeMonth() {
	formField.excludeMonth = true
}

// ExcludeMonth returns if the field's date should not show and use the month part.
func (formField *formField) ExcludeMonth() bool {
	return formField.excludeMonth
}

// SetExcludeDay marks the field's date to not show and use the day part.
func (formField *formField) SetExcludeDay() {
	formField.excludeDay = true
}

// ExcludeDay returns if the field's date should not show and use the day part.
func (formField *formField) ExcludeDay() bool {
	return formField.excludeDay
}

// SetExcludeHour marks the field's date to not show and use the hour part.
func (formField *formField) SetExcludeHour() {
	formField.excludeHour = true
}

// ExcludeHour returns if the field's date should not show and use the hour part.
func (formField *formField) ExcludeHour() bool {
	return formField.excludeHour
}

// SetExcludeMinute marks the field's date to not show and use the minute part.
func (formField *formField) SetExcludeMinute() {
	formField.excludeMinute = true
}

// ExcludeMinute returns if the field's date should not show and use the minute part.
func (formField *formField) ExcludeMinute() bool {
	return formField.excludeMinute
}

// SetExcludeSecond marks the field's date to not show and use the second part.
func (formField *formField) SetExcludeSecond() {
	formField.excludeMinute = true
}

// ExcludeSecond returns if the field's date should not show and use the second part.
func (formField *formField) ExcludeSecond() bool {
	return formField.excludeSecond
}
