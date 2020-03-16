package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 			Structures related to Changes
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
type Change interface {
	apply(mi *MixtapeIndex)
}

type NewPlaylistChange struct {
	Playlist 		Playlist	`json:"playlist"`
}

type RemovePlaylistChange struct {
	RemovePlaylistId	string	`json:"remove_playlist_id"`
}

type UpdatePlaylistChange struct {
	PlaylistId	string	`json:"playlist_id"`
	SongId		string	`json:"song_id"`
}

// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 		Functions for receiver of type Change (Interface)
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~

// NOTE -> 	We're passing in pointers to the below Change objects so that we don't need to receive a copy of them since
//			struct types are non-reference (a.k.a. value) types not reference types. Golang will create copies of the
//			structure then passing them around as arguments.

func (npc *NewPlaylistChange) apply(mi *MixtapeIndex) {
	// Check if User exists
	// If User does not exist, don't apply the change
	if !mi.userExists(npc.Playlist.UserId) {
		fmt.Printf("*WARNING* \n Can not apply New Playlist Change! User - DNE -> %v\n\n", npc.Playlist.UserId)
		return
	}

	// Check if Songs exist
	// If Song does not exist, don't apply the change
	if !mi.songsExist(npc.Playlist.SongIds) {
		fmt.Printf("*WARNING* \n Can not apply New Playlist Change! Song - DNE -> %v\n\n", npc.Playlist.SongIds)
		return
	}

	// Generate a Playlist Id
	npc.Playlist.Id = strconv.Itoa(mi.generatePlaylistId())
	fmt.Println(npc.Playlist.Id)
	// Add the playlist
	mi.Playlists[npc.Playlist.Id] = &npc.Playlist
}

func (rpc *RemovePlaylistChange) apply(mi *MixtapeIndex) {
	// Check if Playlist exists
	// Playlist does not exist, don't apply the change
	if !mi.playlistExists(rpc.RemovePlaylistId) {
		fmt.Printf("*WARNING* \n Can not apply Change! Playlist - DNE -> %v\n\n", rpc.RemovePlaylistId)
		return
	}

	// Remove the playlist from the MixtapeIndex
	delete(mi.Playlists, rpc.RemovePlaylistId)
}

func (upc *UpdatePlaylistChange) apply(mi *MixtapeIndex) {
	// Check if Songs exist
	// If Song does not exist, don't apply the change
	if !mi.songsExist([]string{upc.SongId}) {
		fmt.Printf("*WARNING* \n Can not apply Playlist Update Change! Song - DNE -> %v\n\n", upc.SongId)
		return
	}

	// Check if Playlist exists
	// Playlist does not exist, don't apply the change
	if !mi.playlistExists(upc.PlaylistId) {
		fmt.Printf("*WARNING* \n Can not apply Playlist Update Change! Playlist  - DNE -> %v\n\n", upc.PlaylistId)
		return
	}
	// Appending the song blindly onto the playlist since there is no requirement to prevent duplicates
	mi.Playlists[upc.PlaylistId].SongIds = append(mi.Playlists[upc.PlaylistId].SongIds, upc.SongId)
}

// With this approach we can test for they change type and apply them concurrently
func applyChangeFile(filename string, mi MixtapeIndex) {
	f := openFile(filename)
	defer f.Close()
	dec := json.NewDecoder(f)
	decodeAndApplyChanges(dec, &mi)
}

func (mi MixtapeIndex) persistChanges(filename string) {
	// TODO We could check to see if an Output file exists, if so back it up by renaming it w/ a timestamp
	// TODO	However, this comes with a caveat, if there is nothing to clean up the backups it could eventually consume
	// TODO all the free disk space of the system.
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println("*ERROR* Creating to Output File -->", filename)
		os.Exit(1)
	}

	// Start JSON Object (opening bracket)
	n, err := io.WriteString(f,"{ ")
	writeUserStream(mi.Users, f)
	f.WriteString(",")
	writePlaylistStream(mi.Playlists, f)
	f.WriteString(",")
	writeSongsStream(mi.Songs, f)
	// Close the JSON Object
	n, _ = io.WriteString(f,"}")
	if err != nil {
		fmt.Println("*ERROR* Writing to Output File!", n)
	}

	// TODO Handle Error Use Case -> If there are errors writing the Output file we should probably remove the file
	// TODO as it could be partially written. We don't want to pass along bad data.
}