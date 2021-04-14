package escpos

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/text/encoding/charmap"
)

type Printer struct {
	file io.Writer
}

//Create a new printer
func New(file io.Writer) (p *Printer) {
	p = &Printer{file: file}
	return
}

// Write raw bytes to printer
func (p *Printer) WriteRaw(data []byte) (n int, err error) {
	if len(data) > 0 {
		log.Printf("Writing %d bytes\n", len(data))
		p.file.Write(data)
	} else {
		log.Printf("Wrote NO bytes\n")
	}

	return 0, nil
}

// Write text to the printer
func (p *Printer) Write(data string) (n int, err error) {
	return p.WriteRaw([]byte(data))
}

// Write spanish characters (only)
func (p *Printer) WriteAccents(data string) (n int, err error) {
	//The accents only work with lowercase vowels

	//Convert text to byte
	bytes := []byte(data)

	// Decode accents and ñ
	data, e := charmap.CodePage850.NewEncoder().Bytes(bytes)
	if e != nil {
		log.Fatal(e)
	}
	return p.WriteRaw([]byte(data))
}

// Init the printer
func (p *Printer) Init() {
	p.Write("\x1B@")
}

// End output
func (p *Printer) End() {
	p.Write("\xFA")
}

// Feed the printer
func (p *Printer) Feed(n int) {
	p.Write(fmt.Sprintf("\x1Bd%c", n))
}

func (p *Printer) Align(align string) {
	switch align {
	case "l":
		p.Write(fmt.Sprintf("\x1Ba%c", 0))
	case "c":
		p.Write(fmt.Sprintf("\x1Ba%c", 1))
	case "r":
		p.Write(fmt.Sprintf("\x1Ba%c", 2))
	}
}

func (p *Printer) FontSize(width, height int) {
	p.Write(fmt.Sprintf("\x1D!%c", ((width-1)<<4)|(height-1)))
}
