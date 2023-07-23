package app

import (
	"fmt"
	"net/http"

	"qlova.tech/app/data"
	"qlova.tech/app/internal/html"
	"qlova.tech/app/page"
)

func ListenAndServe(port string, home page.Renderer) error {

	var content = html.Render(home.RenderPage(), data.New(home))

	fmt.Println(content)

	return http.ListenAndServe(port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(content))
	}))
}
