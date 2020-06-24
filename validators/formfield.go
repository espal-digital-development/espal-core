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
func (formField *formField) Type() FieldType {
	return formField._type
}

// Name returns the form name.
func (formField *formField) Name() string {
	return formField.name
}

// Value returns the form field value.
func (formField *formField) Value() string {
	return formField.value
}

// PointerValue returns the pointer to the form field value.
func (formField *formField) PointerValue() *string {
	return &formField.value
}

// SetValue sets the form field value.
func (formField *formField) SetValue(value string) {
	formField.value = value
}

// Values returns the form field values.
func (formField *formField) Values() []string {
	return formField.values
}

// SetValues sets the form field values.
func (formField *formField) SetValues(values []string) error {
	if ChoiceFormField != formField._type {
		return errors.Errorf("can only set values on a choice field")
	}
	formField.values = values
	return nil
}

// SetPlaceholder sets the form field placeholder value.
func (formField *formField) SetPlaceholder(placeholder string) {
	formField.placeholder = placeholder
}

// Placeholder returns the field placeholder value.
func (formField *formField) Placeholder() string {
	return formField.placeholder
}

// SetOptional makes the form field value input optional.
func (formField *formField) SetOptional() {
	formField.optional = true
}

// SetTranslatePlural makes the form field label translated plural.
func (formField *formField) SetTranslatePlural() {
	formField.translatePlural = true
}

// TranslatePlural returns if the form field label should be translated plural.
func (formField *formField) TranslatePlural() bool {
	return formField.translatePlural
}

// SetDontTranslate marks the field to not translate it's label.
func (formField *formField) SetDontTranslate() {
	formField.dontTranslate = true
}

// DontTranslate returns if the field's label shouldn't be translated.
func (formField *formField) DontTranslate() bool {
	return formField.dontTranslate
}

// SetTrim marks the field that it should trim the value before it gets processed.
func (formField *formField) SetTrim() {
	formField.trim = true
}

// Trim returns if the field value needs to be trimmed.
func (formField *formField) Trim() bool {
	return formField.trim
}

// SetMinLength sets the form field min length.
func (formField *formField) SetMinLength(minLength uint) {
	formField.minLength = minLength
}

// MinLength returns the minimal required length of the field's value.
func (formField *formField) MinLength() uint {
	return formField.minLength
}

// SetMaxLength sets the form field max length.
func (formField *formField) SetMaxLength(maxLength uint) {
	formField.maxLength = maxLength
}

// MaxLength returns the maximum allowed length of the field's value.
func (formField *formField) MaxLength() uint {
	return formField.maxLength
}

// SetValidate marks the form field should be validated according to it's type.
func (formField *formField) SetValidate() {
	formField.validate = true
}

// Validate returns if the validators for the field should be triggered.
func (formField *formField) Validate() bool {
	return formField.validate
}

// SetHideLabel marks the form field to not render the label.
func (formField *formField) SetHideLabel() {
	formField.hideLabel = true
}

// HideLabel returns if the field's label should be rendered on the page.
func (formField *formField) HideLabel() bool {
	return formField.hideLabel
}

// SetClass sets the class that will be rendered on the HTML field.
func (formField *formField) SetClass(class string) {
	formField.class = class
}

// Class returns the class that will be rendered on the HTML field.
func (formField *formField) Class() string {
	return formField.class
}

// SetID sets the id that will be rendered on the HTML field.
func (formField *formField) SetID(id string) {
	formField.id = id
}

// ID returns the id that will be rendered on the HTML field.
func (formField *formField) ID() string {
	return formField.id
}

// SetCannotBeEqualToField sets which field field this form field cannot be equal to.
func (formField *formField) SetCannotBeEqualToField(field string) {
	formField.cannotBeEqualToField = field
}

// CannotBeEqualToField gets which field field this form field cannot be equal to.
func (formField *formField) CannotBeEqualToField() string {
	return formField.cannotBeEqualToField
}

// SetNeedsToBeEqualToField sets which field field this form field should be equal to.
func (formField *formField) SetNeedsToBeEqualToField(field string) {
	formField.needsToBeEqualToField = field
}

// AddError adds a field-error to the Errors stack.
func (formField *formField) AddError(errorString string) {
	formField.errors = append(formField.errors, errorString)
}

// HasErrors indicates if the field has errors.
func (formField *formField) HasErrors() bool {
	return len(formField.errors) > 0
}

// Errors returns the errors of the field.
func (formField *formField) Errors() []string {
	return formField.errors
}

// RemoveAllErrors purges all errors from the field.
func (formField *formField) RemoveAllErrors() {
	formField.errors = make([]string, 0)
}

// ValueAsBool converts the string-value and returns it as bool.
func (formField *formField) ValueAsBool() bool {
	return formField.value == "1"
}

// ValueAsUint converts the string-value and returns it as uint.
func (formField *formField) ValueAsUint() uint {
	i, err := strconv.ParseUint(formField.value, 10, 64)
	if err != nil {
		return 0
	}
	return uint(i)
}

// ValueAsUint16 converts the string-value and returns it as uint16.
func (formField *formField) ValueAsUint16() uint16 {
	i, err := strconv.ParseUint(formField.value, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(i)
}

// ValueAsTime converts the string-value and returns it as uint16.
func (formField *formField) ValueAsTime() *time.Time {
	formattedTime, err := time.Parse(time.RFC3339[0:9], formField.value)
	if err != nil {
		panic(errors.ErrorStack(err))
	}
	return &formattedTime
}

// SetToBeEqualToField marks the field to be equal to the given targetField name.
func (formField *formField) SetToBeEqualToField(targetField string) {
	formField.needsToBeEqualToField = targetField
}

// NeedsToBeEqualToField returns the field this field should be equal to.
func (formField *formField) NeedsToBeEqualToField() string {
	return formField.needsToBeEqualToField
}

// SetNotToBeEqualToField marks the field to not to be equal to the given targetField name.
func (formField *formField) SetNotToBeEqualToField(targetField string) {
	formField.cannotBeEqualToField = targetField
}

// NeedsToBeEqualToField returns the field this field should not be equal to.
func (formField *formField) CannotToBeEqualToField() string {
	return formField.cannotBeEqualToField
}

// Optional returns an indicator if the field's filling is optional.
func (formField *formField) Optional() bool {
	return formField.optional
}

// SetNoSelectionText sets the no-selection option entry text.
func (formField *formField) SetNoSelectionText(noSelectionText string) {
	formField.noSelectionText = noSelectionText
}

// NoSelectionText returns the no-selection option entry text.
func (formField *formField) NoSelectionText() string {
	return formField.noSelectionText
}

// SetAllowedValues sets the field's allowed values.
func (formField *formField) SetAllowedValues(allowValues []string) {
	formField.allowedValues = allowValues
}

// AllowedValues returns the field's allowed values.
func (formField *formField) AllowedValues() []string {
	return formField.allowedValues
}

// SetDisallowValues sets the field's disallowed values.
func (formField *formField) SetDisallowValues(disallowValues []string) {
	formField.disallowValues = disallowValues
}

// DisallowValues returns the field's disallowed values.
func (formField *formField) DisallowValues() []string {
	return formField.disallowValues
}
