package data

type DataProvider interface {
	GetData() []interface{}
}
