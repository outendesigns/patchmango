package models

type VersionStruct struct {
    Version      string     `json:"version"`
    ReleaseDate  string     `json:"release_date"`
    Files        []string   `json:"files"`
}

type VersionsListStruct struct {
    Versions  []VersionStruct  `json:"versions"`
}