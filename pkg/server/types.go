package server

type PingResponse struct {
	Addr    string `json:"addr,omitempty"`
	PkgSent int    `json:"pkgSent,omitempty"`
	PkgRecv int    `json:"pkgRecv,omitempty"`
	AvgRtt  int    `json:"avgRtt,omitempty"`
	MaxRtt  int    `json:"maxRtt,omitempty"`
	MinRtt  int    `json:"minRtt,omitempty"`
	Rtts    []int  `json:"rtts,omitempty"`
}

type InfraResponse struct {
	Bridgenf              string `json:"bridgenf,omitempty"`
	Bridgenf6             string `json:"bridgenf6,omitempty"`
	Ipv4Forward           string `json:"ipv4Forward,omitempty"`
	Ipv6DefaultForwarding string `json:"ipv6DefaultForwarding,omitempty"`
}
