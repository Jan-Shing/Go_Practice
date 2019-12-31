package cmd

import(
	"github.com/spf13/cobra"
	"github.com/user/db"
	"fmt"
	"os"
	"strings"
)

var cmdAdd = &cobra.Command{
	Use: "add [task string]",
	Short: "Add a new task to your TODO list",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd * cobra.Command, args []string) {
		task := strings.Join(args, " ")
		
		tasks, err := db.AllTasks()
		if err != nil{
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		match := db.FindTask(task)

		if len(tasks) == 0 || match == false  {
			_, err := db.CreateTask(task)
			if err != nil{
				fmt.Println("Something went wrong: ", err)
				return
			}
			fmt.Printf("Added:  \"%s\"  to your task list.\n", task)
		}

		
	},
}

func init(){
	RootCmd.AddCommand(cmdAdd)
}