package html

import (
	"bufio"
	"fmt"
	"html"
	"reflect"
	"strings"

	"qlova.tech/app/data"
	"qlova.tech/app/page"
	"qlova.tech/app/show"
	"qlova.tech/app/then"
	"qlova.tech/app/user"
)

func deref(sync data.Sync, v any) string {
	path := data.Path(sync, v)
	if path != "" {
		return fmt.Sprintf("'%s'", path)
	}
	return fmt.Sprintf("'%v'", v)
}

func renderJS(w *bufio.Writer, step then.Step, sync data.Sync) {
	switch step := step.(type) {
	case then.Steps:
		for _, step := range step {
			renderJS(w, step, sync)
		}
	case then.Append:
		fmt.Fprintf(w, `append(%v, %v);`, deref(sync, step.List), deref(sync, step.Item))
	case then.Set:
		fmt.Fprintf(w, `set(%v, %v);`, deref(sync, step.Variable), deref(sync, step.Value))
	default:
		fmt.Println("unknown type", reflect.TypeOf(step))
	}
}

func render(w *bufio.Writer, node show.Node, sync data.Sync) {
	switch elem := node.(type) {
	case page.View:
		w.WriteString("<DOCTYPE html><html><head></head><body>")
		for _, node := range elem.Nodes {
			render(w, node, sync)
		}
		w.WriteString("</body></html>")
	case show.View:
		w.WriteString("<div")
		if elem.Hints.Row {
			w.WriteString(` class="row"`)
		}
		w.WriteString(">")
		for _, node := range elem.Nodes {
			render(w, node, sync)
		}
		w.WriteString("</div>")
	case show.String:
		tag := "p"
		if elem.Hints.Title {
			tag = "h2"
		}
		w.WriteString("<" + tag + ">")
		w.WriteString(html.EscapeString(elem.Value))
		w.WriteString("</" + tag + ">")
	case show.Picker[string]:
		w.WriteString(`<input type="text" data-edit="` + data.Path(sync, elem.Value) + `">`)
	case show.Choice:
		if elem.Hints.Button != "" {
			w.WriteString("<button")
			w.WriteString(` onclick="`)
			renderJS(w, elem.Steps, sync)
			w.WriteString(`"`)
			w.WriteRune('>')
			w.WriteString(html.EscapeString(elem.Hints.Button))
			w.WriteString("</button>")
		}
	default:
		fmt.Println("unknown type", reflect.TypeOf(elem))
	}
}

func Render(ui user.Interface, sync data.Sync) string {
	var body strings.Builder

	w := bufio.NewWriter(&body)

	render(w, ui, sync)

	w.Flush()

	return body.String()
}
