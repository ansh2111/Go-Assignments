package empapp

import (
	"errors"
	"log"
	"github.com/hashicorp/go-memdb"
)

// EmployeeIface provides operations on employees
type EmployeeIface interface {
	Create(*Employee) (*Employee, error)
	GetByID(int) (*Employee, error)
	Update(int, *Employee) (*Employee, error)
	Delete(int) error
	ListEmployees(int, int) ([]*Employee, error) 
}

// EmployeeSvc is a concrete implementation of EmployeeIface 
type EmployeeSvc struct{
	db *memdb.MemDB
}

func NewEmployeeSvc(db *memdb.MemDB) EmployeeIface {
	return &EmployeeSvc{db:db}
}

func (esvc *EmployeeSvc) Create(e *Employee) (*Employee, error) {
	log.Println("create")
	txn := esvc.db.Txn(true)
	if err := txn.Insert("employee", e); err != nil {
		return nil, err
	}
	txn.Commit()
	return e, nil
}

func (esvc *EmployeeSvc) GetByID(id int) (*Employee, error) {
	log.Println("getbyid")
	txn := esvc.db.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("employee", "id", id)
	if err != nil {
		return nil, err
	}
	if raw==nil{
		return nil, errors.New("404 not found")
	}
	eres:= raw.(*Employee)
	return eres, nil
}

func (esvc *EmployeeSvc) Update(id int, e *Employee) (*Employee, error) {
	log.Println("update")
	eOld, err := esvc.GetByID(id)
	if err!=nil{
		return nil, err
	}
	txn := esvc.db.Txn(true)
	esvc.formatEntityData(eOld, e)
	if err := txn.Insert("employee", e); err != nil {
		return nil, err
	}
	txn.Commit()
	return e, nil
}

func (esvc *EmployeeSvc) formatEntityData(eOld, eNew *Employee) {
	eNew.Id = eOld.Id
	if eNew.Name=="" {
		eNew.Name=eOld.Name
	}
	if eNew.Position == "" {
		eNew.Position = eOld.Position
	}
	if eNew.Salary == 0.0 {
		eNew.Salary = eOld.Salary
	}
}

func (esvc *EmployeeSvc) Delete(id int) error {
	log.Println("delete")
	e, err := esvc.GetByID(id)
	if err!=nil{
		return err
	}
	txn := esvc.db.Txn(true)
	err = txn.Delete("employee", e)
	log.Println(err)
	txn.Commit()
	return err
}

func (esvc *EmployeeSvc) ListEmployees(page int, limit int) ([]*Employee, error) {
	log.Println("list")
	txn := esvc.db.Txn(false)
	it, err := txn.Get("employee", "id")
	if err!=nil{
		return nil, errors.New("404 no records found")
	}
	offset := page*5
	ctr := 0
	obj := it.Next()
	for ctr=0; ctr < offset ; ctr++ {
		if obj == nil{
			break;
		}
		obj = it.Next()
	}
	if obj==nil {
		return nil, errors.New("404 no records found for given pagination variables")
	}
	if limit==0 || limit>5{
		limit=5;
	}
	res:= make([]*Employee,0)
	for ctr=0;obj!=nil && ctr < limit ; ctr++ {
		res = append(res, obj.(*Employee))
		obj = it.Next()
	}
	txn.Abort()
	return res, err
}
