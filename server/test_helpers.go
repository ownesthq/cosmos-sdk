package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"

	"github.com/ownesthq/cosmos-sdk/client/flags"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

// Get a free address for a test tendermint server
// protocol is either tcp, http, etc
func FreeTCPAddr() (addr, port string, err error) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return "", "", err
	}

	closer := func() {
		err := l.Close()
		if err != nil {
			// TODO: Handle with #870
			panic(err)
		}
	}

	defer closer()

	portI := l.Addr().(*net.TCPAddr).Port
	port = fmt.Sprintf("%d", portI)
	addr = fmt.Sprintf("tcp://0.0.0.0:%s", port)
	return
}

// SetupViper creates a homedir to run inside,
// and returns a cleanup function to defer
func SetupViper(t *testing.T) func() {
	rootDir, err := ioutil.TempDir("", "mock-sdk-cmd")
	require.Nil(t, err)
	viper.Set(flags.FlagHome, rootDir)
	viper.Set(flags.FlagName, "moniker")
	return func() {
		err := os.RemoveAll(rootDir)
		if err != nil {
			// TODO: Handle with #870
			panic(err)
		}
	}
}

// DONTCOVER
