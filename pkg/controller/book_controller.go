package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/youthke/go-crud/pkg/domain/model"
	"github.com/youthke/go-crud/pkg/domain/service"
	"log"
	"strconv"
)

const errorHTML = "500.html"

type BookController interface {
	Create(ctx *gin.Context)
	Index(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(b service.BookService) BookController{
	return &bookController{bookService: b}
}

func(b *bookController) Index(ctx *gin.Context){
	books, err := b.bookService.FindAll()
	if err != nil {
		ctx.HTML(500,errorHTML,nil)
	}

	ctx.HTML(200, "book-index.html", gin.H{"books": books })
}

func(b *bookController) Create(ctx *gin.Context){
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")

	//intに変換
	price, err := strconv.Atoi(ctx.PostForm("price"))
	if err != nil {
		log.Println(err)
		ctx.HTML(500, errorHTML, nil)
		return
	}
	book := model.Book{Title: title, Price: price, Author: author}

	err = b.bookService.Create(book)
	if err != nil {
		log.Println(err)
		ctx.HTML(500, errorHTML, nil)
		return
	}

	ctx.Redirect(302, "/books")
}

func (b *bookController) Edit(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.HTML(500, errorHTML, nil)
		return
	}

	book, err := b.bookService.FindOne(id)

	if err != nil {
		ctx.HTML(500, errorHTML, nil)
		return
	}
	ctx.HTML(200, "book-edit.html", gin.H{"book": book})
}

func(b *bookController) Update(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.HTML(500, errorHTML, nil)
		return
	}

	title := ctx.PostForm("title")
	author := ctx.PostForm("author")


	//intに変換
	price, err := strconv.Atoi(ctx.PostForm("price"))
	if err != nil {
		log.Println(err)
		ctx.HTML(500, errorHTML, nil)
		return
	}


	book := model.Book{ID: id, Title: title, Price: price, Author: author}

	err = b.bookService.Update(book)
	if err != nil {
		log.Println(err)
		ctx.HTML(500, errorHTML, nil)
		return
	}

	ctx.Redirect(302, "/books")
}

func (b *bookController) Delete(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.HTML(500, errorHTML, nil)
		return
	}
	err = b.bookService.Delete(id)

	if err != nil {
		ctx.HTML(500 ,errorHTML,nil)
		return
	}

	ctx.Redirect(302,"/books")
}


