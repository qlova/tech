package app_test

import (
	"testing"

	"qlova.tech/app"
	"qlova.tech/app/page"
	"qlova.tech/app/show"
	"qlova.tech/app/user"
	"qlova.tech/app/user/hint"
	"qlova.tech/app/user/they"
)

type TodoList struct {
	NewItemName string
	Items       []string
}

func (list *TodoList) AddTask() user.Steps {
	return user.Steps{
		they.Append(&list.NewItemName, &list.Items),
		they.Set(&list.NewItemName, ""),
	}
}

func (list *TodoList) RenderPage() page.View {
	return page.New(
		user.Read("My Todo List:", hint.Title),
		user.View(hint.Row)(
			user.Read("New Item:"),
			user.Pick(&list.NewItemName),
			user.Path(list.AddTask, hint.Button("Add")),
		),
		user.View()(
			user.Read("Items:"),
			user.List(&list.Items, func(item *string) show.Layout {
				return user.View()(
					user.Show(item),
					user.Path(they.Remove(item, &list.Items), hint.OnClick),
				)
			}),
		),
	)
}

func TestApp(t *testing.T) {
	if err := app.ListenAndServe(":8080", new(TodoList)); err != nil {
		t.Fatal(err)
	}
}
