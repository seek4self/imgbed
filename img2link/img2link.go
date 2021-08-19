/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-16 16:50:41
 * @LastEditTime: 2021-08-19 14:32:14
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

func initGit() []string {
	repo := newRepository()
	gitAddr := repo.Addr()
	start := strings.Index(gitAddr, github) + len(github) + 1
	end := strings.Index(gitAddr, ".git")
	repositoryPath = gitAddr[start:end]
	if commitFirst {
		return repo.Commit()
	}
	return nil
}

func findImages(target string) []string {
	target = filepath.Clean(target)
	if target != "." {
		target = "." + string(os.PathSeparator) + target
	}
	// 遍历文件
	if len(filepath.Ext(target)) >= 1 {
		return []string{target}
	}
	images, err := readDir(target)
	if err != nil {
		fmt.Println("scan dir err: ", err)
		os.Exit(-1)
	}
	return images
}

type link struct {
	domain string
	style  string
	target string
}

func (l *link) convert() {
	images := initGit()
	if !commitFirst {
		images = findImages(l.target)
	}
	// 打印输出
	fmt.Printf("convert img -> %s link:\n", l.style)
	for _, i := range images {
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

func isImg(file string) bool {
	for _, s := range imgSuffix {
		if filepath.Ext(file) == s {
			return true
		}
	}
	return false
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
		if isImg(f.Name()) {
			files = append(files, dir+PthSep+f.Name())
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
