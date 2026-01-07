package navigation

import "github.com/rivo/tview"

// Page is a UI page that can be registered with the Router.
type Page interface {
	GetName() string
	GetPrimitive() tview.Primitive
	FocusTarget() tview.Primitive
	Refresh()
}

const (
	RouteDashboard = "dashboard"
	RouteSplash    = "splash"
	RouteDetail    = "detail"
)

// Router is both the visual pages container and the logical router.
type Router struct {
	*tview.Pages
	pages       map[string]Page
	currentPage string
}

func NewRouter() *Router {
	return &Router{
		Pages: tview.NewPages(),
		pages: make(map[string]Page),
	}
}

func (r *Router) Register(p Page) {
	name := p.GetName()
	r.pages[name] = p
	r.AddPage(name, p.GetPrimitive(), true, false)
}

func (r *Router) NavigateTo(name string) {
	if _, ok := r.pages[name]; !ok {
		return
	}
	r.currentPage = name
	r.SwitchToPage(name)
	r.pages[name].Refresh()
}

func (r *Router) FocusCurrent(app *tview.Application) {
	if app == nil {
		return
	}
	p, ok := r.pages[r.currentPage]
	if !ok || p == nil {
		app.SetFocus(r)
		return
	}
	if ft := p.FocusTarget(); ft != nil {
		app.SetFocus(ft)
		return
	}
	app.SetFocus(p.GetPrimitive())
}

func (r *Router) Current() string { return r.currentPage }

func (r *Router) Page(name string) Page {
	return r.pages[name]
}
