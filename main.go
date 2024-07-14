package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/imroc/req/v3"
)

func handlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func getInfo(id string) []string {
	client := req.C()
	request := client.R()

	request.SetHeaders(map[string]string{
		"accept": "*/*",
		"accept-language": "zh-CN,zh;q=0.9",
		"authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"content-type": "application/json",
		"priority": "u=1, i",
		"referer": "https://x.com/Kunluntalk/status/1812256811199922205",
		"sec-ch-ua": `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`,
		"sec-ch-ua-mobile": "?0",
		"sec-ch-ua-platform": "macOS",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		"x-client-transaction-id": "1HyuKuXbTS0Jl0IAJMEU4bRiQ3lA4LrVIw/QBTgpexcMDPar0QjgrpyZHvrkzl9eAIPHkNZRlflXMs5OGK1DNyzIwr2t1w",
		"x-csrf-token": "36c6b644b7d47c3cf0c6dd20295d2f913385b6799f068aac210b4ea3cb3edb9baac6a582acf6c3f6b64fd996a632a6d0c579631e1788fd2f3d6dcd9e5c1a887e8c282c840215a1f5fbe8f1c6405aab8b",
		"x-twitter-active-user": "yes",
		"x-twitter-auth-type": "OAuth2Session",
		"x-twitter-client-language": "zh-cn",
	})

	cookies := []*http.Cookie{
		{Name: "guest_id", Value: "v1%3A172092576791049011"},
		{Name: "night_mode", Value: "2"},
		{Name: "guest_id_marketing", Value: "v1%3A172092576791049011"},
		{Name: "guest_id_ads", Value: "v1%3A172092576791049011"},
		{Name: "gt", Value: "1812380012072546732"},
		{Name: "kdt", Value: "j8XVU8Vn3sEFrqz4nIaMtUYPK1Ec3xSdMnyAQSQO"},
		{Name: "auth_token", Value: "0aacab7c1ab695cdf1f43296baf67f5b67a2eb83"},
		{Name: "ct0", Value: "36c6b644b7d47c3cf0c6dd20295d2f913385b6799f068aac210b4ea3cb3edb9baac6a582acf6c3f6b64fd996a632a6d0c579631e1788fd2f3d6dcd9e5c1a887e8c282c840215a1f5fbe8f1c6405aab8b"},
		{Name: "att", Value: "1-kEckFrFP1K3OWNp4JxbRwFxdr76AQodMj7xJyVw1"},
		{Name: "lang", Value: "zh-cn"},
		{Name: "twid", Value: "u%3D732225356388667392"},
		{Name: "personalization_id", Value: "v1_NWL+26Wc/SU+2Ljoze+0gg=="},
	}

	request.SetCookies(cookies...)


	var data = map[string]string{
		"variables": fmt.Sprintf(`{"focalTweetId":"%s","with_rux_injections":false,"includePromotedContent":true,"withCommunity":true,"withQuickPromoteEligibilityTweetFields":true,"withBirdwatchNotes":true,"withVoice":true}`,id),
		"features": `{"rweb_tipjar_consumption_enabled":true,"responsive_web_graphql_exclude_directive_enabled":true,"verified_phone_label_enabled":false,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_timeline_navigation_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"communities_web_enable_tweet_community_results_fetch":true,"c9s_tweet_anatomy_moderator_badge_enabled":true,"articles_preview_enabled":true,"tweetypie_unmention_optimization_enabled":true,"responsive_web_edit_tweet_api_enabled":true,"graphql_is_translatable_rweb_tweet_is_translatable_enabled":true,"view_counts_everywhere_api_enabled":true,"longform_notetweets_consumption_enabled":true,"responsive_web_twitter_article_tweet_consumption_enabled":true,"tweet_awards_web_tipping_enabled":false,"creator_subscriptions_quote_tweet_preview_enabled":false,"freedom_of_speech_not_reach_fetch_enabled":true,"standardized_nudges_misinfo":true,"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled":true,"rweb_video_timestamps_enabled":true,"longform_notetweets_rich_text_read_enabled":true,"longform_notetweets_inline_media_enabled":true,"responsive_web_enhance_cards_enabled":false}`,
		"fieldToggles": `{"withArticleRichContentState":true,"withArticlePlainText":false,"withGrokAnalyze":false}`,
	}
	targetUrl := "https://x.com/i/api/graphql/QVo2zKMcLZjXABtcYpi0mA/TweetDetail"
	//要把data，变成form表单的形式然后get出去
	request.SetFormData(data)

	resp, err := request.Get(targetUrl)
	if err != nil {
		handlerError(err)
	}

	var result = map[string]interface{}{}
	err = resp.Unmarshal(&result)
	if err != nil {
		handlerError(err)
	}

	list := result["data"].(map[string]interface{})["threaded_conversation_with_injections_v2"].(map[string]interface{})["instructions"].([]interface{})[0].(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["itemContent"].(map[string]interface{})["tweet_results"].(map[string]interface{})["result"].(map[string]interface{})["legacy"].(map[string]interface{})["entities"].(map[string]interface{})["media"]
	op := list.([]interface{})[0].(map[string]interface{})["video_info"].(map[string]interface{})["variants"]
	// 取出op中的所有url
	resultList := []string{}
	for _, v := range op.([]interface{}) {
		// fmt.Println(v.(map[string]interface{})["url"])
		if v.(map[string]interface{})["content_type"].(string) == "video/mp4" {
			resultList = append(resultList, v.(map[string]interface{})["url"].(string))
		}

	}
	// fmt.Println(resultList)
	return resultList
}

func main() {
	// 获取用户输入的推特网址
	var url string
	fmt.Println("请输入推特网址：")
	fmt.Scanln(&url)
	fmt.Println("您输入的推特网址是：", url)
	// 获取推特id,以status/为分隔符，取最后一个
	id := strings.Split(url, "status/")[1]
	fmt.Println("推特id是：", id)
	list := getInfo(id)
	for _, v := range list {
		fmt.Println(v)
	}

}
