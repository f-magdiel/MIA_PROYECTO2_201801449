package main

type Mbr struct {
	Mbr_tamano         int64
	Mbr_fecha_creacion [16]byte
	Mbr_dsk_signature  int64
	Dsk_fit            [2]byte
	Partition          [4]Particion
}

type Particion struct {
	Part_status byte
	Part_type   byte
	Part_fit    [2]byte
	Part_start  int64
	Part_size   int64
	Part_name   [16]byte
}
