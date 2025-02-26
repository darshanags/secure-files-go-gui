package gui

import (
	"fmt"

	"github.com/darshanags/secure-files-go/pkg/appparser"
	tk "modernc.org/tk9.0"
)

func (me *App) onFilePick() {
	var msg Msg
	me.inputFilePath = extractFilePaths(tk.GetOpenFile(tk.Title("Select File"), tk.Multiple(true)))

	if len(me.inputFilePath) == 1 {
		inputFileExt := appparser.GetFileExtension(me.inputFilePath[0])

		if inputFileExt == ".enc" {
			sig, err := appparser.GetFileSignature(nil, me.inputFilePath[0])
			if err != nil {
				msg.mType = "error"
				msg.msg = err.Error()
				me.updateInfo(msg, true)
			}

			_ = sig

			sigValid, err := appparser.IsValidFileSignature(sig)
			if err != nil {
				msg.mType = "error"
				msg.msg = err.Error()
				me.updateInfo(msg, true)
			}

			if sigValid {
				me.decryptButton.Configure(tk.State("enabled"))
			}
		} else {
			me.decryptButton.Configure(tk.State("disabled"))
		}

		msg.msg = "File selected: " + me.inputFilePath[0]
	} else {
		me.decryptButton.Configure(tk.State("enabled"))
		msg.msg = fmt.Sprintf("%d files selected.", len(me.inputFilePath))
	}

	me.updateInfo(msg, true)
}
