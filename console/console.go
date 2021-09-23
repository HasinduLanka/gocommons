package console

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
