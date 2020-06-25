package validators

import (
	"mime/multipart"
)

// UploadedFile is used for single- or multiple uploaded file entries.
type UploadedFile interface {
	Header() *multipart.FileHeader
	SetSanitizedName(string)
	SanitizedName() string
	SetSavedPath(savedPath string)
	SavedPath() string
}

type uploadedFile struct {
	header        *multipart.FileHeader
	sanitizedName string
	savedPath     string
}

// Header returns the uploaded file's header information.
func (f *uploadedFile) Header() *multipart.FileHeader {
	return f.header
}

// SetSanitizedName sets the file name that is safe for filesystem storage.
func (f *uploadedFile) SetSanitizedName(sanitizedName string) {
	f.sanitizedName = sanitizedName
}

// SanitizedName returns the file name that is safe for filesystem storage.
func (f *uploadedFile) SanitizedName() string {
	return f.sanitizedName
}

// SetSavedPath sets the location path the uploaded file is saved at.
func (f *uploadedFile) SetSavedPath(savedPath string) {
	f.savedPath = savedPath
}

// SavedPath returns the location path the uploaded file is saved at.
func (f *uploadedFile) SavedPath() string {
	return f.savedPath
}
