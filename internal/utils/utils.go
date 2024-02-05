package utils

import (
	"strconv"
	"strings"
)

func calc(in, out int64) int {
	return int(in + out)
}

type Line struct {
	Title       string
	Ip          string
	InputBytes  int64
	OutputBytes int64
}

func ReadLines(file []byte) []Line {
	lines := make([]Line, 0)

	for _, line := range strings.Split(string(file), "\n") {
		data := strings.Split(line, " ")

		in, _ := strconv.ParseInt(data[3], 10, 64)
		out, _ := strconv.ParseInt(data[4], 10, 64)

		lines = append(
			lines, Line{
				Title:       data[1],
				Ip:          data[2],
				InputBytes:  in,
				OutputBytes: out,
			},
		)
	}

	return lines
}
