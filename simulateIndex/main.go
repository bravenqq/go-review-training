// Package main provides ...
package main

import (
	"encoding/json"
	"errors"
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
4. 给文章id建立索引，测试通过ID索引查找文章效率
5. 给useId创建索引，测试通过userID索引查找文章效率
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

type IDIndex struct {
	ID       int
	Position int
}

func sortID(arr []Article) []IDIndex {
	indexs := make([]IDIndex, len(arr))
	for i := 0; i < len(arr); i++ {
		indexs[i] = IDIndex{ID: arr[i].ID, Position: i}
	}
	//选择排序
	for i := 0; i < len(indexs); i++ {
		min := i
		for j := i + 1; j < len(indexs); j++ {
			if indexs[min].ID > indexs[j].ID {
				min = j
			}
		}

		if min != i {
			temp := indexs[i]
			indexs[i] = indexs[min]
			indexs[min] = temp
		}

	}
	return indexs
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
	idIndexs := sortID(articles)
	fmt.Println(idIndexs)
	// 通过Id查找
	m.Group("/index/articles", func(r martini.Router) {
		m.Get("/detail/:id", func(params martini.Params) (int, string) {
			id := params["id"]
			ID, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				return 503, ""
			}

			var quickFind func(arr []IDIndex, ID int) (int, error)
			quickFind = func(arr []IDIndex, ID int) (int, error) {
				if len(arr) == 0 {
					return 0, errors.New("找不到")
				}
				mid := (len(arr) - 1) / 2
				fmt.Println("mid:", mid)
				if arr[mid].ID == ID {
					return arr[mid].Position, nil
				}
				if ID > arr[mid].ID {
					mid += 1
					if mid >= len(arr) {
						return 0, errors.New("找不到")
					}
					arr = arr[mid:]
				}

				if ID < arr[mid].ID {
					if mid < 0 {
						return 0, errors.New("找不到")
					}
					arr = arr[0:mid]
				}
				quickFind(arr, ID)
				return 0, errors.New("找不到")
			}

			index, err := quickFind(idIndexs, ID)

			if err != nil {
				return 405, ""
			}
			return 200, articles[index].Content
		})
	})

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
