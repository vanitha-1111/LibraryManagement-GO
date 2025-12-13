package db_test

import (
	"os"
	"sync"
	"testing"

	dbpkg "library/service/repository/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func setupDBReturn(t *testing.T) *sqlx.DB {
	dsn := os.Getenv("TEST_DB_URL")
	if dsn == "" {
		t.Fatal("TEST_DB_URL not set")
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		t.Fatalf("cannot connect: %v", err)
	}
	return db
}

func TestConcurrentReturn(t *testing.T) {
	db := setupDBReturn(t)
	repo := dbpkg.NewBorrowDetailRepo(db)

	// Step 1: Create book with 0 copies (just borrowed)
	var bookID int
	err := db.Get(&bookID, `
		INSERT INTO book (book_title, category_id, author, book_copies, book_pub, publisher_name, isbn, copyright_year, date_receive, date_added, status)
		VALUES ('Test Book', 1, 'Author', 0, 'pub', 'pub', 'isbn', 2024, NULL, now(), 'available')
		RETURNING book_id;
	`)
	if err != nil {
		t.Fatal(err)
	}

	// Step 2: Create member
	var memberID int
	err = db.Get(&memberID, `
		INSERT INTO member (firstname, lastname, gender, address, contact, type, year_level, status)
		VALUES ('John','Return','M','addr','000','student','1','active')
		RETURNING member_id;
	`)
	if err != nil {
		t.Fatal(err)
	}

	// Step 3: Create borrow header
	var borrowID int
	err = db.Get(&borrowID, `
		INSERT INTO borrow (member_id, date_borrow, due_date)
		VALUES ($1, now(), now() + interval '7 days')
		RETURNING borrow_id;
	`, memberID)
	if err != nil {
		t.Fatal(err)
	}

	// Step 4: Create ONE borrowdetail (it will be returned)
	var borrowDetailID int
	err = db.Get(&borrowDetailID, `
		INSERT INTO borrowdetails (book_id, borrow_id, borrow_status)
		VALUES ($1, $2, 'pending')
		RETURNING borrow_details_id;
	`, bookID, borrowID)
	if err != nil {
		t.Fatal(err)
	}

	// Step 5: Spawn 20 goroutines to return
	wg := sync.WaitGroup{}
	wg.Add(20)

	success := 0
	fail := 0
	mu := sync.Mutex{}

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()

			err := repo.ReturnBorrowDetail(borrowDetailID)

			mu.Lock()
			if err == nil {
				success++
			} else {
				fail++
			}
			mu.Unlock()
		}()
	}

	wg.Wait()

	t.Log("Success:", success, "Failures:", fail)

	// Expect ONLY ONE success
	if success != 1 {
		t.Fatalf("Expected 1 success, got %d", success)
	}

	// Verify book copies incremented exactly once
	var copies int
	_ = db.Get(&copies, "SELECT book_copies FROM book WHERE book_id=$1", bookID)

	if copies != 1 {
		t.Fatalf("Expected book_copies=1, got %d", copies)
	}
}
