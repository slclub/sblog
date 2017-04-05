package core

type Result struct {
	Status  int
	Data    interface{}
	Message string
}

func NewResult(ret interface{}, err error) *Result {
	status := 0
	errm := ""
	if err != nil {
		status = -1
		errm = err.Error()
	}
	return &Result{Status: status, Data: ret, Message: errm}
}
