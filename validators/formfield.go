package validators

import (
	"strconv"
	"time"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/juju/errors"
)

// Save button actions after success.
const (
	SaveAndEdit   = ""
	SaveAndReturn = "1"
	SaveAndCreate = "2"
	SaveAndClone  = "3"

	defaultFieldMinLength = 7
	defaultFieldMaxLength = 255
)

// FormField is the basic definition of every *FormField.
type FormField interface {
	Type() FieldType
	Name() string
	Value() string
	SetValue(value string)
	Values() []string
	SetValues(values []string) error

	PointerValue() *string
	ValueAsBool() bool
	ValueAsUint() uint
	ValueAsUint16() uint16
	ValueAsTime() *time.Time

	SetPlaceholder(placeholder string)
	Placeholder() string

	SetOptional()
	Optional() bool

	SetTranslatePlural()
	TranslatePlural() bool
	SetDontTranslate()
	DontTranslate() bool

	SetTrim()
	Trim() bool

	SetMinLength(minLength uint)
	MinLength() uint
	SetMaxLength(maxLength uint)
	MaxLength() uint

	SetValidate()
	Validate() bool

	SetHideLabel()
	HideLabel() bool

	SetClass(class string)
	Class() string
	SetID(id string)
	ID() string

	AddError(errorString string)
	HasErrors() bool
	Errors() []string
	RemoveAllErrors()

	SetNeedsToBeEqualToField(targetField string)
	NeedsToBeEqualToField() string

	SetCannotBeEqualToField(targetField string)
	CannotBeEqualToField() string

	SetNoSelectionText(noSelectionText string)
	NoSelectionText() string

	SetAllowedValues(allowValues []string)
	AllowedValues() []string
	SetDisallowValues(disallowValues []string)
	DisallowValues() []string

	// Choice Field
	SetSearchable()
	Searchable() bool
	SetSearchableDataPath(path string)
	SearchableDataPath() string

	DetermineNoSelectText(localeID uint16)

	SetCheckValuesInChoices()
	CheckValuesInChoices() bool

	SetMultiple()
	Multiple() bool

	SetChoices(choices []ChoiceOption)
	AddChoice(choice ChoiceOption)
	ChoiceIsSelected(choiceOption ChoiceOption) bool
	Choices() []ChoiceOption

	// DateTime Field

	SetMinYear(minYear uint)
	MinYear() uint
	SetMaxYear(maxYear uint)
	MaxYear() uint
	SetMinMonth(minMonth uint)
	MinMonth() uint
	SetMaxMonth(maxMonth uint)
	MaxMonth() uint
	SetMinDay(minDay uint)
	MinDay() uint
	SetMaxDay(maxDay uint)
	MaxDay() uint
	SetMinHour(minHour uint)
	MinHour() uint
	SetMaxHour(maxHour uint)
	MaxHour() uint
	SetMinMinute(minMinute uint)
	MinMinute() uint
	SetMaxMinute(maxMinute uint)
	MaxMinute() uint
	SetMinSecond(minSecond uint)
	MinSecond() uint
	SetMaxSecond(maxSecond uint)
	MaxSecond() uint

	SetExcludeYear()
	ExcludeYear() bool
	SetExcludeMonth()
	ExcludeMonth() bool
	SetExcludeDay()
	ExcludeDay() bool
	SetExcludeHour()
	ExcludeHour() bool
	SetExcludeMinute()
	ExcludeMinute() bool
	SetExcludeSecond()
	ExcludeSecond() bool

	// File Field
	SetOptimizeImages()
	OptimizeImages() bool

	SetGzipFilesOnSave()
	GzipFilesOnSave() bool
	SetBrotliFilesOnSave()
	BrotliFilesOnSave() bool

	SetAllowedMIMETypes(allowedMIMETypes []string)

	RemoveUploadedFiles() error
	AddUploadedFile(uploadedFile UploadedFile)
	UploadedFiles() []UploadedFile

	SetFileSaveFolder(fileSaveFolder string)
	FileSaveFolder() string
	SaveFiles() error

	// Number FIeld
	SetMinValue(minValue float64)
	MinValue() float64
	SetMaxValue(maxValue float64)
	MaxValue() float64

	SetStorage(storage storage.Storage) error
}

