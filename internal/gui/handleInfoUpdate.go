package gui

import (
	"errors"
	"fmt"

	tk "modernc.org/tk9.0"
)

type Msg struct {
	mType string
	msg   string
}

func (me *App) updateInfo(m Msg, append bool) error {
	me.infoArea.Configure(tk.State("normal"))
	if !append {
		me.infoArea.Clear()
	}
	if m.mType == "" {
		m.mType = "info"
	}
	switch m.mType {
	case "error":
		me.infoArea.TagConfigure("error", tk.Foreground(tk.Red))
		me.infoArea.InsertML(fmt.Sprintf(`%s %s%s%s`, GetTimestamp(), `<error>`, m.msg, `</error><br/>`))
	case "success":
		me.infoArea.TagConfigure("success", tk.Foreground(tk.Green))
		me.infoArea.InsertML(fmt.Sprintf(`%s %s%s%s`, GetTimestamp(), `<success>`, m.msg, `</success><br/>`))
	case "info":
		me.infoArea.InsertML(fmt.Sprintf(`%s %s`, GetTimestamp(), m.msg+`<br/>`))
	default:
		return errors.New("the msg parameter is invalid")
	}
	me.infoArea.Configure(tk.State("disabled"))

	return nil
}
