## Getting go-yandexdisk-client

```
go get github.com/rusystem/go-yandexdisk-client
```

## This project contains the following methods:

- GetDiskInfo() - get info of disk.
- GetFiles() - get information about files on disk.
- UploadFile() - upload file to disk.
- DeleteResource() - delete file or directory from disk.

## Examples:

### GetDiskInfo():

```go
package main

import (
	"context"
	"fmt"
	yandexdiskapi "github.com/rusystem/go-yandexdisk-client"
	"log"
	"time"
)

const (
	apiToken = "enterYourTokenFromYandexDisk"
	timeout  = time.Second * 5
)

func main() {
	c, err := yandexdiskapi.NewClient(apiToken, timeout)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	data, err := c.GetDiskInfo(ctx)
	if err != nil {
		return
	}
	fmt.Println(data.Info())
}
```

###

### GetFiles():

```go
package main

import (
	"context"
	"fmt"
	yandexdiskapi "github.com/rusystem/go-yandexdisk-client"
	"log"
	"time"
)

const (
	apiToken = "enterYourTokenFromYandexDisk"
	timeout  = time.Second * 5
)

func main() {
	c, err := yandexdiskapi.NewClient(apiToken, timeout)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	files, err := c.GetFiles(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Info())
	}
}
```

###

### UploadFile():

```go
package main

import (
	"context"
	yandexdiskapi "github.com/rusystem/go-yandexdisk-client"
	"log"
	"time"
)

const (
	apiToken = "enterYourTokenFromYandexDisk"
	timeout  = time.Second * 5
)

func main() {
	c, err := yandexdiskapi.NewClient(apiToken, timeout)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	filename := "goMem.png"
	url := "https://miro.medium.com/max/1200/1*_c9Ob50Cdmh09wLSzHRKgQ.png"

	if err := c.UploadFile(ctx, filename, url); err != nil {
		log.Fatal(err)
	}
}
```

###

### DeleteResource():

```go
package main

import (
	"context"
	yandexdiskapi "github.com/rusystem/go-yandexdisk-client"
	"log"
	"time"
)

const (
	apiToken = "enterYourTokenFromYandexDisk"
	timeout  = time.Second * 5
)

func main() {
	c, err := yandexdiskapi.NewClient(apiToken, timeout)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	filename := "goMem.png"
	if err := c.DeleteResource(ctx, filename); err != nil {
		log.Fatal(err)
	}
}
```