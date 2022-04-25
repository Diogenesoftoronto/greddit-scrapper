// import { Buffer } from "buffer";
// import * as Request from "request-promise-native";
// import {RequestPromiseOptions} from "request-promise-native";
// import {
// 	IRedditCredentials,
// 	IRequestOptions,
// 	IAPIUrl,
// 	IAccessTokenResult,
// 	IPageListingResult,
// 	IPageListingResults,
// } from "./RedditScraper.types";

package gredditScrapper

import (
	"fmt"
	"os"
	"github.com/thejerf/strinterp"
	"bytes"
	"time"
)

type GredditScraper struct {
	AccessToken  string
	AuthToken    string
	NextPage     string
	PreviousPage string
	// API_URL IAPIURL

	// export class RedditScraper {

	// 	private AccessToken: string;
	// 	private AuthToken: string;
	// 	private NextPage: string | null = null;
	// 	private PreviousPage: string | null = null;

	// 	private static readonly API_URL: IAPIUrl = {
	// 		AccessToken: "https://www.reddit.com/api/v1/access_token",
	// 		GetData: "https://oauth.reddit.com/r/",
	// 	};

	// 	constructor(options: IRedditCredentials) {
	// 		this.AuthToken = RedditScraper.createAuthToken(options);
	// 	}

	// 	private static createAuthToken(options: IRedditCredentials): string {
	// 		return Buffer.from(`${options.AppId}:${options.AppSecret}`).toString("base64");
	// 	}
	
	// 	private async getPage(options: IRequestOptions): Promise<IPageListingResults> {

	// 		const finalUrl: string = `${RedditScraper.API_URL.GetData}${options.SubReddit}/${options.SortType}/`;

	// 		const requestOptions: RequestPromiseOptions = {
	// 			json: true,
	// 			headers: {
		// 				"Authorization": `Bearer ${this.AccessToken}`,
	// 				"User-Agent": "RedditScraper",
	// 			},
	// 		};

	// 		if (options.FullName) {
	// 			requestOptions.qs = {
	// 				after: options.FullName,
	// 			};
	// 		}

	// 		const pageData = await Request.get(finalUrl, requestOptions);
	// 		this.NextPage = pageData.data.after;
	// 		this.PreviousPage = pageData.data.before;

	// 		return pageData.data.children;
	// 	}
	
	// 	private async getAccessToken(): Promise<string> {

		// 		const options: RequestPromiseOptions = {
			// 			formData: {
				// 				grant_type: "client_credentials",
	// 				device_id: "DO_NOT_TRACK_THIS_DEVICE",
	// 			},
	// 			json: true,
	// 			headers: {
	// 				Authorization: `Basic ${this.AuthToken}`,
	// 			},
	// 		};
	
	// 		const accessTokenData: IAccessTokenResult = await Request.post(RedditScraper.API_URL.AccessToken, options);
	// 		return accessTokenData.access_token;
	// 	}
	
	// }
}

func (g GredditScraper) GetAccessToken() string {
	AccessToken := "https://www.reddit.com/api/v1/access_token"
	return AccessToken
}
func (g GredditScraper) createAuthToken(options IRedditCredentials) string {
	
	// instead of a buffer method to hide appid and other info which is why i think this exists i will try something else the string is in base 64 which I am curious as to why and where this method is used.
	authToken, err := strinterp.NewInterpolator().InterpStr("%RAW:%RAW", options.AppId, options.AppSecret)
	if err != nil {
		panic("string interp failed: press f to pay respects")
	}
	return authToken
}

func (g GredditScraper) GetData() string {
	GetData := "https://oauth.reddit.com/r/"
	return GetData
}
// 	public async scrapeData(options: IRequestOptions): Promise<IPageListingResults> {
func (g GredditScraper) scrapeData(options IRequestOptions) map[string]string {

			if (!options.BeforeDate) {
				options.BeforeDate = time.Now(); //probably will break need proper time figure out date function return in js
			}

	// 		this.AccessToken = await this.getAccessToken();
	
	// 		let finalPageListings: IPageListingResults = [];
	// 		let recordCount: number = 0;
	
	// 		do {
	// 			const pageData = await this.getPage(options);
	
	// 			const listingIsBeforeDateRange = (page: IPageListingResult) =>
	// 					(page.data.created_utc < options.BeforeDate.getTime());
	
	// 			const dataBeforeDateRange = pageData.filter(listingIsBeforeDateRange);
	
	// 			finalPageListings = finalPageListings.concat(dataBeforeDateRange);
	// 			options.FullName = this.NextPage;
	// 			recordCount += options.Records;
	
	// 		} while (options.FullName && recordCount < options.Pages * options.Records);
	
	// 		return finalPageListings;
}
// 	}
