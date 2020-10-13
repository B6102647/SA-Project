// Code generated by entc, DO NOT EDIT.

package bookborrow

const (
	// Label holds the string label denoting the bookborrow type in the database.
	Label = "book_borrow"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldADDEDTIME holds the string denoting the added_time field in the database.
	FieldADDEDTIME = "added_time"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "Owner"
	// EdgeBOOK holds the string denoting the book edge name in mutations.
	EdgeBOOK = "BOOK"
	// EdgePURPOSE holds the string denoting the purpose edge name in mutations.
	EdgePURPOSE = "PURPOSE"

	// Table holds the table name of the bookborrow in the database.
	Table = "book_borrows"
	// OwnerTable is the table the holds the Owner relation/edge.
	OwnerTable = "book_borrows"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the Owner relation/edge.
	OwnerColumn = "user_booklist"
	// BOOKTable is the table the holds the BOOK relation/edge.
	BOOKTable = "book_borrows"
	// BOOKInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BOOKInverseTable = "books"
	// BOOKColumn is the table column denoting the BOOK relation/edge.
	BOOKColumn = "book_booklist"
	// PURPOSETable is the table the holds the PURPOSE relation/edge.
	PURPOSETable = "book_borrows"
	// PURPOSEInverseTable is the table name for the Purpose entity.
	// It exists in this package in order to avoid circular dependency with the "purpose" package.
	PURPOSEInverseTable = "purposes"
	// PURPOSEColumn is the table column denoting the PURPOSE relation/edge.
	PURPOSEColumn = "purpose_booklist"
)

// Columns holds all SQL columns for bookborrow fields.
var Columns = []string{
	FieldID,
	FieldADDEDTIME,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the BookBorrow type.
var ForeignKeys = []string{
	"book_booklist",
	"purpose_booklist",
	"user_booklist",
}
