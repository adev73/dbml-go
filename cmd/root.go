/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/adev73/dbml-go/internal/gen-go-model/gen"
	"github.com/spf13/cobra"
)

var (
	from             string
	out              string
	gopackage        string
	fieldtags        []string
	shouldGenTblName bool
	rememberAlias    bool
	recursive        bool
	exclude          string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dbml-go",
	Short: "Generate Go models from DBML files",
	Long: `A CLI tool to generate Go model structs from DBML (Database Markup Language) files.
Supports generating models from single files or directories of DBML files.`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.Generate(gen.Opts{
			From:             from,
			Out:              out,
			Package:          gopackage,
			FieldTags:        fieldtags,
			ShouldGenTblName: shouldGenTblName,
			RememberAlias:    rememberAlias,
			Recursive:        recursive,
			Exclude:          exclude,
		})
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dbml-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().
		StringVarP(&from, "from", "f", "database.dbml", "source of dbml, can be https://dbdiagram.io/... | fire_name.dbml")
	rootCmd.Flags().StringVarP(&out, "out", "o", "model", "output folder")
	rootCmd.Flags().
		StringVarP(&gopackage, "package", "p", "model", "package name for generated files")
	rootCmd.Flags().
		StringArrayVarP(&fieldtags, "fieldtags", "t", []string{"db", "json", "mapstructure"}, "go field tags to generate")
	rootCmd.Flags().
		BoolVar(&shouldGenTblName, "gen-table-name", false, "generate \"TableName\" function for models")
	rootCmd.Flags().
		BoolVar(&rememberAlias, "remember-alias", false, "remember table alias (only when 'from' is a directory)")
	rootCmd.Flags().
		BoolVar(&recursive, "recursive", false, "recursively search directories (only when 'from' is a directory)")
	rootCmd.Flags().
		StringVarP(&exclude, "exclude", "E", "", "regex pattern to exclude files (only when 'from' is a directory)")
}
