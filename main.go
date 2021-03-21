/* SPDX-FileCopyrightText: © 2021-2022 Nadim Kobeissi <nadim@symbolic.software>
* SPDX-License-Identifier: GPL-3.0-only */
package main

import (
	"context"
	"fmt"
	"github.com/machinebox/progress"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "cpr [file] [destination]",
		Example: "cpr myMovie.mp4 /Volumes/NAS/Movies/myMovie.mp4",
		Short:   "Copy a file",
		Args:    cobra.ExactArgs(2),
		Hidden:  false,
		Run: func(cmd *cobra.Command, args []string) {
			mainCopy(args[0], args[1],
				func() {},
				func(p int, c string, s string, r string) {
					fmt.Fprint(os.Stdout, "\r\r\r\r")
					fmt.Printf("%d%s • %s/%s • %s ", p, "%", c, s, r)
				},
				func(err error) {
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Println("done")
					}
				},
			)
		},
	}
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func mainCopy(
	srcFilePath string, destFilePath string, onStart func(),
	onProgress func(int, string, string, string), onFinish func(error),
) {
	srcFileInfo, err := os.Lstat(srcFilePath)
	if err != nil {
		onFinish(err)
		return
	}
	srcFileDescriptor, err := os.Open(srcFilePath)
	if err != nil {
		onFinish(err)
		return
	}
	destFileDescriptor, err := os.OpenFile(
		destFilePath, os.O_RDWR|os.O_CREATE, srcFileInfo.Mode(),
	)
	if err != nil {
		onFinish(err)
		return
	}
	srcFileReader := progress.NewReader(srcFileDescriptor)
	go func() {
		ctx := context.Background()
		progressChan := progress.NewTicker(
			ctx, srcFileReader, srcFileInfo.Size(), 100*time.Millisecond,
		)
		for p := range progressChan {
			onProgress(
				int(p.Percent()),
				mainCopyFileSizeFormat(p.N()),
				mainCopyFileSizeFormat(p.Size()),
				mainCopyDurationFormat(p.Remaining()),
			)
		}
	}()
	onStart()
	_, err = destFileDescriptor.ReadFrom(srcFileReader)
	onFinish(err)
}

func mainCopyFileSizeFormat(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func mainCopyDurationFormat(d time.Duration) string {
	min := int(d.Minutes())
	sec := int(d.Seconds()) - (min * 60)
	return fmt.Sprintf("%dmin%dsec", min, sec)
}
