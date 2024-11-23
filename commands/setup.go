package commands

var Commands = map[string]func([]string) string{
	"PING": ping,
	"ECHO": echo,
}
