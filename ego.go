package ego

import (
	"bytes"
	"github.com/yuin/cidre"
	"io"
	"net/http"
	"reflect"
)

type EgoRenderer struct {
	cidre.BaseRenderer
}

type egoWriter struct {
	layer int
	w     io.Writer
	b     *bytes.Buffer
}

func (e *egoWriter) Write(p []byte) (int, error) {
	if e.layer == 0 {
		v, err := e.b.Write(p)
		return v, err
	} else {
		v, err := e.w.Write(p)
		return v, err
	}
}

func NewEgoRenderer() *EgoRenderer {
	return &EgoRenderer{}
}

func (r *EgoRenderer) RenderTemplateFile(w io.Writer, name string, value interface{}) {
	panic("EgoRenderer does not support the RenderTemplateFile method.")
}

func renderEgo(w io.Writer, args ...interface{}) {
	ew, eok := w.(*egoWriter)
	contents := ""
	if eok {
		ew.layer++
	}
	if ew.layer > 0 {
		contents = ew.b.String()
	}
	rv := reflect.ValueOf(args[0])
	if rv.Kind() != reflect.Func {
		panic("#2 must be a function.")
	}
	rargs := []reflect.Value{reflect.ValueOf(w)}
	if len(args) > 1 {
		rargs = append(rargs, reflect.ValueOf(args[1]))
	}
	if ew.layer > 0 {
		rargs = append(rargs, reflect.ValueOf(contents))
	}
	rets := rv.Call(rargs)
	err, ok := rets[0].Interface().(error)
	if !ok && !rets[0].IsNil() {
		panic("invalid return value(error or nil expected)")
	}
	if err != nil {
		panic(err)
	}
	if eok && ew.layer == 0 {
		if _, err := ew.w.Write(ew.b.Bytes()); err != nil {
			panic(err)
		}
	}
}

func EgoLayout(w io.Writer, args ...interface{}) {
	renderEgo(w, args...)
}

func (r *EgoRenderer) Html(w http.ResponseWriter, args ...interface{}) {
	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	}
	var b bytes.Buffer
	neww := &egoWriter{-1, w, &b}
	renderEgo(neww, args...)
}
