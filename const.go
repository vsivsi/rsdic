package rsdic

const (
	smallBlockSize          = 64
	largeBlockSize          = 1024
	selectBlockSize         = 4096
	useRawLen               = 48
	smallBlockPerLargeBlock = largeBlockSize / smallBlockSize
)
