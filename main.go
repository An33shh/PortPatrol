package main

/*
#cgo LDFLAGS: -L. -libpacket_analysis
#include <stdlib.h>

extern int analyze_packet(const unsigned char* packet_ptr, unsigned int packet_len);
*/

import "C"
import (
	"fmt"
	"net"
	"sync"
	"time"
	"unsafe"
)

const (
	startPort = 1
	endPort   = 1024
	timeout   = 500 * time.Millisecond
)

func main() {
	var wg sync.WaitGroup
	ip := "10.0.0.16"

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			scanPort(ip, port)
		}(port)
	}

	wg.Wait()
	fmt.Println("Scanning complete")
}

func scanPort(ip string, port int) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Printf("Port %d is open\n", port)

	packet := make([]byte, 20) // Dummy packet for demonstration
	packet[12], packet[13], packet[14], packet[15] = 192, 168, 1, 1   // Source IP
	packet[16], packet[17], packet[18], packet[19] = 192, 168, 1, 100 // Destination IP

	// Pass packet data to Rust function
	cPacket := (*C.uchar)(unsafe.Pointer(&packet[0]))
	length := C.uint(len(packet))
	result := C.analyze_packet(cPacket, length)
	if result != 0 {
		fmt.Printf("Packet analysis failed for port %d\n", port)
	} else {
		fmt.Printf("Packet analysis succeeded for port %d\n", port)
	}
}
