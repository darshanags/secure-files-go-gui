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

func updateInfo(iA *tk.TextWidget, m Msg, append bool) error {
	iA.Configure(tk.State("normal"))
	if !append {
		iA.Clear()
	}
	switch m.mType {
	case "error":
		iA.TagConfigure("error", tk.Foreground(tk.Red))
		iA.InsertML(fmt.Sprintf(`%s %s%s%s`, GetTimestamp(), `<error>`, m.msg, `</error><br/>`))
	case "msg":
		iA.InsertML(fmt.Sprintf(`%s %s`, GetTimestamp(), m.msg+`<br/>`))
	default:
		return errors.New("the msg parameter is invalid")
	}
	iA.Configure(tk.State("disabled"))

	return nil
}
