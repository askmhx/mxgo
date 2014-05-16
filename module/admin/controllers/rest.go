package controllers

import "github.com/menghx/mxgo"

type Rest struct {
	*mxgo.Controller
}

func (rest Rest) add() Result{
	return &JSONResult{}
}

func (rest Rest) delete() Result{
	return &JSONResult{}
}

func (rest Rest) edit() Result{
	return &JSONResult{}
}

func (rest Rest) query() Result{
	return &JSONResult{}
}
