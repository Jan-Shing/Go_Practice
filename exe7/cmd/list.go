package cmd

import(
	"github.com/spf13/cobra"
	"github.com/user/db"
	"fmt"
	"os"
)

var cmdList = &cobra.Command{
	Use: "list",
	Short: "List all of your incomplete tasks",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd * cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil{
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		if len(tasks) == 0{
			fmt.Println("You have no task to list....\n")
			return
		}
		fmt.Println("You have the following tasks to do: \n")
		for i, task := range tasks{
			fmt.Printf("%d. %s \n", i+1, task.Value)
		}
	},
}

func init(){
	RootCmd.AddCommand(cmdList)
}