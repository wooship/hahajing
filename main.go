package main

import (
	"fmt"
	"hahajing/com"
	"hahajing/door"
	"hahajing/kad"
	"hahajing/web"
	"os/exec"
	"runtime"
)

var kadInstance kad.Kad
var webInstance web.Web
var doorInstance door.Door
var keywordManager = com.NewKeywordManager()

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func main() {
	kadInstance.Start()
	doorInstance.Start(keywordManager)

	webInstance.Start(kadInstance.SearchReqCh, doorInstance.KeywordCheckReqCh, keywordManager)

	Open("http://localhost:2021/")
}
