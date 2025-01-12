package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PacketType interface {
	Encode() ([]byte, error)
	Decode(data []byte) error
}

const (
	PacketTypeMovement uint8 = iota
	PacketTypeBattle
)

type MovementPacket struct {
	// direction: 0 = up, 1 = right, 2 = down, 3 = left
	MessageType uint8
	CharacterID uint8
	Direction   uint8
}

func NewMovementPacket(characterID uint8, direction uint8) MovementPacket {
	return MovementPacket{
		MessageType: 0,
		CharacterID: characterID,
		Direction:   direction,
	}
}

func (m *MovementPacket) Encode() ([]byte, error) {
	// stub
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, m.MessageType); err != nil {
		return nil, fmt.Errorf("failed to encode message type: %v", err)
	}

	if err := binary.Write(buf, binary.LittleEndian, m.CharacterID); err != nil {
		return nil, fmt.Errorf("failed to encode character ID: %v", err)
	}

	if err := binary.Write(buf, binary.LittleEndian, m.Direction); err != nil {
		return nil, fmt.Errorf("failed to encode direction: %v", err)
	}

	return buf.Bytes(), nil
}

func (m *MovementPacket) Decode(data []byte) error {

	buf := bytes.NewReader(data)

	if err := binary.Read(buf, binary.LittleEndian, &m.MessageType); err != nil {
		return fmt.Errorf("failed to decode message type: %v", err)
	}

	if err := binary.Read(buf, binary.LittleEndian, &m.CharacterID); err != nil {
		return fmt.Errorf("failed to decode character ID: %v", err)
	}

	if err := binary.Read(buf, binary.LittleEndian, &m.Direction); err != nil {
		return fmt.Errorf("failed to decode direction: %v", err)
	}

	return nil
}

type BattlePacket struct {
	// TBD for battle mechanics
}
