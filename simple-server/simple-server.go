package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(rw http.ResponseWriter, req *http.Request) {
	switch strings.TrimPrefix(req.URL.Path, "/players/") {
	case "me":
		fmt.Fprintf(rw, "20")
	case "you":
		fmt.Fprintf(rw, "10")
	}
}
