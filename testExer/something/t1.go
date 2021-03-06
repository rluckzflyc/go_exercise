package main

import (
	"fmt"
	"log"
	"os"
	// "path/filepath"
	"strings"
)

var lg *log.Logger

type paths struct {
	name    string
	dirPath bool
	exist   bool
}

func init() {
	lg = log.New(os.Stdout, "log", log.Lshortfile)
}

func (p *paths) initP(s string) error {
	p.name = s
	fi, fiErr := os.Lstat(s)
	if fiErr != nil {
		lg.Println(fiErr)
	}
	if fi == nil {
		p.exist = false
		if strings.HasSuffix(s, "/") {
			p.dirPath = true
			return nil
		}
		p.dirPath = false
		return nil
	}
	p.exist = true
	if fi.IsDir() {
		p.dirPath = true
		return nil
	}
	p.dirPath = false
	return nil
}

var dP paths

func check(d string, l int) bool {
	// sP := paths{}
	// dP = paths{}
	// sP.initP(s)
	dP.initP(d)
	if dP.dirPath && dP.exist {
		return true
	}
	if (!dP.dirPath) && (!dP.exist) {
		if l == 1 {
			return true
		}
	}
	return false
}

func main() {
	s := "/data/godevifs"
	fmt.Println(s)
	fmt.Println(check(s,1))
}
