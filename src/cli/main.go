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

type CLI struct {
	TUNAddr string `help:"address of the served tun addr" default:"192.168.11.1/32"`
	TCPPort string `help:"adress of the tcp port to dial" default:"default:4663"`
}

func main() {

	var cli CLI

	_ = kong.Parse(
		&cli,
		kong.Name("tcp-to-tun 4test"),
	)
	logfd, err := os.Create("../log/cli_log")
	if err != nil {
		println("failed create cli logfile")
		return
	}
	log := log.New(logfd, log.Logfmt, false)

	conn, err := net.Dial("tcp", cli.TCPPort)
	if err != nil {
		log.Error("cli dial tcp:", "error", err)
		return
	}
	defer conn.Close()
	log.Info("cli tcp conn created")

	tun, err := initTUN(cli.TUNAddr, 1500)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("cli TUN Created")
	defer tun.Close()
	log.Info("cli TUN (192.168.11.1/32) to TCP (default:4663) serving")
	// go func() {
	// 	msg := "message"
	// 	for {
	// 		time.Sleep(1 * time.Second0
	// 		conn.Write([]byte(msg))
	// 		log.Info("msg message sent")
	// 	}
	// }()

	// initializeMeterProvider()

	go ForwardTCPToTUN(log, tun, conn)
	forwardTUNToTCP(log, tun, conn)

}

func initTUN(ip string, mtu int) (*tun.TUN, error) {

	tun, err := tun.NewTUN()
	if err != nil {
		return nil, fmt.Errorf("cli creating tun failed", err)
	}
	if err = tun.Setup(netip.MustParsePrefix(ip), mtu); err != nil {
		tun.Close()
		return nil, fmt.Errorf("cli setup tun failed", err)
	}
	// if err = tun.AddRoute(netip.MustParsePrefix("192.168.10.1/32")); err != nil {
	// 	tun.Close()
	// 	return nil, fmt.Errorf("cli adding route to tun failed", err)
	// }
	return tun, nil
}

func forwardTUNToTCP(log *slog.Logger, tun *tun.TUN, conn net.Conn) {

	// var meter = otel.Meter("benchmark_test/")
	// msgForwarded, err := meter.Int64Counter(
	// 	"msg",
	// 	// metric.WihDescription("Msg forwarded from tun to tcp"),
	// 	// metric.WithUnit("{msg}"),
	// )
	// if err != nil {
	// 	return
	// }
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
		// msgForwarded.Add(context.Background(), 1)
		log.Info("msg from tun to tcp:", "data", string(buf[:n]))
	}
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

// func initializeMeterProvider() bool {

// 	// Create resource.
// 	res, err := newResource()
// 	if err != nil {
// 		return false
// 	}

// 	// Create a meter provider.
// 	// You can pass this instance directly to your instrumented code if it
// 	// accepts a MeterProvider instance.
// 	meterProvider, err := newMeterProvider(res)
// 	if err != nil {
// 		return false
// 	}

// 	// Handle shutdown properly so nothing leaks.
// 	defer func() {
// 		if err := meterProvider.Shutdown(context.Background()); err != nil {
// 			log.Default.Error(err.Error())
// 		}
// 	}()

// 	// Register as global meter provider so that it can be used via otel.Meter
// 	// and accessed using otel.GetMeterProvider.
// 	// Most instrumentation libraries use the global meter provider as default.
// 	// If the global meter provider is not set then a no-op implementation
// 	// is used, which fails to generate data.
// 	otel.SetMeterProvider(meterProvider)
// 	return true
// }

// func newResource() (*resource.Resource, error) {
// 	return resource.Merge(resource.Default(),
// 		resource.NewWithAttributes(semconv.SchemaURL,
// 			semconv.ServiceName("my-service"),
// 			semconv.ServiceVersion("0.1.0"),
// 		))
// }

// func newMeterProvider(res *resource.Resource) (*metric.MeterProvider, error) {
// 	metricExporter, err := stdoutmetric.New()
// 	if err != nil {
// 		return nil, err
// 	}

// 	meterProvider := metric.NewMeterProvider(
// 		metric.WithResource(res),
// 		metric.WithReader(metric.NewPeriodicReader(metricExporter,
// 			// Default is 1m. Set to 3s for demonstrative purposes.
// 			metric.WithInterval(3*time.Second))),
// 	)
// 	return meterProvider, nil
// }
