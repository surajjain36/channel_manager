package web

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/surajjain36/channel_manager/misc"
	"github.com/surajjain36/channel_manager/util"
)

var helper = util.Helper{}

//worker ...
func (s *Service) worker(chData misc.ChannelData) {
	for {
		select {
		case <-chData.Ch:
			return
		default:
			if *chData.IsActive {
				*chData.Counter += chData.Start
				time.Sleep(time.Second * time.Duration(chData.StepTime))
			}
		}
	}

}

//CreateRoutine ...
//Author: Suraj
//date: 05/03/2020
func (s *Service) CreateRoutine(c *gin.Context) {
	statusCode := http.StatusBadRequest
	res := gin.H{"message": "Something went wrong", "data": nil}
	start, _ := strconv.ParseInt(c.DefaultQuery("start", "0"), 10, 64)
	step, _ := strconv.Atoi(c.DefaultQuery("step", "0"))

	if start != 0 && step != 0 {
		var data = make(map[string]interface{})
		var chData misc.ChannelData
		chData.Start = int(start)
		chData.StepTime = step
		chData.ID = helper.GenerateRandomString(8)
		chData.Ch = make(chan string)
		chData.IsActive = new(bool)
		*chData.IsActive = true
		chData.CreatedAt = time.Now()
		chData.Counter = new(int)
		st := "active"
		chData.Status = &st
		go s.worker(chData)

		s.trackChannel = append(s.trackChannel, chData)
		data["ChannelID"] = chData.ID
		res["data"] = data
		res["message"] = "Gorotine created successfully"
		statusCode = http.StatusOK
	}

	s.responseWriter(c, res, statusCode)
	return
}

//CheckRoutine ...
//Author: Suraj
//date: 05/03/2020
func (s *Service) CheckRoutine(c *gin.Context) {
	statusCode := http.StatusBadRequest
	res := gin.H{"message": "Something went wrong", "data": nil}
	rotineID := c.DefaultQuery("id", "")
	//Find gorutine
	if rotineID != "" {
		var chkData misc.CheckData
		//Can be used any other searching algorithoms(like binary search) to make search faster
		for _, chData := range s.trackChannel {
			if chData.ID == rotineID {
				chkData.ChannelID = chData.ID
				chkData.CurrentCounter = *chData.Counter
				chkData.CreatedAt = chData.CreatedAt
				chkData.StepTime = chData.StepTime
				chkData.Status = *chData.Status
				res["message"] = "Goroutine found successfully"
				statusCode = http.StatusOK
				break
			}
		}
		res["data"] = chkData
		if chkData.ChannelID == "" {
			res["message"] = "Goroutine not found."
		}
	} else { //List all the gorutines. Can be paginated.
		var chkDataList []misc.CheckData
		if len(s.trackChannel) != 0 {
			for _, chData := range s.trackChannel {
				var chkData misc.CheckData
				chkData.ChannelID = chData.ID
				chkData.CurrentCounter = *chData.Counter
				chkData.CreatedAt = chData.CreatedAt
				chkData.StepTime = chData.StepTime
				chkData.Status = *chData.Status
				chkDataList = append(chkDataList, chkData)
			}
			res["message"] = "Goroutine found successfully"
		} else {
			res["message"] = "No goroutines are exists"
		}
		res["data"] = chkDataList
		statusCode = http.StatusOK
	}
	s.responseWriter(c, res, statusCode)
	return
}

//PauseRoutine ...
//Author: Suraj
//date: 06/03/2020
func (s *Service) PauseRoutine(c *gin.Context) {
	statusCode := http.StatusBadRequest
	res := gin.H{"message": "Something went wrong", "data": nil}
	stopRotineID := c.DefaultQuery("id", "")

	if stopRotineID != "" {
		var data = make(map[string]interface{})
		res["message"] = "Gorotine not found"
		for _, chData := range s.trackChannel {
			if chData.ID == stopRotineID {
				*chData.IsActive = false // Pause goroutine
				chData.ModifiedAt = time.Now()
				*chData.Status = "paused"
				data["PausedAt"] = chData.ModifiedAt
				res["data"] = data
				res["message"] = "Gorotine paused successfully"
				statusCode = http.StatusOK
			}
		}
	}

	s.responseWriter(c, res, statusCode)
	return
}

//StopRoutine ...
//Author: Suraj
//date: 06/03/2020
func (s *Service) StopRoutine(c *gin.Context) {
	statusCode := http.StatusBadRequest
	res := gin.H{"message": "Something went wrong", "data": nil}
	stopRotineID := c.DefaultQuery("id", "")

	if stopRotineID != "" {
		for _, chData := range s.trackChannel {
			if chData.ID == stopRotineID {
				if *chData.Status != "stopped" {
					*chData.Status = "stopped"
					chData.Ch <- chData.ID // Quit goroutine
					close(chData.Ch)
					res["message"] = "Gorotine stopped successfully"
				} else {
					res["message"] = "Gorotine is already stopped"
				}
				statusCode = http.StatusOK
			}
		}
	}

	s.responseWriter(c, res, statusCode)
	return
}
