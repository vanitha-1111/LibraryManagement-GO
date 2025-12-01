package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"library/service/handler"
	"library/service/libhttp"
	"library/service/repository"
	dbrepo "library/service/repository/db"
)

func main() {

	dsn := "postgres://postgres:root@localhost:5432/LibraryManagement?sslmode=disable"

	// Connect to database
	db, err := dbrepo.NewDBConnection(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// ------------------------------------------
	// BOOK dependencies
	// ------------------------------------------
	var bookRepo repository.BookRepo = dbrepo.NewBookRepo(db)
	bookService := handler.NewBookService(bookRepo)
	bookHTTP := libhttp.NewBookHTTPHandler(bookService)

	// ------------------------------------------
	// MEMBER dependencies
	// ------------------------------------------
	var memberRepo repository.MemberRepo = dbrepo.NewMemberRepo(db)
	memberService := handler.NewMemberService(memberRepo)
	memberHTTP := libhttp.NewMemberHTTPHandler(memberService)

	//Borrow dependencies
	var borrowRepo repository.BorrowRepo = dbrepo.NewBorrowRep(db)
	borrowService := handler.NewBorrowService(borrowRepo)
	borrowHTTP := libhttp.NewBorrowHTTPHandler(borrowService)

	//BorrowDetail dependencies
	var borrowDetailRepo repository.BorrowDetailRepo = dbrepo.NewBorrowDetailRepo(db)
	borrowDetailService := handler.NewBorrowDetailService(borrowDetailRepo, borrowRepo, memberRepo, bookRepo)
	borrowDetailHTTP := libhttp.NewBorrowDetailHTTPHandler(borrowDetailService)

	//category
	categoryRepo := dbrepo.NewCategoryRepo(db)
	categoryService := handler.NewCategoryService(categoryRepo)
	categoryHTTP := libhttp.NewCategoryHTTPHandler(categoryService)

	//user
	var userRepo repository.UserRepo = dbrepo.NewUserRepo(db)
	userService := handler.NewUserService(userRepo)
	authHTTP := libhttp.NewAuthHTTPHandler(userService)

	// ------------------------------------------
	// SETUP ROUTES
	// ------------------------------------------
	r := gin.Default()
	libhttp.RegisterRoutes(r, bookHTTP, memberHTTP, borrowHTTP, borrowDetailHTTP, categoryHTTP, authHTTP)

	// ------------------------------------------
	// START SERVER
	// ------------------------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)

	fmt.Println("Server running at", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
