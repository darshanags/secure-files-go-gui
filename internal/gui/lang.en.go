package gui

type Lang struct {
	Strings map[string]string
}

func (lang *Lang) init() {
	lang.Strings = make(map[string]string)
	lang.Strings["file_select_label"] = "Select File:"
	lang.Strings["file_select_btn"] = "Select Files"
	lang.Strings["file_select_btn_selected"] = "%d Files Selected"
	lang.Strings["password_label"] = "Password:"
	lang.Strings["info_area_label"] = "Log:"
	lang.Strings["encrypt_btn"] = "Encrypt File"
	lang.Strings["encrypt_btn_p"] = "Encrypt Files"
	lang.Strings["decrypt_btn"] = "Decrypt File"
	lang.Strings["decrypt_btn_p"] = "Decrypt Files"
	lang.Strings["exit_btn"] = "Exit"
}

func (lang *Lang) get(key string) string {
	return lang.Strings[key]
}
