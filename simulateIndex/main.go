// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-martini/martini"
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
	size = 100000
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
		a := Article{ID: rand.Intn(size), Content: RandStringRunes(50), UserID: rand.Intn(size / 10)}
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

	m := martini.Classic()
	// idIndexs := sortID(articles)
	// // 通过Id查找
	//
	// m.Group("/index/articles", func(r martini.Route){
	// 	m.Get("/detail/:id", func(params martini.Params) (int, string) {
	// 		id := params["id"]
	// 		ID, err := strconv.Atoi(id)
	// 		if err != nil {
	// 			log.Println(err)
	// 			return 503, ""
	// 		}
	// 		index,err:=quickFind(idIndexs,ID)
	// 		if err != nil{
	// 		return 405, ""
	// 		}
	// })

	m.Group("/articles", func(r martini.Router) {
		m.Get("/detail/:id", func(params martini.Params) (int, string) {
			id := params["id"]
			ID, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				return 503, ""
			}
			for _, a := range articles {
				if a.ID == ID {
					return 200, a.Content
				}
			}
			return 405, ""
		})
		m.Get("/:id", func(params martini.Params) (int, string) {
			id := params["id"]
			ID, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				return 503, ""
			}
			var arts []Article
			for _, a := range articles {
				if a.ID == ID {
					arts = append(arts, a)
				}
			}
			data, err := json.Marshal(arts)
			if err != nil {
				log.Println(err)
				return 503, ""
			}
			return 200, string(data)
		})
		m.Get("/search/:userID", func(params martini.Params) (int, string) {
			uid := params["userID"]
			UID, err := strconv.Atoi(uid)
			if err != nil {
				log.Println(err)
				return 503, ""
			}
			var arts []Article
			for _, a := range articles {
				if a.UserID == UID {
					arts = append(arts, a)
				}
			}
			data, err := json.Marshal(arts)
			if err != nil {
				log.Println(err)
				return 503, ""
			}
			return 200, string(data)

		})
	})
	m.RunOnAddr(":8080")
}
