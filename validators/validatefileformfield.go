package validators

import (
	"fmt"
)

func (f *Form) validateFileFormField(field *formField) {
	filesCount := len(field.UploadedFiles())
	if !field.Optional() && filesCount > 0 {
		field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(), "fieldXCannotBeEmpty"), field.Name()))
	} else if filesCount > 0 {
		for _, uploadedFile := range field.UploadedFiles() {
			if field.MaxLength() > 0 && field.MaxLength() < uint(len(uploadedFile.SanitizedName())) {
				field.AddError(fmt.Sprintf(f.translationsRepository.Formatted(f.language.ID(), "fieldXCannotBeLongerThanXCharacters"), field.Name()+": "+uploadedFile.SanitizedName(), field.MaxLength()))
			}
		}
	}
}
