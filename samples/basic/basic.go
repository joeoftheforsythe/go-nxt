package main

import (
	"fmt"
	"github.com/tonyheupel/go-nxt"
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

	//reply := make(chan *nxt.ReplyTelegram)
	//n.PlayTone(300, 400, reply)
	//fmt.Println("New Code")
	//setOutputState := 0x04

	//reply := make(chan *nxt.ReplyTelegram)
	//message := []byte{0x02, 0x64, 0x07, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00}
	//message := []byte{0x0B, 0x02, 0xF4, 0x01}
	//message := []byte{0x01, 0xF4, 0x02, 0x0B}
	//fmt.Println(message)

	//nxt.NewDirectCommand(0x03, message, reply)
	//foo := <-reply
	//if !foo.IsSuccess() {
	//fmt.Println("%v: \"%s\"", foo.Status, message)
	//}
	//nxt.NewDirectCommand(0x03, message, nil)

	//fmt.Println("End new code")

	fmt.Println("Connected!")

	// Use a more traditional-looking method/check-for-error style
	//methodStyle(n)

	// Pause in between styles to ensure the old commands are done executing
	//time.Sleep(2 * time.Second)

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

	//message := []byte{0x02, 0x64, 0x07, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00} //C
	cMotor := []byte{0x02, 0x64, 0x07, 0x00, 0x00, 0x20, 0xF4, 0x01}
	bMotor := []byte{0x01, 0x64, 0x07, 0x00, 0x00, 0x20, 0xF4, 0x01}
	n.CommandChannel <- nxt.NewDirectCommand(0x04, cMotor, reply)

	moveCMotorReploy := <-reply

	if moveCMotorReploy.IsSuccess() {
		fmt.Println("YAY")
	} else {
		fmt.Println("Nay")
	}

	n.CommandChannel <- nxt.NewDirectCommand(0x04, bMotor, reply)
	moveBMotorReploy := <-reply

	if moveBMotorReploy.IsSuccess() {
		fmt.Println("YAY")
	} else {
		fmt.Println("Nay")
	}

	cMotorStop := []byte{0x02, 0x64, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	bMotorStop := []byte{0x01, 0x64, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n.CommandChannel <- nxt.NewDirectCommand(0x04, cMotorStop, reply)
	moveCMotorReplyStop := <-reply
	if moveCMotorReplyStop.IsSuccess() {
		fmt.Println("YAY")
	} else {
		fmt.Println("Nay")
	}
	n.CommandChannel <- nxt.NewDirectCommand(0x04, bMotorStop, reply)
	moveBMotorReplyStop := <-reply
	if moveBMotorReplyStop.IsSuccess() {
		fmt.Println("YAY")
	} else {
		fmt.Println("Nay")
	}
	//n.CommandChannel <- nxt.GetBatteryLevel(reply)
	//batteryLevelReply := nxt.ParseGetBatteryLevelReply(<-reply)

	//if batteryLevelReply.IsSuccess() {
	//fmt.Println("Battery level (mv):", batteryLevelReply.BatteryLevelMillivolts)
	//} else {
	//fmt.Println("Was unable to get the current battery level")
	//}
}

func methodStyle(n *nxt.NXT) {
	// Normally use StartProgram but we want to see the name of the running program
	// so we need to wait
	startProgramReply, err := n.StartProgramSync("Explorer.rxe")

	if err != nil {
		fmt.Println("Error starting a program:", err)
	}

	fmt.Println("Reply from StartProgram:", startProgramReply)

	runningProgram, err := n.GetCurrentProgramName()

	if err != nil {
		fmt.Println("Error getting current program name:", err)
	} else {
		fmt.Println("Current running program:", runningProgram)
	}

	time.Sleep(3 * time.Second) // Wait 3 seconds before trying to stop

	fmt.Println("Stopping running program...")
	_, err = n.StopProgramSync()

	if err != nil {
		fmt.Println("Error stopping the running program:", err)
	}

	playSoundFileReply, err := n.PlaySoundFileSync("Green.rso", false)

	if err != nil {
		fmt.Println("Error playing the sound file \"Green.rso\":", err)
	}

	fmt.Println("Reply from PlaySoundFile:", playSoundFileReply)

	//fmt.Println("Playing Convert A for 3 seconds...")
	//playToneReply, err := n.PlayToneSync(440, 3000)

	//if err != nil {
	//fmt.Println("Error playing the tone:", err)
	//}

	//fmt.Println("Reply from PlayTone:", playToneReply)

	batteryMillivolts, err := n.GetBatteryLevelMillivolts()

	if err != nil {
		fmt.Println("Error getting the battery level:", err)
	} else {
		fmt.Println("Battery level (mv):", batteryMillivolts)
	}
}
