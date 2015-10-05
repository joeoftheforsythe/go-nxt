package main

import (
	"fmt"
	"github.com/joeoftheforsythe/go-nxt"
	//"time"
)

func main() {
	n := nxt.NewNXT("heupel-home-bot", "/dev/tty.NXT-DevB")

	fmt.Println(n)
	err := n.Connect()

	if err != nil {
		fmt.Println("Could not connect:", err)
		return
	}

	fmt.Println("Connected!")

	// Use the raw channels style
	channelStyle(n)

	n.Disconnect()
}

func channelStyle(n *nxt.NXT) {
	// All reply messages will be sent to this channel
	reply := make(chan *nxt.ReplyTelegram)

	// Normally would pass in nil for the reply channel and not wait,
	//but we want to see the name of the running program so we need to wait
	//n.CommandChannel <- nxt.StartProgram("DREW.rxe", reply)
	//fmt.Println("Reply from StartProgram:", <-reply)

	//n.CommandChannel <- nxt.GetCurrentProgramName(reply)
	//runningProgramReply := nxt.ParseGetCurrentProgramNameReply(<-reply)
	//fmt.Println("Current running program:", runningProgramReply.Filename)

	//time.Sleep(3 * time.Second) // Wait 3 seconds before trying to stop

	//fmt.Println("Stopping running program...")
	//n.CommandChannel <- nxt.StopProgram(reply)

	//stopProgramReply := <-reply

	//if stopProgramReply.IsSuccess() {
	//fmt.Println("Stopped running program successfully!")
	//} else {
	//fmt.Println("Was unable to stop the program.")
	//}

	//fmt.Println("Playing sound file \"Green\"...")
	//n.CommandChannel <- nxt.PlaySoundFile("Green.rso", false, reply)

	//playSoundFileReply := <-reply

	//if playSoundFileReply.IsSuccess() {
	//fmt.Println("Played sound file successfully!")
	//} else {
	//fmt.Println("Was unable to play the sound file:", playSoundFileReply)
	//}

	fmt.Println("Playing Concert A (440Hz) for 3 seconds...")
	n.CommandChannel <- nxt.PlayTone(220, 3000, reply)

	playToneReply := <-reply

	if playToneReply.IsSuccess() {
		fmt.Println("Played Concert A!")
	} else {
		fmt.Println("Was unable to play the tone:", playToneReply)
	}

	bMotor := nxt.Motor{"B"}
	cMotor := nxt.Motor{"C"}

	n.CommandChannel <- bMotor.MoveMotor(reply)
	bMotorReply := <-reply
	if bMotorReply.IsSuccess() {
		fmt.Println("Yay")
	}

	n.CommandChannel <- cMotor.MoveMotor(reply)
	cMotorReply := <-reply
	if cMotorReply.IsSuccess() {
		fmt.Println("Yay")
	}

	n.CommandChannel <- bMotor.StopMotor(reply)
	bMotorReply = <-reply
	if bMotorReply.IsSuccess() {
		fmt.Println("Yay")
	}

	n.CommandChannel <- cMotor.StopMotor(reply)
	cMotorReply = <-reply
	if cMotorReply.IsSuccess() {
		fmt.Println("Yay")
	}
}
