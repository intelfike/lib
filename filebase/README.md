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
Child(...interface{} => string or int) <br>
string => refer map (has not child => return nil/make child) <br>
int => refer array (overflow => panic()/panic()) <br>

### Getter func

```
    func (f Filebase) GetInterface() (*interface{}, error)
    func (f Filebase) String() string
    func (f Filebase) Keys() ([]string, error)
    func (f Filebase) Len() (int, error)
```

GetInterface() => If you want to do type switch then use this.<br>
But do not often use it for eliminate mistake because hard to use.<br>
<br>
String() => You can do type switch with regexp.<br>
[regexp(string value) => type] <br>
".*" => string <br>
[1-9][0-9]* => int <br>
[1-9][0-9]*.[0-9]*[1-9] => float <br>
(true|false) => bool <br>
null => null <br>
nil => [NotHasChild] <br>
<br>
Keys() => map keys (not map => Error!) <br>
Len() => array length (not array => Error!) <br>

### Setter func

```
    func (f *Filebase) Fpush(fb *Filebase)
    func (f *Filebase) Fset(fb *Filebase)
    func (f *Filebase) Push(i interface{})
    func (f *Filebase) Set(i interface{}) error
```
Set() => map appender & value setter<br>
Push() => array appender <br>

### TODO: func

```
    func (f *Filebase) Remove() error
    func (f *Filebase) Empty() error
```

## Licence
MIT(適当)