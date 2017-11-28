package collections

type List interface {
	Enumerable
	Add(item interface{})
	Remove(item interface{})
	AsEnumerable() Enumerable
}

type DefaultList struct {
	*DefaultEnumerable
}

func NewList() List {
	source := make([]interface{}, 0)
	return NewListFromSource(source)
}

func NewListFromSource(source []interface{}) List {
	return &DefaultList{
		DefaultEnumerable: NewEnumerableFromSource(source).(*DefaultEnumerable),
	}
}

func (list *DefaultList) Add(item interface{}) {
	list.source = append(list.source, item)

}

func (list *DefaultList) Remove(item interface{}) {
	for i, v := range list.source {
		if (IsObject(v) && IsObject(item) && AsObject(v).Equal(AsObject(item))) || v == item {
			list.source = append(list.source[:i], list.source[i+1:])
			goto ForEnd
		}
	}
ForEnd:
}

func (list *DefaultList) AsEnumerable() Enumerable {
	return list.DefaultEnumerable
}
