package models

// Videos

type DatabaseVideo struct {
	// Visible metadata
	Id          int64  `bson:"_id" json:"id"`
	Description string `bson:"description" json:"description"`
	Series      string `bson:"series" json:"series"`
	Public      bool   `bson:"public" json:"public"`
	CreatorId   int64  `bson:"creator_id" json:"creator_id"`

	Likes int64 `bson:"likes"`

	// Backend metadata
	Tags       []string `bson:"tags" json:"tags"`
	Modifiers  []string `bson:"modifiers" json:"modifiers"`
	StorageKey string   `bson:"storage_key" json:"storage_key"`
}

type InitialClassificationResult struct {
	BadTopics uint8 `bson:"bad_topics" json:"bad_topics"`
}

// Users

type DatabaseUser struct {
	// Visible metadata
	Id       int64  `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`

	// Backend metadata
	Interests map[string]int64 `bson:"interests" json:"interests"`
}

type UserLogin struct {
	Id       int64  `bson:"id" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

// Generic user interaction

type DatabaseWatchEvent struct {
	VideoId int64 `bson:"video_id" json:"video_id"`
	UserId  int64 `bson:"user_id" json:"user_id"`
}

type DatabaseLikeEvent struct {
	VideoId int64 `bson:"video_id" json:"video_id"`
	UserId  int64 `bson:"user_id" json:"user_id"`
}
