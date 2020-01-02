package ext

import "os"

func HostName() string {
	if host, err := os.Hostname(); err == nil {
		return host
	} else if host, ok := os.LookupEnv("HOSTNAME"); ok {
		return host
	} else {
		return "unknown"
	}
}

func HostIPStr() string {
	if ip, err := HostIP(); err != nil {
		return "unknown"
	} else {
		return ip.String()
	}
}

func Version() string {
	if version, ok := os.LookupEnv("APP_VERSION"); ok {
		return version
	} else {
		return "dev"
	}
}
