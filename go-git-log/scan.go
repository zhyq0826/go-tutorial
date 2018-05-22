package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

//扫描文件夹中的 git
func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	defer f.Close()
	var path string
	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				// 把包含 git 的 folder 加入 slice
				folders = append(folders, path)
				continue
			}

			// 跳过大包
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}

			// 继续扫码 subpackage
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

// return the dot file
// create is and the enclosing folder if it does not exist
func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLineToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringsSliceToFile(repos, filePath)
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}

	return existing
}

// write content to the file in path overwriting existing content
func dumpStringsSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	ioutil.WriteFile(filePath, []byte(content), 0755)
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

//读取文件按行读取
func parseFileLineToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

//打开文件
func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在
			_, err := os.Create(filePath)
			if err != nil {
				panic(err)
			} else {
				panic(err)
			}
		}
	}

	return f
}
