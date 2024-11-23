package commands

func ping(args []string) (value string) {
	if len(args) == 0 {
		return "PONG"
	}

	return args[0]
}
