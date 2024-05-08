package empapp

import (
    "testing"

    "github.com/hashicorp/go-memdb"
)

func TestCRUD(t *testing.T) {
    srv := setup()
    testEmp := &Employee{
	Id: 0,
	Name: "Test",
	Position: "Tester1",
	Salary:0.0,
    }
    _, err := srv.Create(testEmp)
    if err != nil {
        t.Errorf("Create Error: %s", err)
    }
    _, err = srv.GetByID(testEmp.Id)
    if err != nil {
        t.Errorf("Get Error: %s", err)
    }
    testEmpUp := &Employee{Position:"Tester2"}
    _, err = srv.Update(testEmp.Id, testEmpUp)
    if err != nil {
        t.Errorf("Update Error: %s", err)
    }
    empList, err := srv.ListEmployees(0,5)
    if err != nil {
        t.Errorf("Get Error: %s", err)
    }
    if len(empList)!=1 {
	t.Errorf("Record count mismatch")
    }
    err = srv.Delete(testEmp.Id)
    if err != nil {
        t.Errorf("Delete Error: %s", err)
    }
}

func setup() EmployeeIface {
    dbSchema:= &memdb.DBSchema{
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
    db, err := memdb.NewMemDB(dbSchema)
	if err != nil {
		panic(err)
	}
    return NewEmployeeSvc(db)
}
