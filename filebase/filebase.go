// Gopherのためのjson操作パッケージ
package filebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Filebase struct {
	master interface{}
	path   []interface{}
}

// file name to *Filebase.
func NewFile(name string) (*Filebase, error) {
	file, _ := os.Open("./data.txt")
	b := new(bytes.Buffer)
	io.Copy(b, file)
	return New(b.Bytes())
}

// Byte data to *Filebase
func New(b []byte) (*Filebase, error) {
	fb := new(Filebase)
	err := json.Unmarshal(b, &fb.master)
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
	cur := &f.master
	if len(f.path) == 0 {
		*cur = i
		return nil
	}
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
	v, _ := f.GetInterface()
	f.Set(*v)
}

// Push() =? Filebase to Filebase.
func (f *Filebase) Fpush(fb *Filebase) {
	v, _ := f.GetInterface()
	f.Push(*v)
}

// If you want to do type switch then use this.
// Do not use it much.
//
// You can do type switch with regexp too.
// Refer to String().
//
// This function get interface{} pinter.
func (f Filebase) GetInterface() (*interface{}, error) {
	cur := &f.master
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

// If json node is map then return key list & nil.
//
// else then return nil & error.
func (f Filebase) Keys() ([]string, error) {
	v, _ := f.GetInterface()
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

// If json node is array then return len(array) & nil.
//
// else then return -1 & error.
func (f Filebase) Len() (int, error) {
	v, _ := f.GetInterface()
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
