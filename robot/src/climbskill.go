package practica

import (
	"math"
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

func newDirection(direction float64, degrees float64) float64 {
  return math.Mod(direction+degrees, 360)
}
func myStand(){
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

func LegOrientSequential( waitgroup *sync.WaitGroup, leg int,
													j0 float64, j1 float64, j2 float64, speed int){
	waitgroup.Add(1)
	hexabody.MoveJoint(leg, 2, j2, speed)
	hexabody.MoveJoint(leg, 1, j1, speed)
	hexabody.MoveJoint(leg, 0, j0, speed)
	waitgroup.Done()
}

func ready() {
	// for goroutine sync
	var wg sync.WaitGroup
	// stand
	myStand()

	// Using goroutines to make some commands be executed at the same time
	// ******************** 002 *****************************//
	log.Info.Println("2")
	// open front legs
	// LegOrientSequential(0, 120, 81, 133, SLOW_DURATION)
	// LegOrientSequential(1, 60 , 81, 133, SLOW_DURATION)
	//
  // extend back legs

	go LegOrientSequential(&wg, 3, 90, 91, 100, SLOW_DURATION)
	go LegOrientSequential(&wg, 4, 90, 91, 100, SLOW_DURATION)
	wg.Wait()
	// ******************** 003 *****************************//
	log.Info.Println("3")
	// move leg UP and FRONT
	hexabody.MoveJoint(2, 1, 40, SLOW_DURATION)
	hexabody.MoveJoint(2, 0, 120, SLOW_DURATION)
	// ******************** 004 *****************************//
	log.Info.Println("4")
	// move leg DOWN
	LegOrientSequential(&wg, 2, 120, 81, 133, SLOW_DURATION)
	// ******************** 005 *****************************//
	log.Info.Println("5")
	LegOrientSequential(&wg, 4, 90, 40, 133, SLOW_DURATION)
	// // ******************** 006 *****************************//
	log.Info.Println("6")
	LegOrientSequential(&wg, 4, 52, 40, 133, SLOW_DURATION)
	// // ******************** 007 *****************************//
	log.Info.Println("7")
	LegOrientSequential(&wg, 4, 52, 81, 133, SLOW_DURATION)
	// ******************** 008 *****************************//
	log.Info.Println("8")
	// Move 5 leg UP
	LegOrientSequential(&wg, 5, 90, 40, 133, SLOW_DURATION)
	// ******************** 009 *****************************//
	log.Info.Println("9")
	// Move 5 leg FRONT
	LegOrientSequential(&wg, 5, 60, 40, 133, SLOW_DURATION)
	// ******************** 010 *****************************//
	log.Info.Println("10")
	// Move 5 leg DOWN
	LegOrientSequential(&wg, 5, 60, 81, 133, SLOW_DURATION)
	// ******************** 011 *****************************//
	// Move leg 3 UP
	LegOrientSequential(&wg, 3, 90, 40, 133, SLOW_DURATION)
	// ******************** 012 *****************************//
	// Move leg 3 FRONT
	LegOrientSequential(&wg, 3, 128, 40, 133, SLOW_DURATION)
	// ******************** 013 *****************************//
	// Move leg 3 DOWN
	LegOrientSequential(&wg, 3, 128, 81, 133, SLOW_DURATION)
	// ******************** 014 *****************************//
	// Move all legs BACK to crawl a little
	go LegOrientSequential(&wg, 0, 135, 81, 133, SLOW_DURATION)
	go LegOrientSequential(&wg, 1, 49, 81, 133, SLOW_DURATION)
	go LegOrientSequential(&wg, 2, 97, 81, 133, SLOW_DURATION)
	go LegOrientSequential(&wg, 3, 97, 81, 133, SLOW_DURATION)
	go LegOrientSequential(&wg, 4, 90, 81, 133, SLOW_DURATION)
	go LegOrientSequential(&wg, 5, 83, 81, 133, SLOW_DURATION)
	// ******************** 015 *****************************//
	// prepare to move up to obstacle
	// bend back legs
	LegOrientSequential(&wg, 3, 97, 75, 155, SLOW_DURATION)
  LegOrientSequential(&wg, 4, 90, 75, 155, SLOW_DURATION)
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
	//LegOrientSequential(0, 49, 10, 60, SLOW_DURATION)
	//LegOrientSequential(1, 49, 10, 60, SLOW_DURATION)
	// ******************** 017 *****************************//
	go hexabody.MoveJoint(0, 0, 90, SLOW_DURATION)
	go hexabody.MoveJoint(1, 0, 90, SLOW_DURATION)
	// ******************** 018 *****************************//
	// Fron legs down again
	go hexabody.MoveJoint(0, 1, 81, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 81, SLOW_DURATION)
	// ******************** 019 *****************************//
	// left leg up
	LegOrientSequential(&wg, 0, 90, 40, 133, SLOW_DURATION)
	// ******************** 020 *****************************//
	// left leg DOWN
	LegOrientSequential(&wg, 0, 60, 109, 133, SLOW_DURATION)
	// ******************** 021 *****************************//
	// left leg up
	LegOrientSequential(&wg, 1, 90, 40, 133, SLOW_DURATION)
	// ******************** 022 *****************************//
	// left leg DOWN
	LegOrientSequential(&wg, 1, 120, 109, 133, SLOW_DURATION)
	// ******************** 023 *****************************//
	// back right leg UP
	LegOrientSequential(&wg, 3, 97, 47, 133, SLOW_DURATION)
	// ******************** 024 *****************************//
	// back right leg DOWN -while standing higher
	LegOrientSequential(&wg, 3, 90, 97, 133, SLOW_DURATION)
	//LegOrientSequential(4, 90, 148, 69, SLOW_DURATION)
	// ******************** 025 *****************************//
	log.Info.Println("25")
	// back left leg UP
	LegOrientSequential(&wg, 4, 90, 41, 133, SLOW_DURATION)
	// fron legs UP
	go LegOrientSequential(&wg, 0, 60, 70, 133, SLOW_DURATION)
	go LegOrientSequential(&wg, 1, 120, 70, 133, SLOW_DURATION)
	wg.Wait()
	// ******************** 026 *****************************//
	log.Info.Println("26")
	// back left leg down
	hexabody.MoveJoint(4, 1, 97, SLOW_DURATION)
	// ******************** 027 *****************************//
	// side legs UP
	go LegOrientSequential(&wg, 2, 97, 23, 157, SLOW_DURATION)
	go LegOrientSequential(&wg, 5, 97, 23, 157, SLOW_DURATION)
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
	LegOrientSequential(&wg, 3, 69, 76, 58, SLOW_DURATION)
	// ******************** 031 *****************************//
	// back right leg DOWN
	LegOrientSequential(&wg, 3, 69, 109, 154, SLOW_DURATION)
	// ******************** 032 *****************************//
	// back right left leg EXTEND
  hexabody.MoveJoint(3, 2, 121, SLOW_DURATION)
	// ******************** 033 *****************************//
	// back left leg UP
	LegOrientSequential(&wg, 4, 90, 76, 58, SLOW_DURATION)
	hexabody.MoveJoint(4, 0, 112, SLOW_DURATION)
	// ******************** 034 *****************************//
	// back left leg DOWN
	LegOrientSequential(&wg, 4, 112, 109, 154, SLOW_DURATION)
	// ******************** 035 *****************************//
	// back left left leg EXTEND
  hexabody.MoveJoint(4, 2, 121, SLOW_DURATION)

	// LegOrientSequential(0, 60, 70, 133, SLOW_DURATION)
	// LegOrientSequential(1, 120, 70, 133, SLOW_DURATION)
	// LegOrientSequential(2, 135, 152, 82, SLOW_DURATION)
	// LegOrientSequential(3, 68, 109, 121, SLOW_DURATION)
	// LegOrientSequential(4, 112, 109, 121, SLOW_DURATION)
	// LegOrientSequential(5, 45, 152, 82, SLOW_DURATION)

	// ******************** 036 *****************************//
	log.Info.Println("36")
	// side legs UP
	go LegOrientSequential(&wg, 2, 135, 10, 178, SLOW_DURATION)
	go LegOrientSequential(&wg, 5, 45, 10, 178, SLOW_DURATION)
	wg.Wait()
	// ******************** 037 *****************************//
	log.Info.Println("37")
	// side legs FRONT
	go LegOrientSequential(&wg, 2, 135, 56, 85, SLOW_DURATION)
	go LegOrientSequential(&wg, 5, 45, 56, 85, SLOW_DURATION)
	wg.Wait()
	//back legs extend
	go hexabody.MoveJoint(3, 2, 90, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 90, SLOW_DURATION)
	// fron legs bend
	go hexabody.MoveJoint(0, 2, 180, SLOW_DURATION)
	go hexabody.MoveJoint(1, 2, 180, SLOW_DURATION)
	// ******************** 038 *****************************//
	log.Info.Println("38")
	// side legs DOWN
	//LegOrientSequential(3, 68, 109, 90, SLOW_DURATION)
	LegOrientSequential(&wg, 3, 68, 109, 90, SLOW_DURATION) //STAND
	LegOrientSequential(&wg, 4, 112, 109, 90, SLOW_DURATION) //STAND
	go hexabody.MoveJoint(2, 1, 117, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 117, SLOW_DURATION)
	// ******************** 039 *****************************//
	log.Info.Println("39")
	// front legs uP
	LegOrientSequential(&wg, 3, 68, 109, 90, SLOW_DURATION) //STAND
	LegOrientSequential(&wg, 4, 112, 109, 90, SLOW_DURATION) //STAND
	LegOrientSequential(&wg, 0, 60, 37, 88, SLOW_DURATION)
	LegOrientSequential(&wg, 1, 120, 37, 88, SLOW_DURATION)

	// ******************** 040 *****************************//
	log.Info.Println("40")
	// front legs down
	go hexabody.MoveJoint(0, 1, 142, SLOW_DURATION)
	go hexabody.MoveJoint(1, 1, 142, SLOW_DURATION)
	// ******************** 041 *****************************//
	log.Info.Println("41")
	// back right leg UP
	LegOrientSequential(&wg, 3, 68, 31, 180, SLOW_DURATION)
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
	LegOrientSequential(&wg, 4, 112, 31, 180, SLOW_DURATION)
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
	LegOrientSequential(&wg, 0, 60, 142, 88, SLOW_DURATION) // STAND
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
	go LegOrientSequential(&wg, 2, 135, 10, 157, SLOW_DURATION)
	go LegOrientSequential(&wg, 5, 45, 10, 157, SLOW_DURATION)
	wg.Wait()
	// ******************** 049 *****************************//
	log.Info.Println("49")
	// Middle legs DOWN
	go hexabody.MoveJoint(2, 1, 76, SLOW_DURATION)
	go hexabody.MoveJoint(5, 1, 76, SLOW_DURATION)
	// OPEN back legs
	go LegOrientSequential(&wg, 3, 110, 69, 124, SLOW_DURATION)
	go LegOrientSequential(&wg, 4, 70, 69, 124, SLOW_DURATION)
	wg.Wait()

	// ******************** 050 *****************************//
	log.Info.Println("50")
	go hexabody.MoveJoint(3, 1, 128, SLOW_DURATION)
	go hexabody.MoveJoint(4, 1, 128, SLOW_DURATION)
	go hexabody.MoveJoint(3, 2, 71, SLOW_DURATION)
	go hexabody.MoveJoint(4, 2, 71, SLOW_DURATION)

	// ******************** 051 *****************************//
	log.Info.Println("51")
	LegOrientSequential(&wg, 5, 45, 76, 157, SLOW_DURATION) //STAND
	go hexabody.MoveJoint(2, 0, 60, SLOW_DURATION)
	go hexabody.MoveJoint(5, 0, 120, SLOW_DURATION)
	go hexabody.MoveJoint(3, 0, 56, SLOW_DURATION)
	go hexabody.MoveJoint(4, 0, 124, SLOW_DURATION)
	// ******************** 052 *****************************//
	log.Info.Println("52")
	LegOrientSequential(&wg, 5, 120, 76, 157, SLOW_DURATION) //STAND
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
	LegOrientSequential(&wg, 4, 45, 25, 157, SLOW_DURATION) //STAND
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

func legPositionInfo(legPosition *hexabody.LegPosition) {
	if !legPosition.IsValid() {
		log.Info.Println("The position is not valid, means it's unreachale, fit it.")
		legPosition.Fit()
	}
	x, y, z, err := legPosition.Coordinates()
	if err != nil {
		log.Info.Println("Get coordinates of legposition error:", err)
	} else {
		log.Info.Println("The coordinates of legposition are:", x, y, z)
	}
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
	// direction := 0.0
	// height := SIT_DEPTH
	// Start walking forever in 0 direction (front) at speed 1 (max 1.3)
	// hexabody.WalkContinuously(0, 1)
	// for loop
	for {
		dist, _ := distance.Value()
		log.Info.Println("Distance in millimeters: ", dist)
		time.Sleep(time.Second)

		if dist < 200 {
			// hexabody.StopWalkingContinuously()
			// time.Sleep(time.Second)
			// if height < 200 {
				// height = height + SIT_DEPTH
				// hexabody.StandWithHeight(height)
				ready()
			}	else {
				// height = SIT_DEPTH
				// hexabody.StandWithHeight(height)
				// direction = newDirection(direction,45)
				// hexabody.MoveHead(direction, 0)
			}

		}

		time.Sleep(500 * time.Millisecond)

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
