// Package main provides ...
package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()
	m := NewMockDB(ctrl)
	// m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	// m.EXPECT().Get(gomock.Eq("Sam")).Return(26, nil)
	m.EXPECT().Get(gomock.Any()).Return(0, errors.New("not exist")).AnyTimes()
	tests := []struct {
		input string
		out   int
	}{
		{"Tom", -1},
		{"Sam", 26},
		{"nqq", -1},
	}
	for _, test := range tests {
		if v := GetFromDB(m, test.input); v != test.out {
			t.Errorf("input:%s excepted %d, but get %d\n", test.input, test.out, v)
		}
	}
}
