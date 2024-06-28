package xtream_codes_go

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

type ModelBase struct {
	Num         int    `json:"num"`
	Name        string `json:"name"`
	CategoryId  int    `json:"category_id,string"`
	CategoryIds []int  `json:"category_ids"`
}

type ModelStream struct {
	StreamType   string  `json:"stream_type"`
	StreamId     int     `json:"stream_id"`
	StreamIcon   string  `json:"stream_icon"`
	Added        int     `json:"added,string"`
	IsAdult      boolean `json:"is_adult"`
	CustomSid    string  `json:"custom_sid"`
	DirectSource string  `json:"direct_source"`
}

type ModelVideo struct {
	Rating       float   `json:"rating"`
	Rating5Based float   `json:"rating_5based"`
	Tmdb         varchar `json:"tmdb"`
}

type nummeric int

func (n *nummeric) UnmarshalJSON(data []byte) error {
	var x interface{}
	var v int

	if err := json.Unmarshal(data, &x); err != nil {
		return nil
	}

	switch y := x.(type) {
	case string:
		v, _ = strconv.Atoi(y)
	case int8:
		v = int(y)
	case int16:
		v = int(y)
	case int32:
		v = int(y)
	case int64:
		v = int(y)
	case int:
		v = y
	case float32:
		v = int(y)
	case float64:
		v = int(y)
	}

	*n = nummeric(v)

	return nil
}

type boolean bool

func (b *boolean) UnmarshalJSON(data []byte) error {
	var x interface{}
	var v bool

	if err := json.Unmarshal(data, &x); err != nil {
		return nil
	}

	switch y := x.(type) {
	case string:
		v, _ = strconv.ParseBool(y)
	case int8:
	case int16:
	case int32:
	case int64:
	case int:
		if y == 1 {
			v = true
		}
	case float32:
	case float64:
		if y > 0 {
			v = true
		}
	}

	*b = boolean(v)

	return nil
}

type varchar string

func (v *varchar) UnmarshalJSON(data []byte) error {
	var x interface{}

	if err := json.Unmarshal(data, &x); err != nil {
		return nil
	}

	switch y := x.(type) {
	case string:
		*v = varchar(y)
	case int8:
		*v = varchar(strconv.Itoa(int(y)))
	case int16:
		*v = varchar(strconv.Itoa(int(y)))
	case int32:
		*v = varchar(strconv.Itoa(int(y)))
	case int64:
		*v = varchar(strconv.Itoa(int(y)))
	case float32:
		*v = varchar(strconv.Itoa(int(y)))
	case float64:
		*v = varchar(strconv.Itoa(int(y)))
	}

	return nil
}

type float float32

func (f *float) UnmarshalJSON(data []byte) error {
	var x interface{}

	data = bytes.Replace(data, []byte{','}, []byte{'.'}, -1)

	if err := json.Unmarshal(data, &x); err != nil {
		return nil
	}

	switch y := x.(type) {
	case string:
		if len(y) > 0 {
			if x, err := strconv.ParseFloat(strings.TrimSpace(y), 32); err == nil {
				*f = float(x)
			}
		}
	case int8:
		*f = float(y)
	case int16:
		*f = float(y)
	case int32:
		*f = float(y)
	case int64:
		*f = float(y)
	case float32:
		*f = float(y)
	case float64:
		*f = float(y)

	}

	return nil
}
