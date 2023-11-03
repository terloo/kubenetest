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

func (t *PingTester) TestPing(addr string) (*PingTestResult, error) {

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

	stats := pinger.Statistics()

	rtts := make([]int, len(stats.Rtts))
	for i, rtt := range stats.Rtts {
		rtts[i] = int(rtt.Milliseconds())
	}

	result := &PingTestResult{
		Addr:    stats.Addr,
		PkgSent: stats.PacketsSent,
		PkgRecv: stats.PacketsRecv,
		AvgRtt:  int(stats.AvgRtt.Milliseconds()),
		MaxRtt:  int(stats.MaxRtt.Milliseconds()),
		MinRtt:  int(stats.MinRtt.Milliseconds()),
		Rtts:    rtts,
	}

	return result, nil
}
