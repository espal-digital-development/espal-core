package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

type context interface {
	contexts.RequestContext
	contexts.FormContext
}

// HandleFromRequest processes the core Request when a form was posted.
func (form *Form) HandleFromRequest(context context) error {
	for k := range form.fields {
		if form.fields[k].MaxLength() < form.fields[k].MinLength() {
			form.fields[k].AddError("maxLength cannot be smaller than minLength")
		}

		if HiddenFormField != form.fields[k].Type() && HoneypotFormField != form.fields[k].Type() && TokenFormField != form.fields[k].Type() && !form.fields[k].DontTranslate() && form.fields[k].Placeholder() == "" {
			if form.fields[k].TranslatePlural() {
				form.fields[k].SetPlaceholder(form.translationsRepository.Plural(form.language.ID(), form.fields[k].Name()))
			} else {
				form.fields[k].SetPlaceholder(form.translationsRepository.Singular(form.language.ID(), form.fields[k].Name()))
			}

			if form.fields[k].Optional() {
				form.fields[k].SetPlaceholder(form.fields[k].Placeholder() + " (" + form.translationsRepository.Singular(form.language.ID(), "optional") + ")")
			}
		}

		if ChoiceFormField == form.fields[k].Type() && form.fields[k].NoSelectionText() == "" {
			form.fields[k].DetermineNoSelectText(form.language.ID())
		}
	}

	if form.isFormValidator && context.Method() == "POST" {
		if form.preSubmitCallback != nil {
			form.preSubmitCallback(form)
		}
		if err := form.submit(context); err != nil {
			return errors.Trace(err)
		}
		if form.postSubmitCallback != nil {
			form.postSubmitCallback(form)
		}
	}
	if form.isFormValidator && context.Method() != "POST" {
		if form.postloadCallback != nil {
			form.postloadCallback(form)
		}
	}
	return nil
}

// submit processed the data values that are supplied.
func (form *Form) submit(context contexts.FormContext) error {
	if !form.isFormValidator {
		return errors.Errorf("cannot submit a non-Form type Validator")
	}

	// Invalidate on new Submit
	form.isValid = false
	form.isValidated = false

	// Reset the submit-errors
	form.submitErrors = make([]string, 0)

	for k := range form.fields {
		var err error
		var data string
		var multiData []string
		var dataIsSet bool

		if FileFormField == form.fields[k].Type() {
			if !form.isMultipart() {
				return errors.Errorf("submitted field `%s` is a file, but the form is not multipart", k)
			}
			multiPart, err := context.MultipartForm(1024 * 1024 * 128)
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
			if !form.fields[k].Multiple() && validFiles > 1 {
				return errors.Errorf("file field that is not marked as multiple should not contain multiple files")
			}

			for l := range multiData {
				if multiData[l].Size <= 0 {
					continue
				}
				form.fields[k].AddUploadedFile(&uploadedFile{
					header:        multiData[l],
					sanitizedName: govalidator.SafeFileName(multiData[l].Filename),
				})
			}
		} else {
			if form.isMultipart() {
				multiPart, err := context.MultipartForm(1024 * 1024 * 128)
				if err != nil {
					return errors.Trace(err)
				}
				_, dataIsSet = multiPart.Value[k]
				if dataIsSet {
					values := multiPart.Value[k]
					if form.fields[k].Multiple() {
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
				if form.fields[k].Multiple() {
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
			if CheckboxFormField == form.fields[k].Type() {
				if dataIsSet {
					form.fields[k].SetValue("1")
				} else {
					form.fields[k].SetValue("0")
				}
				continue
			}

			// Empty selects don't get submitted, but need to empty when not
			if !dataIsSet && ChoiceFormField == form.fields[k].Type() {
				if form.fields[k].Multiple() {
					if err := form.fields[k].SetValues(make([]string, 0)); err != nil {
						return errors.Trace(err)
					}
				} else {
					form.fields[k].SetValue("")
				}
			}

			if dataIsSet {
				if form.fields[k].Multiple() {
					form.fields[k].SetValue("")
					multiDataStrings := make([]string, len(multiData))
					copy(multiDataStrings, multiData)
					if err := form.fields[k].SetValues(multiDataStrings); err != nil {
						return errors.Trace(err)
					}
				} else {
					form.fields[k].SetValue(data)
				}
			}
		}
	}

	form.isSubmitted = true

	return nil
}
