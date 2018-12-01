package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
3
3
4
4
5`

var testOkResult = `1
2
3
4
5
`

var testFail = `1
2
1
`

func TestOk(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test for OK failed - error")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("test for OK failed - result not match \n %v %v", result, testOkResult)
	}
}

func TestForError(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("test for OK failed - error: %v", err)
	}
}
