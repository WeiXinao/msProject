package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnakeToLowerCamel(t *testing.T) {
	testCases := []struct {
		name   string
		snake  string
		expect string
		err    error
	}{
		{
			name:   "非法 snake name: 情况1",
			snake:  "jiADCJ",
			expect: "",
			err:    ErrIllegalSakeName,
		},
		{
			name:   "非法 snake name: 情况2",
			snake:  "hello__world",
			expect: "",
			err:    ErrIllegalSakeName,
		},
		{
			name:   "开头_",
			snake:  "_hello_world",
			expect: "helloWorld",
			err:    nil,
		},
		{
			name:   "结尾_",
			snake:  "hello_world_",
			expect: "helloWorld",
			err:    nil,
		},
		{
			name:   "正常情况",
			snake:  "hello_world",
			expect: "helloWorld",
			err:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			camel, err := SnakeToLowerCamel(tc.snake)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expect, camel)
		})
	}
}
