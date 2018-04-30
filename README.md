# ElDorito-Stats-Struct
---
This library aims to provide an up-to-date version of the data structure [ElDorito](https://github.com/ElDewrito/ElDorito) publishes at the end of a match.


## Installation

`go get github.com/Calenaur/ElDorito-Stats-Struct`

## Usage
```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Calenaur/Eldorito-Stats-Struct/data"
)

func main(){
	raw, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/Calenaur/Eldorito-Stats-Struct/example/grifball_game.json")

	if (err != nil) {
		panic(err.Error())
	}

	srv := data.Server{}
	json.Unmarshal(raw, &srv)

	fmt.Println(srv.ServerName)
}
```
