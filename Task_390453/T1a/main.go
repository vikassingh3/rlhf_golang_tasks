package main

import (
	"fmt"
)

const MTU = 1500

// Segment structure representing a TCP segment
type Segment struct {
	data   []byte
	offset uint16
}

// Segments structure for handling multiple segments
type Segments struct {
	segments   []Segment
	lastOffset uint16
}

// Segment a large payload into smaller segments
func SegmentData(data []byte) []Segment {
	segments := make([]Segment, 0)
	offset := 0

	for len(data) > 0 {
		segmentLength := min(len(data), int(MTU-20)) // 20 bytes for TCP header
		segment := Segment{
			data:   data[:segmentLength],
			offset: uint16(offset),
		}

		segments = append(segments, segment)
		data = data[segmentLength:]
		offset += segmentLength
	}

	return segments
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Reassemble segments back into a single payload
func ReassembleSegments(segments []Segment) []byte {
	segmentBuffer := Segments{
		segments: make([]Segment, len(segments)),
	}

	// Store segments based on offsets
	for _, segment := range segments {
		if segment.offset > segmentBuffer.lastOffset {
			segmentBuffer.segments[len(segmentBuffer.segments)] = segment
			segmentBuffer.lastOffset = segment.offset + uint16(len(segment.data))
		} else {
			for i, s := range segmentBuffer.segments {
				if s.offset > segment.offset {
					segmentBuffer.segments = append(segmentBuffer.segments[:i], segmentBuffer.segments[i:]...)
					segmentBuffer.lastOffset = segment.offset + uint16(len(segment.data))
					break
				}
			}
		}
	}

	// Combine segments into the final payload
	var payload []byte
	for _, segment := range segmentBuffer.segments {
		payload = append(payload, segment.data...)
	}

	return payload
}

func main() {
	// Example payload
	data := []byte("Hello, this is a very large payload that needs to be segmented and reassembled!")

	// Segment the data
	segments := SegmentData(data)
	fmt.Println("Segments created:", len(segments))

	// Print each segment
	for _, segment := range segments {
		fmt.Println("Offset:", segment.offset, ", Length:", len(segment.data))
	}

	// Reassemble the segments
	reassembledData := ReassembleSegments(segments)
	fmt.Println("Reassembled Data:", string(reassembledData))
}
