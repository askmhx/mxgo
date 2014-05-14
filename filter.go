package mxgo

type Filter interface {
	Execute(action Action)
}

const (
	FILTER_BEFORE_ACTION = iota
	FILTER_AFTER_ACTION
)

type FilterManager struct {
	filters []Filter
}

func NewFilterManager() *FilterManager{
	fm := &FilterManager{}
	fm.filters  = make([]Filter,0)
	return fm
}

func (fm *FilterManager)AddFilter(filters ...Filter){
	for i := range filters{
		append(fm.filters,filters[i])
	}
}

func (fm *FilterManager)BeforeAction(action Action){
	for i := range fm.filters{
		flt := fm.filters[i]
		flt.Execute(action)
	}
}

func (fm *FilterManager)AfterAction(action Action){
	for i := range fm.filters{
		flt := fm.filters[i]
		flt.Execute(action)
	}
}
