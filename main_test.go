package main

import (
	"testing"
)

const ipLinkOut = `1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: enp1s0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc mq state DOWN mode DEFAULT group default qlen 1000
    link/ether 00:8c:fa:6a:09:87 brd ff:ff:ff:ff:ff:ff
3: wlp2s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DORMANT group default qlen 1000
    link/ether e0:b9:a5:1e:e0:88 brd ff:ff:ff:ff:ff:ff
`

func TestPrettyOut(t *testing.T) {
	//arrange
	expectedPrettyOut := []string{
		"1 lo 00:00:00:00:00:00",
		"2 enp1s0 00:8c:fa:6a:09:87",
		"3 wlp2s0 30:b9:a5:1e:e0:88",
	}
	//act
	result := prettyOut([]byte(ipLinkOut))
	//assert
	if result != expectedPrettyOut {
		t.Error(result)
	}
}
