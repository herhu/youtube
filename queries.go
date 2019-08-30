package main

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

//channelsListByUsername Search the channel by username
func channelsListByUsername(service *youtube.Service, part string, forUsername string) {

	call := service.Channels.List(part)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")

	for _, element := range response.Items {
		fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
			"it has %d subscribers. and %d videos",
			element.Id,
			element.Snippet.Title,
			element.Statistics.SubscriberCount,
			element.Statistics.VideoCount))

	}

}

func searchByKeyword(service *youtube.Service, query *string, maxResults *int64) {
	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	handleError(err, "")

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)
	printIDs("Channels", channels)
	printIDs("Playlists", playlists)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
