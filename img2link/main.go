/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-16 16:35:00
 * @LastEditTime: 2021-08-17 15:35:29
 * @Description:
 **************************************/

package main

import (
	"flag"
)

func main() {
	var cmd link
	flag.StringVar(&cmd.domain, "d", "github", "Image link domain, choose 'cdn' to use jsDelivr CDN acceleration")
	flag.StringVar(&cmd.style, "s", "md", "Image link style, 'md' for markdown, 'url' for http")
	flag.StringVar(&cmd.target, "t", ".", "The path or folder of the target image ")
	flag.Parse()
	initGit()
	cmd.convert()
}
