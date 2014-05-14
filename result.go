package mxgo

import (
	"html/template"
	"com.github/menghx/mxgo/httplib"
	"encoding/json"
)

type Result interface {
	Render() error
}

type BaseResult struct {
	 response *httplib.Response
	 Data interface {}
}

type XMLResult struct {
	*BaseResult
}

func (result *XMLResult) Render() error{
	return nil
}


type JSONResult struct {
	*BaseResult
}

func (result *JSONResult) Render() error{
	bytes,err :=json.Marshal(result.Data)
	if err!=nil {
		return err
	}
	result.response.WriteText(string(bytes))
 	return nil
}

type TemplateResult struct {
	*BaseResult
	tplName string
}

func (result *TemplateResult) Render() error{
	tpl := template.New(result.tplName)
	return tpl.Execute(result.response.ResponseWriter,result.Data)
}

type PlainResult struct {
	*BaseResult
}

func (result *PlainResult) Render() error{
	result.response.WriteText(result.Data)
	return nil
}


type RedirectResult struct {
	*BaseResult
}

func (result *RedirectResult) Render() error{
	result.response.WriteText(result.Data)
	return nil
}

type ForwardResult struct {
	*BaseResult
}

func (result *ForwardResult) Render() error{
	result.response.WriteText(result.Data)
	return nil
}
