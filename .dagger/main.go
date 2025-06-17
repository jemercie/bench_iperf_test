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
func (t *TcpTunnel) Build(cli_bin *dagger.File, srv_bin *dagger.File) (string, error) {

	// builder := t.GetContainer("tcp_to_tun", "/src/", tcp_to_tun_dir)
	ctx := context.Background()
	return dag.Container().
		From("alpine").
		WithExec([]string{"apk", "add", "iperf3", "go"}).
		WithMountedFile("/src/cli", cli_bin).
		WithExec([]string{"chmod", "+x", "src/cli"}).
		WithExec([]string{"ls", "-la", "/src"}).
		WithExposedPort(4633).
		WithExec([]string{"./src/cli"}).
		WithExec([]string{"iperf3", "-c", "192.168.11.1", "-B", "192.168.10.1", "-t", "30"}).
		Stdout(ctx)
	// AsService()
	// Stdout(ctx)
	// Terminal()
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
