package arp

import "errors"

var (
	ErrNoIPv4Interface = errors.New("arp: no IPv4 network interface found")
)
