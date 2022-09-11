package cmd

import (
	"TCPScan/pkg"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	h     string
	start int
	end   int
)

var host = &cobra.Command{
	Use:   "host",
	Short: "端口扫描",
	Long:  "针对单一主机的端口扫描",
	Args: func (cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			return errors.New(color.RedString("extra parameter: " + args[0]))
		}
		if end < start {
			return errors.New(color.RedString("请输入正确的端口范围"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		color.Magenta("[address]: %s:%d~%d", h, start, end)
		address := fmt.Sprintf("%s:%%d", h)
		pkg.StartPort(address, start, end)
	},
}

func init() {
	rootCmd.AddCommand(host)
	host.Flags().StringVarP(&h, "address", "a", "127.0.0.1", "请输入主机地址")
	host.Flags().IntVarP(&start, "startPort", "s", 1, "请输入起始端口")
	host.Flags().IntVarP(&end, "endPort", "e", 65535, "请输入终点端口")
}
