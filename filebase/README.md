## usage

Disp json node [Class == A]<br>

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
fb, _ := filebase.NewByFile("data.json")
length, _ := fb.Len()
for n := 0; n < length; n++{
    if fb.Child(n, "class").String() == `"A"`{
        fmt.Println(fb.Child(n)) // ↓output↓
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

## ???
Q. How to change indent?<br>
A. Can't. <br>

type Filebase <br>
    func New(b []byte) (*Filebase, error)<br>
    func NewByFile(name string) (*Filebase, error)<br>
    func NewByReader(reader io.Reader) (*Filebase, error)<br>
    func (f Filebase) Child(path ...interface{}) *Filebase<br>
    func (f *Filebase) Fpush(fb *Filebase)<br>
    func (f *Filebase) Fset(fb *Filebase)<br>
    func (f Filebase) GetInterface() (*interface{}, error)<br>
    func (f Filebase) Keys() ([]string, error)<br>
    func (f Filebase) Len() (int, error)<br>
    func (f *Filebase) Push(i interface{})<br>
    func (f Filebase) Root() *Filebase<br>
    func (f *Filebase) Set(i interface{}) error<br>
    func (f Filebase) String() string<br>
    <br>
