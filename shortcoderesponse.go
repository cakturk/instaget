package main

// ShortcodeQueryResponse represents the JSON data returned by __a=1 shortcode
// media requests
type ShortcodeQueryResponse struct {
	Graphql struct {
		ShortcodeMedia ShortcodeMedia `json:"shortcode_media"`
	} `json:"graphql"`
}
