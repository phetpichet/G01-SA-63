// Code generated by entc, DO NOT EDIT.

package drug

const (
	// Label holds the string label denoting the drug type in the database.
	Label = "drug"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDrugType holds the string denoting the drugtype field in the database.
	FieldDrugType = "drug_type"
	// FieldStrength holds the string denoting the strength field in the database.
	FieldStrength = "strength"
	// FieldInformation holds the string denoting the information field in the database.
	FieldInformation = "information"

	// EdgeForm holds the string denoting the form edge name in mutations.
	EdgeForm = "form"
	// EdgeUnit holds the string denoting the unit edge name in mutations.
	EdgeUnit = "unit"
	// EdgeVolume holds the string denoting the volume edge name in mutations.
	EdgeVolume = "volume"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeDispenses holds the string denoting the dispenses edge name in mutations.
	EdgeDispenses = "dispenses"

	// Table holds the table name of the drug in the database.
	Table = "drugs"
	// FormTable is the table the holds the form relation/edge.
	FormTable = "drugs"
	// FormInverseTable is the table name for the Form entity.
	// It exists in this package in order to avoid circular dependency with the "form" package.
	FormInverseTable = "forms"
	// FormColumn is the table column denoting the form relation/edge.
	FormColumn = "form_id"
	// UnitTable is the table the holds the unit relation/edge.
	UnitTable = "drugs"
	// UnitInverseTable is the table name for the Unit entity.
	// It exists in this package in order to avoid circular dependency with the "unit" package.
	UnitInverseTable = "units"
	// UnitColumn is the table column denoting the unit relation/edge.
	UnitColumn = "unit_id"
	// VolumeTable is the table the holds the volume relation/edge.
	VolumeTable = "drugs"
	// VolumeInverseTable is the table name for the Volume entity.
	// It exists in this package in order to avoid circular dependency with the "volume" package.
	VolumeInverseTable = "volumes"
	// VolumeColumn is the table column denoting the volume relation/edge.
	VolumeColumn = "volume_id"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "drug_id"
	// DispensesTable is the table the holds the dispenses relation/edge.
	DispensesTable = "dispenses"
	// DispensesInverseTable is the table name for the Dispense entity.
	// It exists in this package in order to avoid circular dependency with the "dispense" package.
	DispensesInverseTable = "dispenses"
	// DispensesColumn is the table column denoting the dispenses relation/edge.
	DispensesColumn = "drug_id"
)

// Columns holds all SQL columns for drug fields.
var Columns = []string{
	FieldID,
	FieldDrugType,
	FieldStrength,
	FieldInformation,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Drug type.
var ForeignKeys = []string{
	"dispense_id",
	"form_id",
	"unit_id",
	"volume_id",
}

var (
	// DrugTypeValidator is a validator for the "DrugType" field. It is called by the builders before save.
	DrugTypeValidator func(string) error
	// DefaultStrength holds the default value on creation for the Strength field.
	DefaultStrength string
	// DefaultInformation holds the default value on creation for the Information field.
	DefaultInformation string
)