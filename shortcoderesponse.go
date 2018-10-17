package main

// ShortcodeQueryResponse represents the JSON data returned by __a=1 shortcode
// media requests
type ShortcodeQueryResponse struct {
	Graphql struct {
		ShortcodeMedia struct {
			Typename   string `json:"__typename"`
			ID         string `json:"id"`
			Shortcode  string `json:"shortcode"`
			Dimensions struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"dimensions"`
			GatingInfo       interface{} `json:"gating_info"`
			MediaPreview     interface{} `json:"media_preview"`
			DisplayURL       string      `json:"display_url"`
			DisplayResources []struct {
				Src          string `json:"src"`
				ConfigWidth  int    `json:"config_width"`
				ConfigHeight int    `json:"config_height"`
			} `json:"display_resources"`
			IsVideo               bool   `json:"is_video"`
			ShouldLogClientEvent  bool   `json:"should_log_client_event"`
			TrackingToken         string `json:"tracking_token"`
			EdgeMediaToTaggedUser struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_to_tagged_user"`
			EdgeMediaToCaption struct {
				Edges []struct {
					Node struct {
						Text string `json:"text"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_caption"`
			CaptionIsEdited    bool `json:"caption_is_edited"`
			HasRankedComments  bool `json:"has_ranked_comments"`
			EdgeMediaToComment struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						ID        string `json:"id"`
						Text      string `json:"text"`
						CreatedAt int    `json:"created_at"`
						Owner     struct {
							ID            string `json:"id"`
							ProfilePicURL string `json:"profile_pic_url"`
							Username      string `json:"username"`
						} `json:"owner"`
						ViewerHasLiked bool `json:"viewer_has_liked"`
						EdgeLikedBy    struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_comment"`
			CommentsDisabled     bool `json:"comments_disabled"`
			TakenAtTimestamp     int  `json:"taken_at_timestamp"`
			EdgeMediaPreviewLike struct {
				Count int           `json:"count"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_preview_like"`
			EdgeMediaToSponsorUser struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_to_sponsor_user"`
			Location                   interface{} `json:"location"`
			ViewerHasLiked             bool        `json:"viewer_has_liked"`
			ViewerHasSaved             bool        `json:"viewer_has_saved"`
			ViewerHasSavedToCollection bool        `json:"viewer_has_saved_to_collection"`
			ViewerInPhotoOfYou         bool        `json:"viewer_in_photo_of_you"`
			ViewerCanReshare           bool        `json:"viewer_can_reshare"`
			Owner                      struct {
				ID                string `json:"id"`
				ProfilePicURL     string `json:"profile_pic_url"`
				Username          string `json:"username"`
				BlockedByViewer   bool   `json:"blocked_by_viewer"`
				FollowedByViewer  bool   `json:"followed_by_viewer"`
				FullName          string `json:"full_name"`
				HasBlockedViewer  bool   `json:"has_blocked_viewer"`
				IsPrivate         bool   `json:"is_private"`
				IsUnpublished     bool   `json:"is_unpublished"`
				IsVerified        bool   `json:"is_verified"`
				RequestedByViewer bool   `json:"requested_by_viewer"`
			} `json:"owner"`
			IsAd                       bool `json:"is_ad"`
			EdgeWebMediaToRelatedMedia struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_web_media_to_related_media"`
			EdgeSidecarToChildren struct {
				Edges []struct {
					Node struct {
						Typename   string `json:"__typename"`
						ID         string `json:"id"`
						Shortcode  string `json:"shortcode"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						GatingInfo       interface{} `json:"gating_info"`
						MediaPreview     string      `json:"media_preview"`
						DisplayURL       string      `json:"display_url"`
						DisplayResources []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"display_resources"`
						DashInfo struct {
							IsDashEligible    bool        `json:"is_dash_eligible"`
							VideoDashManifest interface{} `json:"video_dash_manifest"`
							NumberOfQualities int         `json:"number_of_qualities"`
						} `json:"dash_info"`
						VideoURL              string `json:"video_url"`
						VideoViewCount        int    `json:"video_view_count"`
						IsVideo               bool   `json:"is_video"`
						ShouldLogClientEvent  bool   `json:"should_log_client_event"`
						TrackingToken         string `json:"tracking_token"`
						EdgeMediaToTaggedUser struct {
							Edges []interface{} `json:"edges"`
						} `json:"edge_media_to_tagged_user"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_sidecar_to_children"`
		} `json:"shortcode_media"`
	} `json:"graphql"`
}
