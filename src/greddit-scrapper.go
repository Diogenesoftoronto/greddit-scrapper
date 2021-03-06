package gredditScrapper

import (
	"time"

	"github.com/thejerf/strinterp"
)

type GredditScraper struct {
	AccessToken  string
	AuthToken    string
	NextPage     string
	PreviousPage string
	// API_URL IAPIURL

	// export class RedditScraper {

	// 	constructor(options: IRedditCredentials) {
	// 		this.AuthToken = RedditScraper.createAuthToken(options);
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
	
	
}

// 	private async getAccessToken(): Promise<string> {
	func (g GredditScraper) getAccessToken() string {
		const options RequestPromiseOptions = {
			formData: {
				grant_type: "client_credentials",
				device_id: "DO_NOT_TRACK_THIS_DEVICE",
			},
			json: true,
			headers: {
				Authorization: `Basic ${this.AuthToken}`,
			},
			};
			const accessTokenData IAccessTokenResult = Request.post(RedditScraper.API_URL.AccessToken, options);
			return accessTokenData.access_token;
		}
				
	func (g GredditScraper) GetAccess() string {
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
func listingIsBeforeDateRange(page IPageListingResult, options IRequestOptions) bool {
	return page.data.created_utc < options.BeforeDate.getTime()
}


func (g GredditScraper) scrapeData(options IRequestOptions) map[string]string {

	if !options.BeforeDate {
		options.BeforeDate = time.Now() //probably will break need proper time figure out date function return in js
	}

	g.AccessToken = g.GetAccessToken()
	recordCount := 0
	var finalPageListings IPageListingResults

	for options.FullName && recordCount < options.Pages * options.Records {
		const pageData = g.getPage(options)                                   //need to decide type
		filter(pageData, listingIsBeforeDateRange) // need to decide type
		for _, value: range pageData{
			if (listingIsBeforeDateRange(value, options)) {
				// append element at this index to list otherwise continue looping
			}

		}
		const dataBeforeDateRange = 
		finalPageListings = finalPageListings + dataBeforeDateRange
		options.FullName = g.NextPage
		recordCount += options.Records

	}

			return finalPageListings;
}

