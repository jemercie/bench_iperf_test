package main

import (
	"fmt"
	"log/slog"
	"net"
	"net/netip"
	"os"

	"github.com/alecthomas/kong"
	"loftorbital.com/log"
	"loftorbital.com/net/tun"
)

// cli listen port,send port

type CLI struct {
	TUNAddr string `help:"address of the served tun addr" default:"192.168.10.1/32"`
	TCPPort string `help:"number of the listened tcp address" default:"0.0.0.0:4663"`
}

func main() {
	var cli CLI
	_ = kong.Parse(
		&cli,
		kong.Name("tcp-to-tun 4test"),
	)
	logfd, err := os.Create("../log/srv_log")
	if err != nil {
		println("failed create srv logfile")
		return
	}
	log := log.New(logfd, log.Logfmt, false)

	tun, err := InitTUN(cli.TUNAddr, 1500)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("srv TUN Created")
	defer tun.Close()

	listener, conn, err := InitTCPServer(cli.TCPPort)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer listener.Close()
	defer conn.Close()
	log.Info("srv TUN (192.168.10.1/32) to TCP (0.0.0.0:4663) serving")

	go forwardTUNToTCP(log, tun, conn)
	ForwardTCPToTUN(log, tun, conn)
	// args: port to listen, port to open tun on
}

// InitTUN() open a tun on the given port
func InitTUN(ip string, mtu int) (*tun.TUN, error) {

	tun, err := tun.NewTUN()
	if err != nil {
		return nil, fmt.Errorf("srv creating tun failed", err)
	}

	if err = tun.Setup(netip.MustParsePrefix(ip), mtu); err != nil {
		return nil, fmt.Errorf("ssrv etup tun failed", err)
	}
	if err = tun.AddRoute(netip.MustParsePrefix("192.168.11.1/32")); err != nil {
		tun.Close()
		return nil, fmt.Errorf("srv add route tun failed", err)
	}
	return tun, nil
}

// InitTCPServer() open a listener and wait for an incoming connection to accept
// return both the listener and the connection
func InitTCPServer(port string) (net.Listener, net.Conn, error) {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, nil, fmt.Errorf("srv listener creation failed :%w", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		listener.Close()
		return nil, nil, fmt.Errorf("srv accept connection failed", "error", err)
	}

	return listener, conn, nil
}

// ForwardTCPToTUN()  read msg from tcp connection and forward them to the tun
func ForwardTCPToTUN(log *slog.Logger, tun *tun.TUN, conn net.Conn) {

	buf := make([]byte, 1500)
	log.Info("Waiting for messages from tcp conn to forward")
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Error(err.Error())
			return
		}
		n, err = tun.Write(buf[:n])
		if err != nil {
			log.Error(err.Error())
			return
		}
		// one counter for bytes sent
		// one counter for nb messages -> on each thing to compare!
		println("bytes written:", n)
		// packet, _ := tcpparser.ParseTCPPacket(buf[:n])
		// println("port:", packet.TCP.DstPort)
		// ip_hdr, _ := ipv4.ParseHeader(buf[:n])
		// println("src ip:", ip_hdr.Src, " dst ip:", ip_hdr.Dst)
		log.Info("msg from tcp to tun:", "data", string(buf[:n]))

	}
}

func forwardTUNToTCP(log *slog.Logger, tun *tun.TUN, conn net.Conn) {

	buf := make([]byte, 1500)
	log.Info("Waiting for messages from tun conn to forward")
	for {
		n, err := tun.Read(buf)
		if err != nil {
			log.Error(err.Error())
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Info("msg from tun to tcp:", "data", string(buf[:n]))
	}
}
