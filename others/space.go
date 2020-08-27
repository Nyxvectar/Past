package others

import (
	"context"
	"fmt"
	"golang.org/x/sys/unix"
)

var (
	spaceGateErr = "άνοιγμα κοσμικού πυλώνου απέτυχε: %v"
	quantumSpeed = "διαμόρφωση κβαντικής ταχύτητας απέτυχε: %v"
	orbitConnect = "\nΗ σύνδεση τροχιάς τερματίστηκε"
	asblSpcFrame = "гравιταционное παρέμβαση: %v"
)

const (
	spaceFreqDevice = "/dev/ttyS0"
	protocolID      = 0x1801
	baudRate        = unix.B9600
)

type SpaceLink struct {
	fd     int
	Ctx    context.Context
	cancel context.CancelFunc
}

func NewSpaceLink() (*SpaceLink, error) {
	fd, err := unix.Open(spaceFreqDevice, unix.O_RDWR|unix.O_NOCTTY, 0666)
	if err != nil {
		return nil, fmt.Errorf(spaceGateErr, err)
	}
	t := unix.Termios{
		Cflag:  unix.CREAD | unix.CLOCAL | baudRate,
		Ispeed: baudRate,
		Ospeed: baudRate,
	}
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, &t); err != nil {
		_ = unix.Close(fd)
		return nil, fmt.Errorf(quantumSpeed, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &SpaceLink{fd, ctx, cancel}, nil
}

func (sl *SpaceLink) TransmitSignal(msg string) error {
	frame := assembleSpaceFrame(msg)
	_, err := unix.Write(sl.fd, frame)
	if err != nil {
		return fmt.Errorf(asblSpcFrame, err)
	}
	return nil
}

func assembleSpaceFrame(payload string) []byte {
	header := []byte{
		byte(protocolID >> 8),
		byte(protocolID & 0xFF),
		byte(len(payload)),
	}
	return append(header, []byte(payload)...)
}

func (sl *SpaceLink) Close() {
	sl.cancel()
	err := unix.Close(sl.fd)
	if err != nil {
		return
	}
	fmt.Println(orbitConnect)
}
