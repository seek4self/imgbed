/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-16 16:50:41
 * @LastEditTime: 2021-08-17 15:07:31
 * @Description:
 **************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

var imgSuffix = []string{".png", ".jpg", ".jpeg"}

type style struct {
	begin, end string
}

var mapStyle = map[string]style{
	"md":  {"![](", ")"},
	"url": {"", ""},
}

var (
	github = "github.com"
	branch = "main"
)

var repositoryPath string

func initGit() {
	r, err := git.PlainOpen(".git")
	if err != nil {
		fmt.Println("git init", err)
		os.Exit(-1)
	}
	o, err := r.Remote("origin")
	if err != nil {
		fmt.Println("git remote", err)
		os.Exit(-1)
	}
	gitAddr := o.Config().URLs[0]
	// fmt.Println("git remote:", gitAddr)
	start := strings.Index(gitAddr, github) + len(github) + 1
	end := strings.Index(gitAddr, ".git")
	repositoryPath = gitAddr[start:end]
}

type link struct {
	domain string
	style  string
	target string
}

func (l *link) convert() {
	var (
		imgs []string
		err  error
	)
	l.target = filepath.Clean(l.target)
	if l.target != "." {
		l.target = "." + string(os.PathSeparator) + l.target
	}
	// 遍历文件
	if filepath.Ext(l.target) == "" || filepath.Ext(l.target) == "." {
		imgs, err = readDir(l.target)
		if err != nil {
			fmt.Println("scan dir err: ", err)
			os.Exit(-1)
		}
	} else {
		imgs = append(imgs, l.target)
	}
	// 打印输出
	fmt.Printf("img -> %s link:\n", l.style)
	for _, i := range imgs {
		link := l.format(filepath.ToSlash(strings.TrimLeft(i, ".")))
		fmt.Printf("%s -> %s\n", i, link)
	}
}

func (l link) url() string {
	if l.domain == "cdn" {
		return "https://cdn.jsdelivr.net/gh/" + repositoryPath + "@" + branch
	}
	return "https://raw.githubusercontent.com/" + repositoryPath + "/" + branch
}

func (l link) format(img string) string {
	return mapStyle[l.style].begin + l.url() + img + mapStyle[l.style].end
}

func readDir(dir string) (files []string, err error) {
	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, f := range fileList {
		if f.IsDir() {
			continue
		}
		for _, s := range imgSuffix {
			if filepath.Ext(f.Name()) == s {
				files = append(files, dir+PthSep+f.Name())
			}
		}
	}
	return
}

func WalkPath(root string) (files []string, err error) {
	err = filepath.Walk(root, func(pth string, info os.FileInfo, err error) error {
		log.Printf(pth)
		if info.IsDir() {
			log.Println("is dir")
		}
		if path.Ext(info.Name()) == ".go" {
			log.Println("is file")
			files = append(files, pth)
		}
		return nil
	})
	return
}
