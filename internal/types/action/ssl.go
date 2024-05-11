package action

const (
	SSLRenewID = "ssl-renew"
)

// SSLRenew is an action that renews SSL certificates.
type SSLRenew struct{}

func (s SSLRenew) ID() string {
	return SSLRenewID
}

func (s SSLRenew) Name() string {
	return "SSL Renew"
}

func (s SSLRenew) Description() string {
	return "Renew SSL certificates"
}
