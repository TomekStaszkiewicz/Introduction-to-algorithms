package dataTypes

import (
	"crypto/sha256"
)


// Function calculating Marke tree root hash
func GetRoot(dataBlocks [][]byte) []byte {
	if len(dataBlocks) == 1 {
		return dataBlocks[0]
	}

	var nextIterationData [][]byte 
	for index := 0; index < len(dataBlocks); index += 2 {
		data1 := dataBlocks[index]
		var data2 []byte; 
		if len(dataBlocks) <= index + 1 {
			data2 = dataBlocks[index]
		} else {
			data2 = dataBlocks[index+1]
		}
		nextIterationData = append(nextIterationData, getPairHash(data1, data2))
	}

	if len(nextIterationData) == 1 {
		return nextIterationData[0]
	}
	return GetRoot(nextIterationData)
}

func getPairHash(dataBlock1, dataBlock2 []byte) []byte {
	h1 := getSingleHash(dataBlock1)
	h2 := getSingleHash(dataBlock2)
	result := getSingleHash(append(h1, h2...))
	return result
}

func getSingleHash(dataBlock []byte) []byte {
	h := sha256.New()
	h.Write(dataBlock)

	return h.Sum(nil)
}