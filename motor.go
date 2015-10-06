package nxt

type Motor struct {
	Port string
}

func (m *Motor) MoveMotor(replyChannel chan *ReplyTelegram) *Command {
	moveForward := append([]byte{m.ByteCode()}, []byte{0x64, 0x05, 0x02, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00}...)
	return NewDirectCommand(0x04, moveForward, replyChannel)
}

func (m *Motor) StopMotor(replyChannel chan *ReplyTelegram) *Command {
	stop := append([]byte{m.ByteCode()}, []byte{0x64, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}...)
	return NewDirectCommand(0x04, stop, replyChannel)
}

func (m *Motor) ByteCode() byte {
	if m.Port == "A" {
		return 0x00
	}
	if m.Port == "B" {
		return 0x01
	}
	return 0x02
}
