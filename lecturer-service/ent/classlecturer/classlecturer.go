// Code generated by ent, DO NOT EDIT.

package classlecturer

const (
	// Label holds the string label denoting the classlecturer type in the database.
	Label = "class_lecturer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeLecturer holds the string denoting the lecturer edge name in mutations.
	EdgeLecturer = "lecturer"
	// Table holds the table name of the classlecturer in the database.
	Table = "class_lecturers"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "class_lecturers"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "class_class_lecturers"
	// LecturerTable is the table that holds the lecturer relation/edge.
	LecturerTable = "class_lecturers"
	// LecturerInverseTable is the table name for the Lecturer entity.
	// It exists in this package in order to avoid circular dependency with the "lecturer" package.
	LecturerInverseTable = "lecturers"
	// LecturerColumn is the table column denoting the lecturer relation/edge.
	LecturerColumn = "lecturer_class_lecturers"
)

// Columns holds all SQL columns for classlecturer fields.
var Columns = []string{
	FieldID,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "class_lecturers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_class_lecturers",
	"lecturer_class_lecturers",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}