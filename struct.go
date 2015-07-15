package pipl

type PiplResponse struct {
	HttpStatusCode   int      `json:"@http_status_code"`
	VisibleSources   int      `json:"@visible_sources"`
	AvailableSources int      `json:"@available_sources"`
	SearchID         string   `json:"@search_id"`
	Person           Person   `json:"person"`
	PossiblePersons  []Person `json:"possible_persons"`
	Query            Person   `json:"query"`
	Warnings         []string `json:"warnings,omitempty"`
	Error            string   `json:"error,omitempty"`
}

type Person struct {
	Id             string         `json:"@id"`
	Search_pointer string         `json:"@search_pointer"`
	Match          float32        `json:"@match"`
	Dob            Dob            `json:"dob"`
	Gender         Gender         `json:"gender"`
	Ethnicity      Ethnicity      `json:"ethnicity"`
	Language       Language       `json:"language"`
	OriginCountry  OriginCountry  `json:"Country"`
	Addresses      []Address      `json:"addresses"`
	Relationships  []Relationship `json:"relationships"`
	Educations     []Education    `json:"educations"`
	Emails         []Email        `json:"emails"`
	Images         []Image        `json:"images"`
	Jobs           []Job          `json:"jobs"`
	Names          []Name         `json:"names"`
	Phones         []Phone        `json:"phones"`
	URLs           []URL          `json:"urls"`
	Usernames      []Username     `json:"usernames"`
	UserIDs        []UserID       `json:"user_ids"`
}

type DateRange struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

type Dob struct {
	DateRange DateRange `json:"date_range"`
	Display   string    `json:"display"`
}

type Gender struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Content  string `json:"content"`
}

type OriginCountry struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Country  string `json:"content"`
}

type Ethnicity struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Content  string `json:"content"`
}

type Language struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Language string `json:"language"`
	Region   string `json:"region"`
	Display  string `json:"display"`
}

type Username struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Content  string `json:"content"`
}

type UserID struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Content  string `json:"content"`
}

type Relationship struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Type     string `json:"@type"`
	Subtype  string `json:"@subtype"`
	Person
}

type Address struct {
	Inferred  bool   `json:"@inferred,omitempty"`
	Type      string `json:"@type"`
	Country   string `json:"country"`
	State     string `json:"state"`
	City      string `json:"city"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Apartment string `json:"apartment"`
	ZipCode   string `json:"zip_code"`
	POBox     string `json:"po_box"`
	Raw       string `json:"raw"`
	Display   string `json:"display"`
}

type Education struct {
	Inferred  bool      `json:"@inferred,omitempty"`
	DateRange DateRange `json:"date_range"`
	Degree    string    `json:"degree"`
	Display   string    `json:"display"`
	School    string    `json:"school"`
}

type Email struct {
	Inferred      bool   `json:"@inferred,omitempty"`
	Type          string `json:"@type"`
	Address       string `json:"address"`
	AddressMd5    string `json:"address_md5"`
	Disposable    string `json:"@disposable"`
	EmailProvider bool   `json:"@email_provider"`
}

type Image struct {
	Inferred       bool   `json:"@inferred,omitempty"`
	ThumbnailToken string `json:"thumbnail_token"`
	URL            string `json:"url"`
}

type Job struct {
	Inferred     bool      `json:"@inferred,omitempty"`
	Title        string    `json:"title"`
	Organization string    `json:"organization"`
	Industry     string    `json:"industry"`
	DateRange    DateRange `json:"date_range"`
	Display      string    `json:"display"`
}

type Name struct {
	Inferred bool   `json:"@inferred,omitempty"`
	Type     string `json:"@type"`
	First    string `json:"first"`
	Middle   string `json:"middle"`
	Last     string `json:"last"`
	Prefix   string `json:"prefix"`
	Suffix   string `json:"suffix"`
	Raw      string `json:"raw"`
	Display  string `json:"display"`
}

type Tag struct {
	Classification string `json:"@classification"`
	Content        string `json:"content"`
}

type URL struct {
	Inferred bool   `json:"@inferred,omitempty"`
	SourceID string `json:"@source_id"`
	Name     string `json:"@name"`
	Category string `json:"@category"`
	Domain   string `json:"@domain"`
	URL      string `json:"url"`
}

type Phone struct {
	Inferred             bool   `json:"@inferred,omitempty"`
	Type                 string `json:"@type"`
	CountryCode          int    `json:"country_code"`
	Number               int    `json:"number"`
	Extension            string `json:"extension"`
	Raw                  string `json:"raw"`
	Display              string `json:"display"`
	DisplayInternational string `json:"display_international"`
}
