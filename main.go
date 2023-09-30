package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
	"github.com/mmcdole/gofeed"
	"os"
	"strconv"
	"strings"
	"time"
)

func handler(_ context.Context, _ events.LambdaFunctionURLRequest) error {
	parser := gofeed.NewParser()
	feedURL := "https://aws.amazon.com/about-aws/whats-new/recent/feed/"
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		fmt.Println("Error parsing RSS feed:", err)
		return fmt.Errorf("error parsing RSS feed: %w", err)
	}

	for _, item := range feed.Items {
		if IsSameDay(item.PublishedParsed.String(), time.Now().Format("2006-01-02")) {

			value, mapperErr := mapper(item)
			if err != nil {
				return fmt.Errorf("error parsing RSS feed: %w", mapperErr)
			}

			sendDiscordMessageErr := sendDiscordMessage(value)
			if err != nil {
				return fmt.Errorf("error parsing RSS feed: %w", sendDiscordMessageErr)
			}
		}
	}

	return nil
}

func mapper(item *gofeed.Item) (map[string]interface{}, error) {
	var message map[string]interface{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(item.Description))

	if err != nil {
		return message, fmt.Errorf("error parsing RSS feed: %w", err)
	}

	// FIELDS MAP
	var fields []map[string]interface{}
	for _, v := range item.Categories {
		s := strings.Split(v, ":")

		fields = append(fields, map[string]interface{}{
			"name":  s[0],
			"value": s[1],
		})
	}

	message = map[string]interface{}{
		"content":    item.Content,
		"username":   "AWS news",
		"avatar_url": PayloadFill(os.Getenv(strings.ToUpper("avatar_url")), "avatar_url"),
		"embeds": []map[string]interface{}{
			{
				"title":       item.Title,
				"description": doc.Text(),
				"color":       strconv.Atoi(PayloadFill(os.Getenv(strings.ToUpper("color")), "color")),
				"fields":      fields,
				"footer": map[string]interface{}{
					"text": item.Published,
				},
			},
		},
	}
	return message, nil
}
func sendDiscordMessage(message interface{}) error {
	discordWebhookURL := os.Getenv("DISCORD_WEBHOOK")
	client := resty.New()

	jsonMessage, _ := json.Marshal(message)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonMessage).
		Post(discordWebhookURL)

	if err != nil {
		return fmt.Errorf("error sending message to Discord: %w", err)
	}

	fmt.Printf("Discord webhook response: %v\n", resp.StatusCode())
	fmt.Printf("Discord webhook response: %v\n", resp.Status())
	fmt.Printf("Discord webhook response: %v\n", string(resp.Body()))
	return nil
}

func IsSameDay(date1, date2 string) bool {
	time1, err := time.Parse("2006-01-02", date1[:10])
	if err != nil {
		return false
	}
	time2, err := time.Parse("2006-01-02", date2)
	if err != nil {
		return false
	}
	return time1.Equal(time2)
}

func PayloadFill(payload, field string) string {

	switch field {
	case "avatar_url":
		if payload == "" {
			return "https://appmaster.io/cdn-cgi/image/width=1024,quality=83,format=auto/api/_files/ZC5TGYLhKnWGwvWUFtVyk7/download/"
		}
	case "color":
		if payload == "" {
			return "16753920"
		}
	}
	return payload
}

func main() {
	lambda.Start(handler)
}
