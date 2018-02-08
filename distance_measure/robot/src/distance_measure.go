package  distance_measure

import (	
	"time"
	
	"mind/core/framework/drivers/distance"
	"mind/core/framework/log"
	"mind/core/framework/skill"	
)

type  distance_measure struct {
	skill.Base
	stop      chan bool
}

func (d * distance_measure) distance() float64 {
	d_val, err := distance.Value()	
	
	if err != nil {
		log.Error.Println(err)
	}
	
	return d_val
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.
	
	return & distance_measure{}
}

func (d * distance_measure) OnStart() {
	// Use this method to do something when this skill is starting.
	log.Info.Println("Starting")
	
	if !distance.Available() {
		log.Error.Println("Distance sensor is not available")
	} else {
		distance.Start()
		
		for {
			select {
			case <-d.stop:
				return
			default:		
				log.Info.Println(d.distance())
				time.Sleep(5000 * time.Millisecond)
			}						
		}
		
		distance.Close()
	}
}

func (d * distance_measure) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d * distance_measure) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d * distance_measure) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d * distance_measure) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d * distance_measure) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
