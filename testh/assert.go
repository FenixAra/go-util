package testh

import (
	"reflect"
	"testing"
)

func AssertEqual(msg string, expected, got interface{}, t *testing.T) {
	if expected != got {
		t.Errorf("%s. Expected value: %v, Got: %v", msg, expected, got)
		t.FailNow()
	}
}

func AssertDeepEqual(msg string, expected, got interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s. Expected value: %v, Got: %v", msg, expected, got)
		t.FailNow()
	}
}

func AssertNoErr(msg string, err error, t *testing.T) {
	if err != nil {
		t.Errorf("%s. Got Err: %v", msg, err)
		t.FailNow()
	}
}

func AssertErr(msg string, err error, t *testing.T) {
	if err == nil {
		t.Errorf("%s. Got no Err: %v", msg, err)
		t.FailNow()
	}
}

func AssertContainsObj(msg string, arr []interface{}, v interface{}, t *testing.T) {
	for _, a := range arr {
		if a == v {
			return
		}
	}

	t.Errorf("%s. The array: %v does not contain object: %v", msg, arr, v)
	t.FailNow()
}

func AssertDeepContainsObj(msg string, arr []interface{}, v interface{}, t *testing.T) {
	for _, a := range arr {
		if reflect.DeepEqual(a, v) {
			return
		}
	}

	t.Errorf("%s. The array: %v does not contain object: %v", msg, arr, v)
	t.FailNow()
}

func AssertContainsAny(msg string, x []interface{}, y []interface{}, t *testing.T) {
	v := make(map[interface{}]struct{})
	for _, yv := range y {
		v[yv] = struct{}{}
	}

	for _, xv := range x {
		if _, ok := v[xv]; ok {
			return
		}
	}

	t.Errorf("%s. The array: %v does not contain any of the array's object: %v", msg, x, y)
	t.FailNow()
}

func AssertContainsAll(msg string, x []interface{}, y []interface{}, t *testing.T) {
	v := make(map[interface{}]struct{})
	for _, xv := range x {
		v[xv] = struct{}{}
	}

	for _, yv := range y {
		if _, ok := v[yv]; !ok {
			t.Errorf("%s. The array: %v does not contain all of the array's object: %v. Missing object: %v", msg, x, y, yv)
			t.FailNow()
		}
	}
}
