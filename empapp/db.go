package empapp

import (
	"github.com/hashicorp/go-memdb"
)

// Create the DB schema
func GetEmployeeSchema() *memdb.DBSchema{
empSchema := &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"employee": &memdb.TableSchema{
			Name: "employee",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "Id"},
				},
			},
		},
	},
}
return empSchema
}