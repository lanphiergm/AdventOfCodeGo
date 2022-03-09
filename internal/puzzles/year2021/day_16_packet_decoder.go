package year2021

import (
	"io/ioutil"
	"strconv"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Packet Decoder Part 1 computes the sum of the version numbers in all packets
func PacketDecoderPart1(filename string) interface{} {
	bits := parseBits(filename)
	packet, _, _ := parsePacket(bits, 0, 63)
	return packet.versionSum()
}

// Packet Decoder Part 2 evaluates the top-level packet
func PacketDecoderPart2(filename string) interface{} {
	bits := parseBits(filename)
	packet, _, _ := parsePacket(bits, 0, 63)
	return packet.evaluate()
}

type packet struct {
	version       int
	typeId        int
	isLengthCount bool
	length        int
	literal       int
	subpackets    []packet
}

func (p packet) versionSum() int {
	sum := p.version
	for _, sub := range p.subpackets {
		sum += sub.versionSum()
	}
	return sum
}

func (p packet) evaluate() int {
	value := 0

	switch p.typeId {
	case 0: // sum
		for _, sub := range p.subpackets {
			value += sub.evaluate()
		}

	case 1: // product
		value = 1
		for _, sub := range p.subpackets {
			value *= sub.evaluate()
		}

	case 2: // minimum
		value = utils.MaxInt
		for _, sub := range p.subpackets {
			subVal := sub.evaluate()
			if subVal < value {
				value = subVal
			}
		}

	case 3: // maximum
		value = utils.MinInt
		for _, sub := range p.subpackets {
			subVal := sub.evaluate()
			if subVal > value {
				value = subVal
			}
		}

	case 4: // literal
		value = p.literal

	case 5: // greater than
		if p.subpackets[0].evaluate() > p.subpackets[1].evaluate() {
			value = 1
		}

	case 6: // less than
		if p.subpackets[0].evaluate() < p.subpackets[1].evaluate() {
			value = 1
		}

	case 7: // equal to
		if p.subpackets[0].evaluate() == p.subpackets[1].evaluate() {
			value = 1
		}
	}

	return value
}

func parsePacket(bits []uint64, bit int, index int) (packet, int, int) {
	p := packet{}
	p.version, bit, index = parseInt(bits, bit, index, 3)
	p.typeId, bit, index = parseInt(bits, bit, index, 3)
	if p.typeId == 4 {
		p.literal, bit, index = parseLiteral(bits, bit, index)
	} else {
		p.isLengthCount, bit, index = getBit(bits, bit, index)
		if p.isLengthCount {
			p.length, bit, index = parseInt(bits, bit, index, 11)
			for i := 0; i < p.length; i++ {
				var sub packet
				sub, bit, index = parsePacket(bits, bit, index)
				p.subpackets = append(p.subpackets, sub)
			}
		} else {
			p.length, bit, index = parseInt(bits, bit, index, 15)
			start := getPos(bit, index)
			for getPos(bit, index)-start < p.length {
				var sub packet
				sub, bit, index = parsePacket(bits, bit, index)
				p.subpackets = append(p.subpackets, sub)
			}
		}
	}
	return p, bit, index
}

func parseLiteral(bits []uint64, bit int, index int) (int, int, int) {
	literal := 0
	keepGoing := true
	for keepGoing {
		keepGoing, bit, index = getBit(bits, bit, index)
		for i := 3; i >= 0; i-- {
			isSet := false
			isSet, bit, index = getBit(bits, bit, index)
			if isSet {
				literal |= (1 << i)
			}
		}
		if keepGoing {
			literal = literal << 4
		}
	}
	return literal, bit, index
}

func parseInt(bits []uint64, bit int, index int, count int) (int, int, int) {
	val := 0
	for i := count - 1; i >= 0; i-- {
		isSet := false
		isSet, bit, index = getBit(bits, bit, index)
		if isSet {
			val |= (1 << i)
		}
	}
	return val, bit, index
}

func getBit(bits []uint64, bit int, index int) (bool, int, int) {
	val := bits[bit] & (1 << index)
	index--
	if index < 0 {
		index = 63
		bit++
	}
	return val > 0, bit, index
}

func getPos(bit, index int) int {
	return bit*64 + 63 - index
}

func parseBits(filename string) []uint64 {
	contents, _ := ioutil.ReadFile(filename)
	bits := make([]uint64, 0)
	for i := 0; i < len(contents); i += 16 {
		max := utils.Min(i+16, len(contents))
		num, _ := strconv.ParseUint(string(contents[i:max]), 16, 64)
		if max < i+16 {
			num = num << ((i + 16 - max) * 4)
		}
		bits = append(bits, num)
	}
	return bits
}
