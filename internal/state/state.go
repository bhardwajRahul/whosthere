package state

import (
	"sync"

	"github.com/ramonvermeulen/whosthere/internal/discovery"
)

// AppState holds application-level state shared across views and
// orchestrated by the App. Scanners do not write here directly.
type AppState struct {
	mu sync.RWMutex

	devices    map[string]discovery.Device
	selectedIP string
	listeners  []func(discovery.Device)
}

func NewAppState() *AppState {
	return &AppState{
		devices: make(map[string]discovery.Device),
	}
}

// AddListener registers a callback invoked when a device is upserted.
func (s *AppState) AddListener(fn func(discovery.Device)) {
	if fn == nil {
		return
	}
	s.mu.Lock()
	s.listeners = append(s.listeners, fn)
	s.mu.Unlock()
}

// UpsertDevice merges a device into the canonical device map.
func (s *AppState) UpsertDevice(d *discovery.Device) {
	if d.IP == nil {
		return
	}
	key := d.IP.String()
	if key == "" {
		return
	}

	s.mu.Lock()
	if existing, ok := s.devices[key]; ok {
		existing.Merge(d)
		s.devices[key] = existing
	} else {
		s.devices[key] = *d
	}
	updated := s.devices[key]
	listeners := append([]func(discovery.Device){}, s.listeners...)
	s.mu.Unlock()

	for _, fn := range listeners {
		fn(updated)
	}
}

// DevicesSnapshot returns a copy of all devices for rendering.
func (s *AppState) DevicesSnapshot() []discovery.Device {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]discovery.Device, 0, len(s.devices))
	for _, d := range s.devices {
		out = append(out, d)
	}
	return out
}

// SetSelectedIP stores the currently selected device IP.
func (s *AppState) SetSelectedIP(ip string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.selectedIP = ip
}

// Selected returns the currently selected device, if any.
func (s *AppState) Selected() (discovery.Device, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.selectedIP == "" {
		return discovery.Device{}, false
	}
	d, ok := s.devices[s.selectedIP]
	return d, ok
}

// SelectedIP returns the currently selected device IP, if any.
func (s *AppState) SelectedIP() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.selectedIP
}
