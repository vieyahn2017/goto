goµÄ±àÂë×ª»»

```go
"github.com/axgle/mahonia"
"strings"



decode := mahonia.NewDecoder("gbk")
if decode == nil {
	fmt.Println("(ErrorCode:0002) error is " + err.Error())
	return workas, errors.New("tmahonia.NewDecoder")
}
strings.NewReader(decode.ConvertString(string(buf)))
```



