package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
)

type Person struct {
	Name       string `yaml:"name"`
	Age        int    `yaml:"age"`
	Occupation string `yaml:"occupation"`
}

var LongText = `
Lookup a field in a YAML file by name.

<name> is the name of the person to look up.
<output_field> is the field to look up (age, occupation, etc.).
`
var ExampleText = `
  lookup-cli John age
  lookup-cli Jane occupation
`
var (
	yamlFile string
	rootCmd  = &cobra.Command{
		Use:     "lookup-cli <name> <output_field>",
		Short:   "Lookup a field in a YAML file by name",
		Long:    LongText,
		Example: ExampleText,
		//Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			// 参数不足 2 个,输出帮助并退出
			if len(args) != 2 {
				_ = cmd.Help()
				// 返回一个错误,使得退出代码不会为 0 给 CI 管道中运行的时候可以报错
				// 题目要求没有说此时要不要给报错,只说了给出帮助文本,所以下面代码视情况而定
				// cmd.PrintErrln("Error: accepts 2 arg(s), received", len(args))
				return
			}

			name := args[0]
			outputField := args[1]

			data, err := os.ReadFile(yamlFile)
			if err != nil {
				fmt.Println("Error reading YAML file:", err)
				return
			}

			var people []Person
			err = yaml.Unmarshal(data, &people)
			if err != nil {
				fmt.Println("Error parsing YAML file:", err)
				return
			}

			for _, person := range people {
				if person.Name == name {
					// 此处一定是找到了名字,所以进入后续字段判断
					switch outputField {
					case "age":
						if person.Age != 0 {
							fmt.Println(person.Age)
						} else {
							fmt.Println("Field not found")
						}
					case "occupation":
						if person.Occupation != "" {
							fmt.Println(person.Occupation)
						} else {
							fmt.Println("Field not found")
						}
					default:
						fmt.Println("Field not found")
					}
					return
				}
			}
			// 名字作为最先判断的,没有就退出
			fmt.Println("Name not found")
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&yamlFile, "file", "f", "data.yml", "Path to the YAML file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
