package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestLookupCLI(t *testing.T) {
	// 因为不知道 IDE/CI 管道会在什么目录执行这个测试,所以先写出这个文件
	yamlFile = "testdata.yml"

	testData := `
- name: Alice
  age: 18
  occupation: student
- name: Bob
  age: 33
  occupation: unemployed
- name: Charlie
  age: 65
- name: David
  age: 25
  occupation: software engineer
`
	err := os.WriteFile(yamlFile, []byte(testData), 0644)
	if err != nil {
		t.Fatalf("Failed to create test YAML file: %v", err)
	}
	defer os.Remove(yamlFile)

	// 准备测试用例
	tests := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
	}{
		{"Valid lookup - Alice age", []string{"Alice", "age"}, "18\n", ""},
		{"Valid lookup - Bob occupation", []string{"Bob", "occupation"}, "unemployed\n", ""},
		{"Field not found - Charlie occupation", []string{"Charlie", "occupation"}, "Field not found\n", ""},
		{"Name not found - Eve age", []string{"Eve", "age"}, "Name not found\n", ""},
		//{"Missing arguments", []string{}, "Usage: lookup-cli <name> <output_field>\n", "Error: accepts 2 arg(s), received 0\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 准备捕获 stdout 和 stderr
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			oldStderr := os.Stderr
			rErr, wErr, _ := os.Pipe()
			os.Stderr = wErr

			// 执行并收集结果
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()

			// 关闭先前的管道并恢复 stdout 和 stderr
			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			buf.ReadFrom(r)
			actualOut := buf.String()

			wErr.Close()
			os.Stderr = oldStderr
			var bufErr bytes.Buffer
			bufErr.ReadFrom(rErr)
			actualErr := bufErr.String()

			// 输出的验证
			if actualOut != tt.expectedOut {
				t.Errorf("Expected output %q, got %q", tt.expectedOut, actualOut)
			}

			// 错误信息的验证
			if actualErr != tt.expectedErr {
				t.Errorf("Expected error %q, got %q", tt.expectedErr, actualErr)
			}

			// 出错的时候,验证返回的错误
			if (err != nil) != (tt.expectedErr != "") {
				t.Errorf("Expected error %v, got %v", tt.expectedErr != "", err)
			}
		})
	}
}

func TestLookupCLIWithoutArgs(t *testing.T) {
	// 准备捕获 stdout 和 stderr
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldStderr := os.Stderr
	rErr, wErr, _ := os.Pipe()
	os.Stderr = wErr

	// 执行并收集结果
	rootCmd.SetArgs([]string{})
	err := rootCmd.Execute()

	// 关闭先前的管道并恢复 stdout 和 stderr
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	buf.ReadFrom(r)
	actualOut := buf.String()

	wErr.Close()
	os.Stderr = oldStderr
	var bufErr bytes.Buffer
	bufErr.ReadFrom(rErr)
	actualErr := bufErr.String()

	// 输出的验证,应当包含帮助信息
	if !strings.Contains(actualOut, ExampleText) {
		t.Errorf("Expected output %q, got %q", "Error: accepts 2 arg(s), received 0\n", actualOut)
	}

	// 错误信息的验证
	if actualErr != "" {
		t.Errorf("Expected error %q, got %q", "", actualErr)
	}

	// 出错的时候,验证返回的错误
	if (err != nil) != (actualErr != "") {
		t.Errorf("Expected error %v, got %v", actualErr != "", err)
	}
}
