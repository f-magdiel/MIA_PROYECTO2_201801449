package main

type Ebr struct {
	Part_status byte
	Part_fit    [2]byte
	Part_start  int64
	Part_size   int64
	Part_next   int64
	Part_name   [16]byte
}

type Logica struct {
	Logic [24]Ebr
}
