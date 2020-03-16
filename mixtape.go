package main

import (
	"fmt"
	"strconv"
)

// MixtapeIndex structure is a collection of maps; map key is the id of the respective Type (User, Playlist, Song)
type MixtapeIndex struct {
	Users					map[string]*User
	Playlists				map[string]*Playlist
	Songs					map[string]*Song
	PlaylistIdUpperBound 	int
}

// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 		Structures related to the Mixtape input
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~

// Convenience structure if we are to Read the entire file into memory and then unmarshal contents
type Mixtape struct {
	Users 		[]User 		`json:"users"`
	Playlists 	[]Playlist	`json:"playlists"`
	Songs 		[]Song		`json:"songs"`
}

type User struct {
	Id 		string	`json:"id"`
	Name 	string	`json:"name"`
}

type Playlist struct {
	Id       string   `json:"id"`
	UserId  string   `json:"user_id"`
	SongIds []string `json:"song_ids"`
}

type Song struct {
	Id 		string	`json:"id"`
	Artist 	string	`json:"artist"`
	Title 	string	`json:"title"`
}


// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 		Functions for receiver of type Mixtape
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~

func (m *Mixtape) init() {
	m.Users = []User{}
	m.Playlists = []Playlist{}
	m.Songs = []Song{}
}

// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 		Functions for receiver of type MixtapeIndex
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~

// Initialize the variables of a new MixtapeIndex structure so that they aren't nil
func (m *MixtapeIndex) init() {
	m.Users = make(map[string]*User)
	m.Playlists = make(map[string]*Playlist)
	m.Songs = make(map[string]*Song)
	m.PlaylistIdUpperBound = -1
}

// Update the Upper Bound in the MixtapeIndex
func (mi *MixtapeIndex) generatePlaylistId() int {
	mi.PlaylistIdUpperBound += 1
	return mi.PlaylistIdUpperBound
}


// Determine if the specified User Id exists in the Mixtape
func (mi *MixtapeIndex) userExists(id string) bool {
	if uid, ok := mi.Users[id]; !ok {
		fmt.Println("User DNE ->", uid)
		return false
	}
	return true
}

// Determine if the specified Song Ids (Slice of type String) exist in the Mixtape
func (mi *MixtapeIndex) songsExist(sids []string) bool {
	for _, sid := range sids {
		if id, ok := mi.Songs[sid]; !ok {
			fmt.Println("Song DNE ->", id)
			// Song does not exist
			return false
		}
	}
	return true
}

// Determine if the specified Playlist Id exists in the Mixtape
func (mi *MixtapeIndex) playlistExists(id string)  bool{
	if _, ok := mi.Playlists[id]; !ok {
		fmt.Println("Playlist DNE -->", id)
		// Playlist does not exist, don't apply the change
		return false
	}
	return true
}

// Enumerate over the Mixtape data and create some index structures (maps)
func (m Mixtape) buildIndex() MixtapeIndex{
	mi := MixtapeIndex{}
	// Initialize the MixtapeIndex structure so its variables aren't nil
	mi.init()

	// Note for the below, as per the spec, `:=` re-uses the variable in the loop's scope. In the context of pointers,
	// it means it will be the same address in each operation.
	// Index the Users
	for i, _ := range m.Users {
		u := m.Users[i]
		mi.Users[u.Id] = &u
	}

	// Index the Playlists
	for i, _ := range m.Playlists {
		p := m.Playlists[i]
		mi.Playlists[p.Id] = &p

		// Keep the playlist upper bound index updated
		id, _ := strconv.Atoi(p.Id)
		fmt.Println(id)
		if id > mi.PlaylistIdUpperBound {
			mi.PlaylistIdUpperBound = id
		}
	}

	// Index the Songs
	for i, _ := range m.Songs{
		s := m.Songs[i]
		mi.Songs[s.Id] = &s
	}

	fmt.Printf("Mixtape Index: \n%#v\n", mi)

	return mi
}