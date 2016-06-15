// Package loadbar implements a gr (Go React) component that wraps other component(s) to provide visual state changes.
// This can be used to create loader animations on data loading etc.
package loadbar

import (
	"github.com/bep/gr"
	"github.com/bep/gr/el"
)

// LoadState represents the load state of the component.
type LoadState int

const (
	// StateInitial is the inital loading state.
	StateInitial LoadState = iota
	// StateLoading is the state set when the component is loading.
	StateLoading
	// StateLoaded is the state set when the component is loaded.
	StateLoaded
)

const loadStateKey = "loadState"

type loadBarRenderer struct {
	*gr.This
}

// Loader is a convenience composite of the component and the component implementation needed to return status updates.
// See SetStatus.
type Loader struct {
	*gr.ReactComponent
	delegate *loadBarRenderer
}

// NewLoader creates a new load bar with the optional options.
func NewLoader(options ...gr.Option) *Loader {
	lb := &loadBarRenderer{}
	rc := gr.New(lb, options...)

	return &Loader{ReactComponent: rc, delegate: lb}
}

func (c *loadBarRenderer) Render() gr.Component {

	var loadingBarWrapperCSS gr.Modifier

	state := c.loadState()

	switch state {
	case StateInitial:
		loadingBarWrapperCSS = gr.CSS("gr-lb-wrapper gr-lb-initial")
	case StateLoading:
		loadingBarWrapperCSS = gr.CSS("gr-lb-wrapper gr-lb-loading")
	case StateLoaded:
		loadingBarWrapperCSS = gr.CSS("gr-lb-wrapper gr-lb-loaded")
	}

	bar := el.Div(
		gr.CSS("gr-lb"),
		el.Div(gr.CSS("gr-lb-bar")),
		el.Div(gr.CSS("gr-lb-bar")),
		el.Div(gr.CSS("gr-lb-bar")),
		el.Div(gr.CSS("gr-lb-bar")),
	)

	return el.Div(loadingBarWrapperCSS,
		bar,
		el.Div(gr.CSS("gr-lb-component"), c.Children().Element()),
	)

}

// SetLoadState is the callback used by event handler to trigger a state change.
func (lb *Loader) SetLoadState(s LoadState) {
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
