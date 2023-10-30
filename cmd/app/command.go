package app

import (
	"errors"
	"net/netip"

	"github.com/spf13/cobra"
	"github.com/terloo/kubenetest/pkg/ping"
	"github.com/terloo/kubenetest/pkg/server"
)

func NewKubenetestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubentest",
		Short: "network test for k8s",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(server.LocalIP) == 0 {
				return errors.New("flag --local-ip is needed")
			}

			_, err := netip.ParseAddr(server.LocalIP)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := server.RunServer()
			return err
		},
	}

	fs := cmd.Flags()

	fs.IntVar(&ping.PingCount, "ping-count", 3, "numbers of ping")
	fs.IntVar(&ping.PingTimeOut, "ping-timeout", 3, "timeout of ping (s)")
	fs.StringVar(&server.LocalIP, "local-ip", "", "local ip")

	return cmd
}
