package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all your completed tasks for today",
	Long: `Lists all your completed tasks for today. For example:

$ task completed
You have finished the following tasks today:
- wash the dishes
- clean the car`,
	Run: completed,
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completed(cmd *cobra.Command, args []string) {
	fmt.Println("completed callback called")
}
