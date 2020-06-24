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
func (uploadedFile *uploadedFile) Header() *multipart.FileHeader {
	return uploadedFile.header
}

// SetSanitizedName sets the file name that is safe for filesystem storage.
func (uploadedFile *uploadedFile) SetSanitizedName(sanitizedName string) {
	uploadedFile.sanitizedName = sanitizedName
}

// SanitizedName returns the file name that is safe for filesystem storage.
func (uploadedFile *uploadedFile) SanitizedName() string {
	return uploadedFile.sanitizedName
}

// SetSavedPath sets the location path the uploaded file is saved at.
func (uploadedFile *uploadedFile) SetSavedPath(savedPath string) {
	uploadedFile.savedPath = savedPath
}

// SavedPath returns the location path the uploaded file is saved at.
func (uploadedFile *uploadedFile) SavedPath() string {
	return uploadedFile.savedPath
}
