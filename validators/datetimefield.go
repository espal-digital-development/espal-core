package validators

// NewDateTimeField returns a new instance of FormField with the type DateTime.
func (f *Form) NewDateTimeField(name string) FormField {
	return f.defaultChecks(&formField{
		name:  name,
		_type: DateTimeFormField,
	})
}

// SetMinYear sets the form field min year.
func (f *formField) SetMinYear(minYear uint) {
	f.minYear = minYear
}

// MinYear gets the form field min year.
func (f *formField) MinYear() uint {
	return f.minYear
}

// SetMaxYear sets the form field max year.
func (f *formField) SetMaxYear(maxYear uint) {
	f.maxYear = maxYear
}

// MaxYear gets the form field max year.
func (f *formField) MaxYear() uint {
	return f.maxYear
}

// SetMinMonth sets the form field min month.
func (f *formField) SetMinMonth(minMonth uint) {
	f.minMonth = minMonth
}

// MinMonth gets the form field min month.
func (f *formField) MinMonth() uint {
	return f.minMonth
}

// SetMaxMonth sets the form field max month.
func (f *formField) SetMaxMonth(maxMonth uint) {
	f.maxMonth = maxMonth
}

// MaxMonth gets the form field max month.
func (f *formField) MaxMonth() uint {
	return f.maxMonth
}

// SetMinDay sets the form field min day.
func (f *formField) SetMinDay(minDay uint) {
	f.minDay = minDay
}

// MinDay gets the form field min day.
func (f *formField) MinDay() uint {
	return f.minDay
}

// SetMaxDay sets the form field max day.
func (f *formField) SetMaxDay(maxDay uint) {
	f.maxDay = maxDay
}

// MaxDay gets the form field max day.
func (f *formField) MaxDay() uint {
	return f.maxDay
}

// SetMaxHour sets the form field max hour.
func (f *formField) SetMaxHour(maxHour uint) {
	f.maxHour = maxHour
}

// MaxDay gets the form field max hour.
func (f *formField) MaxHour() uint {
	return f.maxHour
}

// SetMinHour sets the form field min hour.
func (f *formField) SetMinHour(minHour uint) {
	f.minDay = minHour
}

// MinHour gets the form field min hour.
func (f *formField) MinHour() uint {
	return f.minHour
}

// SetMaxMinute sets the form field max minute.
func (f *formField) SetMaxMinute(maxMinute uint) {
	f.maxMinute = maxMinute
}

// MaxMinute gets the form field max minute.
func (f *formField) MaxMinute() uint {
	return f.maxMinute
}

// SetMinMinute sets the form field min minute.
func (f *formField) SetMinMinute(minMinute uint) {
	f.minMinute = minMinute
}

// MinMinute gets the form field min minute.
func (f *formField) MinMinute() uint {
	return f.minMinute
}

// SetMaxSecond sets the form field max second.
func (f *formField) SetMaxSecond(maxSecond uint) {
	f.maxSecond = maxSecond
}

// MaxSecond gets the form field max second.
func (f *formField) MaxSecond() uint {
	return f.maxSecond
}

// SetMinSecond sets the form field min second.
func (f *formField) SetMinSecond(minSecond uint) {
	f.minSecond = minSecond
}

// MinSecond gets the form field min second.
func (f *formField) MinSecond() uint {
	return f.minSecond
}

// SetExlucdeYear marks the field's date to not show and use the year part.
func (f *formField) SetExcludeYear() {
	f.excludeYear = true
}

// ExcludeYear returns if the field's date should not show and use the year part.
func (f *formField) ExcludeYear() bool {
	return f.excludeYear
}

// SetExcludeMonth marks the field's date to not show and use the month part.
func (f *formField) SetExcludeMonth() {
	f.excludeMonth = true
}

// ExcludeMonth returns if the field's date should not show and use the month part.
func (f *formField) ExcludeMonth() bool {
	return f.excludeMonth
}

// SetExcludeDay marks the field's date to not show and use the day part.
func (f *formField) SetExcludeDay() {
	f.excludeDay = true
}

// ExcludeDay returns if the field's date should not show and use the day part.
func (f *formField) ExcludeDay() bool {
	return f.excludeDay
}

// SetExcludeHour marks the field's date to not show and use the hour part.
func (f *formField) SetExcludeHour() {
	f.excludeHour = true
}

// ExcludeHour returns if the field's date should not show and use the hour part.
func (f *formField) ExcludeHour() bool {
	return f.excludeHour
}

// SetExcludeMinute marks the field's date to not show and use the minute part.
func (f *formField) SetExcludeMinute() {
	f.excludeMinute = true
}

// ExcludeMinute returns if the field's date should not show and use the minute part.
func (f *formField) ExcludeMinute() bool {
	return f.excludeMinute
}

// SetExcludeSecond marks the field's date to not show and use the second part.
func (f *formField) SetExcludeSecond() {
	f.excludeMinute = true
}

// ExcludeSecond returns if the field's date should not show and use the second part.
func (f *formField) ExcludeSecond() bool {
	return f.excludeSecond
}
