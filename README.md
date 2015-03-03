# mail
Email related module

# How to use

get this library with `go get`

```
$ go get github.com/mokelab-go/mail
```

Create client object and call `Send` method

```
client := mailgun.New(apiKey, domain, from)
err := client.Send(to, subject, body)
```

# example

```
package main

import (
	"fmt"
	"github.com/mokelab-go/mail/mailgun"
)

const (
	apiKey  = "<Input your API Key>"
	domain  = "<Input your domain in mailgun>"
	from    = "<Input from address>"
	to      = "<Input to address>"
	subject = "メールガンから送信"
	body    = "テストコードから本文"
)

func main() {
	client := mailgun.New(apiKey, domain, from)
	err := client.Send(to, subject, body)
	if err != nil {
		fmt.Printf("Unexepcted Error : %s", err)
	}
}
```

