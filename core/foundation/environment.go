package foundation

type Environment interface {
	Data
	MakeUID() uint64
	GetEntity(entUID uint64) (Entity, bool)
}

type EnvironmentWhole interface {
	Environment
	AddEntity(entity Entity)
	RemoveEntity(entUID uint64)
}
