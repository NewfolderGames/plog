# pLog

**pLog**는 이쁜 로그를 출력하는 라이브러리 입니다.

### 설치

```shell
go get github.com/NewfolderGames/plog
```

### 사용방법

```go
package main

import "github.com/NewfolderGames/plog"

func main() {

	plog.Info("%s", "hello")
	plog.Debug("%s", "hello")
	plog.Warn("%s", "hello")
	plog.Error("%s", "hello")

}
```
