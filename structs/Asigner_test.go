package structs

import "testing"

func TestAsigner(t *testing.T) {

	var asigner Assigner

	b3, _, _, err := asigner.Assign(300)
	if b3 != 1 || err != nil {

		t.Log("Se esperaba b3 == 1 y err == nil, se obtuvo b3 = ", b3, " y err = ", err.Error())
	}

	_, b5, _, err := asigner.Assign(500)
	if b5 != 1 || err != nil {

		t.Log("Se esperaba b5 == 1 y err == nil, se obtuvo b5 = ", b5, " y err = ", err.Error())
	}

	_, _, b7, err := asigner.Assign(700)
	if b7 != 1 || err != nil {

		t.Log("Se esperaba b7 == 1 y err == nil, se obtuvo b7 = ", b7, " y err = ", err.Error())
	}

	b3, b5, b7, err = asigner.Assign(700)
	if b3 != 1 && b7 != 1 || err != nil {

		t.Log("Se esperaba b7 == 1, b3 == 1 y err == nil, se obtuvo b7 = ", b7, ", b3 = ", b3, ", b5 = ", b5, "  y err = ", err.Error())
	}

	_, _, _, err = asigner.Assign(3000)
	if err != nil {

		t.Log("Se esperaba err == nil, se obtuvo err = ", err.Error())
	}

	_, _, _, err = asigner.Assign(6700)
	if err != nil {

		t.Log("Se esperaba err == nil, se obtuvo err = ", err.Error())
	}

}
