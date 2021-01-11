package shell

type Serializable interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}
