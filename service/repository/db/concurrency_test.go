package db_test

import (
	"os"
	"sync"
	"testing"

	"library/service/models"
	dbpkg "library/service/repository/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	dsn := os.Getenv("TEST_DB_URL")
	if dsn == "" {
		t.Fatal("TEST_DB_URL not set")
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	return db
}

func TestConcurrentBorrow(t *testing.T) {
	db := setupTestDB(t)
	repo := dbpkg.NewBorrowDetailRepo(db)

	// Create test book with 1 copy
	var bookID int
	err := db.Get(&bookID, `
		INSERT INTO book (book_title, category_id, author, book_copies, book_pub, publisher_name, isbn, copyright_year, date_receive, date_added, status)
		VALUES ('Test Book', 1, 'Author', 1, 'pub', 'pub', 'isbn', 2024, NULL, now(), 'available')
		RETURNING book_id;
	`)
	if err != nil {
		t.Fatal(err)
	}
	//Create a member
	var memberID int
	err = db.Get(&memberID, `
		INSERT INTO member (firstname, lastname, gender, address, contact, type, year_level, status)
		VALUES ('John','Doe','M','A','123','student','1','active')
		RETURNING member_id;
	`)
	if err != nil {
		t.Fatal(err)
	}
	//create borrow record
	var borrowID int
	err = db.Get(&borrowID, `
		INSERT INTO borrow (member_id, date_borrow, due_date)
		VALUES ($1, now(), now() + interval '7 days')
		RETURNING borrow_id;
	`, memberID)
	if err != nil {
		t.Fatal(err)
	}

	//20 goroutines try to borrow same book
	wg := sync.WaitGroup{}
	wg.Add(20)

	success := 0
	fail := 0
	mu := sync.Mutex{}

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()

			_, err := repo.CreateBorrowDetail(&models.BorrowDetail{
				BorrowID: borrowID,
				BookID:   bookID,
			})
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
	if success != 1 {
		t.Fatalf("Expected only 1 success, got %d", success)
	}

}
