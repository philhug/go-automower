package automower

const fleetUrl = "https://bffmobile-live.hfs2.dss.husqvarnagroup.net"
const authUrl = "https://iam-api.dss.husqvarnagroup.net"
const trackUrl = "https://amc-api.dss.husqvarnagroup.net"
const authUrlPath = authUrl + "/api/v3/"
const trackUrlPath = trackUrl + "/v1/"
const fleetUrlPath = fleetUrl + "/app/v1/"
const authUrlToken = authUrlPath + "token"
const trackUrlMowers = trackUrlPath + "mowers"
const fleetUrlAssets = fleetUrlPath + "assets"

const (
	mowerActionPark  = "park"
	mowerActionStop  = "stop"
	mowerActionStart = "start"
)

//MOWER_ACTION_PARK_UNTIL_NEXT_START = "park"

// STATUSlist
// ERROR
// UNKNOWN
// PAUSED
// STOPPED
// PARKED
// DISCONNECTED
// CONNECTING
// UNDEFINED
// DISABLED
// UPDATING
// IN_OPERATION_IN_CHARGING_STATION
// IN_OPERATION_IN_MOWING_AREA

type loginRequestAttributes struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type loginRequestBody struct {
	Attributes loginRequestAttributes `json:"attributes"`
	Type       string                 `json:"type"`
}

type loginRequest struct {
	Data loginRequestBody `json:"data"`
}

type messageData struct {
	ID         string                 `json:"id"`
	Attributes map[string]interface{} `json:"attributes"`
	Type       string                 `json:"type"`
}

type message struct {
	Data messageData `json:"data"`
}

type MowerLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	GPSStatus string  `json:"gpsStatus"`
}

type MowerStatus struct {
	MowerStatus            string `json:"mowerStatus"`
	ValueFound             bool   `json:"valueFound"`
	NextStartSource        string `json:"nextStartSource"`
	NextStartTimestamp     int64  `json:"nextStartTimestamp"`
	BatteryPercent         int    `json:"batteryPercent"`
	StoredTimestamp        int64  `json:"storedTimestamp"`
	LastErrorCode          int    `json:"lastErrorCode"`
	LastErrorCodeTimestamp int    `json:"lastErrorCodeTimestamp"`
	OperatingMode          string `json:"operatingMode"`
	Connected              bool   `json:"connected"`
	ShowAsDisconnected     bool   `json:"showAsDisconnected"`

	// full status only /status
	CachedSettingsUUID string          `json:"cachedSettingsUUID,omitempty"`
	LastLocations      []MowerLocation `json:"lastLocations,omitempty"`
}

type Mower struct {
	ID              string      `json:"id"`
	ValueFound      bool        `json:"valueFound"`
	Status          MowerStatus `json:"status"`
	Model           string      `json:"model"`
	Name            string      `json:"name"`
	SoftwareVersion string      `json:"swPackageVersionString"`
}

type actionRequest struct {
	Action   string `json:"action"`
	Duration string `json:"duration"`
}

type mowerGeoFence struct {
}
