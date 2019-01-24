package main

import (
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	d, err := time.ParseDuration(string(b) + "s")
	if err != nil {
		return err
	}
	t.Time = time.Unix(0, 0).Add(d)
	t.Time.In(time.FixedZone("UTC+8", 8*60*60))
	return nil
}

type Metadata struct {
	IdentDataIPAddrDecrypted    string `json:"identData_ip_addr_decrypted"`
	IdentDataUserAgentDecrypted string `json:"identData_user_agent_decrypted"`
	Pid                         string `json:"pid"`
	Host                        string `json:"host"`
	Program                     string `json:"program"`
	Time                        Time   `json:"time"`
	Rid                         string `json:"rid"`
}

type IdentData struct {
	Clienttimestamp string `json:"clienttimestamp"`
	BuildProduct    string `json:"build_product"`
	Ident           string `json:"ident"`
	IPAddr          string `json:"ip_addr"`
	ConnCountry     string `json:"conn_country"`
	Country         string `json:"country"`
	ConnID          string `json:"conn_id"`
	CacheID         string `json:"cache_id"`
	Productmap      []struct {
		Value string `json:"value"`
		Key   string `json:"key"`
	} `json:"productmap"`
	Clientoffset string   `json:"clientoffset"`
	Platform     string   `json:"platform"`
	Version      string   `json:"version"`
	Products     []string `json:"products"`
	UserAgent    string   `json:"user_agent"`
	TCPPort      string   `json:"tcp_port"`
	Line         string   `json:"line"`
	AppID        string   `json:"app_id"`
	Revision     string   `json:"revision"`
}

type Download struct {
	Metadata
	IdentData `json:"identData"`
	Message   struct {
		CdnMinLatency        string  `json:"cdn_min_latency"`
		PlaybackID           string  `json:"playback_id"`
		CdnInitialLatency    string  `json:"cdn_initial_latency"`
		Cdn64KMedianLatency  string  `json:"cdn_64k_median_latency"`
		WasteFromAp          string  `json:"waste_from_ap"`
		MessageVersion       string  `json:"message_version"`
		CdnSocketReuse       string  `json:"cdn_socket_reuse"`
		CdnMaxLatency        string  `json:"cdn_max_latency"`
		CdnInitialBwEstimate int     `json:"cdn_initial_bw_estimate"`
		ApAvgBw              float64 `json:"ap_avg_bw"`
		ApMinLatency         string  `json:"ap_min_latency"`
		ApMaxLatency         string  `json:"ap_max_latency"`
		ApAvgLatency         float64 `json:"ap_avg_latency"`
		ContentSize          string  `json:"content_size"`
		ReqsFromAp           string  `json:"reqs_from_ap"`
		BytesFromEthernet    string  `json:"bytes_from_ethernet"`
		Cdn64KAvgLatency     float64 `json:"cdn_64k_avg_latency"`
		Cdn64KInitialLatency string  `json:"cdn_64k_initial_latency"`
		BytesFromCdn         string  `json:"bytes_from_cdn"`
		CdnAvgLatency        float64 `json:"cdn_avg_latency"`
		CdnAvgBw             float64 `json:"cdn_avg_bw"`
		TotalTime            string  `json:"total_time"`
		CdnDomain            string  `json:"cdn_domain"`
		BytesFromWifi        string  `json:"bytes_from_wifi"`
		MessageName          string  `json:"message_name"`
		CdnURIScheme         string  `json:"cdn_uri_scheme"`
		RequestType          string  `json:"request_type"`
		FileID               string  `json:"file_id"`
		ContentType          string  `json:"content_type"`
		CdnXCache            string  `json:"cdn_x_cache"`
		BytesFromCache       string  `json:"bytes_from_cache"`
		Bitrate              string  `json:"bitrate"`
		BytesFromUnknown     string  `json:"bytes_from_unknown"`
		NumCacheError        string  `json:"num_cache_error"`
		BytesFromCarrier     string  `json:"bytes_from_carrier"`
		Cdn64KMinLatency     string  `json:"cdn_64k_min_latency"`
		ErrorFromAp          string  `json:"error_from_ap"`
		Cdn64KMaxLatency     string  `json:"cdn_64k_max_latency"`
		ApInitialLatency     string  `json:"ap_initial_latency"`
		BytesFromAp          string  `json:"bytes_from_ap"`
		ApMedianLatency      string  `json:"ap_median_latency"`
		ErrorFromCdn         string  `json:"error_from_cdn"`
		WasteFromCdn         string  `json:"waste_from_cdn"`
		CdnMedianLatency     string  `json:"cdn_median_latency"`
		ReqsFromCdn          string  `json:"reqs_from_cdn"`
	} `json:"message"`
}

