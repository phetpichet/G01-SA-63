// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// DispensesColumns holds the columns for the "dispenses" table.
	DispensesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "note", Type: field.TypeString},
		{Name: "drug_id", Type: field.TypeInt, Nullable: true},
	}
	// DispensesTable holds the schema information for the "dispenses" table.
	DispensesTable = &schema.Table{
		Name:       "dispenses",
		Columns:    DispensesColumns,
		PrimaryKey: []*schema.Column{DispensesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dispenses_drugs_dispenses",
				Columns: []*schema.Column{DispensesColumns[2]},

				RefColumns: []*schema.Column{DrugsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DrugsColumns holds the columns for the "drugs" table.
	DrugsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "drug_type", Type: field.TypeString},
		{Name: "strength", Type: field.TypeString, Default: "unknown"},
		{Name: "information", Type: field.TypeString, Default: "unknown"},
		{Name: "dispense_id", Type: field.TypeInt, Nullable: true},
		{Name: "form_id", Type: field.TypeInt, Nullable: true},
		{Name: "unit_id", Type: field.TypeInt, Nullable: true},
		{Name: "volume_id", Type: field.TypeInt, Nullable: true},
	}
	// DrugsTable holds the schema information for the "drugs" table.
	DrugsTable = &schema.Table{
		Name:       "drugs",
		Columns:    DrugsColumns,
		PrimaryKey: []*schema.Column{DrugsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "drugs_dispenses_drugs",
				Columns: []*schema.Column{DrugsColumns[4]},

				RefColumns: []*schema.Column{DispensesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drugs_forms_drugs",
				Columns: []*schema.Column{DrugsColumns[5]},

				RefColumns: []*schema.Column{FormsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drugs_units_drugs",
				Columns: []*schema.Column{DrugsColumns[6]},

				RefColumns: []*schema.Column{UnitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drugs_volumes_drugs",
				Columns: []*schema.Column{DrugsColumns[7]},

				RefColumns: []*schema.Column{VolumesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FormsColumns holds the columns for the "forms" table.
	FormsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "form_type", Type: field.TypeString},
	}
	// FormsTable holds the schema information for the "forms" table.
	FormsTable = &schema.Table{
		Name:        "forms",
		Columns:     FormsColumns,
		PrimaryKey:  []*schema.Column{FormsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "position", Type: field.TypeString},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:        "positions",
		Columns:     PositionsColumns,
		PrimaryKey:  []*schema.Column{PositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TitlesColumns holds the columns for the "titles" table.
	TitlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
	}
	// TitlesTable holds the schema information for the "titles" table.
	TitlesTable = &schema.Table{
		Name:        "titles",
		Columns:     TitlesColumns,
		PrimaryKey:  []*schema.Column{TitlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UnitsColumns holds the columns for the "units" table.
	UnitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "unit_type", Type: field.TypeString},
	}
	// UnitsTable holds the schema information for the "units" table.
	UnitsTable = &schema.Table{
		Name:        "units",
		Columns:     UnitsColumns,
		PrimaryKey:  []*schema.Column{UnitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "dispense_id", Type: field.TypeInt, Nullable: true},
		{Name: "drug_id", Type: field.TypeInt, Nullable: true},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "position_id", Type: field.TypeInt, Nullable: true},
		{Name: "title_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_dispenses_users",
				Columns: []*schema.Column{UsersColumns[4]},

				RefColumns: []*schema.Column{DispensesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_drugs_users",
				Columns: []*schema.Column{UsersColumns[5]},

				RefColumns: []*schema.Column{DrugsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_genders_users",
				Columns: []*schema.Column{UsersColumns[6]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_positions_users",
				Columns: []*schema.Column{UsersColumns[7]},

				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_titles_users",
				Columns: []*schema.Column{UsersColumns[8]},

				RefColumns: []*schema.Column{TitlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VolumesColumns holds the columns for the "volumes" table.
	VolumesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "volume_type", Type: field.TypeString},
	}
	// VolumesTable holds the schema information for the "volumes" table.
	VolumesTable = &schema.Table{
		Name:        "volumes",
		Columns:     VolumesColumns,
		PrimaryKey:  []*schema.Column{VolumesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DispensesTable,
		DrugsTable,
		FormsTable,
		GendersTable,
		PositionsTable,
		TitlesTable,
		UnitsTable,
		UsersTable,
		VolumesTable,
	}
)

func init() {
	DispensesTable.ForeignKeys[0].RefTable = DrugsTable
	DrugsTable.ForeignKeys[0].RefTable = DispensesTable
	DrugsTable.ForeignKeys[1].RefTable = FormsTable
	DrugsTable.ForeignKeys[2].RefTable = UnitsTable
	DrugsTable.ForeignKeys[3].RefTable = VolumesTable
	UsersTable.ForeignKeys[0].RefTable = DispensesTable
	UsersTable.ForeignKeys[1].RefTable = DrugsTable
	UsersTable.ForeignKeys[2].RefTable = GendersTable
	UsersTable.ForeignKeys[3].RefTable = PositionsTable
	UsersTable.ForeignKeys[4].RefTable = TitlesTable
}
