package api

import (
	"strings"
	"strconv"
)

func (a *api) Start(port int) error {
	var addressBuilder strings.Builder
	addressBuilder.WriteString(":")
	addressBuilder.WriteString(strconv.Itoa(port))

	return a.App.Listen(addressBuilder.String())
}
