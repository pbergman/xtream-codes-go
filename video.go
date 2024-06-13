package xtream_codes_go

type Video struct {
	Index              int          `json:"index"`
	CodecName          string       `json:"codec_name"`
	CodecLongName      string       `json:"codec_long_name"`
	Profile            string       `json:"profile"`
	CodecType          string       `json:"codec_type"`
	CodecTagString     string       `json:"codec_tag_string"`
	CodecTag           string       `json:"codec_tag"`
	Width              int          `json:"width"`
	Height             int          `json:"height"`
	CodedWidth         int          `json:"coded_width"`
	CodedHeight        int          `json:"coded_height"`
	ClosedCaptions     int          `json:"closed_captions"`
	FilmGrain          int          `json:"film_grain"`
	HasBFrames         int          `json:"has_b_frames"`
	SampleAspectRatio  string       `json:"sample_aspect_ratio"`
	DisplayAspectRatio string       `json:"display_aspect_ratio"`
	PixFmt             string       `json:"pix_fmt"`
	Level              int          `json:"level"`
	ChromaLocation     string       `json:"chroma_location"`
	FieldOrder         string       `json:"field_order"`
	Refs               int          `json:"refs"`
	IsAvc              string       `json:"is_avc"`
	NalLengthSize      string       `json:"nal_length_size"`
	RFrameRate         string       `json:"r_frame_rate"`
	AvgFrameRate       string       `json:"avg_frame_rate"`
	TimeBase           string       `json:"time_base"`
	StartPts           int          `json:"start_pts"`
	StartTime          string       `json:"start_time"`
	BitsPerRawSample   string       `json:"bits_per_raw_sample"`
	Disposition        *Disposition `json:"disposition"`
	Tags               Tags         `json:"tags"`
	Bitrate            int          `json:"bitrate"`
	Rating             string       `json:"rating"`
	// Series
	ExtradataSize int `json:"extradata_size"`
	// Video
	ColorRange     string `json:"color_range"`
	ColorSpace     string `json:"color_space"`
	ColorTransfer  string `json:"color_transfer"`
	ColorPrimaries string `json:"color_primaries"`
}
