/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-16 16:35:00
 * @LastEditTime: 2021-08-19 13:49:58
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
	flag.StringVar(&gitPath, "g", ".", "The path of the .git folder ")
	flag.StringVar(&gitCommitMsg, "m", "upload images", "The message of git commit ")
	flag.BoolVar(&commitFirst, "c", false, "The option to do commit before convert")
	flag.Parse()
	cmd.convert()
}
