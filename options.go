package go_netcat

type NetcatOptions func(nc *netcat)

// WithAddr option set listener address.
func WithAddr(v string) NetcatOptions {
	return func(nc *netcat) {
		nc.addr = v
	}
}

// WithLogger option set logger.
func WithLogger(o ncLogger) NetcatOptions {
	return func(nc *netcat) {
		nc.logger = o
	}
}

// WithDeadlineInSec option set client conn deadline in seconds.
func WithDeadlineInSec(v uint64) NetcatOptions {
	return func(nc *netcat) {
		nc.deadline = v
	}
}
