package model

type Modeli interface {
	GetSource(...string) string

	GetAttr() []string

	Save(Modeli)

	Create(Modeli) (int, error)

	Delete(string, ...interface{})

	Update(Modeli, ...interface{})

	Find(...interface{})

	ID(...int) int

	GetObjectValues(Modeli) []interface{}
	GetObject() map[string]interface{}
}
