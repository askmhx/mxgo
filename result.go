package mxgo

type Result interface {
	render() bool
}

type XMLResult struct {

}

func (xml *XMLResult) render() {

}


type JSONResult struct {

}

func (json *JSONResult) render(){
}

type TemplateResult struct {

}

func (template *TemplateResult) render(){
}

type PlainResult struct {

}


func (plain *PlainResult) render(){
}
