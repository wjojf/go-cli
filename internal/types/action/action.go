package action

var (
	DefaultActions = List{
		&DockerMonitor{},
		&DockerRestart{},
		&SSLRenew{},
	}
)

type Action interface {
	ID() string
	Name() string
	Description() string
}

type List []Action
