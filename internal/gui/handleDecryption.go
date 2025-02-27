package gui

import (
	"fmt"

	"github.com/darshanags/secure-files-go/pkg/appparser"
)

func (me *App) onDecryptButton() {
	var (
		msg             Msg
		compatibleFiles []string
		compatFilesLen  int
		inputFilesLen   = len(me.inputFiles.inputFilePaths)
	)

	if inputFilesLen == 0 {
		msg.mType = "error"
		msg.msg = "no files selected"
		me.updateInfo(msg, true)
		return
	}

	for _, fp := range me.inputFiles.inputFilePaths {
		sig, err := appparser.GetFileSignature(nil, fp)
		if err != nil {
			continue
		}
		sigValid, err := appparser.IsValidFileSignature(sig)
		if err != nil {
			continue
		}

		if sigValid {
			compatibleFiles = append(compatibleFiles, fp)
		}
	}

	compatFilesLen = len(compatibleFiles)

	if compatFilesLen > 0 {
		me.inputFiles.inputFilePaths = compatibleFiles

		if compatFilesLen < inputFilesLen {
			msg.msg = fmt.Sprintf("Out of the selected %d files, only %d can be decrypted. Proceeding with decryption...", inputFilesLen, compatFilesLen)
			me.updateInfo(msg, true)
		}
	} else {
		msg.mType = "error"
		msg.msg = "the selected files are not valid for decryption"
		me.updateInfo(msg, true)
		return
	}

	err := me.handleActions("dec", me.inputFiles.inputFilePaths, me.passwordField.Textvariable())
	if err != nil {
		msg.mType = "error"
		msg.msg = err.Error()
		me.updateInfo(msg, true)
	}
}
