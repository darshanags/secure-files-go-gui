package gui

import (
	"fmt"

	"github.com/darshanags/secure-files-go/pkg/appparser"
)

func (me *App) onEcncryptButton() {
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
		if _, err := appparser.IsValidFileSignature(sig); err == nil {
			continue
		}

		compatibleFiles = append(compatibleFiles, fp)
	}

	compatFilesLen = len(compatibleFiles)

	if compatFilesLen > 0 {
		me.inputFiles.inputFilePaths = compatibleFiles

		if compatFilesLen < inputFilesLen {
			msg.msg = fmt.Sprintf("Out of the selected %d files, only %d can be encrypted. Proceeding with encryption...", inputFilesLen, compatFilesLen)
			me.updateInfo(msg, true)
		}
	} else {
		msg.mType = "error"
		msg.msg = "the selected files are already encrypted"
		me.updateInfo(msg, true)
		return
	}

	err := me.handleActions("enc", me.inputFiles.inputFilePaths, me.passwordField.Textvariable())
	if err != nil {
		msg.mType = "error"
		msg.msg = err.Error()
		me.updateInfo(msg, true)
	}
}
