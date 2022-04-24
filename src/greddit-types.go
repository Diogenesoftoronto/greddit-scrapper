package gredditTypes

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

// export interface IRequestOptions {
// 	SubReddit: string;
// 	SortType: "hot" | "new" | "rising" | "top";
// 	Pages: number;
// 	Records: number;
// 	BeforeDate?: Date;
// 	FullName?: string;
// }

type IRedditCredentials struct {
	SubReddit  string
	sortType   SortType
	Pages      int
	Records    int
	BeforeDate Date
	FullName   string
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
	return errors.New("Invalid leave type")
}

func (st SortType) IsValid() error {
	switch st {
	case hot, top, latest, rising:
		return nil
	}
	return errors.New("Invalid leave type")
}
// hot string "hot"
// top string "new"
// rising string "new"
// new string "new"

// export interface IAPIUrl {
// 	AccessToken: string;
// 	GetData: string;
// }

// export interface IAccessTokenResult {
// 	access_token: string;
// 	expires_in: number;
// }

// export interface IPageListingResult {
// 	kind: string;
// 	data: {
// 		created: number;
// 		created_utc: number;
// 		title: string;
// 		selftext: string;
// 	};
// }

// export interface IPageListingResults extends Array<IPageListingResult> {}
