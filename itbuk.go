package itbuk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Books []Book
}

type Book struct {
	Author      string
	Description string
	ID          int64
	Image       string
	Title       string
}
type BookDetail struct {
	Book
	ISBN     string `json: "isbn"`
	Download string
}

var (
	api_url    = "http://it-ebooks-api.info/v1/"
	configPath = ""
)

func (bd BookDetail) String() string {
	format := "[Book]\n"
	format += "Author: %s\n"
	format += "Title: %s\n"
	format += "Description: %s\n"
	format += "Download: %s\n"
	format += "ISBN: %s\n"

	return fmt.Sprintf(format,
		bd.Author, bd.Title,
		bd.Description,
		bd.Download, bd.ISBN)
}

func Search(topic string, page int) (books []BookDetail, err error) {
	IDch := make(chan int64)
	BDch := make(chan BookDetail)
	var search_url string

	books = make([]BookDetail, page*10)
	for p := 1; p < page+1; p++ {
		search_url = fmt.Sprintf(api_url+"search/"+topic+"/page/%s", strconv.Itoa(p))
		resp, err := http.Get(search_url)

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("URL(%s) error: %s", search_url, err)
			return nil, err
		}

		if err != nil {
			log.Fatalf("error: %s", err)
			return nil, err
		}

		defer resp.Body.Close()

		to_json := new(Response)
		err = json.NewDecoder(resp.Body).Decode(to_json)
		if err != nil {
			log.Fatalf("error decode: %s", err)
			return nil, err
		}

		for _, b := range to_json.Books {
			go detailBook(IDch, BDch)
			IDch <- b.ID
			if err != nil {
				log.Fatalf("error book detail: %s", err)
				return nil, err
			}
		}

		for i, _ := range to_json.Books {
			index := ((p * 10) - 10) + i
			books[index] = <-BDch
		}
	}

	return books, nil
}
func detailBook(IDch chan int64, BDch chan BookDetail) {
	bd, err := BookDetailed(<-IDch)
	if err != nil {
		log.Fatalf("error book detail: %s", err)
	}
	BDch <- bd
}
func BookDetailed(ID int64) (bookDetail BookDetail, err error) {
	detail_url := api_url + "book/" + strconv.Itoa(int(ID))
	resp, err := http.Get(detail_url)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("URL(%s) error: %s", detail_url, err)
		return BookDetail{}, err
	}

	if err != nil {
		log.Fatalf("error: %s", err)
		return BookDetail{}, err
	}

	defer resp.Body.Close()

	to_json := new(BookDetail)
	err = json.NewDecoder(resp.Body).Decode(to_json)
	if err != nil {
		log.Fatalf("error decode: %s", err)
		return BookDetail{}, err
	}

	bookDetail = *to_json
	return bookDetail, nil
}
