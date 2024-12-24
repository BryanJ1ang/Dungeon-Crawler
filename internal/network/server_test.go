package network

import (
	"fmt"
	"testing"
)

func TestMovementPacketEncodeDecode(t *testing.T) {
	fmt.Println("Starting Movement Packet Decode Test")

	sendPacket := NewMovementPacket(1, 2)
	recievePacket := &MovementPacket{}

	data, err := sendPacket.Encode()

	if err != nil {
		t.Errorf("Error encoding packet: %v", err)
	}

	recievePacket.Decode(data)

	if sendPacket.MessageType != recievePacket.MessageType {
		t.Errorf("Expected MessageType %d, got %d", sendPacket.MessageType, recievePacket.MessageType)
	}

	if sendPacket.CharacterID != recievePacket.CharacterID {
		t.Errorf("Expected CharacterID %d, got %d", sendPacket.CharacterID, recievePacket.CharacterID)
	}

	if sendPacket.Direction != recievePacket.Direction {
		t.Errorf("Expected Direction %d, got %d", sendPacket.Direction, recievePacket.Direction)
	}
}
