package host

var (
	AllHosts = []Host{
		Host("site.legocy.online"),
	}
)

type Host string

func (h Host) IP() (string, bool) {
	ip, ok := ips[string(h)]
	return ip, ok
}
