package sightengine

// Response ...
type Response struct {
	Status  string  `json:"status"`
	Summary Summary `json:"summary,omitempty"`
	Request Request `json:"request"`
	Media   Media   `json:"media,omitempty"`

	// Nudity
	Nudity Nudity `json:"nudity,omitempty"`

	// WAD
	Weapon  float64 `json:"weapon,omitempty"`
	Alcohol float64 `json:"alcohol,omitempty"`
	Drugs   float64 `json:"drugs,omitempty"`

	// Face
	Face Face `json:"faces,omitempty"`
}

// Nudity ...
type Nudity struct {
	Raw        float64 `json:"raw"`
	Partial    float64 `json:"partial"`
	PartialTag string  `json:"partial_tag,omitempty"`
	Safe       float64 `json:"safe"`
}

// Face ...
type Face struct {
	LeftX     float64   `json:"x1"`
	TopY      float64   `json:"y1"`
	RightX    float64   `json:"x2"`
	BottomY   float64   `json:"y2"`
	Feature   Feature   `json:"features"`
	Attribute Attribute `json:"attributes"`
}

// Feature ...
type Feature struct {
	LeftEye struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"left_eye"`

	RightEye struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"right_eye"`

	NoseTip struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"nose_tip"`

	LeftMouthCorner struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"left_mouth_corner"`

	RightMouthCorner struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"right_mouth_corner"`
}

// Attribute ...
type Attribute struct {
	Female     float64 `json:"female"`
	Male       float64 `json:"male"`
	Minor      float64 `json:"minor"`
	Sunglasses float64 `json:"sunglasses"`
}

// Summary ...
type Summary struct {
	Action       string   `json:"action"`
	RejectProb   float64  `json:"reject_prob"`
	RejectReason []string `json:"reject_reason"`
}

// Request ...
type Request struct {
	ID         string  `json:"id"`
	Timestamp  float64 `json:"timestamp"`
	Operations int     `json:"operations"`
}

// Media ...
type Media struct {
	ID  string `json:"id"`
	URL string `json:"uri"`
}
