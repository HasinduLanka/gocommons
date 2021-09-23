package osargs

import (
	"os"
	"strings"
)

// Map command line arguments in form of
//
// 'myprog hello --foo bar -voo --help --output path/to/file'
//
// to a map[string]string in the form of
//
// { "myprog":"", "hello":"", "foo":"bar -voo", "help":"", "output":"path/to/file" }
//
func OSArgsToMap() map[string]string {

	return ArgsToMap(os.Args)
}

// Map array of arguments in form of
//
// [hello, --foo, bar, -voo, --help, --output, path/to/file]
//
// to a map[string]string in the form of
//
// { "myprog":"", "hello":"", "foo":"bar -voo", "help":"", "output":"path/to/file" }
//
func ArgsToMap(oa []string) map[string]string {

	args := make(map[string]string)
	for i := 0; i < len(oa); i++ {
		arg := oa[i]
		if strings.HasPrefix(arg, "--") {

			argValues := ""

			for j := i + 1; j < len(oa); j++ {
				if strings.HasPrefix(oa[j], "--") {
					break
				} else {
					argValues += oa[j] + " "
					i++
				}
			}

			args[strings.TrimPrefix(arg, "--")] = strings.TrimSpace(argValues)
		} else {
			args[arg] = ""
		}

	}
	return args
}
