package iframe

import (
	"github.com/bep/gr"
	"github.com/bep/gr/attr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
	"github.com/bep/grcomponents/loadbar"
)


// IFrame renders the IFrame.
type IFrame struct {
	*gr.This

	userProvidedModifers gr.Modifiers
}

// IFrameLoader is a composite that constructs the iFrame with the loader.
// Once constructed by NewIFrame, new URLs can be set in the Frame method.
type IFrameLoader struct {
	iFrame *gr.ReactComponent
	loader *loadbar.Loader
}

const (
	urlKey = "url"
)

// NewIFrame creates a new IFrame with a loader wrapper. Modifiers such as width, height and style can be provided.
// If none, defaults will be set.
//
// See http://github.com/bep/grcomponents/loadbar for how to style the loader.
func NewIFrame(modifiers ...gr.Modifier) *IFrameLoader {
	i := &IFrame{userProvidedModifers: modifiers}
	ic := gr.New(i)
	return &IFrameLoader{iFrame: ic, loader: loadbar.NewLoader()}
}

// Frame loads the given URL in the iframe.
func (i *IFrameLoader) Frame(url string) *gr.Element {
	return i.loader.CreateElement(nil, i.iFrame.CreateElement(gr.Props{urlKey: url, "IFrameLoaded": i.loader.SetLoadState}))
}

func (g IFrame) Render() gr.Component {

	url := g.Props().String(urlKey)

	// Just render an empty Div by default.
	var frame *gr.Element = el.Div()

	if url != "" {
		frame = el.InlineFrame(
			evt.Load(g.iFrameLoaded), attr.Src(url))

		if len(g.userProvidedModifers) > 0 {
			g.userProvidedModifers.Modify(frame)

		} else {
			gr.Modifiers{attr.Ref("iframe"), attr.Width("100%"), gr.Style("border", "none"),
				attr.Height("800px")}.Modify(frame)
		}
	}
	//
	return el.Div(frame)

}

func (g IFrame) iFrameLoaded(event *gr.Event) {
	g.Props().Call("IFrameLoaded", loadbar.StateLoaded)
}

func (g IFrame) ShouldComponentUpdate(next gr.Cops) bool {
	if  g.Props().HasChanged(next.Props, urlKey) {
		g.Props().Call("IFrameLoaded", loadbar.StateLoading)

	}
	return true
}
