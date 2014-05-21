package mxgo


type Filter interface {
	Execute(ctrl *Controller)
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

func (fm *FilterManager)BeforeAction(ctrl *Controller){
	for _,filter:= range fm.filters{
		filter.Execute(ctrl)
	}
}

func (fm *FilterManager)AfterAction(ctrl *Controller){
	for _,filter:= range fm.filters{
		filter.Execute(ctrl)
	}
}
