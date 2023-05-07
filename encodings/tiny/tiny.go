package tiny

var (
	_decoder = &Decoder{}
	_encoder = &Encoder{}
)

func GetDecoder() *Decoder {
	return _decoder
}

func GetEncoder() *Encoder {
	return _encoder
}
