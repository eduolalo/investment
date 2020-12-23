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
