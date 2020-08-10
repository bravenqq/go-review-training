package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var l sync.Mutex

func main() {
	var root string
	if len(os.Args) == 1 {
		root = "."
	} else {
		root = os.Args[1]
	}
	dirsSize := make(map[string]int64)
	seen := make(map[string]bool)
	done := make(chan struct{})
	go func() {
		walkDir(root, dirsSize)
		done <- struct{}{}
	}()

	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-done:
			timer.Stop()
			return
		case <-timer.C:
			l.Lock()
			for dir, size := range dirsSize {
				if !seen[dir] {
					seen[dir] = true
					fmt.Printf("dir:%s size:%.001fGB\n", dir, float64(size)/1e9)
				}
			}
			l.Unlock()
		}

	}
}

func walkDir(dir string, dirsSize map[string]int64) int64 {
	//get all data of dir
	//get fileSize
	//get all sub dir size
	var total int64
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			total += walkDir(filepath.Join(dir, entry.Name()), dirsSize)
		} else {
			total += entry.Size()
		}
	}
	l.Lock()
	dirsSize[dir] = total
	l.Unlock()
	return total
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du:%v\n", err)
		return nil
	}
	return entries
}
