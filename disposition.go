package xtream_codes_go

type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
	// Series...
	Captions     int `json:"captions,omitempty"`
	Descriptions int `json:"descriptions,omitempty"`
	Metadata     int `json:"metadata,omitempty"`
	Dependent    int `json:"dependent,omitempty"`
	StillImage   int `json:"still_image,omitempty"`
}
