package gui

import (
	"errors"
	"path/filepath"
	"sync"

	"github.com/darshanags/secure-files-go/pkg/appparser"
	decryptfile "github.com/darshanags/secure-files-go/pkg/decryptFile"
	encryptfile "github.com/darshanags/secure-files-go/pkg/encryptFile"
	"github.com/darshanags/secure-files-go/pkg/kdf"
	"github.com/darshanags/secure-files-go/pkg/utilities"
)

func (me *App) handleActions(directive string, iF []string, pass string) error {
	var (
		wg         sync.WaitGroup
		outputPath string
	)

	if len(pass) < 8 {
		return errors.New("password needs to be at least 8 characters long")
	}

	me.activeCh = make(chan utilities.AsyncResult, len(iF))

	switch directive {
	case "enc":
		for _, inputFile := range iF {
			salt, key := kdf.Kdf(pass, nil)
			outputPath = appparser.GetOutputPath(directive, inputFile)

			file := &encryptfile.LocalFileInfo{
				InputFilename:  filepath.Base(inputFile),
				InputPath:      inputFile,
				OutputFilename: filepath.Base(outputPath),
				OutputPath:     outputPath,
			}

			go file.EncryptFileAsync(key, salt, &wg, me.activeCh)
			wg.Add(1)
		}
	case "dec":
		for _, inputFile := range iF {
			outputPath = appparser.GetOutputPath(directive, inputFile)

			file := &decryptfile.LocalFileInfo{
				InputFilename:  filepath.Base(inputFile),
				InputPath:      inputFile,
				OutputFilename: filepath.Base(outputPath),
				OutputPath:     outputPath,
			}

			go file.DecryptFileAsync(pass, &wg, me.activeCh)
			wg.Add(1)
		}
	default:
		return errors.New("invalid directive")
	}

	return nil

}
