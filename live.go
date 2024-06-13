package xtream_codes_go

import (
	"strconv"
)

type LiveStream struct {
	model_base
	model_stream

	EpgChannelId      string   `json:"epg_channel_id"`
	TvArchive         boolean  `json:"tv_archive"`
	TvArchiveDuration nummeric `json:"tv_archive_duration"`
}

func (a *ApiClient) GetLiveCategories() ([]CategoryInterface, error) {
	return a.getCategories(CategoryTypeLive)
}

func (a *ApiClient) GetLiveStreams(category int) ([]*LiveStream, error) {
	var streams []*LiveStream
	var values map[string]string

	if category > -1 {
		values = map[string]string{"category_id": strconv.Itoa(category)}
	}

	if err := a.fetch(a.context("get_live_streams", values), playerApi, &streams); err != nil {
		return nil, err
	}

	return streams, nil
}

func (a *ApiClient) GetLiveStreamUri(streamId int, extension string) string {
	return a.streamUrl("live", streamId, extension)
}
