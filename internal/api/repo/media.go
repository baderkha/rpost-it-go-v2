package repo

const (
	MediaTypePicture       = "photo"
	MediaTypeVideo         = "video"
	MediaTypeAudio         = "audio"
	MediaMaxAcceptableSize = 20 * 1024 * 1024 // 20mb max size
)

// Abstraction for media , this will manage where files are stored and provide with logical route path
type Media struct {
	ID             string // uuid
	AccountOwnerId string // accountid
	Type           string // video , audio ... etc
	Path           string // where it's put either through route , external service ...etc
	SizeBytes      uint64
}

// IMediaRepo : media repository that models basic crud for a media object , we use this infront of the actual media for meta info
// this way client can have the freedom to pull any media they want
type IMediaRepo interface {
	// FindById : find a sepcific media given the id
	FindById(id string) *Media
	// FindByAccountOwnerId : find all media owned by a person
	FindByAccountOwnerId(accountId string) *[]Media
	// CreateMedia : create media logical in the db
	CreateMedia(media *Media) (*Media, bool)
	// UpdateMedia : update media
	UpdateMedia(media *Media) (*Media, bool)
	// DeleteMediaByIdAndAccountId : delete specific media that belongs to someone
	DeleteMediaByIdAndAccountId(id string, accountId string) bool
}
