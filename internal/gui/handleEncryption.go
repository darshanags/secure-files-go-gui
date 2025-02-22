package gui

import tk "modernc.org/tk9.0"

func (me *App) onEcncryptButton() {
	var msg Msg
	message, err := handleActions("enc", me.inputFilePath, me.passwordField.Textvariable())
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
