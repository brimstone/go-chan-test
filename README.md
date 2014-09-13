# go-chan-test

This is a program for me to learn how to write a package with channels.

## Usage

```go
package main

import (
	"github.com/brimstone/go-chan-test/channel"
	"github.com/brimstone/go-chan-test/dir"
	"log"
	"time"
)

func main() {
	cf := make(chan channel.File)
	dir1, err := dir.New("/tmp/dir1")
	if err != nil {
		log.Fatal("Can't open /tmp/dir1?")
	}
	dir2, err := dir.New("/tmp/dir2")
	if err != nil {
		log.Fatal("Can't open /tmp/dir1?")
	}

	go dir1.Sync(cf)
	go dir2.Sync(cf)

	log.Println("Setup finish")
	time.Sleep(time.Hour)
}
```