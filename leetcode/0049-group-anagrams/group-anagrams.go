package _049_group_anagrams

func groupAnagrams(strs []string) [][]string {
	hashTable := make(map[uint64][]string, len(strs)/2)

	for _, s := range strs {
		hashValue := encode(s)
		hashTable[hashValue] = append(hashTable[hashValue], s)
	}

	var output = make([][]string, 0, len(hashTable))
	for _, v := range hashTable {
		output = append(output, v)
	}

	return output
}

var prime []uint64 = []uint64{
	2, 3, 5, 7, 11, 13, 17,
	19, 23, 29, 31, 37, 41, 43,
	47, 53, 57, 59, 61, 67,
	71, 73, 79, 83, 89, 101,
}

func encode(s string) uint64 {
	var value uint64 = 1
	for _, c := range s {
		value = value * prime[c - 'a']
	}
	return value
}