# ITBuk

A library implemeting [It-ebooks](http://it-ebooks.info/) API.


### Overview

Even the API is very small, I found it quite to simple to tinker with. Unfortunately the API has a limitation on how many requests per day you do. Below are the main points I covered in this implementation:

- [x] Get book info by topic
- [x] Go routines
- [x] Pagination


```
Search(topic string, page int) (books []BookDetail, err error)
```
```
BookDetailed(ID int64) (bookDetail BookDetail, err error)
```

### Installation

I'd recommend to install using Go ways, although you can always clone and hack around.
```
$ go get -u github.com/sgmac/itbuk
```
A small example can be found on the bin directory.

```
usage: itb -p [npag] <topic>
-p Number of pages per search, defaults to one.
```

### Todo

- Add a flag in order to download the links of each book.
- Definitely I should start doing testing, seriously.
- TravisCI

### License

MIT
