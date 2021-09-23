package console

import "encoding/json"

func Print(msg string) {
	println(msg)
}

func CheckError(err error) bool {
	return err != nil
}

func PrintError(err error) bool {
	if CheckError(err) {
		Print(err.Error())
		return true
	}
	return false
}

func PrintJson(obj interface{}) {
	jb, je := json.MarshalIndent(obj, "", "  ")
	if CheckError(je) {
		Print("Printing this object as JSON failed")
		return
	}

	Print("\n" + string(jb) + "\n")
}
