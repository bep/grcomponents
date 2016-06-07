// Package loadbar implements a gr (Go React) component that wraps other component(s) to provide visual state changes.
// This can be used to create loader animations on data loading etc.

package loadbar

import (
	"github.com/bep/gr"
	"github.com/bep/gr/el"
)

type LoadState int

const (
	StateInitial LoadState = iota
	StateLoading
	StateLoaded
)

const loadStateKey = "loadState"

type loadBarRenderer struct {
	*gr.This
}

// LoadBar is a convenience composite of the component and the component implementation needed to return status updates.
// See SetStatus.
type LoadBar struct {
	*gr.ReactComponent
	delegate *loadBarRenderer
}

// NewLoadBar creates a new load bar with the optional options.
func NewLoadBar(options ...gr.Option) *LoadBar {
	lb := &loadBarRenderer{}
	rc := gr.New(lb, options...)

	return &LoadBar{ReactComponent: rc, delegate: lb}
}

func (c *loadBarRenderer) Render() gr.Component {

	var loadingBarWrapperCss gr.Modifier

	state := c.loadState()

	switch state {
	case StateInitial:
		loadingBarWrapperCss = gr.CSS("gr-lb-wrapper gr-lb-initial")
	case StateLoading:
		loadingBarWrapperCss = gr.CSS("gr-lb-wrapper gr-lb-loading")
	case StateLoaded:
		loadingBarWrapperCss = gr.CSS("gr-lb-wrapper gr-lb-loaded")
	}

	bar := el.Div(
		gr.CSS("gr-lb"),
		el.Div(gr.CSS("gr-lb-bar")),
		el.Div(gr.CSS("gr-lb-bar")),
		el.Div(gr.CSS("gr-lb-bar")),
		el.Div(gr.CSS("gr-lb-bar")),
	)

	return el.Div(loadingBarWrapperCss,
		bar,
		el.Div(gr.CSS("gr-lb-component"), c.Children().Element()),
	)

}

// SetLoadState is the callback used by event handler to trigger a state change.
func (lb *LoadBar) SetLoadState(s LoadState) {
	lb.delegate.setLoadState(s)
}

func (c *loadBarRenderer) loadState() LoadState {
	return LoadState(c.State().Int(loadStateKey))
}

func (c *loadBarRenderer) setLoadState(s LoadState) {
	if c.IsMounted() {
		c.SetState(gr.State{loadStateKey: s})
	}
}

func (c *loadBarRenderer) GetInitialState() gr.State {
	return gr.State{loadStateKey: StateInitial}
}

