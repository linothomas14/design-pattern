package main

type samsungAdapter struct {
	sstv *SamsungTV
}

func (ss *samsungAdapter) turnOn() {
	ss.sstv.setOnState(true)
}

func (ss *samsungAdapter) turnOff() {
	ss.sstv.setOnState(false)
}

func (ss *samsungAdapter) volumeUp() int {
	vol := ss.sstv.getVolume() + 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *samsungAdapter) volumeDown() int {
	vol := ss.sstv.getVolume() - 1
	ss.sstv.setVolume(vol)
	return vol
}

func (ss *samsungAdapter) channelUp() int {
	ch := ss.sstv.getChannel() + 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *samsungAdapter) channelDown() int {
	ch := ss.sstv.getChannel() - 1
	ss.sstv.setChannel(ch)
	return ch
}

func (ss *samsungAdapter) goToChannel(ch int) {
	ss.sstv.setChannel(ch)
}
