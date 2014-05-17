package mxgo


type Filter interface {
	Execute(action *Action)
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
	for _,filter := range filters{
		fm.filters = append(fm.filters,filter)
	}
}

func (fm *FilterManager)BeforeAction(action *Action){
	for _,filter:= range fm.filters{
		filter.Execute(action)
	}
}

func (fm *FilterManager)AfterAction(action *Action){
	for _,filter:= range fm.filters{
		filter.Execute(action)
	}
}
