# A Youdao Fanyi Library
## 有道翻译库

This library emplemented Eng<->Chn translation via Youdao online fanyi service.

##Install
```bash
go get github.com/meoow/youdaofanyi

# build command line tool
make
```

##Usage
```go
package main

import (
	"fmt"
	fy "github.com/meoow/youdaofanyi"
	"os"
	"strings"
)

func main() {
	// fy.Plain : print result in plain text
	// fy.HTML  : print result in HTML(e.g. invoked via GoldenDict)
	// If the input text contains non-ascii characters,
	// it will do Chn->Eng translation,
	// otherwise perform Eng->Chn translation.
	out, _ := fy.Fanyi(strings.Join(os.Args[1:], " "), fy.HTML)
	fmt.Print(out)
}
```

##Command Line
```
./fanyi good job
```
> 好工作