type formField struct {
	_type            FieldType
	uploadedFiles    []UploadedFile
	choices          []ChoiceOption
	allowedMIMETypes []string // TODO :: Implement
	allowedValues    []string
	disallowValues   []string
	values           []string
	errors           []string
	name             string
	value            string
	// TODO :: 7777 label doesn't seem to be get/set and used?
	// label string
	class string
	id    string
	// noSelectionText is the text that will be displayed to
	// imply the user wants to choose none of the choices.
	noSelectionText       string
	searchableDataPath    string
	needsToBeEqualToField string
	// TODO :: 77 This should actually be []string as it could be desirable to
	// have a field not be equal to multiple fields.
	cannotBeEqualToField string
	placeholder          string
	fileSaveFolder       string
	minValue             float64
	maxValue             float64
	minLength            uint
	maxLength            uint
	minYear              uint
	maxYear              uint
	minMonth             uint
	maxMonth             uint
	minDay               uint
	maxDay               uint
	minHour              uint
	maxHour              uint
	minMinute            uint
	maxMinute            uint
	minSecond            uint
	maxSecond            uint
	// checkValuesInChoices will force a check if all choices
	// are actually present in the choices list.
	// Keep in mind that this only works if the choice(s) is/are
	// already filled. So for a combination like Searchable+Multiple
	// it might not work as expected.
	checkValuesInChoices bool
	searchable           bool
	multiple             bool
	optional             bool
	validate             bool
	hideLabel            bool
	// dontTranslate will make sure the Placeholder
	// wont get auto-translated.
	dontTranslate bool
	// translatePlural will choose FastPlural over
	// fast for the Placeholder.
	translatePlural   bool
	trim              bool
	optimizeImages    bool
	excludeYear       bool
	excludeMonth      bool
	excludeDay        bool
	excludeHour       bool
	excludeMinute     bool
	excludeSecond     bool
	gzipFilesOnSave   bool
	brotliFilesOnSave bool
	storage           storage.Storage
}

// Type returns the form type.
func (f *formField) Type() FieldType {
	return f._type
}

// Name returns the form name.
func (f *formField) Name() string {
	return f.name
}

// Value returns the form field value.
func (f *formField) Value() string {
	return f.value
}

// PointerValue returns the pointer to the form field value.
func (f *formField) PointerValue() *string {
	return &f.value
}

// SetValue sets the form field value.
func (f *formField) SetValue(value string) {
	f.value = value
}

// Values returns the form field values.
func (f *formField) Values() []string {
	return f.values
}

// SetValues sets the form field values.
func (f *formField) SetValues(values []string) error {
	if ChoiceFormField != f._type {
		return errors.Errorf("can only set values on a choice field")
	}
	f.values = values
	return nil
}

// SetPlaceholder sets the form field placeholder value.
func (f *formField) SetPlaceholder(placeholder string) {
	f.placeholder = placeholder
}

// Placeholder returns the field placeholder value.
func (f *formField) Placeholder() string {
	return f.placeholder
}

// SetOptional makes the form field value input optional.
func (f *formField) SetOptional() {
	f.optional = true
}

// SetTranslatePlural makes the form field label translated plural.
func (f *formField) SetTranslatePlural() {
	f.translatePlural = true
}

// TranslatePlural returns if the form field label should be translated plural.
func (f *formField) TranslatePlural() bool {
	return f.translatePlural
}

// SetDontTranslate marks the field to not translate it's label.
func (f *formField) SetDontTranslate() {
	f.dontTranslate = true
}

// DontTranslate returns if the field's label shouldn't be translated.
func (f *formField) DontTranslate() bool {
	return f.dontTranslate
}

// SetTrim marks the field that it should trim the value before it gets processed.
func (f *formField) SetTrim() {
	f.trim = true
}

// Trim returns if the field value needs to be trimmed.
func (f *formField) Trim() bool {
	return f.trim
}

// SetMinLength sets the form field min length.
func (f *formField) SetMinLength(minLength uint) {
	f.minLength = minLength
}

// MinLength returns the minimal required length of the field's value.
func (f *formField) MinLength() uint {
	return f.minLength
}

// SetMaxLength sets the form field max length.
func (f *formField) SetMaxLength(maxLength uint) {
	f.maxLength = maxLength
}

// MaxLength returns the maximum allowed length of the field's value.
func (f *formField) MaxLength() uint {
	return f.maxLength
}

// SetValidate marks the form field should be validated according to it's type.
func (f *formField) SetValidate() {
	f.validate = true
}

