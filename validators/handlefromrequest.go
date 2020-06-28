package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

const multiPartMaxMemory = 1024 * 1024 * 128

type context interface {
	contexts.RequestContext
	contexts.FormContext
}

// HandleFromRequest processes the core Request when a form was posted.
func (f *Form) HandleFromRequest(context context) error {
	for k := range f.fields {
		if f.fields[k].MaxLength() < f.fields[k].MinLength() {
			f.fields[k].AddError("maxLength cannot be smaller than minLength")
		}

		if HiddenFormField != f.fields[k].Type() && HoneypotFormField != f.fields[k].Type() &&
			TokenFormField != f.fields[k].Type() && !f.fields[k].DontTranslate() && f.fields[k].Placeholder() == "" {
			if f.fields[k].TranslatePlural() {
				f.fields[k].SetPlaceholder(f.translationsRepository.Plural(f.language.ID(), f.fields[k].Name()))
			} else {
				f.fields[k].SetPlaceholder(f.translationsRepository.Singular(f.language.ID(), f.fields[k].Name()))
			}

			if f.fields[k].Optional() {
				f.fields[k].SetPlaceholder(f.fields[k].Placeholder() + " (" +
					f.translationsRepository.Singular(f.language.ID(), "optional") + ")")
			}
		}

		if ChoiceFormField == f.fields[k].Type() && f.fields[k].NoSelectionText() == "" {
			f.fields[k].DetermineNoSelectText(f.language.ID())
		}
	}

	if f.isFormValidator && context.Method() == "POST" {
		if f.preSubmitCallback != nil {
			f.preSubmitCallback(f)
		}
		if err := f.submit(context); err != nil {
			return errors.Trace(err)
		}
		if f.postSubmitCallback != nil {
			f.postSubmitCallback(f)
		}
	}
	if f.isFormValidator && context.Method() != "POST" {
		if f.postloadCallback != nil {
			f.postloadCallback(f)
		}
	}
	return nil
}

// submit processed the data values that are supplied.
// nolint:gocyclo,funlen
func (f *Form) submit(context contexts.FormContext) error {
	if !f.isFormValidator {
		return errors.Errorf("cannot submit a non-Form type Validator")
	}

	// Invalidate on new Submit
	f.isValid = false
	f.isValidated = false

	// Reset the submit-errors
	f.submitErrors = make([]string, 0)

	for k := range f.fields {
		var err error
		var data string
		var multiData []string
		var dataIsSet bool

		if FileFormField == f.fields[k].Type() {
			if !f.isMultipart() {
				return errors.Errorf("submitted field `%s` is a file, but the form is not multipart", k)
			}
			multiPart, err := context.MultipartForm(multiPartMaxMemory)
			if err != nil {
				return errors.Trace(err)
			}
			multiData, dataIsSet := multiPart.File[k]
			if !dataIsSet {
				continue
			}
			var validFiles int
			for l := range multiData {
				if multiData[l].Size > 0 {
					validFiles++
				}
			}
			if validFiles == 0 {
				continue
			}
			if !f.fields[k].Multiple() && validFiles > 1 {
				return errors.Errorf("file field that is not marked as multiple should not contain multiple files")
			}

			for l := range multiData {
				if multiData[l].Size <= 0 {
					continue
				}
				f.fields[k].AddUploadedFile(&uploadedFile{
					header:        multiData[l],
					sanitizedName: govalidator.SafeFileName(multiData[l].Filename),
				})
			}
		} else {
			if f.isMultipart() {
				multiPart, err := context.MultipartForm(multiPartMaxMemory)
				if err != nil {
					return errors.Trace(err)
				}
				_, dataIsSet = multiPart.Value[k]
				if dataIsSet {
					values := multiPart.Value[k]
					if f.fields[k].Multiple() {
						multiData = make([]string, 0, len(values))
						for l := range values {
							multiData = append(multiData, values[l])
						}
					} else {
						if len(values) > 1 {
							return errors.Errorf("choices field that is not marked as multiple should not contain multiple values")
						}
						if len(values) == 1 {
							data = values[0]
						}
					}
				}
			} else {
				if f.fields[k].Multiple() {
					multiDataCheck, err := context.FormValues(k)
					if err != nil {
						return errors.Trace(err)
					}
					dataIsSet = multiDataCheck != nil && len(multiData) > 0
					if dataIsSet {
						multiData = multiDataCheck
					}
				} else {
					data, err = context.FormValue(k)
					if err != nil {
						return errors.Trace(err)
					}
					dataIsSet = len(data) > 0
				}
			}

			// Can never know if a checkbox wasn't really on the page,
			// as it is not submitted at all as a field when it's unchecked.
			if CheckboxFormField == f.fields[k].Type() {
				if dataIsSet {
					f.fields[k].SetValue("1")
				} else {
					f.fields[k].SetValue("0")
				}
				continue
			}

			// Empty selects don't get submitted, but need to empty when not
			if !dataIsSet && ChoiceFormField == f.fields[k].Type() {
				if f.fields[k].Multiple() {
					if err := f.fields[k].SetValues(make([]string, 0)); err != nil {
						return errors.Trace(err)
					}
				} else {
					f.fields[k].SetValue("")
				}
			}

			if dataIsSet {
				if f.fields[k].Multiple() {
					f.fields[k].SetValue("")
					multiDataStrings := make([]string, len(multiData))
					copy(multiDataStrings, multiData)
					if err := f.fields[k].SetValues(multiDataStrings); err != nil {
						return errors.Trace(err)
					}
				} else {
					f.fields[k].SetValue(data)
				}
			}
		}
	}

	f.isSubmitted = true

	return nil
}
