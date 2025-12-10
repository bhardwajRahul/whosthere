package arp

import (
	"context"
	"sync"

	"github.com/ramonvermeulen/whosthere/internal/discovery"
	"go.uber.org/zap"
)

var _ discovery.Scanner = (*Scanner)(nil)

// Scanner implements ARP-based discovery.
// On first run, it triggers ARP requests for the entire /24 subnet.
// On subsequent runs, it triggers a lightweight refresh before reading the ARP cache.
type Scanner struct {
	firstRunDone bool
	firstRunLock sync.Mutex

	logger *zap.Logger
}

func (s *Scanner) Name() string { return "arp" }

// Scan performs ARP discovery.
// On first call: triggers ARP for entire /24 subnet, then reads cache.
// On subsequent calls: performs a short trigger, waits briefly, then reads cache.
func (s *Scanner) Scan(ctx context.Context, out chan<- discovery.Device) error {
	if s.logger == nil {
		s.logger = zap.L().With(zap.String("scanner", s.Name()))
	}

	s.firstRunLock.Lock()
	isFirstRun := !s.firstRunDone
	if isFirstRun {
		s.firstRunDone = true
	}
	s.firstRunLock.Unlock()

	if isFirstRun {
		s.logger.Info("First ARP scan - triggering full subnet sweep")
		go s.triggerFullSubnetSweep()
	}

	return s.readARPCache(ctx, out)
}
