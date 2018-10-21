package main

// ProfilePostPage is 2896 bytes
type ProfilePostPage struct {
	ActivityCounts interface{} `json:"activity_counts"`
	Config         struct {
		CsrfToken string      `json:"csrf_token"`
		Viewer    interface{} `json:"viewer"`
		ViewerID  interface{} `json:"viewerId"`
	} `json:"config"`
	SupportsEs6  bool   `json:"supports_es6"`
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	EntryData    struct {
		ProfilePage []struct {
			LoggingPageID         string `json:"logging_page_id"`
			ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
			Graphql               struct {
				User struct {
					Biography              string `json:"biography"`
					BlockedByViewer        bool   `json:"blocked_by_viewer"`
					CountryBlock           bool   `json:"country_block"`
					ExternalURL            string `json:"external_url"`
					ExternalURLLinkshimmed string `json:"external_url_linkshimmed"`
					EdgeFollowedBy         struct {
						Count int `json:"count"`
					} `json:"edge_followed_by"`
					FollowedByViewer bool `json:"followed_by_viewer"`
					EdgeFollow       struct {
						Count int `json:"count"`
					} `json:"edge_follow"`
					FollowsViewer        bool        `json:"follows_viewer"`
					FullName             string      `json:"full_name"`
					HasChannel           bool        `json:"has_channel"`
					HasBlockedViewer     bool        `json:"has_blocked_viewer"`
					HighlightReelCount   int         `json:"highlight_reel_count"`
					HasRequestedViewer   bool        `json:"has_requested_viewer"`
					ID                   string      `json:"id"`
					IsBusinessAccount    bool        `json:"is_business_account"`
					BusinessCategoryName interface{} `json:"business_category_name"`
					BusinessEmail        interface{} `json:"business_email"`
					BusinessPhoneNumber  interface{} `json:"business_phone_number"`
					BusinessAddressJSON  interface{} `json:"business_address_json"`
					IsPrivate            bool        `json:"is_private"`
					IsVerified           bool        `json:"is_verified"`
					EdgeMutualFollowedBy struct {
						Count int           `json:"count"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_mutual_followed_by"`
					ProfilePicURL            string      `json:"profile_pic_url"`
					ProfilePicURLHd          string      `json:"profile_pic_url_hd"`
					RequestedByViewer        bool        `json:"requested_by_viewer"`
					Username                 string      `json:"username"`
					ConnectedFbPage          interface{} `json:"connected_fb_page"`
					EdgeOwnerToTimelineMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Typename           string `json:"__typename"`
								ID                 string `json:"id"`
								EdgeMediaToCaption struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								Shortcode          string `json:"shortcode"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								Dimensions       struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayURL  string `json:"display_url"`
								EdgeLikedBy struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								GatingInfo   interface{} `json:"gating_info"`
								MediaPreview string      `json:"media_preview"`
								Owner        struct {
									ID       string `json:"id"`
									Username string `json:"username"`
								} `json:"owner"`
								ThumbnailSrc       string `json:"thumbnail_src"`
								ThumbnailResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"thumbnail_resources"`
								IsVideo        bool `json:"is_video"`
								VideoViewCount int  `json:"video_view_count"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
					EdgeSavedMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_saved_media"`
					EdgeMediaCollections struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_collections"`
				} `json:"user"`
			} `json:"graphql"`
			FelixOnboardingVideoResources struct {
				Mp4    string `json:"mp4"`
				Poster string `json:"poster"`
			} `json:"felix_onboarding_video_resources"`
		} `json:"ProfilePage"`
		PostPage []struct {
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
					Location struct {
						ID            string `json:"id"`
						HasPublicPage bool   `json:"has_public_page"`
						Name          string `json:"name"`
						Slug          string `json:"slug"`
						AddressJSON   string `json:"address_json"`
					} `json:"location"`
					ViewerHasLiked             bool `json:"viewer_has_liked"`
					ViewerHasSaved             bool `json:"viewer_has_saved"`
					ViewerHasSavedToCollection bool `json:"viewer_has_saved_to_collection"`
					ViewerInPhotoOfYou         bool `json:"viewer_in_photo_of_you"`
					ViewerCanReshare           bool `json:"viewer_can_reshare"`
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
								AccessibilityCaption  interface{} `json:"accessibility_caption"`
								IsVideo               bool        `json:"is_video"`
								ShouldLogClientEvent  bool        `json:"should_log_client_event"`
								TrackingToken         string      `json:"tracking_token"`
								EdgeMediaToTaggedUser struct {
									Edges []interface{} `json:"edges"`
								} `json:"edge_media_to_tagged_user"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_sidecar_to_children"`
					EncodingStatus interface{} `json:"encoding_status"`
					IsPublished    bool        `json:"is_published"`
					ProductType    string      `json:"product_type"`
					Title          string      `json:"title"`
					VideoDuration  float64     `json:"video_duration"`
					ThumbnailSrc   string      `json:"thumbnail_src"`
				} `json:"shortcode_media"`
			} `json:"graphql"`
		} `json:"PostPage"`
	} `json:"entry_data"`
	Gatekeepers struct {
		Sf      bool `json:"sf"`
		Ld      bool `json:"ld"`
		Seo     bool `json:"seo"`
		Seoht   bool `json:"seoht"`
		Saa     bool `json:"saa"`
		PhoneQp bool `json:"phone_qp"`
		Nt      bool `json:"nt"`
	} `json:"gatekeepers"`
	Knobs struct {
		AcctNtb int `json:"acct:ntb"`
		Cb      int `json:"cb"`
		Captcha int `json:"captcha"`
		Fr      int `json:"fr"`
	} `json:"knobs"`
	Qe struct {
		EarlyFlush struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"early_flush"`
		FormNavigationDialog struct {
			G string `json:"g"`
			P struct {
				ShowSignInNavigationDialog string `json:"show_sign_in_navigation_dialog"`
			} `json:"p"`
		} `json:"form_navigation_dialog"`
		CredMan struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"cred_man"`
		Iab struct {
			G string `json:"g"`
			P struct {
				HasOpenAppAndroid string `json:"has_open_app_android"`
			} `json:"p"`
		} `json:"iab"`
		AppUpsellLi struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"app_upsell_li"`
		AppUpsell struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"app_upsell"`
		StaleFix struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stale_fix"`
		ProfileHeaderName struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"profile_header_name"`
		Bc3L struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"bc3l"`
		DirectConversationReporting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"direct_conversation_reporting"`
		GeneralReporting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"general_reporting"`
		Reporting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"reporting"`
		AccRecoveryLink struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"acc_recovery_link"`
		Notif struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"notif"`
		FbUnlink struct {
			G string `json:"g"`
			P struct {
				HasNewFlow string `json:"has_new_flow"`
			} `json:"p"`
		} `json:"fb_unlink"`
		MobileStoriesDoodling struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"mobile_stories_doodling"`
		ShowCopyLink struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"show_copy_link"`
		PEdit struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"p_edit"`
		Four04AsReact struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"404_as_react"`
		AccRecovery struct {
			G string `json:"g"`
			P struct {
				HasAccountRecoveryRedesign string `json:"has_account_recovery_redesign"`
			} `json:"p"`
		} `json:"acc_recovery"`
		AsyncUnreadCounts struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"async_unread_counts"`
		Collections struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"collections"`
		CommentTa struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"comment_ta"`
		Su struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"su"`
		DiscPpl struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"disc_ppl"`
		EbdUl struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ebd_ul"`
		EbdsimLi struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ebdsim_li"`
		EbdsimLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ebdsim_lo"`
		EmptyFeed struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"empty_feed"`
		Bundles struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"bundles"`
		ExitStoryCreation struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"exit_story_creation"`
		Appsell struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"appsell"`
		Imgopt struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"imgopt"`
		FollowButton struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"follow_button"`
		LogCont struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"log_cont"`
		Msisdn struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"msisdn"`
		BgSync struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"bg_sync"`
		Onetaplogin struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"onetaplogin"`
		LoginPoe struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"login_poe"`
		PrivateLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"private_lo"`
		ProfileTabs struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"profile_tabs"`
		PushNotifications struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"push_notifications"`
		QplFln struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"qpl_fln"`
		Reg struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"reg"`
		EmSig struct {
			G string `json:"g"`
			P struct {
				HasMultiStepEmailPrefill      string `json:"has_multi_step_email_prefill"`
				HasMultiStepEmailPrefillLeven string `json:"has_multi_step_email_prefill_leven"`
			} `json:"p"`
		} `json:"em_sig"`
		MultiregIter struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"multireg_iter"`
		RegVp struct {
			G string `json:"g"`
			P struct {
				HideValueProp string `json:"hide_value_prop"`
			} `json:"p"`
		} `json:"reg_vp"`
		ReportMedia struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"report_media"`
		ReportProfile struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"report_profile"`
		ScrollLog struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"scroll_log"`
		SidecarSwipe struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"sidecar_swipe"`
		SuUniverse struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"su_universe"`
		Stale struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stale"`
		LoStories struct {
			G string `json:"g"`
			P struct {
				ContextualLogin string `json:"contextual_login"`
			} `json:"p"`
		} `json:"lo_stories"`
		Stories struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stories"`
		TpPblshr struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"tp_pblshr"`
		Video struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"video"`
		GdprEuTos struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"gdpr_eu_tos"`
		GdprRowTos struct {
			G string `json:"g"`
			P struct {
				TosVersion string `json:"tos_version"`
			} `json:"p"`
		} `json:"gdpr_row_tos"`
		FdGr struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"fd_gr"`
		Felix struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix"`
		FelixClearFbCookie struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_clear_fb_cookie"`
		FelixCreationDurationLimits struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_creation_duration_limits"`
		FelixCreationEnabled struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_creation_enabled"`
		FelixCreationFbCrossposting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_creation_fb_crossposting"`
		FelixCreationFbCrosspostingV2 struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_creation_fb_crossposting_v2"`
		FelixCreationValidation struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_creation_validation"`
		FelixCreationVideoUpload struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_creation_video_upload"`
		FelixEarlyOnboarding struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"felix_early_onboarding"`
		UnfollowConfirm struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"unfollow_confirm"`
		ProfileEnhanceLi struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"profile_enhance_li"`
		ProfileEnhanceLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"profile_enhance_lo"`
		PhoneConfirm struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"phone_confirm"`
		CommentEnhance struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"comment_enhance"`
		MwebTopicalExplore struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"mweb_topical_explore"`
		WebNametag struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_nametag"`
		ImageDowngrade struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"image_downgrade"`
		ImageDowngradeLite struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"image_downgrade_lite"`
		FollowAllFb struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"follow_all_fb"`
		LiteDirectUpsell struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"lite_direct_upsell"`
		WebLoggedoutNoop struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_loggedout_noop"`
		StoriesVideoPreload struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stories_video_preload"`
		LiteStoriesVideoPreload struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"lite_stories_video_preload"`
		A2HsHeuristicUc struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"a2hs_heuristic_uc"`
		A2HsHeuristicNonUc struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"a2hs_heuristic_non_uc"`
		WebHashtag struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_hashtag"`
		HeaderScroll struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"header_scroll"`
		Rout struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"rout"`
		Websr struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"websr"`
		WebLoFollow struct {
			G string `json:"g"`
			P struct {
				FollowAfterLogin string `json:"follow_after_login"`
			} `json:"p"`
		} `json:"web_lo_follow"`
		WebShare struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_share"`
		LiteRating struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"lite_rating"`
		WebEmbedsShare struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_embeds_share"`
		WebShareLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_share_lo"`
		WebEmbedsLoggedOut struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_embeds_logged_out"`
		Sl struct {
			G string `json:"g"`
			P struct {
				ShowLogo string `json:"show_logo"`
			} `json:"p"`
		} `json:"sl"`
		RegNux struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"reg_nux"`
		WebDatasaverMode struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"web_datasaver_mode"`
		LiteDatasaverMode struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"lite_datasaver_mode"`
		LiteVideoUpload struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"lite_video_upload"`
		IgAat struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ig_aat"`
		LoReturnurl struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"lo_returnurl"`
		ForgotPass struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"forgot_pass"`
	} `json:"qe"`
	Hostname        string `json:"hostname"`
	DeploymentStage string `json:"deployment_stage"`
	Platform        string `json:"platform"`
	RhxGis          string `json:"rhx_gis"`
	Nonce           string `json:"nonce"`
	ServerChecks    struct {
	} `json:"server_checks"`
	ZeroData struct {
	} `json:"zero_data"`
	RolloutHash    string `json:"rollout_hash"`
	BundleVariant  string `json:"bundle_variant"`
	ProbablyHasApp bool   `json:"probably_has_app"`
}

func (p *ProfilePostPage) listURLs() []string {
	var urls []string
	sm := &p.EntryData.PostPage[0].Graphql.ShortcodeMedia
	switch {
	case len(sm.EdgeSidecarToChildren.Edges) > 0:
		edges := sm.EdgeSidecarToChildren.Edges
		for i := range edges {
			r := &edges[i]
			urls = append(urls, r.Node.DisplayResources[2].Src)
		}
	case sm.Typename == "GraphImage":
		urls = append(urls, sm.DisplayResources[2].Src)
	case sm.Typename == "GraphVideo":
		urls = append(urls, sm.VideoURL)
	}
	return urls
}
