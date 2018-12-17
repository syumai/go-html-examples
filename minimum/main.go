package main

import "syscall/js"

type Node struct {
	Name     string
	Attrs    map[string]string
	Children []interface{} // Node or string
}

var document = js.Global().Get("document")

func CreateElement(i interface{}) js.Value {
	var el js.Value
	switch v := i.(type) {
	case string:
		el = document.Call("createTextNode", v)
	case Node:
		el = document.Call("createElement", v.Name)
		for k, v := range v.Attrs {
			el.Set(k, v)
		}
		for _, childNode := range v.Children {
			childEl := CreateElement(childNode)
			el.Call("appendChild", childEl)
		}
	default:
		el = document.Call("createTextNode", "")
	}
	return el
}

func Render(node Node, container js.Value) {
	container.Call("appendChild", CreateElement(node))
}

func main() {
	html := Node{
		Name: "div",
		Children: []interface{}{
			Node{
				Name: "h2",
				Children: []interface{}{
					"HTML Example",
				},
			},
			Node{
				Name:  "a",
				Attrs: map[string]string{"href": "https://golang.org"},
				Children: []interface{}{
					"Go to golang.org",
				},
			},
		},
	}

	body := document.Get("body")
	Render(html, body)
}
