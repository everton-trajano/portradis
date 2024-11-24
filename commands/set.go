package commands

import "github.com/everton-trajano/portradis/tools"

func set(args []string) string {
	tools.Memory.SET(args[0], args[1])
	return "done"
}
