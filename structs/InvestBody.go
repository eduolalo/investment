package structs

import (
	"log"
)

// InvestBody estructura que representa al cuepo del request
type InvestBody struct {
	Investment int32 `json:"investment"`
}

// Unmarshal obtiene los datos del body y los mete en la estructura
func (s *InvestBody) Unmarshal(body []byte) {

	if err := json.Unmarshal(body, s); err != nil {

		log.Println("*** structs.InvestBody.Unmarshal ***")
		log.Println(err.Error())
		log.Println("--- structs.InvestBody.Unmarshal ---")
	}
}

// String regresa el string del monto la inversión
func (s InvestBody) String() string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(s.Investment)

	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {

			return string(buf[pos:])
		}
	}
}
