package  shrug

import (	
	"time"
	
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"	
)

type  shrug struct {
	skill.Base
}

const (
	FAST_DURATION = 80
	SLOW_DURATION = 1000
)

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return & shrug{}
}

func (d * shrug) OnStart() {
	log.Info.Println("Starting Skill")
	err := hexabody.Start()
	
	if err != nil {
		log.Error.Println("Hexabody start err:", err)
		return
	} 
	
	right_leg := 0	
	left_leg  := 1	
	
	hexabody.Stand()
	hexabody.MoveHead(0, 10)
	head_position := hexabody.Direction()	
	
	if head_position > 45 {
		left_leg  = int(head_position / 45)
		right_leg = int((head_position / 45) - 1)
	}
	
	log.Info.Println(head_position)
	log.Info.Println(left_leg)
	log.Info.Println(right_leg)
	
	go hexabody.MoveJoint(left_leg, 2, 0, FAST_DURATION)
	go hexabody.MoveJoint(left_leg, 1, 90, FAST_DURATION)	
	
	go hexabody.MoveJoint(right_leg, 2, 0, FAST_DURATION)
	go hexabody.MoveJoint(right_leg, 1, 90, FAST_DURATION)
	
	time.Sleep(SLOW_DURATION * time.Millisecond)
	go hexabody.MoveJoint(left_leg, 1, 10, SLOW_DURATION)
	go hexabody.MoveJoint(right_leg, 1, 10, SLOW_DURATION)	
	
	time.Sleep(SLOW_DURATION * time.Millisecond)
	go hexabody.MoveJoint(left_leg, 1, 80, SLOW_DURATION)
	go hexabody.MoveJoint(right_leg, 1, 80, SLOW_DURATION)	
	
	time.Sleep(SLOW_DURATION * time.Millisecond)
	go hexabody.MoveJoint(left_leg, 1, 10, SLOW_DURATION)
	go hexabody.MoveJoint(right_leg, 1, 10, SLOW_DURATION)		
	
	time.Sleep(SLOW_DURATION * time.Millisecond)
	hexabody.Stand()
}

func (d * shrug) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d * shrug) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d * shrug) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d * shrug) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d * shrug) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
