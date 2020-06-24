package validators

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"strings"

	"github.com/espal-digital-development/espal-core/storage"
	"github.com/espal-digital-development/espal-core/text"
	"github.com/juju/errors"
)

// NewFileField returns a new instance of FormField with the type File.
func (form *Form) NewFileField(name string, storage storage.Storage) FormField {
	return form.defaultChecks(&formField{
		name:              name,
		_type:             FileFormField,
		storage:           storage,
		gzipFilesOnSave:   form.configService.AssetsGZipFiles(),
		brotliFilesOnSave: form.configService.AssetsBrotliFiles(),
	})
}

// SetOptimizeImages marks that uploaded images should be optimized as best
// as possible for file size. This may vary per platform's available optimizers
// and aren't quality compressors, but only encoder efficiency and meta optimizations.
func (formField *formField) SetOptimizeImages() {
	formField.optimizeImages = true
}

// OptimizeImages returns if the uploaded file(s) should be optimized.
func (formField *formField) OptimizeImages() bool {
	return formField.optimizeImages
}

// SetGzipFilesOnSave marks the files to also be gzip compressed on save.
func (formField *formField) SetGzipFilesOnSave() {
	formField.gzipFilesOnSave = true
}

// GzipFilesOnSave returns if the files should be gzip compressed on save.
func (formField *formField) GzipFilesOnSave() bool {
	return formField.gzipFilesOnSave
}

// SetGzipFilesOnSave marks the files to also be brotli compressed on save.
func (formField *formField) SetBrotliFilesOnSave() {
	formField.brotliFilesOnSave = true
}

// BrotliFilesOnSave returns if the files should be brotli compressed on save.
func (formField *formField) BrotliFilesOnSave() bool {
	return formField.brotliFilesOnSave
}

// SetAllowedMIMETypes sets the allowed MIME types that can be uploaded through the form.
func (formField *formField) SetAllowedMIMETypes(allowedMIMETypes []string) {
	formField.allowedMIMETypes = allowedMIMETypes
}

// RemoveUploadedFiles removes any files that were uploaded on
// the targeted field. This is useful when a form fails after
// the upload and it should be cleaned up.
// It won't delete images already linked on the loaded data.
func (formField *formField) RemoveUploadedFiles() error {
	for k := range formField.uploadedFiles {
		var fullPath string
		if formField.fileSaveFolder != "" {
			fullPath = formField.fileSaveFolder + string(os.PathSeparator)
		}
		fullPath += formField.uploadedFiles[k].SanitizedName()
		if !formField.storage.Exists(fullPath) {
			continue
		}
		if err := formField.storage.Delete(fullPath); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// AddUploadedFile adds a new UploadedFile to the internal files list.
func (formField *formField) AddUploadedFile(uploadedFile UploadedFile) {
	if formField.uploadedFiles == nil {
		formField.uploadedFiles = make([]UploadedFile, 0)
	}
	formField.uploadedFiles = append(formField.uploadedFiles, uploadedFile)
}

// UploadedFiles returns all the field's uploaded files.
func (formField *formField) UploadedFiles() []UploadedFile {
	return formField.uploadedFiles
}

// SetFileSaveFolder sets the specific folder (different than default)
// where to save the file(s) too.
func (formField *formField) SetFileSaveFolder(fileSaveFolder string) {
	formField.fileSaveFolder = fileSaveFolder
}

// FileSaveFolder gets the specific folder (different than default)
// where to save the file(s) too.
func (formField *formField) FileSaveFolder() string {
	return formField.fileSaveFolder
}

// SaveFiles saves the file(s) that were uploaded to the private
// or public folder based on the field.PublicFile setting.
func (formField *formField) SaveFiles() error {
	if FileFormField != formField._type {
		return errors.Errorf("cannot run SaveFiles on a non-FileFormField")
	}
	if formField.storage == nil {
		return errors.Errorf("cannot run SaveFiles without a storage being set")
	}

	// TODO :: 7 When updating, but replacing a/the file(s) should delete the old ones too.
	// This is harder with multiple files tho

	var loopErr error
	savedFiles := make([]string, 0)

	for k := range formField.uploadedFiles {
		var fullPath string
		if formField.fileSaveFolder != "" {
			fullPath = formField.fileSaveFolder + string(os.PathSeparator)
		}
		fullPath += formField.uploadedFiles[k].SanitizedName()

		var generationSuccessful bool
		if formField.storage.Exists(fullPath) {
			var extension string
			var baseName string
			dotSlices := strings.Split(formField.uploadedFiles[k].SanitizedName(), ".")
			length := len(dotSlices)
			if length > 1 && dotSlices[length-1] != "" {
				extension = dotSlices[length-1]
				for i := 0; i < length-1; i++ {
					baseName += dotSlices[i]
				}
			} else {
				baseName = formField.uploadedFiles[k].SanitizedName()
			}

			for i := 0; i < 25; i++ {
				extra := text.RandomString(4)
				newName := baseName + "_" + extra + "." + extension
				var checkPath string
				if formField.fileSaveFolder != "" {
					checkPath = formField.fileSaveFolder + string(os.PathSeparator)
				}
				checkPath += newName
				if formField.storage.Exists(checkPath) {
					continue
				}
				formField.uploadedFiles[k].SetSanitizedName(newName)
				fullPath = checkPath
				generationSuccessful = true
				break
			}
			if !generationSuccessful {
				loopErr = errors.Errorf("attempted 25 extended filenames, but all were taken. Aborting")
				break
			}
		}
		if loopErr != nil {
			break
		}

		// if f.OptimizeImages {
		// 	// TODO :: Apply image optimizations `pngquant` etc. when it works on Windows
		// }

		file, err := formField.uploadedFiles[k].Header().Open()
		if err != nil {
			loopErr = errors.Trace(err)
			break
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			loopErr = errors.Trace(err)
			break
		}
		if err := formField.storage.Set(fullPath, fileBytes); err != nil {
			loopErr = errors.Trace(err)
			break
		}
		savedFiles = append(savedFiles, fullPath)

		if !formField.HasErrors() && formField.gzipFilesOnSave {
			var b bytes.Buffer
			gw, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
			if err != nil {
				loopErr = errors.Trace(err)
				break
			}
			_, err = gw.Write(fileBytes)
			if err != nil {
				loopErr = errors.Trace(err)
				break
			}
			if err := gw.Close(); err != nil {
				loopErr = errors.Trace(err)
				break
			}
			if !formField.HasErrors() {
				if err := formField.storage.Set(fullPath+".gz", b.Bytes()); err != nil {
					loopErr = errors.Trace(err)
					break
				}
				savedFiles = append(savedFiles, fullPath+".gz")
			}
		}
		// if f.configService.BrotliFiles() {
		// 	// TODO :: Implement if Brotli would ever work on Windows
		// }

		formField.uploadedFiles[k].SetSavedPath(fullPath)
	}

	if loopErr != nil {
		// TODO :: 777 When doing update on multiple files, it should
		// not accidentally delete files that already existed
		for k := range savedFiles {
			// Don't abrubtly stop if one fails, as with multiple files it might
			// leave way more resedue data than needed
			if err := formField.storage.Delete(savedFiles[k]); err != nil {
				loopErr = errors.Wrap(loopErr, err)
			}
		}

		return errors.Trace(loopErr)
	}

	return nil
}

// SetStorage sets the storage where uploaded files get saved to.
func (formField *formField) SetStorage(storage storage.Storage) error {
	if FileFormField != formField._type {
		return errors.Errorf("cannot run SetStorage on a non-FileFormField")
	}
	formField.storage = storage
	return nil
}
