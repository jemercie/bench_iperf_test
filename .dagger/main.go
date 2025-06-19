// A generated module for TcpTunnel functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/tcp-tunnel/internal/dagger"
)

type TcpTunnel struct{}

// Returns a container that echoes whatever string argument is provided
func (m *TcpTunnel) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *TcpTunnel) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

// func (t *TcpTunnel) Build(tcp_to_tun_dir *dagger.Directory, tun_to_tcp_dir *dagger.Directory) (string, error) {
func (t *TcpTunnel) Build(cli_bin *dagger.File, srv_bin *dagger.File, cli_daemon_conf *dagger.File, srv_daemon_conf *dagger.File) (*dagger.Service, error) {

	// builder := t.GetContainer("tcp_to_tun", "/src/", tcp_to_tun_dir)
	ctx := context.Background()

	srv := dag.Container().
		From("debian").
		WithExposedPort(4663, dagger.ContainerWithExposedPortOpts{Protocol: "TCP"}).
		WithMountedFile("/bin/srv", srv_bin).
		WithMountedFile("etc/init.d/monitor_iperf3_srv", srv_daemon_conf).
		WithExec([]string{"apt", "update"}).
		WithExec([]string{"apt", "install", "-y", "iperf3"}).
		WithExec([]string{"apt", "install", "-y", "iproute2"}).
		WithExec([]string{"apt", "install", "-y", "openrc"}).
		WithExec([]string{"apt", "install", "-y", "sudo"}).
		WithExec([]string{"apt", "install", "-y", "golang"}).
		WithExec([]string{"chmod", "+x", "/bin/srv"}).
		WithExec([]string{"mkdir", "-p", "/dev/net"}).
		WithExec([]string{"touch", "/dev/net/tun"}).
		WithExec([]string{"chmod", "+x", "etc/init.d/monitor_iperf3_srv"}).
		WithExec([]string{"openrc", "default"}).
		// WithExec([]string{"mkdir", "-p", "dev/net/tun"}).
		// WithExec([]string{"chmod", "600", "dev/net/tun"}).
		WithExec([]string{"touch", "/run/openrc/softlevel"}). // to force bcs it's read only filesystem (maybe not a good idea)
		WithExec([]string{"rc-update", "add", "monitor_iperf3_srv", "default"}).
		WithExec([]string{"service", "monitor_iperf3_srv", "start"}).
		WithEntrypoint([]string{"sudo", "srv"}).
		AsService().
		WithHostname("srv")

	_, err := srv.Start(ctx)
	if err != nil {
		return nil, err
	}

	cli := dag.Container().
		From("debian").
		WithExposedPort(4663, dagger.ContainerWithExposedPortOpts{Protocol: "TCP"}).
		WithMountedFile("/bin/cli", cli_bin).
		WithMountedFile("etc/init.d/monitor_iperf3_cli", cli_daemon_conf).
		WithExec([]string{"apt", "update"}).
		WithExec([]string{"apt", "install", "-y", "iperf3"}).
		WithExec([]string{"apt", "install", "-y", "iproute2"}).
		WithExec([]string{"apt", "install", "-y", "openrc"}).
		WithExec([]string{"apt", "install", "-y", "sudo"}).
		WithExec([]string{"mkdir", "install", "-y", "sudo"}).
		WithExec([]string{"apt", "install", "mkdir", "golang"}).
		WithExec([]string{"chmod", "+x", "/bin/cli"}).
		WithExec([]string{"mkdir", "-p", "/dev/net"}).
		WithExec([]string{"touch", "/dev/net/tun"}).
		WithExec([]string{"chmod", "+x", "etc/init.d/monitor_iperf3_cli"}).
		WithExec([]string{"openrc", "default"}).
		// WithExec([]string{"mkdir", "-p", "dev/net/net"}).
		// WithExec([]string{"chmod", "600", "dev/net/net"}).
		WithExec([]string{"touch", "/run/openrc/softlevel"}). // to force bcs it's read only filesystem (maybe not a good idea)
		WithExec([]string{"rc-update", "add", "monitor_iperf3_cli", "default"}).
		WithExec([]string{"service", "monitor_iperf3_cli", "start"}).
		WithEntrypoint([]string{"sudo", "cli"}).
		AsService().
		WithHostname("cli")

	_, err = cli.Start(ctx)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

// func (t *TcpTunnel) GetContainer(binary_name string, workdir string, source *dagger.Directory) *dagger.Container {
// return dag.Container().
// From("alpine:latest").
// WithMountedDirectory("/src/", source).
// WithWorkdir(workdir).
// WithExec([]string{"apk", "add", "go"}).
// WithExec([]string{"go", "build", "-o", binary_name})
// WithExposedPort().
// }
