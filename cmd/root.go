package cmd

import (
	"TCPScan/pkg"

	"github.com/spf13/cobra"
)

// func Demo() {
// 	app := &cli.App{
// 		Name:  "TCPScan",
// 		Usage: "let network discover easy.",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println(c.Args().Len())
// 			return nil
// 		},
// 	}

// 	err := app.Run(os.Args)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

var rootCmd = &cobra.Command{
	Use: "TCPScan",
	Short: "简易网络嗅探器",
	Run: func(cmd *cobra.Command, args []string) {
		segments := []string{"10", "1", "1"}
		pkg.Start(segments, 70, 90, 100)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
