// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

/*
模拟索引的设计
1.设计一个存储器Array来存储文章
2. 通过Id查找文章
3. 通过userId查找文章
4. 将文章按id排序，测试通过ID查找文章效率
5. 给useId创建索引，测试通过userID的查找文章效率
*/

const (
	size = 1000
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Article struct {
	ID      int
	Content string
	UserID  int
}

func main() {

	articles := make([]Article, 0, size)
	start := time.Now()
	for i := 0; i < size; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		a := Article{ID: r.Intn(10000), Content: RandStringRunes(20), UserID: r.Intn(200)}
		articles = append(articles, a)
	}
	fmt.Println("generate take:", time.Now().Sub(start))
	start = time.Now()
	data, err := json.Marshal(articles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("marsh take:", time.Now().Sub(start))
	start = time.Now()
	f, err := os.OpenFile("resource", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(data)
	fmt.Println("write take:", time.Now().Sub(start))
}
