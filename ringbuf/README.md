## 使い方
New()で作成する。引数はスライス。<br>
Write()で書き込む。<br>
Get()で取り出す。<br>
IndexOld(), IndexNew()は引数に0を指定するとそれぞれ最も古いデータ、新しいデータを取得できる。
Clear()でデータをすべて削除(見かけ上)
## 特徴
Write()でいくら書き込んでも、スライスのサイズを変更しない<br>
Write()回数がCap()より小さい　なら返ってくるスライスのサイズはWrite()の回数分<br>
Write()回数がCap()より大きい　なら返ってくるスライスのサイズはCap()分<br>
Write()によってCapを超えた場合、古いデータが自動的に上書きされる<br>

## コード例

```
rb, _ := ringbuf.New(make([]int, 2))

rb.Write(333)
rb.Write(22)
rb.Write(1)

fmt.Println(rb) // [22 1]
```
