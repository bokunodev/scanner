package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/google/gousb"
)

const (
	vid                 = 0x08ff
	pid                 = 0x0009
	manufacturer        = "Sycreader RFID Technology Co., Ltd"
	product             = "SYC ID&IC USB Reader"
	productmanufacturer = "Sycreader RFID Technology Co., Ltd SYC ID&IC USB Reader"
)

func main() {
	usbctx := gousb.NewContext()
	defer usbctx.Close()

	usbctx.Debug(0)

	devices, err := usbctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return true
	})
	if err != nil {
		panic(err)
	}

	var scanner *gousb.Device

	for _, dev := range devices {
		man, err := dev.Manufacturer()
		if err != nil {
			panic(err)
		}

		prod, err := dev.Product()
		if err != nil {
			panic(err)
		}

		if fmt.Sprint(man, " ", prod) == productmanufacturer {
			scanner = dev
			continue
		}

		dev.Close()
	}

	if scanner == nil {
		log.Fatal("no scanner found")
	}

	defer scanner.Close()

	scanner.SetAutoDetach(true)

	iface, cleanup, err := scanner.DefaultInterface()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	in, err := iface.InEndpoint(1)
	if err != nil {
		panic(err)
	}

	var bb bytes.Buffer
	for {
		bb.Reset()

		// inner_loop:
		for i := 0; i < 11; i++ {
			var buf [16]byte
			_, err := in.Read(buf[:])
			if err != nil {
				panic(err)
			}

			//if buf[2] == 0x00 {
			//	break inner_loop
			//}

			bb.WriteByte(scancodes[buf[2]])
		}

		fmt.Println(bb.String())
	}
	/*

	   	if *bufSize > 1 {
	   		log.Print("Creating buffer...")
	   		s, err := ep.NewStream(*size, *bufSize)
	   		if err != nil {
	   			log.Fatalf("ep.NewStream(): %v", err)
	   		}
	   		defer s.Close()
	   		rdr = s
	   	}

	   opCtx := context.Background()

	   	if *timeout > 0 {
	   		var done func()
	   		opCtx, done = context.WithTimeout(opCtx, *timeout)
	   		defer done()
	   	}

	   buf := make([]byte, *size)
	   log.Print("Reading...")

	   	for i := 0; *num == 0 || i < *num; i++ {
	   		num, err := rdr.ReadContext(opCtx, buf)
	   		if err != nil {
	   			log.Fatalf("Reading from device failed: %v", err)
	   		}
	   		os.Stdout.Write(buf[:num])
	   	}
	*/
}

var scancodes = map[byte]byte{
	0x04: 'A', // Keyboard a and A
	0x05: 'B', // Keyboard b and B
	0x06: 'C', // Keyboard c and C
	0x07: 'D', // Keyboard d and D
	0x08: 'E', // Keyboard e and E
	0x09: 'F', // Keyboard f and F
	0x0a: 'G', // Keyboard g and G
	0x0b: 'H', // Keyboard h and H
	0x0c: 'I', // Keyboard i and I
	0x0d: 'J', // Keyboard j and J
	0x0e: 'K', // Keyboard k and K
	0x0f: 'L', // Keyboard l and L
	0x10: 'M', // Keyboard m and M
	0x11: 'N', // Keyboard n and N
	0x12: 'O', // Keyboard o and O
	0x13: 'P', // Keyboard p and P
	0x14: 'Q', // Keyboard q and Q
	0x15: 'R', // Keyboard r and R
	0x16: 'S', // Keyboard s and S
	0x17: 'T', // Keyboard t and T
	0x18: 'U', // Keyboard u and U
	0x19: 'V', // Keyboard v and V
	0x1a: 'W', // Keyboard w and W
	0x1b: 'X', // Keyboard x and X
	0x1c: 'Y', // Keyboard y and Y
	0x1d: 'Z', // Keyboard z and Z

	0x1e: '1', // Keyboard 1 and !
	0x1f: '2', // Keyboard 2 and @
	0x20: '3', // Keyboard 3 and #
	0x21: '4', // Keyboard 4 and $
	0x22: '5', // Keyboard 5 and %
	0x23: '6', // Keyboard 6 and ^
	0x24: '7', // Keyboard 7 and &
	0x25: '8', // Keyboard 8 and *
	0x26: '9', // Keyboard 9 and (
	0x27: '0', // Keyboard 0 and )
}
