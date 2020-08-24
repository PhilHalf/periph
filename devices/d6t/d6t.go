// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package d6t

import (
	"periph.io/x/periph/conn"
	"periph.io/x/periph/conn/i2c"
)

// New opens a handle to the device.
func New(i i2c.Bus) (*Dev, error) {
	d := &Dev{&i2c.Dev{Bus: i, Addr: 0x0a}}

	if err := d.c.Tx([]byte{0x4C}, nil); err != nil {
		return nil, err
	}

	return d, nil
}

// Dev is a handle to the device.
type Dev struct {
	c conn.Conn
}

// Read 35 bytes from the D6T44L sensor
func (d *Dev) Read() ([]byte, error) {
	var b [35]byte
	if err := d.c.Tx(nil, b[:]); err != nil {
		return nil, err
	}
	return b[:], nil
}
