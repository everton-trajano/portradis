package commands

import "strings"

func echo(args []string) string {
	return strings.Join(args, " ")
}
