/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var limit int;

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_anime_cli",
	Short: "Command line application to fetch your favorite anime details",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    query := args[0]

    config := Config{
      Limit: fmt.Sprint(limit),  
    }
    
    results := Search(query, config)

    // Table Configs
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.SetStyle(table.StyleBold)

    t.AppendHeader(table.Row{"Title", "Episodes", "Type", "Status"})
  
    for _, anime := range results {
      t.AppendRow(table.Row{anime.Title, anime.Episodes, anime.Type, anime.Status})
    } 
  
    t.Render()
  },
}

func init(){
  rootCmd.Flags().IntVar(&limit, "limit", 10, "Limit the number of search results" )
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
  }
}
