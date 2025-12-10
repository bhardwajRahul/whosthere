package arp

import (
	"context"
	"fmt"

	"github.com/ramonvermeulen/whosthere/internal/discovery"
	"go.uber.org/zap"
)

// should be implemented later, currently only darwin is supported
func (s *Scanner) readLinuxARPCache(ctx context.Context, out chan<- discovery.Device) error {
	log := zap.L().With(zap.String("scanner", "arp"))
	log.Warn("Linux ARP cache reader is not implemented; skipping")
	return fmt.Errorf("linux ARP cache reader not implemented")
}
