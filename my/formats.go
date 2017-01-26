package my

import (
	"bufio"
	"bytes"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/textproto"
	"strings"
)

//win表示可能なパイプを作成する ReadLine()で行を取得できる
func SjisReader(stdout io.ReadCloser) *textproto.Reader {
	rio := transform.NewReader(stdout, japanese.ShiftJIS.NewDecoder())
	return textproto.NewReader(bufio.NewReader(rio))
}

//stringのsjis→utf8変換
func FromSjis(str string) (string, error) {
	rio := transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
	b, err := ioutil.ReadAll(rio)
	return string(b), err
}

//自動でエンコードsjis or utf-8
func AutoEnc(body string) (string, error) {
	var f []byte
	encodings := []string{"sjis", "utf-8"}
	for _, enc := range encodings {
		if enc != "" {
			ee, _ := charset.Lookup(enc)
			if ee == nil {
				continue
			}
			var buf bytes.Buffer
			ic := transform.NewWriter(&buf, ee.NewDecoder())
			_, err := ic.Write([]byte(body))
			if err != nil {
				continue
			}
			err = ic.Close()
			if err != nil {
				continue
			}
			f = buf.Bytes()
			break
		}
	}
	return string(f), nil
}
