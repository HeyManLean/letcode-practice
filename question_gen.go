package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func executeGraphql(params map[string]interface{}) map[string]interface{} {
	jsonData, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	url := "https://leetcode.cn/graphql/"
	buffer := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest(http.MethodPost, url, buffer)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Cookie", "aliyungf_tc=6f9a3413a01a041586ec41010e964479bc2d7f65b02dc8f9fe17cf37bf2979bb; gr_user_id=3b562268-fd5b-401f-8f44-254a5e2fa010; HMACCOUNT=D280242294C93994; _bl_uid=bzme618qfwUtks66878ygk7jXwe0; Hm_lvt_f0faad39bcf8471e3ab3ef70125152c3=1727582343; messages=.eJyLjlaKj88qzs-Lz00tLk5MT1XSMdAxMtVRiik1M0i0iCk1TUkziik1T01OA5Jm5klAEcO0RKVYHbprjAUAsM027Q:1t13dd:u2aR0wBTeLFZ5tNOTxR45son-L9-pMOiEJ6h5vDkPic; csrftoken=W560COhT3vDJ4auxAt6su33a3RzCp20IUjz6MNafjCvRkvhiP3v4xkeTojIHWEse; tfstk=gIDrBrwEdJ0jgGJ_coeUbsl7Suw8d8YsEvaQxDm3Vz4oR3BUx20xdbiHyXuqR2H7PuiQxy4mozgSNDgV2xuNdTTJwDXUdJY65d9s23eLKFsfIvnlyuElrpau-imRPgJxpd9s20AKMuMwCYgJI54_-JVutS40klZuK9VH0rqUX6bo-Jx20krdZgVutKV0Vkw3-2qH0mD_0y4l3lFkfZkij2XHZS4iqrWhQGUuZEoROTWPTuV4SS4VKuk447ziqYEQ9MZEsAP-SE17rfkK7lgWnTuZxch3gv7DuJnj_2rZEhfa84nmFSDkA1EYpXD7nYYDISFnnqqmPH1L0XonB8FRv93EochTEYYdFzhnmv4jEH6USfiT3oHBkT4EiXD7wJQk8yHqmYcF4FbLmPXrpb7hT7qY0Pt202N-sl4XBx0l96FmMoz6V31d9741d_kB46CLi0q453t5.; a2873925c34ecbd2_gr_last_sent_cs1=heymanlean; sl-session=wzOPYa+uGGcVyecKhXb/Gg==; _gid=GA1.2.667320483.1729587298; a2873925c34ecbd2_gr_session_id=610ab41f-097d-4c2a-bd2a-cd1ba22ad394; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=610ab41f-097d-4c2a-bd2a-cd1ba22ad394; a2873925c34ecbd2_gr_session_id_sent_vst=610ab41f-097d-4c2a-bd2a-cd1ba22ad394; LEETCODE_SESSION=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfYXV0aF91c2VyX2lkIjoiMTg2MzE2IiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiZGphbmdvLmNvbnRyaWIuYXV0aC5iYWNrZW5kcy5Nb2RlbEJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI5OTk1NTE3N2VkODAyMDA4N2FmY2Q0MThjNTNiMGU2NDEzYmEyMzkxODQzNjU3NjM1ZGQ4NWUxMTQ2MmNmYzc4IiwiaWQiOjE4NjMxNiwiZW1haWwiOiIiLCJ1c2VybmFtZSI6ImhleW1hbmxlYW4iLCJ1c2VyX3NsdWciOiJoZXltYW5sZWFuIiwiYXZhdGFyIjoiaHR0cHM6Ly9hc3NldHMubGVldGNvZGUuY24vYWxpeXVuLWxjLXVwbG9hZC9kZWZhdWx0X2F2YXRhci5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiZGV2aWNlX2lkIjoiNTBiMDgwZWRjOTRlYmFjOGE4ZWZjNDBkZTg1MGUzOTMiLCJpcCI6IjYxLjE0NC4yMDEuMjMzIiwiX3RpbWVzdGFtcCI6MTcyOTA4MzY3MS4wMTU5NjMzLCJleHBpcmVkX3RpbWVfIjoxNzMxNjEwODAwLCJ2ZXJzaW9uX2tleV8iOjB9.G0acQQ8mroyedpFcaYzR3vlYQDIry9FyH0wB3k27H28; Hm_lpvt_f0faad39bcf8471e3ab3ef70125152c3=1729653351; _ga=GA1.1.517057270.1722862630; a2873925c34ecbd2_gr_cs1=heymanlean; _ga_PDVPZYN3CW=GS1.1.1729651657.29.1.1729653554.24.0.0")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var ret map[string]interface{}
	err = json.Unmarshal(respBody, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func queryTitleContent(titleSlug string) (string, string) {
	ret := executeGraphql(map[string]interface{}{
		"query": `query questionTranslations($titleSlug: String!) { question(titleSlug: $titleSlug) {
			translatedTitle
			translatedContent
		}}`,
		"variables":     map[string]string{"titleSlug": titleSlug},
		"operationName": "questionTranslations",
	})
	retData := ret["data"].(map[string]interface{})
	questionData := retData["question"].(map[string]interface{})
	return questionData["translatedTitle"].(string), html2md(questionData["translatedContent"].(string))
}

func queryCode(titleSlug string) string {
	ret := executeGraphql(map[string]interface{}{
		"query": `query questionEditorData($titleSlug: String!) {question(titleSlug: $titleSlug) {
			questionId
			questionFrontendId
			codeSnippets {
				lang
				langSlug
				code
			}
			envInfo
			enableRunCode
			hasFrontendPreview
			frontendPreviews
		}}`,
		"variables": map[string]string{
			"titleSlug": titleSlug,
		},
		"operationName": "questionEditorData",
	})
	retData := ret["data"].(map[string]interface{})
	questionData := retData["question"].(map[string]interface{})
	snippets := questionData["codeSnippets"].([]interface{})

	goCode := ""
	for _, snippet := range snippets {
		snippetData := snippet.(map[string]interface{})
		if snippetData["lang"].(string) == "Go" {
			goCode = snippetData["code"].(string)
			break
		}
	}
	return goCode

}

func html2md(html string) string {
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		panic(err)
	}
	return markdown
}

type CategorySlug struct {
	Slug      string      `json:"slug"`
	Questions []*Question `json:"questions"`
}

type Question struct {
	TitleSlug string `json:"titleSlug"`
	ID        int    `json:"id"`
}

func loadQuestions(filename string) []*CategorySlug {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var categorySlugList []*CategorySlug
	if err = json.Unmarshal(fileData, &categorySlugList); err != nil {
		panic(err)
	}
	return categorySlugList
}

const questionTemplate = `
package main

/*
{{title}}

{{content}}

{{url}}
*/

{{code}}


func main() {
}
`

func main() {
	for cno, categorySlug := range loadQuestions("question_list.json") {
		slug := categorySlug.Slug

		dirName := fmt.Sprintf("%02d_%s", cno+1, strings.ReplaceAll(slug, "-", "_"))

		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			err := os.Mkdir(dirName, 0755)
			if err != nil {
				panic(err)
			}
		}

		for qno, question := range categorySlug.Questions {
			code := queryCode(question.TitleSlug)
			title, content := queryTitleContent(question.TitleSlug)
			url := fmt.Sprintf("https://leetcode.cn/problems/%s/description", question.TitleSlug)

			fileContent := questionTemplate
			fileContent = strings.ReplaceAll(fileContent, "{{title}}", fmt.Sprintf("%d. %s", question.ID, title))
			fileContent = strings.ReplaceAll(fileContent, "{{content}}", content)
			fileContent = strings.ReplaceAll(fileContent, "{{code}}", code)
			fileContent = strings.ReplaceAll(fileContent, "{{url}}", url)

			toFileName := fmt.Sprintf("%s/%02d_%s_%d.go", dirName, qno+1, strings.ReplaceAll(question.TitleSlug, "-", "_"), question.ID)
			if _, err := os.Stat(toFileName); err == nil {
				continue
			}

			os.WriteFile(toFileName, []byte(fileContent), 0755)

			fmt.Println(toFileName)
		}
	}
}
