/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"chat/singleTest/agent"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 连接服务器
		conn, err := net.Dial("tcp", "192.168.1.13:9999")
		if err != nil {
			fmt.Println("连接失败", err)
			return
		}
		defer conn.Close()

		// 创建agent
		client := agent.NewAgent(conn)
		client.InitDirectives()

		// 监听消息接收
		go client.Reader()

		buf := bufio.NewReader(os.Stdin)
		// 监听输入
		for {
			fmt.Print("> ")
			directives, err := buf.ReadBytes('\n')
			if err != nil {
				fmt.Println("命令操作", err)
				return
			}
			// 去除换行符
			str := strings.Replace(string(directives), "\n", "", -1)
			// 处理指令
			handeler, ok := client.Directives[str]
			if !ok {
				fmt.Println("指令不存在")
			} else {
				err := handeler.Do(client)
				if err != nil {
					fmt.Println("handler err:", err)
					return
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
