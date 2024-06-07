package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/alexaandru/go-sitter-forest/sqlite"
	sitter "github.com/alexaandru/go-tree-sitter-bare"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hardgo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const (
	code = `
PRAGMA schema.auto_vacuum;
PRAGMA schema.auto_vacuum = 0;
PRAGMA schema.auto_vacuum(NONE);

SELECT EXISTS (SELECT cname FROM tblname);
SELECT NOT EXISTS (SELECT cname FROM tblname);

SELECT ALL count(*), max(a) FROM t1 WHERE b>'one' GROUP BY b;
SELECT ALL count(*), max(a) FROM t1 WHERE a!='b' GROUP BY b HAVING count(*)=1;
SELECT ALL count(*), max(a) FROM t1 WHERE 0 GROUP BY b HAVING count(*)=2;
`
	expected = "(sql_stmt_list (sql_stmt (pragma_stmt (PRAGMA) (identifier) (identifier))) (sql_stmt (pragma_stmt (PRAGMA) (identifier) (identifier) (pragma_value (signed_number (numeric_literal))))) (sql_stmt (pragma_stmt (PRAGMA) (identifier) (identifier) (pragma_value (identifier)))) (sql_stmt (select_stmt (SELECT) (EXISTS) (select_stmt (SELECT) (identifier) (from_clause (FROM) (table_or_subquery (identifier)))))) (sql_stmt (select_stmt (SELECT) (NOT) (EXISTS) (select_stmt (SELECT) (identifier) (from_clause (FROM) (table_or_subquery (identifier)))))) (sql_stmt (select_stmt (SELECT) (ALL) (function_name (identifier)) (function_name (identifier)) (identifier) (from_clause (FROM) (table_or_subquery (identifier))) (where_clause (WHERE) (identifier) (string_literal)) (group_by_clause (GROUP) (BY) (identifier)))) (sql_stmt (select_stmt (SELECT) (ALL) (function_name (identifier)) (function_name (identifier)) (identifier) (from_clause (FROM) (table_or_subquery (identifier))) (where_clause (WHERE) (identifier) (string_literal)) (group_by_clause (GROUP) (BY) (identifier) (HAVING) (function_name (identifier)) (numeric_literal)))) (sql_stmt (select_stmt (SELECT) (ALL) (function_name (identifier)) (function_name (identifier)) (identifier) (from_clause (FROM) (table_or_subquery (identifier))) (where_clause (WHERE) (numeric_literal)) (group_by_clause (GROUP) (BY) (identifier) (HAVING) (function_name (identifier)) (numeric_literal)))))"
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hardgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	n, err := sitter.ParseCtx(context.Background(), []byte(code), sqlite.GetLanguage())
	if err != nil {
		fmt.Printf("Expected no error got %v", err)
		os.Exit(1)

	}

	if act := n.String(); act != expected {
		fmt.Printf("Expected %q got %q", expected, act)
		os.Exit(1)
	}
}
