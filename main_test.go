package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ipLinkOut = `1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: enp1s0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc mq state DOWN mode DEFAULT group default qlen 1000
    link/ether 00:8c:fa:6a:09:87 brd ff:ff:ff:ff:ff:ff
3: wlp2s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DORMANT group default qlen 1000
    link/ether e0:b9:a5:1e:e0:88 brd ff:ff:ff:ff:ff:ff
`

const firstAdapter = "1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000 link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00"

var ipLinkOutCutted = []string{
	"1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000 link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00", "2: enp1s0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc mq state DOWN mode DEFAULT group default qlen 1000 link/ether 00:8c:fa:6a:09:87 brd ff:ff:ff:ff:ff:ff", "3: wlp2s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DORMANT group default qlen 1000 link/ether e0:b9:a5:1e:e0:88 brd ff:ff:ff:ff:ff:ff",
}

func TestCutter(t *testing.T) {
	//act
	result := cutter([]byte(ipLinkOut))
	//assert
	assert.Equal(t, ipLinkOutCutted, result)
}

func TestWordExtractor01(t *testing.T) {
	//arrange
	expected := "1"
	//act
	result := wordExtractor(firstAdapter, 0)
	//assert
	assert.Equal(t, expected, result)
}

func TestWordExtractor02(t *testing.T) {
	//arrange
	expected := "lo"
	//act
	result := wordExtractor(firstAdapter, 1)
	//assert
	assert.Equal(t, expected, result)
}

func TestWordExtractor03(t *testing.T) {
	//arrange
	expected := "00:00:00:00:00:00"
	//act
	result := wordExtractor(firstAdapter, 16)
	//assert
	assert.Equal(t, expected, result)
}

func TestWordExtractor04(t *testing.T) {
	//arrange
	expected := "DEFAULT"
	//act
	result := wordExtractor(firstAdapter, 10)
	//assert
	assert.Equal(t, expected, result)
}

func TestAggregator(t *testing.T) {
	//arrange
	var expected = []adapter{{1, "lo", "00:00:00:00:00:00", "DEFAULT"}, {2, "enp1s0", "00:8c:fa:6a:09:87", "DEFAULT"}, {3, "wlp2s0", "e0:b9:a5:1e:e0:88", "DORMANT"}}
	//act
	result := aggregator(ipLinkOutCutted)
	//assert
	assert.Equal(t, expected, result)
}
