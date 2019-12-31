package cmd

import(
	"github.com/spf13/cobra"
	"github.com/user/db"
	"fmt"
)

var cmdRemove = &cobra.Command{
	Use: "Remove",
	Short: "Rmove all of your incomplete tasks",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd * cobra.Command, args []string) {
		err := db.DeleteAllTask()
		if err != nil{
			fmt.Println("Rmove all task fail....")
		}else{
			fmt.Println(" Remove all task... ")
		}
	},
}

func init(){
	RootCmd.AddCommand(cmdRemove)
}