package scale

import "strings"

var intervalMap = map[string]int{
	"A": 3,
	"M": 2,
	"m": 1,
}

var sharps = [12]string{
	"A", "A#", "B", "C", "C#", "D",
	"D#", "E", "F", "F#", "G", "G#",
}

var flats = [12]string{
	"A", "Bb", "B", "C", "Db", "D",
	"Eb", "E", "F", "Gb", "G", "Ab",
}

/**
* Use Sharps:
G, D, A, E, B, F# major
e, b, f#, c#, g#, d# minor

Use Flats:
F, Bb, Eb, Ab, Db, Gb major
d, g, c, f, bb, eb minor
*/

var sharpKeys = [14]string{
	"C", "G", "D", "A", "E", "B", "F#",
	"a", "e", "b", "f#", "c#", "g#", "d#",
}

func capitalize(note string) string {
	if len(note) == 2 && string(note[1]) == "b" {
		return strings.ToUpper(string(note[0])) + "b"
	}
	return strings.ToUpper(note)
}

func createScale(tonic string, intervals string) []string {
	notes := organizeNotes(tonic)
	scale := make([]string, len(intervals))
	baseIndex := 0

	scale[0] = capitalize(tonic)

	for idx, interval := range intervals {
		intervalValue := intervalMap[string(interval)]

		baseIndex = baseIndex + intervalValue

		if baseIndex < len(notes) {
			scale[idx+1] = notes[baseIndex]
		}
	}

	return scale
}

func organizeNotes(tonic string) []string {
	chromatic := []string{}
	hasSharps := false
	var notes [12]string

	for _, note := range sharpKeys {
		if tonic == note {
			hasSharps = true
			break
		}
	}

	if hasSharps == true {
		notes = sharps
	} else {
		notes = flats
	}

	for idx, note := range notes {
		if capitalize(tonic) == note {
			chromatic = append(chromatic, notes[idx:]...)
			chromatic = append(chromatic, notes[:idx]...)
		}
	}

	return chromatic
}

func Scale(tonic, intervals string) []string {
	if intervals == "" {
		return organizeNotes(tonic)
	}

	return createScale(tonic, intervals)
}
