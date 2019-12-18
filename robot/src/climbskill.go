package practica

import (
	"mind/core/framework/skill"
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/drivers/distance"
	"mind/core/framework/log"
	"time"
	// for goroutine sync
  "sync"
)

const (
	STAND_DEPTH         = 200
	SIT_DEPTH           = 50.0
	// speed to move the joints (ms)
	FAST_DURATION = 80
	SLOW_DURATION = 500
)

type practica struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return &practica{}
}

func myStand(){
	// This function makes the robot stand still in the same position than
	// the simulator does
	// All the joints are moving at the same time
	// ************* stand like in the simulator ********************//
	go hexabody.MoveJoint(0, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(0, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(0, 2, 133, SLOW_DURATION)

	go hexabody.MoveJoint(1, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(1, 2, 133, SLOW_DURATION)

	go hexabody.MoveJoint(2, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(2, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(2, 2, 133, SLOW_DURATION)

	go hexabody.MoveJoint(3, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(3, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(3, 2, 133, SLOW_DURATION)

	go hexabody.MoveJoint(4, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(4, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 133, SLOW_DURATION)

	go hexabody.MoveJoint(5, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(5, 2, 133, SLOW_DURATION)
}

func MoveLegSequential( waitgroup *sync.WaitGroup, leg int,
													 jointdeg [3]float64, sequence [3]int, speed int){
	// MoveLegSequential variables:
	//		leg: leg number to actuate
	//		jointdeg: array if three elements with the angles for each joint [j0,j1,j2]
	//		sequence: ordering of the movement sequence. F.ex. [0, 1, 2] to move
	//							joint 0 first, then joint 1 and then joint 2
	//		speed: speed at to which move each joint
	waitgroup.Add(1)
	for s := range sequence {
		hexabody.MoveJoint(leg, s, jointdeg[s], speed)
	}
	waitgroup.Done()
}

func climb() {
	// Sequence that climbs an obstacle that is in front of the robot.
	// for goroutine sync
	var wg sync.WaitGroup
	// stand
	myStand()

	// Using goroutines to make some commands be executed at the same time
	// ******************** 002 *****************************//
	log.Info.Println("2")
	// open front legs
	// MoveLegSequential(0, 120, 81, 133, SLOW_DURATION)
	// MoveLegSequential(1, 60 , 81, 133, SLOW_DURATION)
	//
  // extend back legs
	angles := [3]float64{90, 91, 100}
	seq    := [3]int		{2, 1, 0}

	go MoveLegSequential(&wg, 3, angles, seq, SLOW_DURATION)
	go MoveLegSequential(&wg, 4, angles, seq, SLOW_DURATION)
	wg.Wait()
	// ******************** 003 *****************************//
	log.Info.Println("3")
	// move leg UP and FRONT
	hexabody.MoveJoint(2, 1, 40, SLOW_DURATION)
	hexabody.MoveJoint(2, 0, 120, SLOW_DURATION)
	// ******************** 004 *****************************//
	log.Info.Println("4")
	// move leg DOWN
	MoveLegSequential(&wg, 2, [3]float64{120, 81, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 005 *****************************//
	log.Info.Println("5")
	MoveLegSequential(&wg, 4, [3]float64{90, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// // ******************** 006 *****************************//
	log.Info.Println("6")
	MoveLegSequential(&wg, 4, [3]float64{52, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// // ******************** 007 *****************************//
	log.Info.Println("7")
	MoveLegSequential(&wg, 4, [3]float64{52, 81, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 008 *****************************//
	log.Info.Println("8")
	// Move 5 leg UP
	MoveLegSequential(&wg, 5, [3]float64{90, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 009 *****************************//
	log.Info.Println("9")
	// Move 5 leg FRONT
	MoveLegSequential(&wg, 5, [3]float64{60, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 010 *****************************//
	log.Info.Println("10")
	// Move 5 leg DOWN
	MoveLegSequential(&wg, 5, [3]float64{60, 81, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 011 *****************************//
	// Move leg 3 UP
	MoveLegSequential(&wg, 3, [3]float64{90, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 012 *****************************//
	// Move leg 3 FRONT
	MoveLegSequential(&wg, 3, [3]float64{90, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 013 *****************************//
	// Move leg 3 DOWN
	MoveLegSequential(&wg, 3, [3]float64{128, 81, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 014 *****************************//
	// Move all legs BACK to crawl a little
	go MoveLegSequential(&wg, 0, [3]float64{135, 81, 133},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	go MoveLegSequential(&wg, 1, [3]float64{49, 81, 133},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	go MoveLegSequential(&wg, 2, [3]float64{97, 81, 133},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	go MoveLegSequential(&wg, 3, [3]float64{97, 81, 133},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	go MoveLegSequential(&wg, 4, [3]float64{90, 81, 133},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	go MoveLegSequential(&wg, 5, [3]float64{83, 81, 133},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 015 *****************************//
	// prepare to move up to obstacle
	// bend back legs
	MoveLegSequential(&wg, 3, [3]float64{97, 75, 155},
											[3]int{0, 1, 2}, SLOW_DURATION)
  MoveLegSequential(&wg, 4, [3]float64{90, 75, 155},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// extend middle legs
	go hexabody.MoveJoint(2, 1, 148, SLOW_DURATION)
	go hexabody.MoveJoint(2, 2, 69, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 148, SLOW_DURATION)
	go hexabody.MoveJoint(5, 2, 69, SLOW_DURATION)
	// ******************** 016 *****************************//
	hexabody.MoveJoint(0, 1, 10, SLOW_DURATION)
	hexabody.MoveJoint(1, 1, 10, SLOW_DURATION)

	//go hexabody.MoveJoint(0, 2, 60, SLOW_DURATION)
	//go hexabody.MoveJoint(1, 2, 60, SLOW_DURATION)
	//MoveLegSequential(0, 49, 10, 60, SLOW_DURATION)
	//MoveLegSequential(1, 49, 10, 60, SLOW_DURATION)
	// ******************** 017 *****************************//
	go hexabody.MoveJoint(0, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(1, 0, 90, SLOW_DURATION)
	// ******************** 018 *****************************//
	// Fron legs down again
	go hexabody.MoveJoint(0, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 81, SLOW_DURATION)
	// ******************** 019 *****************************//
	// left leg up
	MoveLegSequential(&wg, 0, [3]float64{90, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 020 *****************************//
	// left leg DOWN
	MoveLegSequential(&wg, 0, [3]float64{60, 109, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 021 *****************************//
	// left leg up
	MoveLegSequential(&wg, 1, [3]float64{90, 40, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 022 *****************************//
	// left leg DOWN
	MoveLegSequential(&wg, 1, [3]float64{120, 109, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 023 *****************************//
	// back right leg UP
	MoveLegSequential(&wg, 3, [3]float64{97, 47, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 024 *****************************//
	// back right leg DOWN -while standing higher
	MoveLegSequential(&wg, 3, [3]float64{90, 97, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	//MoveLegSequential(4, 90, 148, 69, SLOW_DURATION)
	// ******************** 025 *****************************//
	log.Info.Println("25")
	// back left leg UP
	MoveLegSequential(&wg, 4, [3]float64{90, 41, 133},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// fron legs UP
	go MoveLegSequential(&wg, 0,[3]float64{ 60, 70, 133},
												[3]int{0, 1, 2}, SLOW_DURATION)
	go MoveLegSequential(&wg, 1, [3]float64{ 120, 70, 133},
													[3]int{0, 1, 2}, SLOW_DURATION)
	wg.Wait()
	// ******************** 026 *****************************//
	log.Info.Println("26")
	// back left leg down
	hexabody.MoveJoint(4, 1, 97, SLOW_DURATION)
	// ******************** 027 *****************************//
	// side legs UP
	go MoveLegSequential(&wg, 2, [3]float64{97, 23, 157},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	go MoveLegSequential(&wg, 5, [3]float64{97, 23, 157},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 028 *****************************//
	// side legs FRONT
	hexabody.MoveJoint(2, 0, 135, SLOW_DURATION)
	hexabody.MoveJoint(5, 0, 45, SLOW_DURATION)
	// ******************** 029 *****************************//
	// all legs UUUP (excepting two front legs)
	//side legs
	hexabody.MoveJoint(2, 2, 23, SLOW_DURATION) //STAND
	go hexabody.MoveJoint(2, 2, 82, SLOW_DURATION)
	go hexabody.MoveJoint(5, 2, 82, SLOW_DURATION)

	go hexabody.MoveJoint(2, 1, 152, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 152, SLOW_DURATION)

	//back legs.
	go hexabody.MoveJoint(4, 1, 130, SLOW_DURATION)
	go hexabody.MoveJoint(3, 1, 130, SLOW_DURATION)

	go hexabody.MoveJoint(4, 2, 58, SLOW_DURATION)
	go hexabody.MoveJoint(3, 2, 58, SLOW_DURATION)

	// ******************** 030 *****************************//
	// back right leg UP
	MoveLegSequential(&wg, 3, [3]float64{69, 76, 58},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 031 *****************************//
	// back right leg DOWN
	MoveLegSequential(&wg, 3, [3]float64{69, 109, 154},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 032 *****************************//
	// back right left leg EXTEND
  hexabody.MoveJoint(3, 2, 121, SLOW_DURATION)
	// ******************** 033 *****************************//
	// back left leg UP
	MoveLegSequential(&wg, 4, [3]float64{90, 76, 58},
											[3]int{0, 1, 2}, SLOW_DURATION)
	hexabody.MoveJoint(4, 0, 112, SLOW_DURATION)
	// ******************** 034 *****************************//
	// back left leg DOWN
	MoveLegSequential(&wg, 4, [3]float64{112, 109, 154},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 035 *****************************//
	// back left left leg EXTEND
  hexabody.MoveJoint(4, 2, 121, SLOW_DURATION)


	// ******************** 036 *****************************//
	log.Info.Println("36")
	// side legs UP
	MoveLegSequential(&wg, 2, [3]float64{135, 10, 178},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	MoveLegSequential(&wg, 5, [3]float64{45, 10, 178},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	//wg.Wait()
	// ******************** 037 *****************************//
	log.Info.Println("37")
	// side legs FRONT
	MoveLegSequential(&wg, 2, [3]float64{135, 56, 85},
												 [3]int{0, 1, 2}, SLOW_DURATION)
	MoveLegSequential(&wg, 5, [3]float64{45, 56, 85 },
												 [3]int{0, 1, 2}, SLOW_DURATION)
	//wg.Wait()
	//back legs extend
	go hexabody.MoveJoint(3, 2, 90, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 90, SLOW_DURATION)
	// fron legs bend
	go hexabody.MoveJoint(0, 2, 180, SLOW_DURATION)
	go hexabody.MoveJoint(1, 2, 180, SLOW_DURATION)
	// ******************** 038 *****************************//
	log.Info.Println("38")
	// side legs DOWN
	//MoveLegSequential(3, 68, 109, 90, SLOW_DURATION)
	MoveLegSequential(&wg, 3, [3]float64{68, 109, 90},
											[3]int{0, 1, 2}, SLOW_DURATION) //STAND
	MoveLegSequential(&wg, 4, [3]float64{112, 109, 90},
											[3]int{0, 1, 2}, SLOW_DURATION) //STAND
	go hexabody.MoveJoint(2, 1, 117, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 117, SLOW_DURATION)
	// ******************** 039 *****************************//
	log.Info.Println("39")
	// front legs uP
	MoveLegSequential(&wg, 3, [3]float64{68, 109, 90},
											[3]int{0, 1, 2}, SLOW_DURATION) //STAND
	MoveLegSequential(&wg, 4, [3]float64{112, 109, 90},
											[3]int{0, 1, 2}, SLOW_DURATION) //STAND
	MoveLegSequential(&wg, 0, [3]float64{60, 37, 88},
											[3]int{0, 1, 2}, SLOW_DURATION)
	MoveLegSequential(&wg, 1, [3]float64{120, 37, 88},
											[3]int{0, 1, 2}, SLOW_DURATION)

	// ******************** 040 *****************************//
	log.Info.Println("40")
	// front legs down
	go hexabody.MoveJoint(0, 1, 142, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 142, SLOW_DURATION)
	// ******************** 041 *****************************//
	log.Info.Println("41")
	// back right leg UP
	MoveLegSequential(&wg, 3, [3]float64{68, 31, 180},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 042 *****************************//
	log.Info.Println("42")
	// back right leg DOWN
	hexabody.MoveJoint(3, 2, 130, SLOW_DURATION)
	hexabody.MoveJoint(3, 1, 149, SLOW_DURATION)
	// ******************** 043 *****************************//
	log.Info.Println("43")
	// back right leg DOWN
	go hexabody.MoveJoint(3, 1, 108, SLOW_DURATION)
	go hexabody.MoveJoint(3, 2, 83, SLOW_DURATION)
	// ******************** 044 *****************************//
	log.Info.Println("44")
	// back right leg UP
	MoveLegSequential(&wg, 4, [3]float64{112, 31, 180},
											[3]int{0, 1, 2}, SLOW_DURATION)
	// ******************** 045 *****************************//
	log.Info.Println("45")
	// back right leg DOWN
	hexabody.MoveJoint(4, 2, 130, SLOW_DURATION)
	hexabody.MoveJoint(4, 1, 149, SLOW_DURATION)
	// ******************** 046 *****************************//
	log.Info.Println("46")
	// back right leg DOWN
	go hexabody.MoveJoint(4, 1, 108, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 83, SLOW_DURATION)
	// ******************** 047 *****************************//
	log.Info.Println("47")
	MoveLegSequential(&wg, 0, [3]float64{60, 142, 88},
											[3]int{0, 1, 2}, SLOW_DURATION) // STAND
	// Front legs UP
	go hexabody.MoveJoint(0, 1, 112, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 112, SLOW_DURATION)
  // Middle legs BACK
	go hexabody.MoveJoint(2, 0, 68, SLOW_DURATION)
	go hexabody.MoveJoint(5, 0, 112, SLOW_DURATION)
	go hexabody.MoveJoint(2, 2, 126, SLOW_DURATION)
	go hexabody.MoveJoint(5, 2, 126, SLOW_DURATION)

	// ******************** 048 *****************************//
	log.Info.Println("48")
	// Middle legs UP
	MoveLegSequential(&wg, 2, [3]float64{135, 10, 157},
												 [3]int{2, 1, 0}, SLOW_DURATION)
	MoveLegSequential(&wg, 5, [3]float64{45, 10, 157},
													[3]int{2, 1, 0}, SLOW_DURATION)
	//wg.Wait()
	// ******************** 049 *****************************//
	log.Info.Println("49")
	// Middle legs DOWN
	go hexabody.MoveJoint(2, 1, 76, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 76, SLOW_DURATION)
	// OPEN back legs
	MoveLegSequential(&wg, 3, [3]float64{110, 69, 124},
													[3]int{2, 1, 0}, SLOW_DURATION)
	MoveLegSequential(&wg, 4, [3]float64{70, 69, 124},
													[3]int{2, 1, 0}, SLOW_DURATION)
	//wg.Wait()

	// ******************** 050 *****************************//
	log.Info.Println("50")
	go hexabody.MoveJoint(3, 1, 128, SLOW_DURATION)
	go hexabody.MoveJoint(4, 1, 128, SLOW_DURATION)
	go hexabody.MoveJoint(3, 2, 71, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 71, SLOW_DURATION)

	// ******************** 051 *****************************//
	log.Info.Println("51")
	MoveLegSequential(&wg, 5, [3]float64{45, 76, 157},
											[3]int{0, 1, 2}, SLOW_DURATION) //STAND
	go hexabody.MoveJoint(2, 0, 60, SLOW_DURATION)
	go hexabody.MoveJoint(5, 0, 120, SLOW_DURATION)
	go hexabody.MoveJoint(3, 0, 56, SLOW_DURATION)
	go hexabody.MoveJoint(4, 0, 124, SLOW_DURATION)
	// ******************** 052 *****************************//
	log.Info.Println("52")
	MoveLegSequential(&wg, 5, [3]float64{120, 76, 157},
											[3]int{0, 1, 2}, SLOW_DURATION) //STAND
	go hexabody.MoveJoint(0, 1, 78, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 78, SLOW_DURATION)

	go hexabody.MoveJoint(2, 1, 25, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 25, SLOW_DURATION)

	go hexabody.MoveJoint(3, 1, 25, SLOW_DURATION)
	go hexabody.MoveJoint(4, 1, 25, SLOW_DURATION)
	go hexabody.MoveJoint(3, 2, 157, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 157, SLOW_DURATION)
	// ******************** 053 *****************************//
	log.Info.Println("53")
	go hexabody.MoveJoint(2, 0, 135, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 45, SLOW_DURATION)

	go hexabody.MoveJoint(3, 0, 135, SLOW_DURATION)
	go hexabody.MoveJoint(4, 1, 45, SLOW_DURATION)
	// ******************** 054 *****************************//
	log.Info.Println("54")
	MoveLegSequential(&wg, 4, [3]float64{45, 25, 157},
											[3]int{0, 1, 2}, SLOW_DURATION) //STAND
	go hexabody.MoveJoint(0, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(2, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(3, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(4, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 81, SLOW_DURATION)

	go hexabody.MoveJoint(0, 2, 133, SLOW_DURATION)
	go hexabody.MoveJoint(1, 2, 133, SLOW_DURATION)
	go hexabody.MoveJoint(2, 2, 133, SLOW_DURATION)
	go hexabody.MoveJoint(3, 2, 133, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 133, SLOW_DURATION)
	go hexabody.MoveJoint(5, 2, 133, SLOW_DURATION)
  // ******************** 055 *****************************//
	myStand()

}

func (d *practica) OnStart() {
	// Use this method to do something when this skill is starting.
	hexabody.Start()
	distance.Start()
	hexabody.Stand()
}

func (d *practica) OnClose() {
	// Use this method to do something when this skill is closing.
	hexabody.Close()
	distance.Close()
}

func (d *practica) OnConnect() {
	// Use this method to do something when the remote connected.
	// Move head to 0 position (always start on same position)
	hexabody.MoveHead(0,0)
	// Start walking forever in 0 direction (front) at speed 1 (max 1.3)
	hexabody.WalkContinuously(0, 1)
	// for loop
	for {
		// read distance to potential obstacles
		dist, _ := distance.Value()
		log.Info.Println("Distance in millimeters: ", dist)
		time.Sleep(time.Second)

		// If we are very close to an obstacle
		if dist < 200 {
			// Stop walking
			hexabody.StopWalkingContinuously()
			// Climb obstacle
			climb()

			time.Sleep(500 * time.Millisecond)

			}	else { // If there's no obstacles
				// Keep walking
				hexabody.WalkContinuously(0, 1)
			}

		}

	}


func (d *practica) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
	hexabody.StopWalkingContinuously()
	hexabody.Relax()
}

func (d *practica) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d *practica) OnRecvString(data string) {
	// Use this method to do something when skill receive string from remote client.
}
