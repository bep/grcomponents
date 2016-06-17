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

	o Options
}

// Options provides optional options. All of these have sane defaults.
type Options struct {
	OnClass           string
	OffClass          string
	OnText            string
	OffText           string
	OffPanelClass     string
	OnPanelClass      string
	OffPanelBodyClass string
	OnPanelBodyClass  string

	// Panel will, by default, start in its closed state. Set this flag to change that.
	StartOpen bool
}

// New creates a new toggle wrapper with sane defaults.
func New() *gr.ReactComponent {
	return gr.New(&toggle{})
}

// NewWith cretes a new toggle wrapper with the given options.
// TODO(bep) options: wrapper element.
func NewWith(options Options) *gr.ReactComponent {
	return gr.New(&toggle{o: options})
}

func (c *toggle) Render() gr.Component {

	var (
		toggleStyle gr.Modifier
		buttonText  string

		panelClass     gr.Modifier
		panelBodyClass gr.Modifier = gr.Discard
	)

	enabled := c.State().Bool(toggleKey)

	if enabled {
		if c.o.OnClass != "" {
			toggleStyle = gr.CSS(c.o.OnClass)
		} else {
			toggleStyle = gr.Style("display", "block")
		}

		buttonText = def(c.o.OnText, "Hide")

		panelClass = gr.CSS(def(c.o.OnPanelClass, "panel-default"))

		if c.o.OnPanelBodyClass != "" {
			panelBodyClass = gr.CSS(c.o.OnPanelBodyClass)
		}

	} else {
		if c.o.OffClass != "" {
			toggleStyle = gr.CSS(c.o.OffClass)
		} else {
			toggleStyle = gr.Style("display", "none")
		}
		buttonText = def(c.o.OffText, "Show")

		if c.o.OffPanelBodyClass != "" {
			panelBodyClass = gr.CSS(c.o.OffPanelBodyClass)
		}

		panelClass = gr.CSS(def(c.o.OffPanelClass, "panel-default"))
	}

	return el.Div(
		gr.CSS("panel"), panelClass,
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
				panelBodyClass,
				c.Children().Element(),
			),
		),
	)

}

func (c *toggle) GetInitialState() gr.State {
	return gr.State{toggleKey: c.o.StartOpen}
}

func (c *toggle) onClick(event *gr.Event) {
	c.SetState(gr.State{toggleKey: !c.State().Bool(toggleKey)})
}

func def(v, d string) string {
	if v != "" {
		return v
	}
	return d
}
