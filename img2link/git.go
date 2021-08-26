/**************************************
 * @Author: mazhuang
 * @Date: 2021-08-18 18:21:56
 * @LastEditTime: 2021-08-26 09:49:39
 * @Description:
 **************************************/

package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

var (
	commitFirst  = false
	gitPath      = "."
	gitCommitMsg = "upload images"
)

type repository struct {
	r     *git.Repository
	wt    *git.Worktree
	files []string
}

func newRepository() *repository {
	r, err := git.PlainOpen(gitPath)
	if err != nil {
		color.Red("git init", err)
		os.Exit(-1)
	}
	wt, err := r.Worktree()
	if err != nil {
		color.Red("git worktree ", err)
		os.Exit(-1)
	}
	return &repository{
		r:     r,
		wt:    wt,
		files: make([]string, 0),
	}
}

// Addr return the origin address of the repository
func (r *repository) Addr() string {
	o, err := r.r.Remote("origin")
	if err != nil {
		color.Red("git remote", err)
		os.Exit(-1)
	}
	return o.Config().URLs[0]
}

func (r *repository) Status() {
	status, err := r.wt.Status()
	if err != nil {
		color.Red("git status err", err)
		os.Exit(-1)
	}
	r.files = make([]string, 0)
	for file := range status {
		r.files = append(r.files, file)
	}
	fmt.Println("git status:")
	fmt.Println(status)
}

func (r *repository) Commit() (images []string) {
	r.Status()
	images = make([]string, 0)
	for _, f := range r.files {
		if isImg(f) {
			images = append(images, f)
		}
		fmt.Println("git add", f)
		_, err := r.wt.Add(f)
		if err != nil {
			color.Yellow("git Add err: ", err)
		}
	}
	fmt.Println("")
	r.Status()
	commit, err := r.wt.Commit(gitCommitMsg, &git.CommitOptions{})
	if err != nil {
		color.Red("git commit err: ", err)
	}
	obj, err := r.r.CommitObject(commit)
	if err != nil {
		color.Red("git commit object err: ", err)
	}
	fmt.Println(obj)
	r.Push()
	return
}

func (r *repository) Push() {
	fmt.Println("start pushing ...")
	err := r.r.Push(&git.PushOptions{})
	if err != nil {
		color.Red("!!! push to remote err, Please push manually ")
		return
	}
	fmt.Println("push done.")
}
