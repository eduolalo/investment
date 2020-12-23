package structs

import (
	"invest/redis"
	"log"
	"strconv"
	"sync"
)

// Statistic estructura de las respuestas de statics
type Statistic struct {
	Success           int     `json:"successful_assignments"`
	AverageFailInvest float32 `json:"average_fail_invest"`
	Fail              int     `json:"failed_assignments"`
	Total             int     `json:"total_assignments"`
	AverageInvest     float32 `json:"average_invest"`
	OkInvest          int
	FailInvest        int
}

// Calculate obtiene los datos de las inversiones hechas
func (s *Statistic) Calculate() {

	var wg sync.WaitGroup
	// sección asíncrona
	{
		wg.Add(4)
		go s.getFailInvest(&wg)
		go s.getOkInvest(&wg)
		go s.getSuccess(&wg)
		go s.getFail(&wg)

		wg.Wait()
	}
	s.Total = s.Fail + s.Success
	if s.Success != 0 && s.OkInvest != 0 {

		s.AverageInvest = float32(s.OkInvest) / float32(s.Success)
	}
	if s.Fail != 0 && s.FailInvest != 0 {

		s.AverageFailInvest = float32(s.FailInvest) / float32(s.Fail)
	}
}

// getSuccess obtiene las asignaciones exitosas
func (s *Statistic) getSuccess(wg *sync.WaitGroup) {

	defer wg.Done()

	res := redis.Get("success")
	if res == "" {
		return
	}
	success, err := strconv.Atoi(res)
	if err != nil {

		log.Println("*** structs.Statistic.GetSuccess ***")
		log.Println(err.Error())
		log.Println("*** structs.Statistic.GetSuccess ***")
		return
	}
	s.Success = success
}

// getFail obtiene las asignaciones fallidas
func (s *Statistic) getFail(wg *sync.WaitGroup) {

	defer wg.Done()

	res := redis.Get("fail")
	if res == "" {
		return
	}
	fail, err := strconv.Atoi(res)
	if err != nil {

		log.Println("*** structs.Statistic.GetFail ***")
		log.Println(err.Error())
		log.Println("*** structs.Statistic.GetFail ***")
		return
	}
	s.Fail = fail
}

// getOkInvest obtiene las asignaciones exitosas
func (s *Statistic) getOkInvest(wg *sync.WaitGroup) {

	defer wg.Done()

	res := redis.Get("success-investment")
	if res == "" {
		return
	}
	success, err := strconv.Atoi(res)
	if err != nil {

		log.Println("*** structs.Statistic.GetOkInvest ***")
		log.Println(err.Error())
		log.Println("*** structs.Statistic.GetOkInvest ***")
		return
	}
	s.OkInvest = success
}

// getFailInvest obtiene las asignaciones exitosas
func (s *Statistic) getFailInvest(wg *sync.WaitGroup) {

	defer wg.Done()

	res := redis.Get("fail-investment")
	if res == "" {
		return
	}
	fail, err := strconv.Atoi(res)
	if err != nil {

		log.Println("*** structs.Statistic.GetFailInvest ***")
		log.Println(err.Error())
		log.Println("*** structs.Statistic.GetFailInvest ***")
		return
	}
	s.FailInvest = fail
}
