package  distance_scan

import (
	"time"
	"math"
	
	"mind/core/framework/drivers/distance"
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"	
)

const (
	MOVE_HEAD_DURATION    = 100  // milliseconds
)

type  distance_scan struct {
	skill.Base
	stop      			   chan bool
	start_head_direction   float64
	current_head_direction float64
}

func (d * distance_scan) distance() float64 {
	d_val, err := distance.Value()	
	
	if err != nil {
		log.Error.Println(err)
	}
	
	return d_val
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return & distance_scan{}
}

func (d * distance_scan) OnStart() {
	log.Info.Println("Starting Skill")
	err := hexabody.Start()
	
	if err != nil {
		log.Error.Println("Hexabody start err:", err)
		return
	}	
	
	if !distance.Available() {
		log.Error.Println("Distance sensor is not available")
		return
	} else {
		log.Info.Println("Starting Scan")
		
		d.start_head_direction = hexabody.Direction()	
		log.Info.Println(d.start_head_direction)
		d.current_head_direction = math.Mod(d.start_head_direction+1, 360) 
		log.Info.Println(d.current_head_direction)
		hexabody.MoveHead(d.current_head_direction, MOVE_HEAD_DURATION)
		distance.Start()		
		
		for {
			select {
			case <-d.stop:
				return
			default:
				log.Info.Println(d.distance())
				d.current_head_direction = math.Mod(d.current_head_direction+1, 360) 
				hexabody.MoveHead(d.current_head_direction, MOVE_HEAD_DURATION)		
				
				
				if d.current_head_direction == d.start_head_direction {
					d.stop <- true	
				}
			}
		}

		time.Sleep(250 * time.Millisecond)
		distance.Close()
		hexabody.Close()
	}
}

func (d * distance_scan) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d * distance_scan) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d * distance_scan) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d * distance_scan) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d * distance_scan) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
