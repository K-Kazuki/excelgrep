/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCmd(t *testing.T) {
	type result struct {
		pattern string
		paths   []string
	}

	tests := []struct {
		name string
		args string
		want result
	}{
		{
			name: "",
			args: "aaa bbb ccc",
			want: result{
				pattern: "aaa",
				paths:   []string{"bbb", "ccc"},
			},
		},
	}

	for _, tt := range tests {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()
		cmd.SetOutput(buf)

		cmdArgs := strings.Split(tt.args, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs)
		cmd.Execute()

		rawGot := buf.String()
		r := strings.Split(rawGot, "\n")

		for _, v := range r {
			fmt.Printf("\ttype: %T, val: %v\n", v, v)
		}
		fmt.Printf("got: %v", r)

		t.Run(tt.name, func(t *testing.T) {
			t.Errorf("format string")
			// if "" != tt.want {
			// 	t.Errorf("NewCmdRoot() = %v, want %v", got, tt.want)
			// }
		})
	}
}
