package validators

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/espal-digital-development/espal-core/storage"
	"github.com/espal-digital-development/espal-core/text"
	"github.com/juju/errors"
)

const fileExtraHashLength = 4

// NewFileField returns a new instance of FormField with the type File.
func (f *Form) NewFileField(name string, storage storage.Modifyable) FormField {
	return f.defaultChecks(&formField{
		name:              name,
		_type:             FileFormField,
		storage:           storage,
		brotliFilesOnSave: f.configService.AssetsBrotliFiles(),
		gzipFilesOnSave:   f.configService.AssetsGZipFiles(),
	})
}

// SetOptimizeImages marks that uploaded images should be optimized as best
// as possible for file size. This may vary per platform's available optimizers
// and aren't quality compressors, but only encoder efficiency and meta optimizations.
func (f *formField) SetOptimizeImages() {
	f.optimizeImages = true
}

// OptimizeImages returns if the uploaded file(s) should be optimized.
func (f *formField) OptimizeImages() bool {
	return f.optimizeImages
}

// SetGzipFilesOnSave marks the files to also be brotli compressed on save.
func (f *formField) SetBrotliFilesOnSave() {
	f.brotliFilesOnSave = true
}

// BrotliFilesOnSave returns if the files should be brotli compressed on save.
func (f *formField) BrotliFilesOnSave() bool {
	return f.brotliFilesOnSave
}

// SetGzipFilesOnSave marks the files to also be gzip compressed on save.
func (f *formField) SetGzipFilesOnSave() {
	f.gzipFilesOnSave = true
}

// GzipFilesOnSave returns if the files should be gzip compressed on save.
func (f *formField) GzipFilesOnSave() bool {
	return f.gzipFilesOnSave
}

// SetAllowedMIMETypes sets the allowed MIME types that can be uploaded through the form.
func (f *formField) SetAllowedMIMETypes(allowedMIMETypes []string) {
	f.allowedMIMETypes = allowedMIMETypes
}

// RemoveUploadedFiles removes any files that were uploaded on
// the targeted field. This is useful when a form fails after
// the upload and it should be cleaned up.
// It won't delete images already linked on the loaded data.
func (f *formField) RemoveUploadedFiles() error {
	for k := range f.uploadedFiles {
		var fullPath string
		if f.fileSaveFolder != "" {
			fullPath = f.fileSaveFolder + string(os.PathSeparator)
		}
		fullPath += f.uploadedFiles[k].SanitizedName()
		if !f.storage.Exists(fullPath) {
			continue
		}
		if err := f.storage.Delete(fullPath); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// AddUploadedFile adds a new UploadedFile to the internal files list.
func (f *formField) AddUploadedFile(uploadedFile UploadedFile) {
	if f.uploadedFiles == nil {
		f.uploadedFiles = make([]UploadedFile, 0)
	}
	f.uploadedFiles = append(f.uploadedFiles, uploadedFile)
}

// UploadedFiles returns all the field's uploaded files.
func (f *formField) UploadedFiles() []UploadedFile {
	return f.uploadedFiles
}

// SetFileSaveFolder sets the specific folder (different than default)
// where to save the file(s) too.
func (f *formField) SetFileSaveFolder(fileSaveFolder string) {
	f.fileSaveFolder = fileSaveFolder
}

// FileSaveFolder gets the specific folder (different than default)
// where to save the file(s) too.
func (f *formField) FileSaveFolder() string {
	return f.fileSaveFolder
}

// SaveFiles saves the file(s) that were uploaded to the private
// or public folder based on the field.PublicFile setting.
// nolint:funlen
func (f *formField) SaveFiles() error {
	if FileFormField != f._type {
		return errors.Errorf("cannot run SaveFiles on a non-FileFormField")
	}
	if f.storage == nil {
		return errors.Errorf("cannot run SaveFiles without a storage being set")
	}

	// TODO :: 7 When updating, but replacing a/the file(s) should delete the old ones too.
	// This is harder with multiple files tho

	var loopErr error
	savedFiles := make([]string, 0)

	for k := range f.uploadedFiles {
		var fullPath string
		if f.fileSaveFolder != "" {
			fullPath = f.fileSaveFolder + string(os.PathSeparator)
		}
		fullPath += f.uploadedFiles[k].SanitizedName()

		var generationSuccessful bool
		if f.storage.Exists(fullPath) {
			var extension string
			var baseName string
			dotSlices := strings.Split(f.uploadedFiles[k].SanitizedName(), ".")
			length := len(dotSlices)
			if length > 1 && dotSlices[length-1] != "" {
				extension = dotSlices[length-1]
				for i := 0; i < length-1; i++ {
					baseName += dotSlices[i]
				}
			} else {
				baseName = f.uploadedFiles[k].SanitizedName()
			}

			for i := 0; i < 25; i++ {
				extra := text.RandomString(fileExtraHashLength)
				newName := baseName + "_" + extra + "." + extension
				var checkPath string
				if f.fileSaveFolder != "" {
					checkPath = f.fileSaveFolder + string(os.PathSeparator)
				}
				checkPath += newName
				if f.storage.Exists(checkPath) {
					continue
				}
				f.uploadedFiles[k].SetSanitizedName(newName)
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

		file, err := f.uploadedFiles[k].Header().Open()
		if err != nil {
			loopErr = errors.Trace(err)
			break
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			loopErr = errors.Trace(err)
			break
		}
		if err := f.storage.Set(fullPath, fileBytes); err != nil {
			loopErr = errors.Trace(err)
			break
		}
		savedFiles = append(savedFiles, fullPath)

		// TODO :: 7 Move the runtime.GOOS check to a dedicated file portion
		if !f.HasErrors() && f.brotliFilesOnSave && runtime.GOOS != "windows" {
			var b bytes.Buffer
			gw := brotli.NewWriterLevel(&b, gzip.BestCompression)
			_, err = gw.Write(fileBytes)
			if err != nil {
				loopErr = errors.Trace(err)
				break
			}
			if err := gw.Close(); err != nil {
				loopErr = errors.Trace(err)
				break
			}
			if !f.HasErrors() {
				if err := f.storage.Set(fullPath+".br", b.Bytes()); err != nil {
					loopErr = errors.Trace(err)
					break
				}
				savedFiles = append(savedFiles, fullPath+".br")
			}
		}
		if !f.HasErrors() && f.gzipFilesOnSave {
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
			if !f.HasErrors() {
				if err := f.storage.Set(fullPath+".gz", b.Bytes()); err != nil {
					loopErr = errors.Trace(err)
					break
				}
				savedFiles = append(savedFiles, fullPath+".gz")
			}
		}

		f.uploadedFiles[k].SetSavedPath(fullPath)
	}

	if loopErr != nil {
		// TODO :: 777 When doing update on multiple files, it should
		// not accidentally delete files that already existed
		for k := range savedFiles {
			// Don't abrubtly stop if one fails, as with multiple files it might
			// leave way more resedue data than needed
			if err := f.storage.Delete(savedFiles[k]); err != nil {
				loopErr = errors.Wrap(loopErr, err)
			}
		}

		return errors.Trace(loopErr)
	}

	return nil
}

// SetStorage sets the storage where uploaded files get saved to.
func (f *formField) SetStorage(storage storage.Modifyable) error {
	if FileFormField != f._type {
		return errors.Errorf("cannot run SetStorage on a non-FileFormField")
	}
	f.storage = storage
	return nil
}
