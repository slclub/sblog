package dispatcher

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//Param type
type Context gin.Context

//Inject function type
type InjectFn func(c *Context)
type LoopFuncs []*box

//Dispatcher
type dispatcher struct {
	context       *Context
	loop_func     map[int8]LoopFuncs
	executedNames []string
}

type box struct {
	fn     InjectFn
	name   string
	status bool
}

var print = fmt.Println

func New() *dispatcher {
	return &dispatcher{&Context{}, map[int8]LoopFuncs{}, []string{}}
}

//Can be rewritten from the outside world
//Example: dispatcher.HandleFunc = func(c *dispatcher.Context){}
var HandleFunc InjectFn

func (di *dispatcher) D(c *Context) *dispatcher {
	di.context = c
	return di
}

//Injection
func (di *dispatcher) Di(fn InjectFn) InjectFn {
	//di.Bind(fn, ROUTE_FUNC)
	return di.Handle(fn)
}

//Organize the injected functions in order
func (di *dispatcher) Handle(fn InjectFn) InjectFn {
	HandleFunc = func(c *Context) {
		c = di.context
		//Begin
		di.Deal(c, BEGIN)
		di.Deal(c, ROUTE_FUNC)
		fn(c)
		di.Deal(c, END)
	}
	return HandleFunc
}

//Bind function Binding function to the corresponding location
func (di *dispatcher) Bind(name string, fn InjectFn, pos int8) {
	Location := []int8{BEGIN, ROUTE_FUNC, END}
	for _, l := range Location {
		if l&pos <= 0 {
			continue
		}
		di.loop_func[BEGIN].Push(name, fn)
	}
}

func (di *dispatcher) Deal(c *Context, pos int8) {
	for _, item := range di.loop_func[BEGIN] {
		if item.status == false {
		}
		item.fn(c)
	}
}

//Push function to the inject queue.
func (lf LoopFuncs) Push(name string, fn InjectFn) {
	newBox := &box{fn, name, true}
	lf = append(lf, newBox)
}
