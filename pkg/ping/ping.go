package ping

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

var (
	PingCount   int
	PingTimeOut int
)

type PingTester struct {
	LocalIP string
	IcmpID  int16
}

func NewPingTester(localIP string, icmpID int16) *PingTester {
	p := &PingTester{
		LocalIP: localIP,
		IcmpID:  icmpID,
	}
	return p
}

func (t *PingTester) Ping(addr string) (*probing.Statistics, error) {

	pinger, err := probing.NewPinger(addr)
	if err != nil {
		return nil, err
	}
	pinger.Timeout = time.Duration(PingTimeOut) * time.Second
	pinger.Count = PingCount
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	return pinger.Statistics(), nil
}
