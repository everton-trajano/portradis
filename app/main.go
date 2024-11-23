package main

import (
	"errors"
	"io"
	"log/slog"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/everton-trajano/portradis/commands"
	"github.com/everton-trajano/portradis/tools"
)

const (
	receiveBuf = 1024
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379") //nolint:gosec // temporary
	if err != nil {
		slog.Error("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer func() {
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error("Error accepting connection: ", "error", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	buffer := make([]byte, receiveBuf)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				slog.Error("Reading", "error", err)
			}
			os.Exit(1)
		}

		cmd := strings.Split(strings.Replace(string(buffer[:n]), "*", "", 1), "\r\n")
		argsCount, _ := strconv.Atoi(cmd[0])
		command := cmd[2]

		args := tools.CleanArgs(cmd, argsCount)
		slog.Info("[RECEIVED]", "argsCount", argsCount, "command", strings.ToUpper(command), "args", args)
		if command == "COMMAND" {
			conn.Write([]byte("+OK\r\n"))
			continue
		}

		response := commands.Commands[strings.ToUpper(command)](args)
		slog.Info("[RESPONDING]", "args", response)
		conn.Write([]byte("+" + response + "\r\n"))
	}
}
