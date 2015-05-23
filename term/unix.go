package term

import (
	"syscall"
	"unsafe"
)

type State syscall.Termios

func ioctl(fd, ioctl int, data unsafe.Pointer) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(ioctl), uintptr(data))
	if err == 0 {
		return nil
	}
	return err
}

func GetState(fd int) (state *State, err error) {
	state = new(State)
	err = ioctl(fd, syscall.TCGETS, unsafe.Pointer(state))
	return
}

func SetState(fd int, state *State) error {
	return ioctl(fd, syscall.TCSETS, unsafe.Pointer(state))
}

func (state *State) MakeRaw() {
	state.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	state.Oflag &^= syscall.OPOST
	state.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	state.Cflag &^= syscall.CSIZE | syscall.PARENB
	state.Cflag |= syscall.CS8
}

func MakeRaw(fd int) (original *State, err error) {
	original, err = GetState(fd)
	if err != nil {
		return nil, err
	}
	current := *original
	current.MakeRaw()
	err = SetState(fd, &current)
	return
}

