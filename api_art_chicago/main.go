package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ArtWork struct {
	Pagination struct {
		Total       int64  `json:"total"`
		Limit       int64  `json:"limit"`
		Offset      int64  `json:"offset"`
		TotalPage   int64  `json:"total_page"`
		CurrentPage int64  `json:"current_page"`
		NextUrl     string `json:"next_url"`
	} `json:"pagination"`
	Data []struct {
		ID        int         `json:"id"`
		APIModel  string      `json:"api_model"`
		APILink   string      `json:"api_link"`
		IsBoosted bool        `json:"is_boosted"`
		Title     string      `json:"title"`
		AltTitles interface{} `json:"alt_titles"`
		Thumbnail struct {
			Lqip    string `json:"lqip"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			AltText string `json:"alt_text"`
		} `json:"thumbnail"`
		MainReferenceNumber         string      `json:"main_reference_number"`
		HasNotBeenViewedMuch        bool        `json:"has_not_been_viewed_much"`
		BoostRank                   interface{} `json:"boost_rank"`
		DateStart                   int         `json:"date_start"`
		DateEnd                     int         `json:"date_end"`
		DateDisplay                 string      `json:"date_display"`
		DateQualifierTitle          string      `json:"date_qualifier_title"`
		DateQualifierID             int         `json:"date_qualifier_id"`
		ArtistDisplay               string      `json:"artist_display"`
		PlaceOfOrigin               string      `json:"place_of_origin"`
		Dimensions                  string      `json:"dimensions"`
		MediumDisplay               string      `json:"medium_display"`
		Inscriptions                interface{} `json:"inscriptions"`
		CreditLine                  string      `json:"credit_line"`
		CatalogueDisplay            interface{} `json:"catalogue_display"`
		PublicationHistory          interface{} `json:"publication_history"`
		ExhibitionHistory           interface{} `json:"exhibition_history"`
		ProvenanceText              interface{} `json:"provenance_text"`
		PublishingVerificationLevel string      `json:"publishing_verification_level"`
		InternalDepartmentID        int         `json:"internal_department_id"`
		FiscalYear                  int         `json:"fiscal_year"`
		FiscalYearDeaccession       interface{} `json:"fiscal_year_deaccession"`
		IsPublicDomain              bool        `json:"is_public_domain"`
		IsZoomable                  bool        `json:"is_zoomable"`
		MaxZoomWindowSize           int         `json:"max_zoom_window_size"`
		CopyrightNotice             string      `json:"copyright_notice"`
		HasMultimediaResources      bool        `json:"has_multimedia_resources"`
		HasEducationalResources     bool        `json:"has_educational_resources"`
		HasAdvancedImaging          bool        `json:"has_advanced_imaging"`
		Colorfulness                float64     `json:"colorfulness"`
		Color                       struct {
			H          int     `json:"h"`
			L          int     `json:"l"`
			S          int     `json:"s"`
			Percentage float64 `json:"percentage"`
			Population int     `json:"population"`
		} `json:"color"`
		Latitude               interface{}   `json:"latitude"`
		Longitude              interface{}   `json:"longitude"`
		Latlon                 interface{}   `json:"latlon"`
		IsOnView               bool          `json:"is_on_view"`
		OnLoanDisplay          interface{}   `json:"on_loan_display"`
		GalleryTitle           interface{}   `json:"gallery_title"`
		GalleryID              interface{}   `json:"gallery_id"`
		ArtworkTypeTitle       string        `json:"artwork_type_title"`
		ArtworkTypeID          int           `json:"artwork_type_id"`
		DepartmentTitle        interface{}   `json:"department_title"`
		DepartmentID           interface{}   `json:"department_id"`
		ArtistID               int           `json:"artist_id"`
		ArtistTitle            string        `json:"artist_title"`
		AltArtistIds           []interface{} `json:"alt_artist_ids"`
		ArtistIds              []int         `json:"artist_ids"`
		ArtistTitles           []string      `json:"artist_titles"`
		CategoryIds            []string      `json:"category_ids"`
		CategoryTitles         []string      `json:"category_titles"`
		TermTitles             []string      `json:"term_titles"`
		StyleID                interface{}   `json:"style_id"`
		StyleTitle             interface{}   `json:"style_title"`
		AltStyleIds            []interface{} `json:"alt_style_ids"`
		StyleIds               []interface{} `json:"style_ids"`
		StyleTitles            []interface{} `json:"style_titles"`
		ClassificationID       string        `json:"classification_id"`
		ClassificationTitle    string        `json:"classification_title"`
		AltClassificationIds   []string      `json:"alt_classification_ids"`
		ClassificationIds      []string      `json:"classification_ids"`
		ClassificationTitles   []string      `json:"classification_titles"`
		SubjectID              string        `json:"subject_id"`
		AltSubjectIds          []string      `json:"alt_subject_ids"`
		SubjectIds             []string      `json:"subject_ids"`
		SubjectTitles          []string      `json:"subject_titles"`
		MaterialID             string        `json:"material_id"`
		AltMaterialIds         []interface{} `json:"alt_material_ids"`
		MaterialIds            []string      `json:"material_ids"`
		MaterialTitles         []string      `json:"material_titles"`
		TechniqueID            interface{}   `json:"technique_id"`
		AltTechniqueIds        []interface{} `json:"alt_technique_ids"`
		TechniqueIds           []interface{} `json:"technique_ids"`
		TechniqueTitles        []interface{} `json:"technique_titles"`
		ThemeTitles            []string      `json:"theme_titles"`
		ImageID                string        `json:"image_id"`
		AltImageIds            []interface{} `json:"alt_image_ids"`
		DocumentIds            []interface{} `json:"document_ids"`
		SoundIds               []interface{} `json:"sound_ids"`
		VideoIds               []interface{} `json:"video_ids"`
		TextIds                []interface{} `json:"text_ids"`
		SectionIds             []interface{} `json:"section_ids"`
		SectionTitles          []interface{} `json:"section_titles"`
		SiteIds                []interface{} `json:"site_ids"`
		SuggestAutocompleteAll []struct {
			Input    []string `json:"input"`
			Contexts struct {
				Groupings []string `json:"groupings"`
			} `json:"contexts"`
			Weight int `json:"weight,omitempty"`
		} `json:"suggest_autocomplete_all"`
		SourceUpdatedAt string `json:"source_updated_at"`
		UpdatedAt       string `json:"updated_at"`
		Timestamp       string `json:"timestamp"`
	} `json:"data"`
}

func main() {

	log.Print("Starting application")
	os.Setenv("APP_PORT", ":8080")
	APP_PORT := os.Getenv("APP_PORT")
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/artwork", GetArtworkHandler)
	http.ListenAndServe(APP_PORT, serveMux)
}

func GetArtworkHandler(response http.ResponseWriter, request *http.Request) {

	page := request.URL.Query().Get("page")
	artwork, error := GetArtwork(page)

	if error != nil {
		log.Fatalf("Could not get artwork, error %v", error)
		return
	}

	encoder := json.NewEncoder(response)
	encoder.Encode(artwork)
}

func GetArtwork(page string) (*ArtWork, error) {

	request, requestError := http.Get(fmt.Sprintf("https://api.artic.edu/api/v1/artworks?page=%v", page))

	if requestError != nil {
		log.Fatal("Could not sending http request")
		return nil, errors.New("could not request artwork data")
	}

	defer request.Body.Close()
	bytes, readError := io.ReadAll(request.Body)

	if readError != nil {
		log.Fatal("Error reading response body")
		return nil, errors.New("could not read json message")
	}

	artwork := ArtWork{}
	error := json.Unmarshal(bytes, &artwork)

	if error != nil {
		log.Fatal("Error parsing json")
		return nil, errors.New("could not parse json")
	}

	return &artwork, nil
}
