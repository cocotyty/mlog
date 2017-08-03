package main

import (
	"bytes"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filedata, _ := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/cocotyty/mlog/log.go")
	lines := bytes.Split(filedata, []byte("\n"))
	buffer := bytes.NewBuffer(nil)
	engine := otto.New()
	currentScript := ""
	start := -1
	codeLines := false
	outputLines := bytes.NewBuffer(nil)
	for i, line := range lines {
		trimedLine := bytes.TrimSpace(line)
		if bytes.HasPrefix(trimedLine, []byte("//#code")) {
			outputLines.Write(line)
			outputLines.WriteByte('\n')
			if codeLines {
				codeLines = false
				continue
			} else {
				codeLines = true
				outputLines.Write(buffer.Bytes())
				outputLines.WriteByte('\n')
				continue
			}
		}
		if codeLines {
			continue
		}

		if bytes.HasPrefix(trimedLine, []byte("//#tpl")) {
			if start == -1 {
				currentScript = string(trimedLine[6:])
				start = i
			} else {
				tpl := bytes.Join(lines[start+1:i], []byte{'\n'})
				engine.Set("tpl", string(tpl))
				log.Println(currentScript)
				v, err := engine.Eval(currentScript)
				if err != nil {
					log.Println(err)
					continue
				}
				buffer.WriteString(v.String())
				start = -1
			}
		}
		outputLines.Write(line)
		outputLines.WriteByte('\n')
	}
	ioutil.WriteFile(os.Getenv("GOPATH")+"/src/github.com/cocotyty/mlog/log.go", outputLines.Bytes(), 0755)
}
