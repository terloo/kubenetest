package ping

type PingTestResult struct {
	Addr    string
	PkgSent int
	PkgRecv int
	AvgRtt  int
	MaxRtt  int
	MinRtt  int
	Rtts    []int
}
