package cmd

import(
	"strconv"
	"github.com/spf13/cobra"
	"github.com/user/db"
	"fmt"
)

var cmdDo = &cobra.Command{
	Use: "do [task index]",
	Short: "Mark a task on your TODO list as complete",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd * cobra.Command, args []string) {
		var ids []int
		for _, arg := range args{
			id, err := strconv.Atoi(arg)
			if err != nil{
				fmt.Println("Failed to parse the argument: ", arg)
			}else{
				ids = append(ids, id)
			}

			tasks, err := db.AllTasks()
			if err != nil{
				fmt.Println("Something went wrong: ", err)
				return
			}
			for _, id = range ids{
				if id <= 0 || id > len(tasks) {
					fmt.Println("Invalid task number: ", id)
					continue
				}
				task := tasks[id-1]
				err := db.DeleteTask(task.Key)
				if err != nil{
					fmt.Printf("Failed to mark \"%d\" as completed. Eooro: %s \n", id, err)
				}else{
					fmt.Printf("Marked \" %d \" as completed.\n",id)
				}
			}
		}
	
	},
}

func init(){
	RootCmd.AddCommand(cmdDo)
}