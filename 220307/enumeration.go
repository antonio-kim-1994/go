package main

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

// 위의 선언과 동일하다.
const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

type Color int

const (
	RED Color = iota
	ORANGE
	YELLOW
	GREEN
)
