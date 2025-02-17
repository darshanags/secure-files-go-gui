package gui

import tk "modernc.org/tk9.0"

func (me *App) onDecryptButton() {
	var msg Msg
	var btnState string = me.decryptButton.State()

	if btnState == "disabled" {
		return
	}

	message, err := handleActions("dec", me.inputFilePath, me.passwordField.Textvariable())

	if err != nil {
		msg.mType = "error"
		msg.msg = err.Error()
		updateInfo(me.infoArea, msg, true)
	} else {
		me.passwordField.Configure(tk.Textvariable(""))
		msg.mType = "msg"
		msg.msg = message
		updateInfo(me.infoArea, msg, true)
	}
}
