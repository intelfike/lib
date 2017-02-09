# update/refer to json like firebase(web).
This is golang package.<br>
firebase(web)っぽくjsonを加工・参照できるgolangのパッケージです。

## install
command

```go get github.com/intelfike/lib/filebase```

## usage

Disp json node if [class == A]<br>

```
jsonData := `
[
    {"id": 1,"name": "タカハシ","class": "A" },
    {"id": 2,"name": "スズキ","class": "A" },
    {"id": 3,"name": "タナカ","class": "B"},
    {"id": 4,"name": "イシバシ","class": "B"},
    {"id": 5,"name": "ナカヤマ","class": "B"} 
]
`
fb, _ := filebase.New([]byte(jsonData))
length, _ := fb.Len()
for n := 0; n < length; n++{
    if fb.Child(n, "class").String() == `"A"`{
        fmt.Println(fb.Child(n)) // ↓[output]↓
    }
}
```

output

```
{
        "class": "A",
        "id": 1,
        "name": "タカハシ"
}
{
        "class": "A",
        "id": 2,
        "name": "スズキ"
}
```

## type and func

```
type Filebase 
    func New(b []byte) (*Filebase, error)
    func NewByFile(name string) (*Filebase, error)
    func NewByReader(reader io.Reader) (*Filebase, error)
    func (f Filebase) Child(path ...interface{}) *Filebase
    func (f *Filebase) Fpush(fb *Filebase)
    func (f *Filebase) Fset(fb *Filebase)
    func (f Filebase) GetInterface() (*interface{}, error)
    func (f Filebase) Keys() ([]string, error)
    func (f Filebase) Len() (int, error)
    func (f *Filebase) Push(i interface{})
    func (f Filebase) Root() *Filebase
    func (f *Filebase) Set(i interface{}) error
    func (f Filebase) String() string
```

### Maker func

```
    func New(b []byte) (*Filebase, error)
    func NewByFile(name string) (*Filebase, error)
    func NewByReader(reader io.Reader) (*Filebase, error)
```

### Referer func

```
    func (f Filebase) Child(path ...interface{}) *Filebase
    func (f Filebase) Root() *Filebase
```

### Getter func

```
    func (f Filebase) GetInterface() (*interface{}, error)
    func (f Filebase) String() string
    func (f Filebase) Keys() ([]string, error)
    func (f Filebase) Len() (int, error)
```

### Setter func

```
    func (f *Filebase) Fpush(fb *Filebase)
    func (f *Filebase) Fset(fb *Filebase)
    func (f *Filebase) Push(i interface{})
    func (f *Filebase) Set(i interface{}) error
```

### TODO: func

```
    func (f *Filebase) Remove() error
    func (f *Filebase) Empty() error
```