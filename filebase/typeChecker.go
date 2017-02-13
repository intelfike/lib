package filebase

func (f Filebase) IsString() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	_, ok := (*i).(string)
	return ok
}
func (f Filebase) IsBool() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	_, ok := (*i).(bool)
	return ok
}
func (f Filebase) IsInt() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	_, ok := (*i).(int)
	return ok
}
func (f Filebase) IsUint() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	_, ok := (*i).(uint)
	return ok
}
func (f Filebase) IsFloat() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	_, ok := (*i).(float64)
	return ok
}
func (f Filebase) IsNull() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	return *i == nil
}
func (f Filebase) Exists() bool {
	_, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	return true
}
func (f Filebase) IsArray() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	_, ok := (*i).([]interface{})
	return ok
}
func (f Filebase) IsMap() bool {
	i, err := f.GetInterfacePt()
	if err != nil {
		return false
	}
	_, ok := (*i).(map[string]interface{})
	return ok
}
