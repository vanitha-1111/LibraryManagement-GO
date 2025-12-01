-- sql/schema.sql
-- PostgreSQL schema for LibraryManagement

-- 1) category
CREATE TABLE IF NOT EXISTS category (
  category_id SERIAL PRIMARY KEY,
  classname VARCHAR(100) NOT NULL
);

-- 2) users
CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(200) NOT NULL,
  firstname VARCHAR(100),
  lastname VARCHAR(100),
  role VARCHAR(50) DEFAULT 'user'
);

-- 3) member
CREATE TABLE IF NOT EXISTS member (
  member_id SERIAL PRIMARY KEY,
  firstname VARCHAR(100) NOT NULL,
  lastname VARCHAR(100) NOT NULL,
  gender VARCHAR(10),
  address VARCHAR(255),
  contact VARCHAR(100),
  type VARCHAR(100),
  year_level VARCHAR(100),
  status VARCHAR(100) DEFAULT 'Active'
);

-- 4) book
CREATE TABLE IF NOT EXISTS book (
  book_id SERIAL PRIMARY KEY,
  book_title VARCHAR(200) NOT NULL,
  category_id INT NOT NULL,
  author VARCHAR(100),
  book_copies INT NOT NULL DEFAULT 0,
  book_pub VARCHAR(200),
  publisher_name VARCHAR(200),
  isbn VARCHAR(50),
  copyright_year INT,
  date_receive VARCHAR(50),
  date_added TIMESTAMP NOT NULL DEFAULT now(),
  status VARCHAR(50) DEFAULT 'New',
  CONSTRAINT fk_book_category FOREIGN KEY (category_id) REFERENCES category(category_id)
);

-- 5) borrow
CREATE TABLE IF NOT EXISTS borrow (
  borrow_id SERIAL PRIMARY KEY,
  member_id INT NOT NULL,
  date_borrow DATE NOT NULL,
  due_date DATE,
  CONSTRAINT fk_borrow_member FOREIGN KEY (member_id) REFERENCES member(member_id)
);

-- 6) borrowdetails
CREATE TABLE IF NOT EXISTS borrowdetails (
  borrow_details_id SERIAL PRIMARY KEY,
  book_id INT NOT NULL,
  borrow_id INT NOT NULL,
  borrow_status VARCHAR(50) DEFAULT 'pending',
  date_return TIMESTAMP,
  CONSTRAINT fk_bd_book FOREIGN KEY (book_id) REFERENCES book(book_id),
  CONSTRAINT fk_bd_borrow FOREIGN KEY (borrow_id) REFERENCES borrow(borrow_id)
);

-- Optional: a type table (if you used it)
CREATE TABLE IF NOT EXISTS type (
  id SERIAL PRIMARY KEY,
  borrowertype VARCHAR(100)
);

-- Indexes for performance (optional)
CREATE INDEX IF NOT EXISTS idx_book_category ON book(category_id);
CREATE INDEX IF NOT EXISTS idx_borrow_member ON borrow(member_id);
CREATE INDEX IF NOT EXISTS idx_bd_borrow ON borrowdetails(borrow_id);

---seed data
-- sql/seed.sql (or appended to schema.sql)

-- Add categories
INSERT INTO category (classname) VALUES ('Science') ON CONFLICT DO NOTHING;
INSERT INTO category (classname) VALUES ('General') ON CONFLICT DO NOTHING;

-- Add admin user (plain password here for quick testing, hash in production)
INSERT INTO users (username, password, firstname, lastname, role)
VALUES ('admin','admin','System','Admin','admin') ON CONFLICT (username) DO NOTHING;

-- Add a couple of members
INSERT INTO member (firstname, lastname, gender, address, contact, type, year_level, status)
VALUES ('Mark','Sanchez','Male','Talisay','212010','Teacher','Faculty','Active') RETURNING member_id;

INSERT INTO member (firstname, lastname, gender, address, contact, type, year_level, status)
VALUES ('April','Aguilar','Female','E.B. Magalona','00','Student','Second Year','Active');

-- Add books
INSERT INTO book (book_title, category_id, author, book_copies, book_pub, publisher_name, isbn, copyright_year, date_added, status)
VALUES ('Physics Textbook', 1, 'Author A', 3, 'Pub A', 'Publisher A', 'ISBN-001', 2020, now(), 'New')
RETURNING book_id;

INSERT INTO book (book_title, category_id, author, book_copies, book_pub, publisher_name, isbn, copyright_year, date_added, status)
VALUES ('Science in our World', 1, 'Brian Knapp', 2, 'Regency', 'Prentice Hall', 'ISBN-002', 1996, now(), 'New');

-- Create a borrow transaction for member_id = 1 (adjust if RETURNING ids differ)
INSERT INTO borrow (member_id, date_borrow, due_date) VALUES (1, '2025-02-28', '2025-03-05') RETURNING borrow_id;
-- Suppose the returned borrow_id is 2 (verify on your DB)

-- Add borrow details (books under that borrow)
INSERT INTO borrowdetails (book_id, borrow_id, borrow_status) VALUES (1, 1, 'pending');
INSERT INTO borrowdetails (book_id, borrow_id, borrow_status, date_return) VALUES (2, 1, 'returned', now());
