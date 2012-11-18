package main

import (
	"testing"
)

func Test_findFiles(t *testing.T) {
	files := findFiles([]string{"./testdata"})
	if file := <-files; file != "testdata/01tcb10.txt" {
		t.Errorf("First file was wrong %v.", file)
	}
	if file := <-files; file != "testdata/02tcb10.txt" {
		t.Errorf("Second file was wrong %v.", file)
	}
	if _, ok := <-files; ok {
		t.Errorf("Should only find two files.")
	}
}