// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DigestsColumns holds the columns for the "digests" table.
	DigestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "subject_digests", Type: field.TypeInt, Nullable: true},
	}
	// DigestsTable holds the schema information for the "digests" table.
	DigestsTable = &schema.Table{
		Name:       "digests",
		Columns:    DigestsColumns,
		PrimaryKey: []*schema.Column{DigestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "digests_subjects_digests",
				Columns:    []*schema.Column{DigestsColumns[3]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DssesColumns holds the columns for the "dsses" table.
	DssesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gitbom_sha256", Type: field.TypeString, Unique: true},
		{Name: "payload_type", Type: field.TypeString},
		{Name: "dsse_statement", Type: field.TypeInt, Nullable: true},
	}
	// DssesTable holds the schema information for the "dsses" table.
	DssesTable = &schema.Table{
		Name:       "dsses",
		Columns:    DssesColumns,
		PrimaryKey: []*schema.Column{DssesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dsses_statements_statement",
				Columns:    []*schema.Column{DssesColumns[3]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DsseSignaturesColumns holds the columns for the "dsse_signatures" table.
	DsseSignaturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// DsseSignaturesTable holds the schema information for the "dsse_signatures" table.
	DsseSignaturesTable = &schema.Table{
		Name:       "dsse_signatures",
		Columns:    DsseSignaturesColumns,
		PrimaryKey: []*schema.Column{DsseSignaturesColumns[0]},
	}
	// StatementsColumns holds the columns for the "statements" table.
	StatementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// StatementsTable holds the schema information for the "statements" table.
	StatementsTable = &schema.Table{
		Name:       "statements",
		Columns:    StatementsColumns,
		PrimaryKey: []*schema.Column{StatementsColumns[0]},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
	}
	// StatementSubjectsColumns holds the columns for the "statement_subjects" table.
	StatementSubjectsColumns = []*schema.Column{
		{Name: "statement_id", Type: field.TypeInt},
		{Name: "subject_id", Type: field.TypeInt},
	}
	// StatementSubjectsTable holds the schema information for the "statement_subjects" table.
	StatementSubjectsTable = &schema.Table{
		Name:       "statement_subjects",
		Columns:    StatementSubjectsColumns,
		PrimaryKey: []*schema.Column{StatementSubjectsColumns[0], StatementSubjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "statement_subjects_statement_id",
				Columns:    []*schema.Column{StatementSubjectsColumns[0]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "statement_subjects_subject_id",
				Columns:    []*schema.Column{StatementSubjectsColumns[1]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DigestsTable,
		DssesTable,
		DsseSignaturesTable,
		StatementsTable,
		SubjectsTable,
		StatementSubjectsTable,
	}
)

func init() {
	DigestsTable.ForeignKeys[0].RefTable = SubjectsTable
	DssesTable.ForeignKeys[0].RefTable = StatementsTable
	StatementSubjectsTable.ForeignKeys[0].RefTable = StatementsTable
	StatementSubjectsTable.ForeignKeys[1].RefTable = SubjectsTable
}
