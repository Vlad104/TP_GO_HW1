package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"bytes"
	"sort"
)

// интерфейс для сортировки папок и файлов
type byName []os.FileInfo

func (data byName) Len() int { 
	return len(data) 
}

func (data byName) Swap(i, j int) { 
	data[i], data[j] = data[j], data[i] 
}

func (data byName) Less(i, j int) bool { 
	return data[i].Name() < data[j].Name() 
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	var buffer bytes.Buffer // до конца обхода вся ифномация заисывается в этот буфер
	err := makeTree(&buffer, path, "", printFiles)
	fmt.Fprint(out, buffer.String())
	return err
}

// рекурсивная функция обхода дирректорий
func makeTree(out io.Writer, path string, prefix string, printFiles bool) error {
	newFile, err := os.Open(path)
	if err != nil {
		return err
	}
	fileInfo, err := newFile.Readdir(-1)
	newFile.Close()
	if err != nil {
		return err
	}

	sort.Sort(byName(fileInfo))

	// вычисляем количество папок
	// если указан флаг -f, то папок + файлов
	totalFileNumber := calcDirNub(fileInfo, printFiles)

	// currentFileNumber считает кол-во пройденых папок
	// если указан флаг -f, то папок + файлов
	currentFileNumber := 0
	for _, file := range fileInfo {
		isLastDir := currentFileNumber == totalFileNumber - 1	
		var indent string
		if isLastDir {
			indent = "└───"
		} else {
			indent = "├───"
		}

		newPath := path + string(os.PathSeparator) + file.Name()
		if file.IsDir() {
			currentFileNumber++
			printDir(out, file.Name(), indent, prefix)
			if isLastDir {
				if makeTree(out, newPath, prefix + "\t", printFiles) != nil {
					return err
				}
			} else {
				if makeTree(out, newPath, prefix + "│\t", printFiles) != nil {
					return err
				}
			}
		} else {
			if (printFiles) {
				currentFileNumber++
				printFile(out, file.Name(), indent, prefix, int(file.Size()))				
			}
		}
	}
	return nil
}

func printDir(out io.Writer, fileName string, indent string, prefix string) {
	fmt.Fprint(out, prefix)
	fmt.Fprint(out, indent)
	fmt.Fprint(out, fileName)
	fmt.Fprint(out, "\n")
}

func printFile(out io.Writer, fileName string, indent string, prefix string, fileSize int) {
	fmt.Fprint(out, prefix)
	fmt.Fprint(out, indent)
	fmt.Fprint(out, fileName)
	var sizeStr string
	if fileSize > 0 {
		sizeStr = " (" + strconv.Itoa(fileSize) + "b)"	
	} else {
		sizeStr = " (empty)"		
	}
	fmt.Fprint(out, sizeStr)	
	fmt.Fprint(out, "\n")	
}

func calcDirNub(fileInfo []os.FileInfo, printFiles bool) int {
	// если задан флаг -f
	if printFiles {
		return len(fileInfo)
	}

	// если не задан флаг -f
	// считаем кол-во папок
	N := 0
	for _, file := range fileInfo {
		if file.IsDir() {
			N++
		}
	}
	return N
}