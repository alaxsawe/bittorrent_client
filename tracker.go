package main

import (
	"fmt"
)

type Tracker struct {
	Meta MetaInfo
}

type TrackerResponse struct {
	MinInterval string
	Interval string
	Peers string
}

func (tr *TrackerResponse) GetPeerAddresses() []string {
	var result []string
	for i := 0; i < len(tr.Peers); i += 6 {
		ipAndPort := fmt.Sprintf("%d.%d.%d.%d:%d", tr.Peers[i+0], tr.Peers[i+1],
		tr.Peers[i+2],tr.Peers[i+3], (uint16(tr.Peers[i+4])<<8)|uint16(tr.Peers[i+5]))
		result = append(result, ipAndPort)
	}
	return result
}
