// Package collapsible provides a panel that can be toggled on and off.
package collapsible

import (
	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
)

const toggleKey = "toggleState"

type toggle struct {
	*gr.This
}

// New creates a new toggle wrapper.
// TODO(bep) options: text, style, wrapper element.
func New() *gr.ReactComponent {
	return gr.New(&toggle{})
}

func (c *toggle) Render() gr.Component {

	var (
		toggleStyle gr.Modifier
		buttonText  string
	)

	enabled := c.State().Bool(toggleKey)

	if enabled {
		toggleStyle = gr.Style("display", "block")
		buttonText = "Hide"
	} else {
		toggleStyle = gr.Style("display", "none")
		buttonText = "Show"
	}

	return el.Div(gr.CSS("panel panel"),
		el.Div(
			gr.CSS("panel panel"),
			el.Div(
				gr.CSS("panel-heading"),
				el.Header4(
					gr.CSS("panel-title"),
					el.Button(
						gr.CSS("btn", "btn-link"),
						gr.Text(buttonText),
						evt.Click(c.onClick)),
				),
			),
			el.Div(
				toggleStyle,
				el.Div(
					gr.CSS("panel-body"),
					c.Children().Element(),
				),
			),
		),
	)

}

func (c *toggle) GetInitialState() gr.State {
	return gr.State{toggleKey: false}
}

func (c *toggle) onClick(event *gr.Event) {
	c.SetState(gr.State{toggleKey: !c.State().Bool(toggleKey)})
}
