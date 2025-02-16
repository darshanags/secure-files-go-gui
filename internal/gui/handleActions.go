package gui

import (
	"errors"

	"github.com/darshanags/secure-files-go/pkg/appparser"
	decryptfile "github.com/darshanags/secure-files-go/pkg/decryptFile"
	encryptfile "github.com/darshanags/secure-files-go/pkg/encryptFile"
	"github.com/darshanags/secure-files-go/pkg/kdf"
)

func handleActions(directive string, iF string, pass string) (string, error) {
	message := ""

	if len(iF) < 1 {
		return message, errors.New("input file is invalid")
	}

	if len(pass) < 8 {
		return message, errors.New("password needs to be at least 8 characters long")
	}

	switch directive {
	case "enc":
		salt, key := kdf.Kdf(pass, nil)
		outputPath := appparser.GetOutputPath(directive, iF)

		message, err := encryptfile.EncryptFile(iF, outputPath, key, salt)

		return message, err

	case "dec":
		outputPath := appparser.GetOutputPath(directive, iF)
		message, err := decryptfile.DecryptFile(iF, outputPath, pass)
		return message, err
	default:
		return message, errors.New("invalid directive")
	}
}
