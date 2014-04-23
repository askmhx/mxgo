package router

type Router struct {

}

func NewRouter() *Router{
	route := &Router{}
	return route
}

func (route *Router)ParserConfig(configPath string) error{
	return nil
}

func (route *Router)FindAction(uri string) func(){
	return nil
}

