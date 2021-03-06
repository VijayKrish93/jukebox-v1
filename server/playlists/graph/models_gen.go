// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package playlists

import (
	"fmt"
	"io"
	"strconv"
)

type AddSongs struct {
	Name          string `json:"name"`
	LinkedAccount string `json:"linkedAccount"`
	PlaylistID    string `json:"playlistId"`
}

type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type NewPlaylist struct {
	Name   string `json:"name"`
	UserID string `json:"userID"`
}

type NewSong struct {
	Name       string `json:"name"`
	PlaylistID string `json:"playlistId"`
}

type Playlist struct {
	ID          string    `json:"id"`
	User        *User     `json:"user"`
	Name        string    `json:"name"`
	Account     *Accounts `json:"account"`
	Images      []*Image  `json:"images"`
	Description *string   `json:"description"`
}

func (Playlist) IsEntity() {}

type User struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Playlists []*Playlist `json:"playlists"`
}

func (User) IsEntity() {}

type Accounts string

const (
	AccountsApple   Accounts = "APPLE"
	AccountsSpotify Accounts = "SPOTIFY"
)

var AllAccounts = []Accounts{
	AccountsApple,
	AccountsSpotify,
}

func (e Accounts) IsValid() bool {
	switch e {
	case AccountsApple, AccountsSpotify:
		return true
	}
	return false
}

func (e Accounts) String() string {
	return string(e)
}

func (e *Accounts) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Accounts(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Accounts", str)
	}
	return nil
}

func (e Accounts) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
