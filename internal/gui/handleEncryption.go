package gui

func (me *App) onEcncryptButton() {
	var msg Msg

	err := me.handleActions("enc", me.inputFilePath, me.passwordField.Textvariable())
	if err != nil {
		msg.mType = "error"
		msg.msg = err.Error()
		me.updateInfo(msg, true)
	}
}
