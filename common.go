package collections

type Object interface {
	Equal(item Object) bool
	ToString() string
}

func IsObject(item interface{}) bool {
	_, ok := item.(Object)
	return ok
}

func AsObject(item interface{}) Object {
	v, ok := item.(Object)
	if ok {
		return v
	}
	return nil
}