// Validate returns if the validators for the field should be triggered.
func (f *formField) Validate() bool {
	return f.validate
}

// SetHideLabel marks the form field to not render the label.
func (f *formField) SetHideLabel() {
	f.hideLabel = true
}

// HideLabel returns if the field's label should be rendered on the page.
func (f *formField) HideLabel() bool {
	return f.hideLabel
}

// SetClass sets the class that will be rendered on the HTML field.
func (f *formField) SetClass(class string) {
	f.class = class
}

// Class returns the class that will be rendered on the HTML field.
func (f *formField) Class() string {
	return f.class
}

// SetID sets the id that will be rendered on the HTML field.
func (f *formField) SetID(id string) {
	f.id = id
}

// ID returns the id that will be rendered on the HTML field.
func (f *formField) ID() string {
	return f.id
}

// SetCannotBeEqualToField sets which field field this form field cannot be equal to.
func (f *formField) SetCannotBeEqualToField(field string) {
	f.cannotBeEqualToField = field
}

// CannotBeEqualToField gets which field field this form field cannot be equal to.
func (f *formField) CannotBeEqualToField() string {
	return f.cannotBeEqualToField
}

// SetNeedsToBeEqualToField sets which field field this form field should be equal to.
func (f *formField) SetNeedsToBeEqualToField(field string) {
	f.needsToBeEqualToField = field
}

// AddError adds a field-error to the Errors stack.
func (f *formField) AddError(errorString string) {
	f.errors = append(f.errors, errorString)
}

// HasErrors indicates if the field has errors.
func (f *formField) HasErrors() bool {
	return len(f.errors) > 0
}

// Errors returns the errors of the field.
func (f *formField) Errors() []string {
	return f.errors
}

// RemoveAllErrors purges all errors from the field.
func (f *formField) RemoveAllErrors() {
	f.errors = make([]string, 0)
}

// ValueAsBool converts the string-value and returns it as bool.
func (f *formField) ValueAsBool() bool {
	return f.value == "1"
}

// ValueAsUint converts the string-value and returns it as uint.
func (f *formField) ValueAsUint() uint {
	i, err := strconv.ParseUint(f.value, 10, 64)
	if err != nil {
		return 0
	}
	return uint(i)
}

// ValueAsUint16 converts the string-value and returns it as uint16.
func (f *formField) ValueAsUint16() uint16 {
	i, err := strconv.ParseUint(f.value, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(i)
}

// ValueAsTime converts the string-value and returns it as uint16.
func (f *formField) ValueAsTime() *time.Time {
	formattedTime, err := time.Parse(time.RFC3339[0:9], f.value)
	if err != nil {
		panic(errors.ErrorStack(err))
	}
	return &formattedTime
}

// SetToBeEqualToField marks the field to be equal to the given targetField name.
func (f *formField) SetToBeEqualToField(targetField string) {
	f.needsToBeEqualToField = targetField
}

// NeedsToBeEqualToField returns the field this field should be equal to.
func (f *formField) NeedsToBeEqualToField() string {
	return f.needsToBeEqualToField
}

// SetNotToBeEqualToField marks the field to not to be equal to the given targetField name.
func (f *formField) SetNotToBeEqualToField(targetField string) {
	f.cannotBeEqualToField = targetField
}

// NeedsToBeEqualToField returns the field this field should not be equal to.
func (f *formField) CannotToBeEqualToField() string {
	return f.cannotBeEqualToField
}

// Optional returns an indicator if the field's filling is optional.
func (f *formField) Optional() bool {
	return f.optional
}

// SetNoSelectionText sets the no-selection option entry text.
func (f *formField) SetNoSelectionText(noSelectionText string) {
	f.noSelectionText = noSelectionText
}

// NoSelectionText returns the no-selection option entry text.
func (f *formField) NoSelectionText() string {
	return f.noSelectionText
}

// SetAllowedValues sets the field's allowed values.
func (f *formField) SetAllowedValues(allowValues []string) {
	f.allowedValues = allowValues
}

// AllowedValues returns the field's allowed values.
func (f *formField) AllowedValues() []string {
	return f.allowedValues
}

// SetDisallowValues sets the field's disallowed values.
func (f *formField) SetDisallowValues(disallowValues []string) {
	f.disallowValues = disallowValues
}

// DisallowValues returns the field's disallowed values.
func (f *formField) DisallowValues() []string {
	return f.disallowValues
}
