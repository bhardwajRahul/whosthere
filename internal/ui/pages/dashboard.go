package pages

import (
	"github.com/derailed/tview"
	"github.com/ramonvermeulen/whosthere/internal/state"
	"github.com/ramonvermeulen/whosthere/internal/ui/components"
	"github.com/ramonvermeulen/whosthere/internal/ui/navigation"
)

var _ navigation.Page = &DashboardPage{}

// DashboardPage is the dashboard showing discovered devices.
type DashboardPage struct {
	*tview.Flex
	deviceTable *components.DeviceTable
	spinner     *components.Spinner
	state       *state.AppState

	navigate func(route string)
}

func NewDashboardPage(s *state.AppState, navigate func(route string)) *DashboardPage {
	t := components.NewDeviceTable()
	spinner := components.NewSpinner()
	spinner.SetSuffix(" Scanning...")

	main := tview.NewFlex().SetDirection(tview.FlexRow)
	main.AddItem(
		tview.NewTextView().
			SetText("whosthere").
			SetTextAlign(tview.AlignCenter),
		0, 1, false,
	)
	main.AddItem(t, 0, 18, true)

	status := tview.NewFlex().SetDirection(tview.FlexColumn)
	status.AddItem(spinner.View(), 0, 1, false)
	status.AddItem(
		tview.NewTextView().
			SetText("j/k: up/down - g/G: top/bottom - Enter: details").
			SetTextAlign(tview.AlignRight),
		0, 1, false,
	)
	main.AddItem(status, 1, 0, false)

	dp := &DashboardPage{
		Flex:        main,
		deviceTable: t,
		spinner:     spinner,
		state:       s,
		navigate:    navigate,
	}

	t.SetSelectedFunc(func(row, col int) {
		ip := t.SelectedIP()
		if ip == "" {
			return
		}
		s.SetSelectedIP(ip)
		if dp.navigate != nil {
			dp.navigate(navigation.RouteDetail)
		}
	})

	return dp
}

func (p *DashboardPage) GetName() string { return navigation.RouteDashboard }

func (p *DashboardPage) GetPrimitive() tview.Primitive { return p }

func (p *DashboardPage) FocusTarget() tview.Primitive { return p.deviceTable }

func (p *DashboardPage) Spinner() *components.Spinner { return p.spinner }

func (p *DashboardPage) RefreshFromState() {
	devices := p.state.DevicesSnapshot()
	p.deviceTable.ReplaceAll(devices)
}

func (p *DashboardPage) Refresh() {
	p.RefreshFromState()
}
