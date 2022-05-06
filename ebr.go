package main

type Ebr struct {
	Part_status byte
	Part_fit    [2]byte
	Part_start  int64
	Part_size   int64
	Part_next   int64
	Part_name   [16]byte
}

type Disco struct {
	id    int64
	size  int64
	path  string
	Part  [4]Diskpart
	Logic [24]Diskpart
}

type Diskpart struct {
	id      string
	size    int64
	start   int64
	name    string
	path    string
	tipo    string
	mostrar string
}
