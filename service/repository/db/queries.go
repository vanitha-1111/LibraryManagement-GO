package db

const InsertBookQuery = `
INSERT INTO book 
(book_title, category_id, author, book_copies, book_pub, publisher_name, isbn, copyright_year, date_receive, date_added, status)
VALUES 
(:book_title, :category_id, :author, :book_copies, :book_pub, :publisher_name, :isbn, :copyright_year, :date_receive, NOW(), :status)
RETURNING book_id;
`
const GetBookByIDQuery = `
SELECT * 
FROM book
WHERE book_id=$1;
`
const ListBooksQuery = `
SELECT *
FROM book
ORDER BY book_id;
`

const ListBooksByCategoryNameQuery = `
SELECT b.*
FROM book b
JOIN category c on b.category_id=c.category_id
WHERE (c.classname) = ($1)
ORDER BY b.book_id;
`

// Member queries
const InsertMemberQuery = `
INSERT INTO member
(firstname, lastname, gender, address, contact, type, year_level, status)
VALUES (:firstname, :lastname, :gender, :address, :contact, :type, :year_level, :status)
RETURNING member_id;
`
const GetMemberByIDQuery = `
SELECT *
FROM member
WHERE member_id=$1;
`
const ListMembersQuery = `
SELECT *
FROM member
ORDER BY member_id;
`

const UpdateMemberQuery = `
UPDATE member
SET firstname = :firstname,
    lastname = :lastname,
    gender = :gender,
    address = :address,
    contact = :contact,
    type = :type,
    year_level = :year_level,
    status = :status
WHERE member_id = :member_id
RETURNING member_id;
`
const DeleteMemberQuery = `
DELETE FROM member where member_id=$1;
`

//Borrow queries

const InsertBorrowQuery = `
INSERT INTO borrow
(member_id, date_borrow, due_date)
VALUES (:member_id, :date_borrow, :due_date)
RETURNING borrow_id;
`

const GetBorrowByIDQuery = `
SELECT *
FROM borrow
WHERE borrow_id = $1;
`

const ListBorrowsQuery = `
SELECT *
FROM borrow
ORDER BY borrow_id;
`

// BorrowDetails
const InsertBorrowDetailQuery = `
INSERT INTO borrowdetails
(book_id, borrow_id, borrow_status, date_return)
VALUES (:book_id, :borrow_id, :borrow_status, :date_return)
RETURNING borrow_details_id;
`

const GetBorrowDetailsByBorrowIDQuery = `
SELECT bd.borrow_details_id, bd.book_id, bd.borrow_id, bd.borrow_status, bd.date_return, b.book_title
FROM borrowdetails bd
LEFT JOIN book b ON bd.book_id = b.book_id
WHERE bd.borrow_id = $1
ORDER BY bd.borrow_details_id;
`
const GetBorrowDetailByIDQuery = `
SELECT * FROM borrowdetails WHERE borrow_details_id = $1;
`
const UpdateBorrowDetailStatusQuery = `
UPDATE borrowdetails
SET borrow_status = $1, date_return = $2
WHERE borrow_details_id = $3;
`
const DecrementBookCopiesQuery = `
UPDATE book
SET book_copies=book_copies-1
where book_id=$1 AND book_copies>0
RETURNING book_copies;
`
const IncrementBookCopiesQuery = `
UPDATE book
SET book_copies = book_copies + 1
WHERE book_id = $1
RETURNING book_copies;
`

// Category queries
const ListCategoriesQuery = `
SELECT *
FROM category
ORDER BY category_id;
`
const InsertCategoryQuery = `
INSERT INTO category(classname)
values($1)
RETURNING category_id;
`
const DeleteCategoryQuery = `
DELETE FROM category
WHERE category_id=$1;
`

// USER queries
const GetUserByUsernameQuery = `
SELECT *
FROM users
WHERE username = $1;
`
const InsertUserQuery = `
INSERT INTO users (username, password, firstname, lastname, role)
        VALUES (:username, :password, :firstname, :lastname, :role)
        RETURNING user_id;
`
const MemberBorrowHistoryQuery = `
SELECT 
    b.borrow_id,
    b.date_borrow,
    b.due_date,
    bd.borrow_details_id,
    bd.book_id,
    bd.borrow_status,
    bd.date_return,
    bk.book_title
FROM borrow b
JOIN borrowdetails bd ON b.borrow_id = bd.borrow_id
JOIN book bk ON bd.book_id = bk.book_id
WHERE b.member_id = $1
ORDER BY b.borrow_id;
`
