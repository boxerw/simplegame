package foundation

type Entity interface {
	Destroy()
	GetDestroyed() bool
	GetEnvironment() Environment
	GetUID() uint64
}

type EntityWhole interface {
	Entity
	GetInheritor() Entity
}
