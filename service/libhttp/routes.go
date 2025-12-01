package libhttp

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.Engine,
	bookHandler *BookHTTPHandler,
	memberHandler *MemberHTTPHandler,
	borrowHandler *BorrowHTTPHandler,
	borrowDetailHandler *BorrowDetailHTTPHandler,
	categoryHandler *CategoryHTTPHandler,
	authHandler *AuthHTTPHandler,
) {

	//----AUTH -----
	r.POST("/auth/login", authHandler.Login)

	auth := r.Group("/")
	auth.Use(AuthMiddleware())

	// ------------------------
	// BOOK ROUTES
	// ------------------------
	//r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.ListBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.GET("/books/category/:name", bookHandler.GetBooksByCategoryName)

	admin := r.Group("/")
	admin.Use(AuthMiddleware(), AdminOnly())
	admin.POST("/books", bookHandler.CreateBook)

	// ------------------------
	// MEMBER ROUTES
	// ------------------------
	r.POST("/members", memberHandler.CreateMember)
	r.GET("/members", memberHandler.ListMembers)
	r.GET("/members/:id", memberHandler.GetMemberByID)
	r.PUT("/members/:id", memberHandler.UpdateMember)

	admin.DELETE("/members/:id", memberHandler.DeleteMember)

	// ------------------------
	// Borrow ROUTES
	// ------------------------
	r.POST("/borrow", borrowHandler.CreateBorrow)
	r.GET("/borrow", borrowHandler.ListBorrows)
	r.GET("/borrow/:id", borrowHandler.GetBorrowByID)

	// ------------------------
	// BorrowDetail ROUTES
	// ------------------------
	r.POST("/borrowdetails", borrowDetailHandler.CreateBorrowDetail)
	r.GET("/borrow/:id/details", borrowDetailHandler.GetBorrowDetailsByBorrowID)
	r.PUT("/borrowdetails/:id/return", borrowDetailHandler.ReturnBorrowDetail)
	r.GET("/members/:id/history", borrowDetailHandler.GetMemberBorrowHistory)

	//---------------------
	//Category Routes
	//----------------------
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories", categoryHandler.GetAllCategories)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)

}
