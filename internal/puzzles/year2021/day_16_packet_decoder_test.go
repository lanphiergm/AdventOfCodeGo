package year2021

import (
	"testing"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Tests Packet Decoder Part 1 with the sample data
func TestPacketDecoderPart1Sample(t *testing.T) {
	expected := 31
	actual := PacketDecoderPart1("../../../data/year2021/day_16_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Packet Decoder Part 1 with the real puzzle data
func TestPacketDecoderPart1Puzzle(t *testing.T) {
	expected := 984
	actual := PacketDecoderPart1("../../../data/year2021/day_16_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Packet Decoder Part 2 with the sample data
func TestPacketDecoderPart2Sample(t *testing.T) {
	expected := 54
	actual := PacketDecoderPart2("../../../data/year2021/day_16_sample_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}

// Tests Packet Decoder Part 2 with the real puzzle data
func TestPacketDecoderPart2Puzzle(t *testing.T) {
	expected := 1015320896946
	actual := PacketDecoderPart2("../../../data/year2021/day_16_puzzle_data.txt")
	utils.AssertAreEqual(t, expected, actual)
}