type EndSong struct {
	Metadata
	IdentData `json:"identData"`
	Message   struct {
		MsSeekfwd        string `json:"ms_seekfwd"`
		NStutter         string `json:"n_stutter"`
		PlaybackID       string `json:"playback_id"`
		PlayTrack        string `json:"play_track"`
		ReferrerVendor   string `json:"referrer_vendor"`
		MsSeekback       string `json:"ms_seekback"`
		PromotionType    string `json:"promotion_type"`
		MessageVersion   string `json:"message_version"`
		PlayerID         string `json:"player_id"`
		Shuffle          bool   `json:"shuffle"`
		DisplayTrack     string `json:"display_track"`
		Audiocodec       string `json:"audiocodec"`
		SourceStart      string `json:"source_start"`
		OfflineKey       bool   `json:"offline_key"`
		PlayContext      string `json:"play_context"`
		UnionPlayed      string `json:"union_played"`
		Provider         string `json:"provider"`
		CachedKey        bool   `json:"cached_key"`
		MaxContinous     string `json:"max_continous"`
		PLowbuffer       string `json:"p_lowbuffer"`
		StreamingRule    string `json:"streaming_rule"`
		AcceptedTc       string `json:"accepted_tc"`
		SourceEnd        string `json:"source_end"`
		MessageName      string `json:"message_name"`
		GaiaDevID        string `json:"gaia_dev_id"`
		Referer          string `json:"referer"`
		FileID           string `json:"file_id"`
		Offline          bool   `json:"offline"`
		BytesPlayed      string `json:"bytes_played"`
		ArtificialDelay  string `json:"artificial_delay"`
		ReasonStart      string `json:"reason_start"`
		NSeekback        string `json:"n_seekback"`
		Transition       string `json:"transition"`
		ReasonEnd        string `json:"reason_end"`
		MsPlayed         string `json:"ms_played"`
		ReferrerVersion  string `json:"referrer_version"`
		OfflineTimestamp string `json:"offline_timestamp"`
		NSeekfwd         string `json:"n_seekfwd"`
		IncognitoMode    bool   `json:"incognito_mode"`
	} `json:"message"`
}
type Gaia struct {
	Metadata
	IdentData `json:"identData"`
	Message   struct {
		DevID             string `json:"dev_id"`
		FallbackLen       string `json:"fallback_len"`
		EventType         string `json:"event_type"`
		Extra             string `json:"extra"`
		RemoteDevID       string `json:"remote_dev_id"`
		ProtoVer          string `json:"proto_ver"`
		NumRemotes        string `json:"num_remotes"`
		Reason            string `json:"reason"`
		MessageVersion    string `json:"message_version"`
		Context           string `json:"context"`
		FeatureIdentifier string `json:"feature_identifier"`
		ErrorCode         string `json:"error_code"`
		MessageName       string `json:"message_name"`
	} `json:"message"`
}
type Interaction struct {
	Metadata
	IdentData `json:"identData"`
	Message   struct {
		MessageName    string `json:"message_name"`
		PageIdentifier string `json:"page_identifier"`
		MessageVersion string `json:"message_version"`
		Intent         string `json:"intent"`
		ItemID         string `json:"item_id"`
		PageURI        string `json:"page_uri"`
		Type           string `json:"type"`
		InteractionID  string `json:"interaction_id"`
	} `json:"message"`
}
