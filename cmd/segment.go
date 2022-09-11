package cmd

import (
	"TCPScan/pkg"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	//segs []string
	segA, segB, segC string
	workers int
)

var segment = &cobra.Command{
	Use: "segment",
	Short: "自定义网段扫描",
	Run: func (cmd *cobra.Command, args []string) {
		color.Magenta("[network segment]: %s.%s.%s.x, workers: %d", segA, segB, segC, workers)
		segments := []string{segA, segB, segC}
		pkg.Start(segments, 70, 90, workers)
	},
}

func init() {
	rootCmd.AddCommand(segment)
	//segment.Flags().StringArrayVarP(&segs, "seg", "s", []string{"10", "1", "1"}, "请输入网段")
	segment.Flags().StringVarP(&segA, "segA", "A", "10", "请输入A网段")
	segment.Flags().StringVarP(&segB, "segB", "B", "1", "请输入B网段")
	segment.Flags().StringVarP(&segC, "segC", "C", "1", "请输入C网段")
	segment.Flags().IntVarP(&workers, "workers", "w", 100, "请输入worker数")
}