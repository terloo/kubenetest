package server

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/terloo/kubenetest/pkg/infra"
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
	mux.HandleFunc("/infra", handleWrapChain(infraHandle))
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
	result, err := ping.NewPingTester(LocalIP, int16(rand.Int())).TestPing(addr)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	resp := &PingResponse{
		Addr:    result.Addr,
		PkgSent: result.PkgSent,
		PkgRecv: result.PkgRecv,
		AvgRtt:  result.AvgRtt,
		MaxRtt:  result.MaxRtt,
		MinRtt:  result.MinRtt,
		Rtts:    result.Rtts,
	}

	json.NewEncoder(w).Encode(resp)
}

func infraHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	result, err := infra.NewInfraTester().TestInfra()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	resp := &InfraResponse{
		Bridgenf:              result.Bridgenf.Error(),
		Bridgenf6:             result.Bridgenf6.Error(),
		Ipv4Forward:           result.Ipv4Forward.Error(),
		Ipv6DefaultForwarding: result.Ipv4Forward.Error(),
	}

	json.NewEncoder(w).Encode(resp)
}
