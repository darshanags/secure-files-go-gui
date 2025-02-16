package gui

import tk "modernc.org/tk9.0"

func (me *App) onQuit() {
	answer := tk.MessageBox(tk.Icon("question"), tk.Msg("Are you sure you want to exit?"), tk.Title("Exiting Program"), tk.Type("yesno"))
	if answer == "yes" {
		tk.Destroy(tk.App)
	}
}
