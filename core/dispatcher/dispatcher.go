package dispatcher

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoopFuncs []*box

//Dispatcher
type dispatcher struct {
	loop_func   map[int8]LoopFuncs
	GlobalNames []string
	permission  *PermissionBox
}

//Inject function struct
type box struct {
	fn     gin.HandlerFunc
	name   string
	status bool
}

//Injection function execution authority structure
//Can be executed according to the specific route
type PermissionBox struct {
	Yes map[string][]string
	No  map[string][]string
}

var print = fmt.Println

var front *dispatcher
var admin *dispatcher

func Front() *dispatcher {
	if front == nil {
		front = New()
	}
	return front
}
func Admin() *dispatcher {
	if admin == nil {
		admin = New()
	}
	return admin
}

func New() *dispatcher {
	var permissionMap = make(map[string][]string)
	return &dispatcher{map[int8]LoopFuncs{}, []string{}, &PermissionBox{permissionMap, permissionMap}}
}

//Can be rewritten from the outside world
//Example: dispatcher.HandleFunc = func(c *dispatcher.gin.Context){}
var HandleFunc gin.HandlerFunc

func (di *dispatcher) D(c *gin.Context) *dispatcher {
	return di
}

//Injection
func (di *dispatcher) Di(fn gin.HandlerFunc) gin.HandlerFunc {
	//di.Bind(fn, ROUTE_FUNC)
	var ret = di.Handle(fn)
	return ret
}

func Fine(c *gin.Context) {
}

//Organize the injected functions in order
func (di *dispatcher) Handle(fn gin.HandlerFunc) gin.HandlerFunc {
	HandleFunc = func(c *gin.Context) {
		//Begin
		di.Deal(c, BEGIN)
		di.Deal(c, ROUTE_FUNC)
		if !c.IsAborted() {
			fn(c)
		}
		di.Deal(c, END)
	}

	Recovery()
	return HandleFunc
}

//Bind function Binding function to the corresponding location
func (di *dispatcher) Bind(name string, fn gin.HandlerFunc, pos int8) *dispatcher {
	Location := []int8{BEGIN, ROUTE_FUNC, END}
	for _, l := range Location {
		if l&pos <= 0 {
			continue
		}
		di.loop_func[l] = append(di.loop_func[l], newBox(name, fn))
	}

	return di
}

//Dispatcher factory.
func (di *dispatcher) Deal(c *gin.Context, pos int8) {
	for _, item := range di.loop_func[pos] {
		if item.status == false {
			continue
		}

		//if inSlice(item.name, di.permission.Yes) == false {
		//	continue
		//}

		if inSlice(item.name, di.permission.No[c.Request.URL.Path]) {
			continue
		}
		item.fn(c)
	}
}

//Set the permissions that can be executed by the injection function
//Service for a single route path·
func (di *dispatcher) Permission(path string, no []string) *dispatcher {
	//di.permission.Yes = yes
	di.permission.No[path] = no
	return di
}

//Relatively high operating authority.
//Can Global permissions can be prevented from running
//Service for a single route path·
func (di *dispatcher) NotAllow(path string, not []string) *dispatcher {
	di.permission.No[path] = not
	return di
}

//Service for all routes
func (di *dispatcher) GloblaAllowed(allowSlice []string) {
	di.GlobalNames = append(di.GlobalNames, allowSlice...)
}

func Recovery() {
	defer func() {
		fmt.Println("Dispatcher error:")
	}()
	gin.Recovery()
}

//The item param exist in the slice.
func inSlice(item string, s []string) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

//New box function to the inject queue.
func newBox(name string, fn gin.HandlerFunc) *box {
	return &box{fn, name, true}
}
