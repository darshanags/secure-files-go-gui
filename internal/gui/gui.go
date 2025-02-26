package gui

import (
	"fmt"
	"log"
	"time"

	"github.com/darshanags/secure-files-go/pkg/utilities"
	tk "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

type App struct {
	passwordField    *tk.TEntryWidget
	infoArea         *tk.TextWidget
	filePickerButton *tk.TButtonWidget
	encryptButton    *tk.TButtonWidget
	decryptButton    *tk.TButtonWidget
	exitButton       *tk.TButtonWidget
	fileSelectLabel  *tk.LabelWidget
	passwordLabel    *tk.LabelWidget
	infoAreaLabel    *tk.LabelWidget
	inputFilePath    []string
	activeCh         chan utilities.AsyncResult
}

func NewApp(appName string) *App {
	app := &App{}
	tk.ActivateTheme("azure light")
	tk.App.WmTitle(appName)
	tk.WmProtocol(tk.App, tk.WM_DELETE_WINDOW, app.onQuit)
	app.makeWidgets()
	app.makeLayout()
	app.makeBindings()
	app.startTicker()
	return app
}

func (me *App) makeWidgets() {
	me.makeUserInputs()
	me.makeInfoArea()
	me.makeActionButtons()
}

func (me *App) makeUserInputs() {
	me.fileSelectLabel = tk.Label(tk.Anchor("w"), tk.Txt("Select File: "))
	me.filePickerButton = tk.TButton(tk.Txt("Select File..."))
	me.passwordLabel = tk.Label(tk.Anchor("w"), tk.Txt("Password: "))
	me.passwordField = tk.TEntry(tk.Justify("left"), tk.Show("*"), tk.Textvariable(""))
}

func (me *App) makeInfoArea() {
	me.infoAreaLabel = tk.Label(tk.Anchor("w"), tk.Txt("Log:"))
	me.infoArea = tk.Text(tk.Font(tk.CourierFont(), 10),
		tk.State("disabled"), tk.Setgrid(true), tk.Undo(false),
		tk.Wrap("word"), tk.Relief("sunken"), tk.Borderwidth(1))
}

func (me *App) makeActionButtons() {
	me.encryptButton = tk.TButton(tk.Txt("Encrypt File"))
	me.decryptButton = tk.TButton(tk.Txt("Decrypt File"), tk.State("disabled"))
	me.exitButton = tk.TButton(tk.Txt("Exit"))
}

func (me *App) makeLayout() {
	me.layoutUserInputs()
	me.layoutInfoArea()
	me.laoutActionArea()
}

func (me *App) layoutUserInputs() {
	opts := tk.Opts{tk.Padx("1m"), tk.Pady("2m"), tk.Ipadx("1m"), tk.Ipady("1m")}
	tk.Grid(me.fileSelectLabel, tk.Row(0), tk.Column(0), tk.Sticky("nws"), opts)
	tk.Grid(me.filePickerButton, tk.Row(0), tk.Column(1), tk.Sticky("news"), tk.Columnspan(2), opts)
	tk.Grid(me.passwordLabel, tk.Row(1), tk.Column(0), tk.Sticky("nws"), opts)
	tk.Grid(me.passwordField, tk.Row(1), tk.Column(1), tk.Sticky("news"), tk.Columnspan(2), opts)
}

func (me *App) layoutInfoArea() {
	tk.Grid(me.infoAreaLabel, tk.Columnspan(3), tk.Sticky("news"))
	tk.Grid(me.infoArea, tk.Columnspan(3), tk.Sticky("news"))
}

func (me *App) laoutActionArea() {
	opts := tk.Opts{tk.Padx("1m"), tk.Pady("2m"), tk.Ipadx("1m"), tk.Ipady("1m")}
	tk.Grid(me.encryptButton, tk.Row(4), tk.Column(0), tk.Sticky("news"), opts)
	tk.Grid(me.decryptButton, tk.Row(4), tk.Column(1), tk.Sticky("news"), opts)
	tk.Grid(me.exitButton, tk.Row(4), tk.Column(2), tk.Sticky(tk.E), opts)
}

func (me *App) makeBindings() {
	tk.Bind(me.filePickerButton, "<ButtonRelease>", tk.Command(me.onFilePick))
	tk.Bind(me.encryptButton, "<ButtonRelease>", tk.Command(me.onEcncryptButton))
	tk.Bind(me.decryptButton, "<ButtonRelease>", tk.Command(me.onDecryptButton))
	tk.Bind(tk.App, "<Escape>", tk.Command(me.onQuit))
	tk.Bind(me.exitButton, "<ButtonRelease>", tk.Command(me.onQuit))
}

func (me *App) startTicker() *tk.Ticker {

	ticker, err := tk.NewTicker(100*time.Millisecond, me.tick)

	if err != nil {
		log.Fatalln(err)
	}

	return ticker
}

func (me *App) tick() {
	var msg Msg

	select {
	case result, ok := <-me.activeCh:
		if !ok {
			fmt.Println("Channel is closed")
			return
		}
		if result.Error != nil {
			msg.mType = "error"
			msg.msg = result.Error.Error()
			me.updateInfo(msg, true)
		} else {
			me.passwordField.Configure(tk.Textvariable(""))
			msg.mType = "success"
			msg.msg = result.Message
			me.updateInfo(msg, true)
		}
	default:
		// No value available in the channel
	}

}

func (me *App) Run() {
	tk.App.SetResizable(false, false)
	tk.App.Center()
	tk.WmDeiconify(tk.App)
	tk.App.Wait()
}
