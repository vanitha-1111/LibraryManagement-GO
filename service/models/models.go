package models

import "time"

// Represents one row in the "book" table
type Book struct {
	BookID        int       `db:"book_id" json:"book_id"`
	BookTitle     string    `db:"book_title" json:"book_title"`
	CategoryID    int       `db:"category_id" json:"category_id"`
	Author        string    `db:"author" json:"author"`
	BookCopies    int       `db:"book_copies" json:"book_copies"`
	BookPub       string    `db:"book_pub" json:"book_pub"`
	PublisherName string    `db:"publisher_name" json:"publisher_name"`
	ISBN          string    `db:"isbn" json:"isbn"`
	CopyrightYear int       `db:"copyright_year" json:"copyright_year"`
	DateReceive   *string   `db:"date_receive" json:"date_receive"`
	DateAdded     time.Time `db:"date_added" json:"date_added"`
	Status        string    `db:"status" json:"status"`
}

// Represents one row in the "member" table
type Member struct {
	MemberID  int    `db:"member_id" json:"member_id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Gender    string `db:"gender" json:"gender"`
	Address   string `db:"address" json:"address"`
	Contact   string `db:"contact" json:"contact"`
	Type      string `db:"type" json:"type"`
	YearLevel string `db:"year_level" json:"year_level"`
	Status    string `db:"status" json:"status"`
}
type Borrow struct {
	BorrowID   int    `db:"borrow_id" json:"borrow_id"`
	MemberID   int    `db:"member_id" json:"member_id"`
	DateBorrow string `db:"date_borrow" json:"date_borrow"`
	DueDate    string `db:"due_date" json:"due_date"`
}
type BorrowDetail struct {
	BorrowDetailsID int     `db:"borrow_details_id" json:"borrow_details_id"`
	BookID          int     `db:"book_id" json:"book_id"`
	BorrowID        int     `db:"borrow_id" json:"borrow_id"`
	BorrowStatus    string  `db:"borrow_status" json:"borrow_status"`
	DateReturn      *string `db:"date_return" json:"date_return"`
	BookTitle       string  `db:"book_title" json:"book_title"`
}
type Category struct {
	CategoryID int    `db:"category_id" json:"category_id"`
	ClassName  string `db:"classname" json:"classname"`
}
type User struct {
	UserId    int    `db:"user_id" json:"user_id"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"` // (hashed later in Go)
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Role      string `db:"role" json:"role"`
}

type BorrowHistoryItem struct {
	BorrowID        int     `db:"borrow_id" json:"borrow_id"`
	DateBorrow      string  `db:"date_borrow" json:"date_borrow"`
	DueDate         string  `db:"due_date" json:"due_date"`
	BorrowDetailsID int     `db:"borrow_details_id" json:"borrow_details_id"`
	BookID          int     `db:"book_id" json:"book_id"`
	BookTitle       string  `db:"book_title" json:"book_title"`
	BorrowStatus    string  `db:"borrow_status" json:"borrow_status"`
	DateReturn      *string `db:"date_return" json:"date_return"`
}
