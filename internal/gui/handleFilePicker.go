package gui

import (
	"strings"

	"github.com/darshanags/secure-files-go/pkg/appparser"
	tk "modernc.org/tk9.0"
)

func (me *App) onFilePick() {
	var msg Msg
	me.inputFilePath = strings.Join(tk.GetOpenFile(tk.Title("Select File"), tk.Multiple(false)), "")
	inputFileExt := appparser.GetFileExtension(me.inputFilePath)

	if inputFileExt == ".enc" {
		sig, err := appparser.GetFileSignature(nil, me.inputFilePath)
		if err != nil {
			msg.mType = "error"
			msg.msg = err.Error()
			updateInfo(me.infoArea, msg, true)
		}

		_ = sig

		sigValid, err := appparser.IsValidFileSignature(sig)
		if err != nil {
			msg.mType = "error"
			msg.msg = err.Error()
			updateInfo(me.infoArea, msg, true)
		}

		if sigValid {
			me.decryptButton.Configure(tk.State("enabled"))
		}
	} else {
		me.decryptButton.Configure(tk.State("disabled"))
	}

	msg.mType = "msg"
	msg.msg = "File selected: " + me.inputFilePath
	updateInfo(me.infoArea, msg, true)
}
