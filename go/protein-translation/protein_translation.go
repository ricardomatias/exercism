package protein

import "errors"

type CodonError error

var STOP CodonError = errors.New("STOP")
var ErrInvalidBase CodonError = errors.New("ErrInvalidBase")

type Codon []string

func (c Codon) find(codon string) bool {
	for _, codonMember := range c {
		if codon == codonMember {
			return true
		}
	}

	return false
}

var codons = [8]Codon{
	Codon{"AUG"},
	Codon{"UUU", "UUC"},
	Codon{"UUA", "UUG"},
	Codon{"UCU", "UCC", "UCA", "UCG"},
	Codon{"UAU", "UAC"},
	Codon{"UGU", "UGC"},
	Codon{"UGG"},
	Codon{"UAA", "UAG", "UGA"},
}

var proteins = [8]string{
	"Methionine",
	"Phenylalanine",
	"Leucine",
	"Serine",
	"Tyrosine",
	"Cysteine",
	"Tryptophan",
	"STOP",
}

func FromCodon(codon string) (string, error) {
	if len(codon) == 0 {
		return "", errors.New("Not a codon")
	}

	for idx, codonGroup := range codons {
		if isMember := codonGroup.find(codon); isMember {
			if idx == 7 {
				return "", STOP
			}

			return proteins[idx], nil
		}
	}
	return "", ErrInvalidBase
}

func createCodonGroups(rnaBytes []byte) (codonGroup []string) {
	prevIndex := 0
	count := 1
	rnaLength := int(len(rnaBytes) / 3)

	for index := 3; count <= rnaLength; index += 3 {
		codonGroup = append(codonGroup, string(rnaBytes[prevIndex:index]))
		prevIndex += 3
		count++
	}
	return codonGroup
}

func FromRNA(rna string) (proteinsFound []string, err error) {
	rnaBytes := []byte(rna)

	codonGroups := createCodonGroups(rnaBytes)

	for _, codonCandidate := range codonGroups {
		codon, err := FromCodon(codonCandidate)

		if err == nil {
			proteinsFound = append(proteinsFound, codon)
			continue
		}

		switch err {
		case STOP:
			return proteinsFound, nil
		case ErrInvalidBase:
			return proteinsFound, err
		}
	}

	return proteinsFound, nil
}
