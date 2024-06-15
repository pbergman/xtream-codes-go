package xtream_codes_go

import "sort"

type CategoryType string

const (
	CategoryTypeLive   CategoryType = "live"
	CategoryTypeVod    CategoryType = "vod"
	CategoryTypeSeries CategoryType = "series"
)

type category struct {
	Id       int    `json:"category_id,string"`
	Name     string `json:"category_name"`
	ParentId int    `json:"parent_id"`
}

func (c *category) GetId() int {
	return c.Id
}

func (c *category) GetName() string {
	return c.Name
}

func (c *category) GetParentId() int {
	return c.ParentId
}

type CategoryInterface interface {
	GetId() int
	GetName() string
	GetParentId() int
}

type categoryWrapper []CategoryInterface

func (c categoryWrapper) Len() int {
	return len(c)
}

func (c categoryWrapper) Less(i, j int) bool {
	return c[i].GetName() < c[j].GetName()
}

func (c categoryWrapper) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (a *ApiClient) getCategories(categoryType CategoryType) ([]CategoryInterface, error) {
	var categories []*category

	if err := a.fetch(a.context("get_"+string(categoryType)+"_categories", nil), playerApi, &categories); err != nil {
		return nil, err
	}

	var coll = make([]CategoryInterface, len(categories))

	for i, c := 0, len(categories); i < c; i++ {
		coll[i] = categories[i]
	}

	sort.Sort(categoryWrapper(coll))

	return coll, nil
}
