package gui

import (
	"fmt"

	tk "modernc.org/tk9.0"
)

func (me *App) onFilePick() {
	var (
		msg        Msg
		noSelFiles int
	)

	me.inputFiles.inputFilePaths = tk.GetOpenFile(tk.Title(me.lang.get("file_select_title")), tk.Multiple(true))
	noSelFiles = len(me.inputFiles.inputFilePaths)
	me.filePickerButton.Configure(tk.Txt(fmt.Sprintf(me.lang.get("file_select_btn_selected"), noSelFiles)))
	me.inputFiles.selected = true
	tk.Focus(me.passwordField)

	if noSelFiles > 1 {
		me.decryptButton.Configure(tk.Txt(me.lang.get("decrypt_btn_p")))
		me.encryptButton.Configure(tk.Txt(me.lang.get("encrypt_btn_p")))
	}
	msg.msg = fmt.Sprintf("%d files selected.", noSelFiles)

	me.updateInfo(msg, true)
}
