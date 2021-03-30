package socketio

import (
	"github.com/loadimpact/k6/js/modules"
)

type SocketIO struct {
	Version string
}

const version = "v0.0.1"

func init() {
	modules.Register("k6/x/socket-io", &SocketIO{
		Version: version,
	})
}
