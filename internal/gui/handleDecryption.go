package gui

func (me *App) onDecryptButton() {
	var (
		msg      Msg
		btnState string = me.decryptButton.State()
	)

	if btnState == "disabled" {
		return
	}

	err := me.handleActions("dec", me.inputFilePath, me.passwordField.Textvariable())
	if err != nil {
		msg.mType = "error"
		msg.msg = err.Error()
		me.updateInfo(msg, true)
	}
}
