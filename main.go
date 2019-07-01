package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"

	"github.com/guesslin/inject/types"
)

const (
	snapshotLength int32         = 65535
	promiscuous    bool          = false
	timeout        time.Duration = time.Microsecond * 10
)

var (
	options gopacket.SerializeOptions = gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
)

var (
	src      = flag.String("src", "", "")
	srcMac   = flag.String("srcMac", "", "")
	dst      = flag.String("dst", "", "")
	dstMac   = flag.String("dstMac", "", "")
	nic      = flag.String("nic", "eth0", "")
	filename = flag.String("filename", "", "")
)

func generatePacket(payload []byte) []byte {
	client, err := types.NewEndpoint(*src, *srcMac)
	exitOnError(err)
	server, err := types.NewEndpoint(*dst, *dstMac)
	exitOnError(err)

	ethernetLayer := &layers.Ethernet{
		SrcMAC:       client.Mac,
		DstMAC:       server.Mac,
		EthernetType: layers.EthernetTypeIPv4,
	}
	ipLayer := &layers.IPv4{
		Version:  4,
		SrcIP:    client.IP,
		DstIP:    server.IP,
		Protocol: layers.IPProtocolTCP,
	}
	customizedTCPLayer := &layers.TCP{
		SrcPort: layers.TCPPort(client.Port),
		DstPort: layers.TCPPort(server.Port),
		SYN:     true,
	}
	customizedTCPLayer.SetNetworkLayerForChecksum(ipLayer)
	// And create the packet with the layers
	buffer := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buffer, options,
		ethernetLayer,
		ipLayer,
		customizedTCPLayer,
		gopacket.Payload(payload),
	)

	exitOnError(err)
	return buffer.Bytes()
}

func main() {
	flag.Parse()
	handle, err := pcap.OpenLive(*nic, snapshotLength, promiscuous, timeout)
	exitOnError(err)
	defer handle.Close()
	if *filename == "" {
		err = handle.WritePacketData(generatePacket(nil))
	} else {
		err = handle.WritePacketData(generatePacket(loadPayloadFrom(*filename)))
	}
	if err != nil {
		log.Println(err)
	}
}

func exitOnError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}

func loadPayloadFrom(fname string) []byte {
	f, err := os.Open(fname)
	exitOnError(err)
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	exitOnError(err)
	return content
}

func loadPcap(fname string) []gopacket.Packet {
	hpcap, err := pcap.OpenOffline(fname)
	exitOnError(err)
	packets := make([]gopacket.Packet, 0)
	psource := gopacket.NewPacketSource(hpcap, hpcap.LinkType())
	for packet := range psource.Packets() {
		packets = append(packets, packet)
	}
	return packets
}
