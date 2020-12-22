package structs

// CreditAsigner Interface para asignar créditos
type CreditAsigner interface {
	Assign(invesment int32) (borrow3, borrow5, borrow7 int32, err error)
}
