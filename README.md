# decodeWechat
## 微信小程序对称解密go版本

安装
```
$ go get github.com/jiyarong/decodeWechat
```

使用
```go
package main

import "github.com/jiyarong/decodeWechat"
import "fmt"

func main () {
	decoder := decodeWechat.DecodeWechatEncrypted{
		AppId: "wx165574a98cb7f957",
		SessionKey: "L1gj7dCykbGngoko/W/uIg==",
	}
	iv := "/WOuqBXBCtSaiImUMn3iwg=="
	encryptedData := "cvGAtGAWc6g+AKP9UcV2i4ygaTl8j+1P0lhHnv4DQD8HvnSmW965o59aDPihXVpDHHvQL6PKcy5sqeyk3WwcHOEtfCK9iIn5pft1aSbYZTV5DtVNPUXpCQybp03idnw+yDQg9Obilz8odiVvhLYS58xW//bgjUZnjGYpmcoe4nMsKKnAcvMLLi7aN+ba720L92/sDB8cCiAQ8HBsW0v0crTjHoa4jLDurTz+6KPg1xXzFmUR5i2alFh49TYEW0Gfxibj9qXL08PuW78nQyVfqyfq/OxTSd5d33qDQa3ZA7GxTPhQptGK37+lE9WnNU+OYgpKndzQGhx+syI5lqF8OVMnPji8Iudk3phC2YjVG5EhLpS9uk8We0UTk3d9TVsRq0c3dHnHWVnMxZaT5oPdfQocsZ3smGGNNMEzzneLCa4eOlr/68d0k732Adze/6E35kcYT0MX2dpe7FXgslnYUKad/9Cm/f60V9bf1jVfRvs="

	result := decoder.DecryptData(encryptedData, iv)
	fmt.Printf(result)
}

```
