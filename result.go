package mxgo

import (
	"html/template"
	"github.com/menghx/mxgo/httplib"
	"encoding/json"
	"errors"
)

type Result interface {
	Render() error
}

type BaseResult struct {
	 Request *httplib.Request
	 Response *httplib.Response
	 Data interface {}
}

type XMLResult struct {
	*BaseResult
}

func (result XMLResult) Render() error{
	return nil
}


type JSONResult struct {
	*BaseResult
}

func (result JSONResult) Render() error{
	bytes,err :=json.Marshal(result.Data)
	if err!=nil {
		return err
	}
	result.Response.WriteText(string(bytes))
 	return nil
}

type TemplateResult struct {
	*BaseResult
	tplName string
}

func (result TemplateResult) Render() error{
	tpl := template.New(result.tplName)
	return tpl.Execute(result.Response.ResponseWriter,result.Data)
}

type PlainResult struct {
	*BaseResult
}

func (result PlainResult) Render() error{
	dt,found := result.Data.(string)
	if found {
		result.Response.WriteText(dt)
	}
	return errors.New("result data type error")
}


type RedirectResult struct {
	*BaseResult
}

func (result RedirectResult) Render() error{
	dt,found := result.Data.(string)
	if found {
		result.Response.WriteText(dt)
	}
	return errors.New("result data type error")
}

type ForwardResult struct {
	*BaseResult
}

func (result ForwardResult) Render() error{
	dt,found := result.Data.(string)
	if found {
		result.Response.WriteText(dt)
	}
	return errors.New("result data type error")
}

type StaticResult struct {
	*BaseResult
}

func (result StaticResult) Render() error{
	dt,found := result.Data.(string)
	if found {
		result.Response.WriteFile(dt)
	}
	return errors.New("result data type error")
}
