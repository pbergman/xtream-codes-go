package xtream_codes_go

type Audio struct {
	Index          int          `json:"index"`
	CodecName      string       `json:"codec_name"`
	CodecLongName  string       `json:"codec_long_name"`
	CodecType      string       `json:"codec_type"`
	CodecTagString string       `json:"codec_tag_string"`
	CodecTag       string       `json:"codec_tag"`
	SampleFmt      string       `json:"sample_fmt"`
	SampleRate     string       `json:"sample_rate"`
	Channels       int          `json:"channels"`
	ChannelLayout  string       `json:"channel_layout"`
	BitsPerSample  int          `json:"bits_per_sample"`
	RFrameRate     string       `json:"r_frame_rate"`
	AvgFrameRate   string       `json:"avg_frame_rate"`
	TimeBase       string       `json:"time_base"`
	StartPts       int          `json:"start_pts"`
	StartTime      string       `json:"start_time"`
	Disposition    *Disposition `json:"disposition"`
	Tags           Tags         `json:"tags"`
	// Vod
	Profile       string `json:"profile"`
	CodecTimeBase string `json:"codec_time_base"`
	// Series
	BitRate string `json:"bit_rate"`
}
