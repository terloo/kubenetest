package server

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/terloo/kubenetest/pkg/ping"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/klog/v2"
)

var (
	LocalIP string
)

func RunServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handleWrapChain(pingHandle))
	healthz.InstallReadyzHandler(mux)

	klog.Info("server running...")
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		return err
	}

	return nil
}

func pingHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	addr := r.URL.Query().Get("addr")
	stats, err := ping.NewPingTester(LocalIP, int16(rand.Int())).Ping(addr)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	rtts := make([]int, len(stats.Rtts))
	for i, rtt := range stats.Rtts {
		rtts[i] = int(rtt.Milliseconds())
	}

	resp := &PingResponse{
		Addr:    stats.Addr,
		PkgSent: stats.PacketsSent,
		PkgRecv: stats.PacketsRecv,
		AvgRtt:  int(stats.AvgRtt.Milliseconds()),
		MaxRtt:  int(stats.MaxRtt.Milliseconds()),
		MinRtt:  int(stats.MinRtt.Milliseconds()),
		Rtts:    rtts,
	}

	json.NewEncoder(w).Encode(resp)
}
