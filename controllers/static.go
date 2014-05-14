package mxgo

type StaticController struct {
	*Controller
}

func (stc *StaticController)handle() Result{
	return nil
}
