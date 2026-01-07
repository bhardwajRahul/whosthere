package pages

import (
	"fmt"
	"strings"

	"github.com/ramonvermeulen/whosthere/internal/ui/navigation"
	"github.com/rivo/tview"
)

var _ navigation.Page = (*SplashPage)(nil)

var LogoBig = []string{
	`Knock Knock..                                                     `,
	`                _               _   _                   ___       `,
	`      __      _| |__   ___  ___| |_| |__   ___ _ __ ___/ _ \      `,
	`      \ \ /\ / / '_ \ / _ \/ __| __| '_ \ / _ \ '__/ _ \// /      `,
	`       \ V  V /| | | | (_) \__ \ |_| | | |  __/ | |  __/ \/       `,
	`        \_/\_/ |_| |_|\___/|___/\__|_| |_|\___|_|  \___| ()       `,
	"\n",
	"\n",
	"\n",
}

// SplashPage adapts the splash logo into a Page.
type SplashPage struct {
	root    *tview.Flex
	version string
}

func (p *SplashPage) Refresh() {}

func NewSplashPage(version string) *SplashPage {
	s := &SplashPage{root: tview.NewFlex().SetDirection(tview.FlexRow), version: version}

	logo := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	logo.SetTextColor(tview.Styles.SecondaryTextColor)

	logoText := strings.Join(LogoBig, "\n")
	_, err := fmt.Fprint(logo, logoText)
	if err != nil {
		return nil
	}

	logoLines := len(strings.Split(logoText, "\n"))

	centeredLogo := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(logo, logoLines, 0, false).
		AddItem(nil, 0, 1, false)

	versionView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	if version != "" {
		_, _ = fmt.Fprintf(versionView, "whosthere - v%s", version)
	}

	// Root layout: logo takes remaining space, version line is always at bottom.
	s.root.
		AddItem(centeredLogo, 0, 1, false).
		AddItem(versionView, 1, 0, false)

	return s
}

func (p *SplashPage) GetName() string { return navigation.RouteSplash }

func (p *SplashPage) GetPrimitive() tview.Primitive { return p.root }

func (p *SplashPage) FocusTarget() tview.Primitive { return p.root }

func (p *SplashPage) RefreshFromState() {}
