package main

import (
	"fmt"
	"github.com/joeoftheforsythe/go-nxt"
	"time"
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

	fmt.Println("Playing Concert A (440Hz) for 1 seconds...")
	n.CommandChannel <- nxt.PlayTone(220, 1000, reply)

	playToneReply := <-reply

	if playToneReply.IsSuccess() {
		fmt.Println("Played Concert A!")
	} else {
		fmt.Println("Was unable to play the tone:", playToneReply)
	}

	time.Sleep(1 * time.Second)

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

	time.Sleep(3 * time.Second) // Wait 3 seconds before trying to stop

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

	n.CommandChannel <- nxt.GetBatteryLevel(reply)
	batteryLevelReply := nxt.ParseGetBatteryLevelReply(<-reply)

	if batteryLevelReply.IsSuccess() {
		fmt.Println("Battery level (mv):", batteryLevelReply.BatteryLevelMillivolts)
	} else {
		fmt.Println("Was unable to get the current battery level")
	}
}
