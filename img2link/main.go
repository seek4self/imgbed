/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-16 16:35:00
 * @LastEditTime: 2021-08-17 11:40:58
 * @Description:
 **************************************/

package main

import (
	"flag"
)

func main() {
	var cmd link
	flag.StringVar(&cmd.domain, "d", "github", "Image link domain, choose 'cdn' to use cdn acceleration")
	flag.StringVar(&cmd.style, "s", "md", "Image link style, 'md' for markdown, 'url' for http")
	flag.StringVar(&cmd.target, "t", ".", "The path or folder of the target image ")
	flag.Parse()
	initGit()
	cmd.convert()
}
