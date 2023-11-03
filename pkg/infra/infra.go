package infra

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	bridgenf              = "/proc/sys/net/bridge/bridge-nf-call-iptables"
	bridgenf6             = "/proc/sys/net/bridge/bridge-nf-call-ip6tables"
	ipv4Forward           = "/proc/sys/net/ipv4/ip_forward"
	ipv6DefaultForwarding = "/proc/sys/net/ipv6/conf/default/forwarding"
)

type InfraTester struct {
}

func NewInfraTester() *InfraTester {
	tester := &InfraTester{}
	return tester
}

func (_ *InfraTester) TestInfra() (*InfraTestResult, error) {

	result := &InfraTestResult{
		Bridgenf:              (&fileContentCheck{Path: bridgenf, Content: []byte{'1'}}).Check(),
		Bridgenf6:             (&fileContentCheck{Path: bridgenf6, Content: []byte{'1'}}).Check(),
		Ipv4Forward:           (&fileContentCheck{Path: ipv4Forward, Content: []byte{'1'}}).Check(),
		Ipv6DefaultForwarding: (&fileContentCheck{Path: ipv6DefaultForwarding, Content: []byte{'1'}}).Check(),
	}

	return result, nil
}

type fileContentCheck struct {
	Path    string
	Content []byte
}

func (c *fileContentCheck) Check() error {
	f, err := os.Open(bridgenf)
	if err != nil {
		return err
	}

	fc, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	if !bytes.Equal(c.Content, fc) {
		return errors.New(fmt.Sprintf("not equal %s", string(c.Content)))
	}

	return nil
}
