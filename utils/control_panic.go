package utils

import (
	"fmt"
)

var nextID = make(chan uint16)

func Counter(num uint16) {
	if num == 0 {
		for i := uint16(1); ; i++ {
			nextID <- i
		}
	}
	for i := uint16(1); ; i++ {
		nextID <- i
		if i == num {
			panic("Many mistakes have happened!")
		}
	}
}

func PanicCtrl() {
	if r := recover(); r != nil {
		num := <-nextID
		Logger(3, fmt.Sprintf("[%v] %v ", num, r))
	}
}
