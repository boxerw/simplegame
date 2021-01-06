package foundation

type Data interface {
	SetValue(name string, value interface{})
	GetValue(name string) interface{}
}
