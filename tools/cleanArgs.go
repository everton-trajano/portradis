package tools

func CleanArgs(cmd []string, argsCount int) (args []string) {
	for index, piece := range cmd[3 : 3+((argsCount-1)*2)] {
		if index%2 == 0 {
			continue
		}
		args = append(args, piece)
	}
	return
}
