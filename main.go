package main

import (
	"embed"
	"net/http"

	"github.com/syumai/workers"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(assetsFS)))
	workers.Serve(nil)
}
