package utils

import (
	"fmt"
	"time"
)

func PrintTime(prefix string) {

	s := time.Now().Local().Format(time.Stamp)
	fmt.Printf("%s %s \n", prefix, s)

}
