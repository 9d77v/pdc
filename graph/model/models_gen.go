// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Character struct {
	CharacterName string `json:"characterName"`
	OriginalName  string `json:"originalName"`
}

type Episode struct {
	ID        int64       `json:"id"`
	VideoID   int64       `json:"videoID"`
	Num       float64     `json:"num"`
	Title     string      `json:"title"`
	Desc      string      `json:"desc"`
	Cover     string      `json:"cover"`
	URL       string      `json:"url"`
	Subtitles []*Subtitle `json:"subtitles"`
	CreatedAt int64       `json:"createdAt"`
	UpdatedAt int64       `json:"updatedAt"`
}

type NewCharacter struct {
	CharacterName string `json:"characterName"`
	OriginalName  string `json:"originalName"`
}

type NewEpisode struct {
	VideoID   int64          `json:"videoID"`
	Num       float64        `json:"num"`
	Title     *string        `json:"title"`
	Desc      *string        `json:"desc"`
	Cover     *string        `json:"cover"`
	URL       string         `json:"url"`
	Subtitles []*NewSubtitle `json:"subtitles"`
}

type NewStaff struct {
	Job     string   `json:"job"`
	Persons []string `json:"persons"`
}

type NewSubtitle struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type NewUpdateEpisode struct {
	ID        int64          `json:"id"`
	Num       *float64       `json:"num"`
	Title     *string        `json:"title"`
	Desc      *string        `json:"desc"`
	Cover     *string        `json:"cover"`
	URL       string         `json:"url"`
	Subtitles []*NewSubtitle `json:"subtitles"`
}

type NewUpdateVideo struct {
	ID         int64           `json:"id"`
	Title      *string         `json:"title"`
	Desc       *string         `json:"desc"`
	PubDate    *int64          `json:"pubDate"`
	Cover      *string         `json:"cover"`
	Characters []*NewCharacter `json:"characters"`
	Staffs     []*NewStaff     `json:"staffs"`
	Tags       []string        `json:"tags"`
	IsShow     *bool           `json:"isShow"`
}

type NewVideo struct {
	Title      string          `json:"title"`
	Desc       *string         `json:"desc"`
	PubDate    *int64          `json:"pubDate"`
	Cover      *string         `json:"cover"`
	Characters []*NewCharacter `json:"characters"`
	Staffs     []*NewStaff     `json:"staffs"`
	Tags       []string        `json:"tags"`
	IsShow     bool            `json:"isShow"`
	VideoURLs  []string        `json:"videoURLs"`
	Subtitles  []string        `json:"subtitles"`
}

type Staff struct {
	Job     string   `json:"Job"`
	Persons []string `json:"Persons"`
}

type Subtitle struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Video struct {
	ID         int64        `json:"id"`
	Title      string       `json:"title"`
	Desc       string       `json:"desc"`
	PubDate    int64        `json:"pubDate"`
	Cover      string       `json:"cover"`
	Episodes   []*Episode   `json:"episodes"`
	Characters []*Character `json:"characters"`
	Staffs     []*Staff     `json:"staffs"`
	Tags       []string     `json:"tags"`
	IsShow     bool         `json:"isShow"`
	CreatedAt  int64        `json:"createdAt"`
	UpdatedAt  int64        `json:"updatedAt"`
}

type VideoConnection struct {
	TotalCount int64    `json:"totalCount"`
	Edges      []*Video `json:"edges"`
}