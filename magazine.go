package magazine

//Address type contians information about the address.
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

//Subscriber type contians information about the subscriber.
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

//Employee type contians information about the employee.
type Employee struct {
	Name   string
	Salary float64
	Address
}
