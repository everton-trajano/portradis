package commands

import "github.com/everton-trajano/portradis/tools"

func get(args []string) string {
	return string(tools.Memory.GET(args[0]))
}
