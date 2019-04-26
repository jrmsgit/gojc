// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	fp "path/filepath"
	"sort"
)

var z *zip.Writer
var zbuf *bytes.Buffer
var loaded []string

var b64 = base64.StdEncoding.EncodeToString
var sprintf = fmt.Sprintf

var pkgName string
var prefix string

func init() {
	cwd, _ := os.Getwd()
	pkgName = fp.Base(cwd)
	flag.StringVar(&pkgName, "package", pkgName, "package `name`")
	flag.StringVar(&prefix, "prefix", "", "install prefix `path`")
}

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Printf("zip gen pkg='%s' prefix='%s' %v\n", pkgName, prefix, args)

	files := make([]string, 0)
	for _, patt := range args {
		l, err := fp.Glob(patt)
		check(err)
		files = append(files, l...)
	}

	sort.Strings(files)
	loaded = make([]string, 0)
	zbuf = new(bytes.Buffer)
	z = zip.NewWriter(zbuf)

	uniq := make(map[string]bool)
	for _, fn := range files {
		_, done := uniq[fn]
		if !done {
			check(load(fn))
			uniq[fn] = true
		}
	}

	check(z.Close())
	check(write())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func load(filename string) error {
	fi, err := os.Stat(filename)
	check(err)
	var src io.ReadCloser
	src, err = os.Open(filename)
	check(err)
	defer src.Close()
	var h *zip.FileHeader
	h, err = zip.FileInfoHeader(fi)
	check(err)
	h.Name = filename
	var zh io.Writer
	zh, err = z.CreateHeader(h)
	check(err)
	var n int64
	n, err = io.Copy(zh, src)
	check(err)
	check(z.Flush())
	fmt.Printf("zip %s %d\n", filename, n)
	loaded = append(loaded, filename)
	return nil
}

func write() error {
	lloaded := len(loaded)
	fmt.Printf("fs.zip %d files\n", lloaded)
	if lloaded == 0 {
		return nil
	}
	zipfile := b64(zbuf.Bytes())
	dst := new(bytes.Buffer)
	_, err := dst.WriteString(sprintf("package %s\n", pkgName))
	check(err)
	_, err = dst.WriteString("\n")
	check(err)
	_, err = dst.WriteString("func init() {\n")
	check(err)
	//~ _, err = dst.WriteString(sprintf("\tprefix = \"%s\"\n", prefix))
	//~ check(err)
	_, err = dst.WriteString(sprintf("\tzipfile = \"%s\"\n", zipfile))
	check(err)
	_, err = dst.WriteString("}\n")
	check(err)
	_, err = dst.WriteString("\n")
	check(err)
	for _, fn := range loaded {
		_, err = dst.WriteString(sprintf("// %s\n", fn))
		check(err)
	}
	fmt.Printf("fs.zip %d\n", zbuf.Len())
	check(ioutil.WriteFile("fs.zip", zbuf.Bytes(), 0640))
	return ioutil.WriteFile("zipfs.go", dst.Bytes(), 0640)
}
