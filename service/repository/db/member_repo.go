package db

import (
	"library/service/models"

	"github.com/jmoiron/sqlx"
)

type MemberRepoImpl struct {
	DB *sqlx.DB
}

func NewMemberRepo(db *sqlx.DB) *MemberRepoImpl {
	return &MemberRepoImpl{DB: db}
}

// CreateMember inserts into DB and returns the full created member
func (r *MemberRepoImpl) CreateMember(member *models.Member) (*models.Member, error) {
	rows, err := r.DB.NamedQuery(InsertMemberQuery, member)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newID int
	if rows.Next() {
		if err := rows.Scan(&newID); err != nil {
			return nil, err
		}
	}

	// Assign ID back to struct
	member.MemberID = newID

	return member, nil
}

// List all members
func (r *MemberRepoImpl) ListMembers() ([]models.Member, error) {
	var list []models.Member
	if err := r.DB.Select(&list, ListMembersQuery); err != nil {
		return nil, err
	}
	return list, nil
}

// Get member by ID
func (r *MemberRepoImpl) GetMemberByID(id int) (*models.Member, error) {
	var m models.Member
	if err := r.DB.Get(&m, GetMemberByIDQuery, id); err != nil {
		return nil, err
	}
	return &m, nil
}

// Put -update member
func (r *MemberRepoImpl) UpdateMember(member *models.Member) (*models.Member, error) {
	rows, err := r.DB.NamedQuery(UpdateMemberQuery, member)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var updatedId int
	if rows.Next() {
		if err := rows.Scan(&updatedId); err != nil {
			return nil, err
		}
	}
	return r.GetMemberByID(updatedId)
}

// Delete member
func (r *MemberRepoImpl) DeleteMember(id int) error {
	_, err := r.DB.Exec(DeleteMemberQuery, id)
	return err
}
