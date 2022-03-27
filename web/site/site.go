package site

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/use/html/body"
	"qlova.tech/use/html/head"
	"qlova.tech/use/html/script"
	"qlova.tech/use/html/template"
	"qlova.tech/use/html/title"
	"qlova.tech/web/data"
	"qlova.tech/web/tree"

	_ "embed"
)

func listenAndServe(title string, handler http.Handler) error {
	var port = ":0"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	var iport uint16
	var browser bool

	//Determine a stable port number from the app's name.

	if port == ":0" {
		var hash = fnv.New64()
		hash.Write([]byte(title))

		iport = uint16(hash.Sum64()%(math.MaxUint16-1000)) + 1000

		port = fmt.Sprint(":", iport)
		browser = true
	}

retry:
	listener, err := net.Listen("tcp", port)
	if err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			iport++
			port = fmt.Sprint(":", iport)
			goto retry
		}
		return err
	}

	if browser {
		go launch("http://" + listener.Addr().String())
	}

	fmt.Printf("\nlaunching %v version %v on http://localhost%v\n", title, "0.1", port)

	return http.Serve(listener, handler)
}

//go:embed site.js
var js string

// Open opens the given root tree renderer
// as a website. If PORT is provided, the
// website will be server over HTTP on the
// port, otherwise it will open the website
// in a browser.
func Open(root any) error {
	rtype := reflect.TypeOf(root)

	if rtype.Kind() == reflect.Ptr {
		data.Register(root)
	}
	for rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}

	name := rtype.Name()
	switch root.(type) {
	case string:
		name = root.(string)
	}

	var seed = tree.NewSeed()

	var content = tree.Node{
		html.New(root),
	}

	if renderer, ok := root.(tree.Renderer); ok {
		content = tree.Render(renderer, seed)
	}

	// Extract templates.
	var templates []any
	for rtype, iface := range seed.TreeRenderers {
		if renderer, ok := iface.(tree.Renderer); ok {
			templates = append(templates, append(renderer.RenderTree(seed), template.Tag, html.Attr("data-type", rtype.Name())))
		}
	}

	var init = make(map[string]any)
	init[strings.TrimPrefix(data.PathOf(root), "/")] = root
	for _, p := range seed.TreeRenderers {
		init[strings.TrimPrefix(data.PathOf(p), "/")] = p
	}

	b, err := json.Marshal(init)
	if err != nil {
		panic(err)
	}

	var handlers = data.HandlersOf(seed.TreeRenderers)

	var document = html.Render(html.New(
		head.New(
			title.Set(name),
			script.New(
				script.Source("/index.js"),
			),
		),
		append(content, append(templates, body.Tag, html.Attr("onload", "data.load();"))...),
	))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/index.js" {
			w.Header().Set("Content-Type", "application/javascript")
			fmt.Fprint(w, js)
			return
		}

		if r.URL.Path == "/data" || r.URL.Path == "/data/" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(b))
			return
		}

		if handler, ok := handlers[r.URL.Path]; ok {
			handler.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, document)
	})

	return listenAndServe(name, handler)
}
