package main

type mbr struct {
	mbr_tamano         int64
	mbr_fecha_creacion [16]byte
	mbr_dsk_signature  int64
	dsk_fit            [2]byte
	partition          [4]particion
}

type particion struct {
	part_status byte
	part_type   byte
	part_fit    [2]byte
	part_start  int64
	part_size   int64
	part_name   [16]byte
}
