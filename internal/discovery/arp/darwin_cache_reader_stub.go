//go:build !darwin && !freebsd && !netbsd && !openbsd

package arp

import (
	"context"

	"github.com/ramonvermeulen/whosthere/internal/discovery"
)

// readDarwinARPCache is a no-op on non-Darwin/BSD platforms; the real implementation
// exists in darwin_cache_reader.go and is gated by build tags.
func (s *Scanner) readDarwinARPCache(ctx context.Context, out chan<- discovery.Device) error {
	return nil
}
