package cmd

import (
	"context"
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/bootstrap"
	"github.com/spf13/cobra"
)

// esIndexCmd represents the esIndex command
var esIndexCmd = &cobra.Command{
	Use:   "es:index",
	Short: "创建index",
	Long:  `创建index.`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.CliInit()

		global.Env = env
		// 执行ES请求需要提供一个上下文对象
		ctx := context.Background()

		// 首先检测下weibo索引是否存在
		exists, err := global.ElasticsearchClient.IndexExists("api_log").Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !exists {
			// weibo索引不存在，则创建一个
			_, err := global.ElasticsearchClient.CreateIndex("api_log").BodyString(mapping).Do(ctx)
			if err != nil {
				// Handle error
				panic(err)
			}
		}
	},
}

var env string

func init() {
	rootCmd.AddCommand(esIndexCmd)

	esIndexCmd.Flags().StringVarP(&env, "env", "e", "", "a")
}

const mapping = `
{
  "mappings": {
    "properties": {
      "created":{
		"type": "date",
		"format": "yyyy-MM-dd HH:mm:ss"
      },
      "request":{
		"type": "text"
      },
      "response":{
		"type": "text"
      },
      "header":{
		"type": "text"
      },
      "code":{
		"type": "integer"
      },
      "query":{
		"type": "text"
      }
    }
  }
}`
