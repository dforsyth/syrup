package syrup

import (
	"fmt"
	"os"
	"reflect"
	"waffle"
)

type StdoutResultWriter struct {
	fields []string
}

func NewStdoutResultWriter(fields []string) *StdoutResultWriter {
	return &StdoutResultWriter{
		fields: fields,
	}
}

func (rw *StdoutResultWriter) WriteResults(w *waffle.Worker) os.Error {
	for _, p := range w.Partitions() {
		for _, v := range p.Vertices() {
			s := reflect.ValueOf(v).Elem()
			fmt.Printf("%s:\n", v.VertexId())
			for _, field := range rw.fields {
				if f := s.FieldByName(field); f.IsValid() {
					fmt.Printf("\t%s: %v\n", field, f.Interface())
				}
			}
		}
	}
	return nil
}
