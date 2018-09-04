package strand

// ToRNA RNA complement (per RNA transcription)
func ToRNA(dna string) (rna string) {
	dnaMap := map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}

	for _, nucleotide := range dna {
		rna += dnaMap[nucleotide]
	}
	return rna
}
