package server

type PingResponse struct {
	Addr    string `json:"addr"`
	PkgSent int    `json:"PkgSent"`
	AvgRtt  int    `json:"avgCost"`
	MaxRtt  int    `json:"MaxRtt"`
	MinRtt  int    `json:"MinRtt"`
	Rtts    []int  `json:"costs"`
}
