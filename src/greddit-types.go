package gredditScrapper

// Go-Reddit Scrapper Schema
// export interface IRedditCredentials {
// 	AppId: string;
// 	AppSecret: string;
// }

import (
	"encoding/json"
	"errors"
)


type IRedditCredentials struct {
	AppId, AppSecret string
}


//  nullability is not offered in go
type IRequestOptions struct {
	SubReddit  string
	sortType   SortType
	Pages      int
	Records    int
	BeforeDate Date
	FullName   string
}


// date type is ambigous
type Date struct {
 Date string
}

type SortType string

const (
	hot    SortType = "hot"
	top    SortType = "top"
	latest SortType = "new"
	rising SortType = "rising"
)
func (st *SortType) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	sortType := SortType(s)
	switch sortType {
	case hot, top, latest, rising:
			*st = sortType
			return nil
	}
	return errors.New("invalid sort type")
}

func (st SortType) IsValid() error {
	switch st {
	case hot, top, latest, rising:
		return nil
	}
	return errors.New("invalid sort type")
}
// export interface IAPIUrl {
// 	AccessToken: string;
// 	GetData: string;
// }

type IAPIUrl struct {
	AccessToken, GetData string
}

// export interface IAccessTokenResult {
// 	access_token: string;
// 	expires_in: number;
// }

type IAccessTokenResult struct {
	access_token, expires_in string
}

// export interface IPageListingResult {
// 	kind: string;
// 	data: {
// 		created: number;
// 		created_utc: number;
// 		title: string;
// 		selftext: string;
// 	};
// }
// unsure if this is a good translation
type IPageListingResult struct {
	kind string
	data PageDataListing
}

type PageDataListing struct {
	created int
	created_utc int
	title string
	selftext string
}


// export interface IPageListingResults extends Array<IPageListingResult> {}

type IPageListingResults struct {
	IPageListingResults []IPageListingResult
}