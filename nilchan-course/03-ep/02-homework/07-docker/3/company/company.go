package company

import (
	"maps"
	"sync"
)

type Company struct {
	Employees map[int]Employee
	ID        int
	mtx       sync.RWMutex
}

func NewCompany() *Company {
	return &Company{
		Employees: make(map[int]Employee),
		ID:        0,
	}
}

func (c *Company) CreateEmployee(fullName, position string) Employee {
	employee := NewEmployee(fullName, position)

	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.Employees[c.ID] = employee
	c.ID++

	return employee
}

func (c *Company) GetEmployees() map[int]Employee {
	tmp := make(map[int]Employee)

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	maps.Copy(tmp, c.Employees)

	return tmp
}

func (c *Company) DeleteEmployee(id int) error {
	if _, ok := c.Employees[id]; !ok {
		return ErrNotFound
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.Employees, id)

	return nil
}
