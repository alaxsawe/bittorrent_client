package main

import (
	"os"
	"math/rand"
	"strconv"
)
func peerId() string {
	sid := "-tt" + strconv.Itoa(os.Getpid()) + "_" +
	strconv.FormatInt(rand.Int63(), 10)
	return sid[0:20]
}
