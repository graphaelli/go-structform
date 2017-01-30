package ubjson

const (
	noMarker byte = 0

	// value markers
	nullMarker     byte = 'Z'
	noopMarker     byte = 'N'
	trueMarker     byte = 'T'
	falseMarker    byte = 'F'
	int8Marker     byte = 'i'
	uint8Marker    byte = 'U'
	int16Marker    byte = 'I'
	int32Marker    byte = 'l'
	int64Marker    byte = 'L'
	float32Marker  byte = 'd'
	float64Marker  byte = 'D'
	highPrecMarker byte = 'H'
	charMarker     byte = 'C'
	stringMarker   byte = 'S'

	objStartMarker byte = '{'
	objEndMarker   byte = '}'
	arrStartMarker byte = '['
	arrEndMarker   byte = ']'

	countMarker byte = '#'
	typeMarker  byte = '$'
)
