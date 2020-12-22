package structs

import (
	"errors"
)

// Assigner implementa métodos para asignar créditos según la inversión recibida
type Assigner struct {
	investment int32
	borrow3    int32
	borrow5    int32
	borrow7    int32
}

// Assign Regresa una propuesta de asignación de créditos según el monto de inversión
func (a *Assigner) Assign(investment int32) (borrow3, borrow5, borrow7 int32, err error) {

	if xix := investment % 100; xix != 0 {

		err = errors.New("La inversión debe ser en múltiplos de 100")
		return
	}

	switch i := investment; {
	case i < 300:
		err = errors.New("No se puede asignar crédito alguno")
		return
	case i == 300:
		borrow3 = 1
		return
	case i == 500:
		borrow5 = 1
		return
	case i == 700:
		borrow7 = 1
		return
	}

	a.reset(investment)
	a.assignFor()
	if a.investment == 0 {
		borrow3 = a.borrow3
		borrow5 = a.borrow5
		borrow7 = a.borrow7
		return
	}
	a.reset(investment)
	a.assignModulus()
	if a.investment == 0 {
		borrow3 = a.borrow3
		borrow5 = a.borrow5
		borrow7 = a.borrow7
		return
	}
	err = errors.New("No es posible asignar esta inversión a los montos asignados")
	return
}

// assignFor Intenta hacer asignación mediante un ciclo para repartir los recursos en
// distintos tipos de préstamo
func (a *Assigner) assignFor() {

	for a.investment >= 300 {

		if a.investment >= 700 {

			a.investment -= 700
			a.borrow7++
		}

		if a.investment >= 500 {

			a.investment -= 500
			a.borrow5++
		}

		if a.investment >= 300 {

			a.investment -= 300
			a.borrow3++
		}
	}
}

// assignModulus intenta asignar los recursos probando si la inversión es múltiple
// de los montos autorizados
func (a *Assigner) assignModulus() {

	if mod := a.investment % 700; mod == 0 {
		a.borrow7 = a.investment / 700
		a.investment = 0
		return
	}
	if mod := a.investment % 500; mod == 0 {
		a.borrow5 = a.investment / 500
		a.investment = 0
		return
	}
	if mod := a.investment % 300; mod == 0 {
		a.borrow3 = a.investment / 300
		a.investment = 0
	}
}

// reset regresa los borrow a su estado inicial
func (a *Assigner) reset(investment int32) {
	a.investment = investment
	a.borrow7 = 0
	a.borrow5 = 0
	a.borrow3 = 0
}
