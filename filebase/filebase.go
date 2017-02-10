// Gopherのためのjson操作パッケージ
package filebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Filebase struct {
	masterBytes []byte
	master      *interface{}
	path        []interface{}
}

var _ fmt.Stringer = Filebase{}

// file name to *Filebase.
func NewByFile(name string) (*Filebase, error) {
	file, _ := os.Open(name)
	return NewByReader(file)
}

func NewByReader(reader io.Reader) (*Filebase, error) {
	b := new(bytes.Buffer)
	io.Copy(b, reader)
	return New(b.Bytes())
}

// Byte data to *Filebase
func New(b []byte) (*Filebase, error) {
	fb := new(Filebase)
	fb.masterBytes = b
	fb.master = new(interface{})
	err := json.Unmarshal(b, fb.master)
	if err != nil {
		return nil, err
	}
	return fb, nil
}

// if has child then
// update child.
//
// else
// make child.
//
// Can't make array.
func (f *Filebase) Set(i interface{}) error {
	if len(f.path) == 0 {
		*f.master = i
		return nil
	}
	cur := new(interface{})
	*cur = *f.master
	prev := *cur
	prevkey := ""
	_ = prevkey
	for n, pathv := range f.path {
		switch pt := pathv.(type) {
		case string:
			switch mas := (*cur).(type) {
			case map[string]interface{}:
				var ok bool
				prev = *cur
				*cur, ok = mas[pt]
				if !ok {
					mas[pt] = nil //map[string]interface{}{pt: i}
				}
				if n == len(f.path)-1 {
					mas[pt] = i
				}
			default:
				paths := []string{prevkey}
				for _, v := range f.path[n:] {
					paths = append(paths, v.(string))
				}
				mapNest(prev.(map[string]interface{}), i, 0, paths...)
				return nil
			}
			prevkey = pt
		case int:
			mas, ok := (*cur).([]interface{})
			if !ok {
				continue
			}
			if n != len(f.path)-1 {
				*cur = mas[pt]
				continue
			}
			// 最後の要素なら
			// fmt.Println(mas, pt, i)
			mas[pt] = i
		default:
			panic("Chilc(...interface{} => string or int)")
		}
	}
	return nil
}
func mapNest(m map[string]interface{}, val interface{}, depth int, s ...string) {
	if depth == len(s)-1 {
		m[s[depth]] = val
		return
	}
	mm := map[string]interface{}{s[depth+1]: nil}
	m[s[depth]] = mm
	mapNest(mm, val, depth+1, s...)
}

// It like append().
//
// If json node is array then add i.
//
// else then set []interface{i} (initialize for array).
func (f *Filebase) Push(i interface{}) {
	_, err := f.Len()
	if err != nil {
		f.Set([]interface{}{i})
		return
	}
	p, _ := f.GetInterface()
	ar := (*p).([]interface{})
	f.Set(append(ar, i))
}

// Set() => Filebase to Filebase.
func (f *Filebase) Fset(fb *Filebase) {
	v, _ := fb.GetInterface()
	f.Set(*v)
}

// Push() => Filebase to Filebase.
func (f *Filebase) Fpush(fb *Filebase) {
	v, _ := fb.GetInterface()
	f.Push(*v)
}

// Remove() remove map or array element
func (f *Filebase) Remove() {
	if len(f.path) == 0 {
		panic("can't remove root.")
	}
	path := f.path[len(f.path)-1]
	i, _ := f.Parent().GetInterface()
	switch t := path.(type) {
	case string:
		delete((*i).(map[string]interface{}), t)
	case int:
		arr := (*i).([]interface{})
		f.Parent().Set(append(arr[:t], arr[t+1:]...))
	default:
		panic("Child()")
	}
}

// If you want to do type switch then use this.
// Do not use it much.
//
// You can do type switch with regexp too.
// Refer to String().
//
// This function get interface{} pinter.
func (f Filebase) GetInterface() (*interface{}, error) {
	cur := new(interface{})
	*cur = *f.master
	for _, pathv := range f.path {
		switch pt := pathv.(type) {
		case string:
			mas, ok := (*cur).(map[string]interface{})
			if !ok {
				continue
			}
			*cur, ok = mas[pt]
			if !ok {
				return nil, errors.New("Data is not found")
			}
		case int:
			mas, ok := (*cur).([]interface{})
			if !ok {
				continue
			}
			*cur = mas[pt]
		default:
			panic("Child(...interface{}.(type) == string or int)")
		}
	}
	return cur, nil
}

// loop map or array
func (f Filebase) Each(fn func(*Filebase)) {
	length, err := f.Len()
	if err == nil {
		for n := 0; n < length; n++ {
			fn(f.Child(n))
		}
	}
	keys, err := f.Keys()
	if err == nil {
		for _, key := range keys {
			fn(f.Child(key))
		}
	}
}

// Get json root.
func (f Filebase) Root() *Filebase {
	f.path = make([]interface{}, 0)
	return &f
}

// Child(...interface{}.(type) == string or int)
func (f Filebase) Child(path ...interface{}) *Filebase {
	f.path = append(f.path, path...)
	return &f
}

// Get json parent.
func (f Filebase) Parent() *Filebase {
	if len(f.path) == 0 {
		panic("root has not parent.")
		return nil
	}
	f.path = f.path[:len(f.path)-1]
	return &f
}

// f location become to new json root
func (f Filebase) Clone() (*Filebase, error) {
	p, err := f.GetInterface()
	if err != nil {
		return nil, err
	}
	newfb := new(Filebase)
	b, err := json.Marshal(*p)
	if err != nil {
		return nil, err
	}
	newfb.master = new(interface{})
	json.Unmarshal(b, newfb.master)
	return newfb, nil
}

// If json node is map then return key list & nil.
//
// else then return nil & error.
func (f Filebase) Keys() ([]string, error) {
	v, _ := f.GetInterface()
	if v == nil {
		return nil, errors.New("json node equal nil or not has.")
	}
	s := []string{}
	switch t := (*v).(type) {
	case map[string]interface{}:
		for key, _ := range t {
			s = append(s, key)
		}
		return s, nil
	}
	return nil, errors.New("KeyList() => json not map")
}

//This get len, check if array.
//
// If json node is array then return len(array) & nil.
//
// else then return -1 & error.
func (f Filebase) Len() (int, error) {
	v, _ := f.GetInterface()
	if v == nil {
		return -1, errors.New("json node equal nil or not has.")
	}
	switch t := (*v).(type) {
	case []interface{}:
		return len(t), nil
	}
	return -1, errors.New("Len() => json not array")
}

// [regexp(string value) => type]
//  ".*" => string
//  [1-9][0-9]* => int
//  [1-9][0-9]*.[0-9]*[1-9] => float
//  (true|false) => bool
//  null => null
//  nil => [NotHasChild]
// example
//  if fb.Child("class") == `"string"`{}
//  if fb.Child("id") == "12"{}
//  if fb.Child("id") == "true"{}
//  if fb.Child("id") == "null"{} // null value
//  if Child("id???") == "nil"{} // not has child!
func (f Filebase) String() string {
	v, err := f.GetInterface()
	if err != nil {
		return "nil"
	}
	b, _ := json.MarshalIndent(*v, "", "\t")
	return string(b)
}
