package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Kind string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind string `json:"kind"`
	Id string `json:"id"`
	Stats Stats `json:"statistics"`
}

type Stats struct {
	Views string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos string `json:"videoCount"`
}

func GetSubscribers() (Item, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	os.Setenv("YOUTUBE_KEY", "AIzaSyAgixMR9QNs0zenJjRjHdZcqBSRKUkUIdI")
	os.Setenv("CHANNEL_ID", "UC-lHJZR3Gqxm24_Vd_AJ5Yw")

	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return Item{},err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	return response.Items[0], nil

}