package filebase

// Get json root.
func (f Filebase) Root() *Filebase {
	f.path = make([]interface{}, 0)
	return &f
}

// Child(...interface{}.(type) == string or int)
func (f Filebase) Child(path ...interface{}) *Filebase {
	for _, v := range path {
		switch v.(type) {
		case string, int:
			f.path = append(f.path, v)
		default:
			panic("Child(...interface{}.(type) == string or int)")
		}
	}
	return &f
}

// Get json parent.
func (f Filebase) Parent() *Filebase {
	if len(f.path) == 0 {
		panic("root has not parent.")
	}
	f.path = f.path[:len(f.path)-1]
	return &f
}
