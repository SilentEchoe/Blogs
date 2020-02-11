package main

import "io"

var w io.Writer

w = os.Stdout

f := w.(*os.File)
c := w.(*bytes.Buffer)