package main

type QueryResponse struct {
	Data struct {
		User struct {
			EdgeOwnerToTimelineMedia struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						Typename   string `json:"__typename"`
						ID         string `json:"id"`
						Dimensions struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayURL       string `json:"display_url"`
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
						AccessibilityCaption interface{} `json:"accessibility_caption"`
						EdgeMediaToCaption   struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						Shortcode          string `json:"shortcode"`
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
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_comment"`
						CommentsDisabled     bool `json:"comments_disabled"`
						TakenAtTimestamp     int  `json:"taken_at_timestamp"`
						EdgeMediaPreviewLike struct {
							Count int `json:"count"`
							Edges []struct {
								Node struct {
									ID            string `json:"id"`
									ProfilePicURL string `json:"profile_pic_url"`
									Username      string `json:"username"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_preview_like"`
						GatingInfo   interface{} `json:"gating_info"`
						MediaPreview string      `json:"media_preview"`
						Owner        struct {
							ID       string `json:"id"`
							Username string `json:"username"`
						} `json:"owner"`
						ViewerHasLiked             bool   `json:"viewer_has_liked"`
						ViewerHasSaved             bool   `json:"viewer_has_saved"`
						ViewerHasSavedToCollection bool   `json:"viewer_has_saved_to_collection"`
						ViewerInPhotoOfYou         bool   `json:"viewer_in_photo_of_you"`
						ViewerCanReshare           bool   `json:"viewer_can_reshare"`
						ThumbnailSrc               string `json:"thumbnail_src"`
						ThumbnailResources         []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"thumbnail_resources"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"data"`
	Status string `json:"status"`
}