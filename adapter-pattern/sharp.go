package main

import "fmt"

type SharpTV struct {
	vol     int
	channel int
	isOn    bool
}

func (st *SharpTV) turnOn() {
	fmt.Println("SharpTV is now on")
	st.isOn = true
}

func (st *SharpTV) turnOff() {
	fmt.Println("SharpTV is now off")
	st.isOn = false
}

func (st *SharpTV) volumeUp() int {
	st.vol++
	fmt.Println("Increasing SharpTV volume to", st.vol)
	return st.vol
}

func (st *SharpTV) volumeDown() int {
	st.vol--
	fmt.Println("Decreasing SharpTV volume to", st.vol)
	return st.vol
}

func (st *SharpTV) channelUp() int {
	st.channel++
	fmt.Println("Decreasing SharpTV channel to", st.channel)
	return st.channel
}

func (st *SharpTV) channelDown() int {
	st.channel--
	fmt.Println("Decreasing SharpTV channel to", st.channel)
	return st.channel
}

func (st *SharpTV) goToChannel(ch int) {
	st.channel = ch
	fmt.Println("Setting SharpTV channel to", st.channel)
}
