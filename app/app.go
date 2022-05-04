package app

import (
	"bytes"
	"encoding/json"
	"html"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"qlova.tech/app/user"

	_ "embed"
)

const HelloWorld helloworld = true

type helloworld bool

func (helloworld) RenderInterface(user.Data) user.Interface {
	return user.Interface{
		user.View("Hello World"),
	}
}

var (
	//go:embed app.js
	js string

	//go:embed app.css
	css string
)

// ListenAndServe starts a web server on the given port and
// starts listening for users, serving them the provided
// user interface. The user interface is crawled for pages
// and the correct page will be served depending on the
// user's request. As a special case, if no port is provided
// a deterministic port will be chosen and ListenAndServe
// will attempt to launch the user interface locally.
func ListenAndServe(port string, ui user.InterfaceRenderer) error {
	return http.ListenAndServe(port, Handler(ui))
}

func nameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr {
		return nameOf(rtype.Elem())
	}
	return rtype.Name()
}

// Handler returns a http.Handler that can be used to serve
// the provided user interface over HTTP, in response to
// requests made by a particular user.
func Handler(ui user.InterfaceRenderer) http.Handler {

	var data = user.DataOf(ui)

	var render = renderer{
		data: data,
	}

	root := ui.RenderInterface()
	html := render.document(root, nameOf(reflect.TypeOf(ui)))

	var resources = make(map[string]http.Handler)
	for _, rtype := range data.Types() {
		resources[nameOf(rtype)] = resource{rtype}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/app.js":
			w.Header().Set("Content-Type", "text/javascript")
			w.Write([]byte(js))
			return
		case "/app.css":
			w.Header().Set("Content-Type", "text/css")
			w.Write([]byte(css))
			return
		}

		if strings.HasPrefix(r.URL.Path, "/data/") {
			var name = strings.TrimPrefix(r.URL.Path, "/data/")
			if handler, ok := resources[name]; ok {
				handler.ServeHTTP(w, r)
				return
			}
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	})
}

type renderer struct {
	bytes.Buffer

	data user.Data

	indent       int
	headingLevel int
}

func (r *renderer) writeLine(line string) {
	for i := 0; i < r.indent; i++ {
		r.WriteByte('\t')
	}
	r.WriteString(line)
	r.WriteByte('\n')
}

type htmlAttribute struct {
	name  string
	value string
}

func (r *renderer) render(ui user.Interface) {
	for i := 0; i < r.indent; i++ {
		r.WriteByte('\t')
	}

	r.WriteString("<")

	// determine the tag and build attributes.
	var tag string
	var void = false
	var attributes []htmlAttribute

	for _, hint := range ui.Hints {
		if len(hint) > 0 && hint[0] == '<' {
			tag = string(hint[1 : len(hint)-1])
			if tag[len(tag)-1] == '/' {
				tag = tag[:len(tag)-1]
				void = true
			}
		} else {
			for _, attr := range strings.Split(string(hint), " ") {
				key, value, _ := strings.Cut(string(attr), "=")

				var found = false
				for _, attr := range attributes {
					if attr.name == key {
						attr.value += value
						found = true
						break
					}
				}
				if !found {
					attributes = append(attributes, htmlAttribute{
						name:  key,
						value: value,
					})
				}
			}
		}
	}

	// special handling of headers.
	if tag == "title" {
		level := r.headingLevel + 1
		if level > 6 {
			level = 6
		}
		tag = "h" + strconv.Itoa(level)
	}

	r.WriteString(tag)

	for _, attr := range attributes {
		r.WriteByte(' ')
		r.WriteString(attr.name)
		if attr.value != "" {
			r.WriteString(`="`)
			r.WriteString(attr.value)
			r.WriteByte('"')
		}
	}

	r.WriteString(">")

	if void {
		return
	}

	r.WriteString(strings.Replace(html.EscapeString(ui.Text), "\n", "<br>", -1))

	if len(ui.Nodes) > 0 {
		r.WriteString("\n")
		r.indent++
		if tag == "title" {
			r.headingLevel++
		}
		for _, node := range ui.Nodes {
			r.render(node)
		}
		if tag == "title" {
			r.headingLevel--
		}
		r.indent--

		for i := 0; i < r.indent; i++ {
			r.WriteByte('\t')
		}
	}
	r.WriteString("</")
	r.WriteString(tag)
	r.WriteString(">\n")
}

func (r *renderer) document(ui user.Interface, name string) []byte {

	r.writeLine("<!DOCTYPE html>")
	r.writeLine("<html>")
	r.indent++
	r.writeLine("<head>")
	r.indent++
	r.writeLine("<title>" + html.EscapeString(name) + "</title>")
	r.writeLine("<script src=app.js></script>")
	r.writeLine("<link href=app.css rel=stylesheet>")
	r.indent--
	r.writeLine("</head>")
	r.writeLine("<body>")
	r.indent++
	r.render(ui)
	r.indent--
	r.writeLine("</body>")
	r.indent--
	r.writeLine("</html>")

	return r.Bytes()
}

type resource struct {
	rtype reflect.Type
}

func (res resource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handle := func(err error) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	MethodNotAllowed := func() {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	BadRequest := func(err error) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	iface := reflect.New(res.rtype).Interface()
	switch r.Method {
	case http.MethodGet:
		handler, ok := iface.(user.DataWithGet)
		if !ok {
			MethodNotAllowed()
			return
		}

		//TODO decode query.

		if err := handler.Get(ctx); err != nil {
			handle(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(iface); err != nil {
			handle(err)
			return
		}
	case "SEARCH":
		handler, ok := iface.(user.DataWithSearch)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := json.NewDecoder(r.Body).Decode(iface); err != nil {
			BadRequest(err)
			return
		}
		if err := handler.Search(ctx); err != nil {
			handle(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(iface); err != nil {
			handle(err)
			return
		}
	case http.MethodPost:
		handler, ok := iface.(user.DataWithPost)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := json.NewDecoder(r.Body).Decode(iface); err != nil {
			BadRequest(err)
			return
		}
		if err := handler.Post(ctx); err != nil {
			handle(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(iface); err != nil {
			handle(err)
			return
		}
	case http.MethodPut:
		handler, ok := iface.(user.DataWithPut)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := json.NewDecoder(r.Body).Decode(iface); err != nil {
			BadRequest(err)
			return
		}
		if err := handler.Put(ctx); err != nil {
			handle(err)
			return
		}
	case http.MethodDelete:
		handler, ok := iface.(user.DataWithDelete)
		if !ok {
			MethodNotAllowed()
			return
		}
		if err := handler.Delete(ctx); err != nil {
			handle(err)
			return
		}
	case http.MethodPatch:
		handler, ok := iface.(user.DataWithPatch)
		if !ok {
			MethodNotAllowed()
			return
		}
		var patch user.DataPatch
		if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
			handle(err)
			return
		}
		if err := handler.Patch(ctx, patch); err != nil {
			handle(err)
			return
		}
	}
}
