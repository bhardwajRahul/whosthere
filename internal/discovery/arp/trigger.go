package arp

import (
	"net"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	maxConcurrentTriggers = 200
	triggerDeadline       = 1200 * time.Millisecond
	tcpDialTimeout        = 300 * time.Millisecond
)

var (
	udpTriggerPorts = []int{9, 1900, 5353}
	tcpTriggerPorts = []int{80, 443, 22, 554, 8009}
)

// triggerFullSubnetSweep sends ARP triggers for all IPs in /24 subnet (non-blocking).
// to send out ARP directly you need root privileges, so we use indirect methods.
func (s *Scanner) triggerFullSubnetSweep() {
	if s.logger == nil {
		s.logger = zap.L().With(zap.String("scanner", s.Name()))
	}

	localIP, subnet, err := getLocalNetwork()
	if err != nil {
		s.logger.Warn("Could not get local network for ARP trigger", zap.Error(err))
		return
	}

	// Only trigger for /24 subnets
	if ones, bits := subnet.Mask.Size(); ones != 24 || bits != 32 {
		s.logger.Debug("Skipping ARP trigger - not a /24 network",
			zap.String("subnet", subnet.String()),
			zap.Int("prefix", ones))
		return
	}

	ips := generateSubnetIPs(subnet, localIP)
	if len(ips) == 0 {
		s.logger.Debug("No IPs to trigger ARP for")
		return
	}

	s.logger.Info("Triggering ARP requests for subnet",
		zap.String("subnet", subnet.String()))

	// Trigger with concurrency limit
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrentTriggers) // Limit concurrent goroutines

	for _, ip := range ips {
		wg.Add(1)
		sem <- struct{}{}

		go func(targetIP net.IP) {
			defer wg.Done()
			defer func() { <-sem }()
			s.sendARPTarget(targetIP)
		}(ip)
	}

	wg.Wait()
	s.logger.Debug("ARP triggering completed")
}

// sendARPTarget tries multiple transport triggers to reliably force an ARP resolution.
// To send ARP requests directly you need root privileges, so we use indirect methods.
func (s *Scanner) sendARPTarget(ip net.IP) {
	deadline := time.Now().Add(triggerDeadline)

	for _, p := range udpTriggerPorts {
		addr := &net.UDPAddr{IP: ip, Port: p}
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			continue
		}
		_ = conn.SetWriteDeadline(deadline)
		_, _ = conn.Write([]byte{0})
		_ = conn.Close()
	}

	for _, p := range tcpTriggerPorts {
		addr := net.JoinHostPort(ip.String(), strconv.Itoa(p))
		c, err := net.DialTimeout("tcp", addr, tcpDialTimeout)
		if err == nil {
			_ = c.Close()
		}
	}
}
