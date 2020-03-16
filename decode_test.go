package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 		Tests related to parsing Input into Maps
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~

// Test decodeUsers with invalid input stream
func TestDecodeUsersStreamInvalidInput(t *testing.T) {
	// Load the Sample Data
	filename := "./test/json_object_empty.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 0

	var actualUsers []User
	c := decodeUsers(dec)
	for u := range c {
		actualUsers = append(actualUsers, *u)
	}
	actualLength := len(actualUsers)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test decodeUsers with empty input array
func TestDecodeUsersStreamEmpty(t *testing.T) {
	// Load the Sample Data
	filename := "./test/json_array_empty.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 0

	var actualUsers []User
	c := decodeUsers(dec)
	for u := range c {
		actualUsers = append(actualUsers, *u)
	}
	actualLength := len(actualUsers)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test decodeUsers with vaild input stream (one user)
func TestDecodeUsersStreamSingleInstance(t *testing.T) {
	// Load the Sample Data
	filename := "./test/users_array_single.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 1
	expectedUserId := "6"
	expectedUserName := "Ryo Daiki"

	var actualUsers []User
	c := decodeUsers(dec)
	for u := range c {
		actualUsers = append(actualUsers, *u)
	}
	actualLength := len(actualUsers)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}

	for _, u := range actualUsers {
		if u.Id == expectedUserId && u.Name == expectedUserName {
			return
		}
	}
	t.Fatalf("Couldn't find the expected User instance in the Array of type User!!!\n")
}

// Test decodeUsers with vaild input stream (multiple users)
func TestDecodeUsersStreamMultipleInstances(t *testing.T) {
	// Load the Sample Data
	filename := "./test/users_array.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 7
	expectedUserId := "3"
	expectedUserName := "Ankit Sacnite"

	var actualUsers []User
	c := decodeUsers(dec)
	for u := range c {
		actualUsers = append(actualUsers, *u)
	}
	actualLength := len(actualUsers)

	var expectedUsers []User

	// Seek back to the begiing of the file
	ret, err := f.Seek(0, 0)
	if err != nil {
		t.Fatalf("Couldn't seek to beginning of file! %s\nERROR -->\n%#v\n", filename, err)
	}
	fmt.Println("Seeking back to :", ret)
	err = dec.Decode(&expectedUsers)
	if err != nil {
		t.Fatalf("Error decoding JSON to Array of Users! \nERROR -->\n%#v\n", err)
	}

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}

	if !reflect.DeepEqual(expectedUsers, actualUsers) {
		t.Fatalf("Expected Users (%v) IS NOT EQUVALENT to Actual Users (%v)!", expectedUsers, actualUsers)
	}

	for _, u := range actualUsers {
		if u.Id == expectedUserId && u.Name == expectedUserName {
			return
		}
	}
	t.Fatalf("Couldn't find the expected User instance in the Array of type User!!!\n")
}

// Test decodePlaylists with invalid input stream
func TestDecodePlaylistStreamInvalidInput(t *testing.T) {
	// Load the Sample Data
	filename := "./test/json_object_empty.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 0

	var actualPlaylists []Playlist
	c := decodePlaylists(dec)
	for p := range c {
		actualPlaylists = append(actualPlaylists, *p)
	}
	actualLength := len(actualPlaylists)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test decodePlaylists with empty  input stream
func TestDecodePlaylistsStreamEmpty(t *testing.T) {
	// Load the Sample Data
	filename := "./test/json_array_empty.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 0

	var actualPlaylists []Playlist
	c := decodePlaylists(dec)
	for p := range c {
		actualPlaylists = append(actualPlaylists, *p)
	}
	actualLength := len(actualPlaylists)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test decodePlaylists with vaild input stream (one playlist)
func TestDecodePlaylistsSingleStreamInstance(t *testing.T) {
	// Load the Sample Data
	filename := "./test/playlists_array_single.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 1
	expectedPlaylistId := "2"
	expectedUserId := "3"
	expectedSongIds := []string{"6", "8", "11"}

	var actualPlaylists []Playlist
	c := decodePlaylists(dec)
	for p := range c {
		actualPlaylists = append(actualPlaylists, *p)
	}
	actualLength := len(actualPlaylists)
	actualPlaylist := actualPlaylists[0]

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
	if len(expectedSongIds) != len(actualPlaylist.SongIds) {
		t.Fatalf("Expected playlist with %d songs, Got: %v", len(expectedSongIds), len(actualPlaylist.SongIds))
	}
	if expectedPlaylistId != actualPlaylist.Id || expectedUserId != actualPlaylist.UserId {
		t.Fatalf("Expected Playlist Id: %s, User Id: %s, Got Playlist Id: %s, User Id: %s",
			expectedPlaylistId, expectedUserId, actualPlaylist.Id, actualPlaylist.UserId)
	}
	if !reflect.DeepEqual(expectedSongIds, actualPlaylist.SongIds) {
		t.Fatalf("Expected Playlist with Songs %s, --> Got Songs: %s",
			expectedSongIds, actualPlaylist.SongIds)
	}
}

// Test decodePlaylists with valid input stream (multiple playlists)
func TestDecodePlaylistsStreamMultipleInstances(t *testing.T) {
	// Load the Sample Data
	filename := "./test/playlists_array.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 3
	expectedPlaylistId := "2"
	expectedUserId := "3"
	expectedPlaylistLength := 3

	var actualPlaylists []Playlist
	c := decodePlaylists(dec)
	for p := range c {
		actualPlaylists = append(actualPlaylists, *p)
	}
	actualLength := len(actualPlaylists)

	var expectedPlaylists []Playlist

	// Seek back to the begiing of the file
	ret, err := f.Seek(0, 0)
	if err != nil {
		t.Fatalf("Couldn't seek to beginning of file! %s\nERROR -->\n%#v\n", filename, err)
	}
	fmt.Println("Seeking back to :", ret)
	err = dec.Decode(&expectedPlaylists)
	if err != nil {
		t.Fatalf("Error decoding JSON to Array of type Playlist! \nERROR -->\n%#v\n", err)
	}

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}

	if !reflect.DeepEqual(expectedPlaylists, actualPlaylists) {
		t.Fatalf("Expected Users (%v) IS NOT EQUVALENT to Actual Users (%v)!", expectedPlaylists, actualPlaylists)
	}

	// Probably unnecessary with the use of DeepEqual
	for _, p := range actualPlaylists {
		if p.Id == expectedPlaylistId && p.UserId == expectedUserId && len(p.SongIds) == expectedPlaylistLength {
			return
		}
	}
	t.Fatalf("Couldn't find the expected Playlist instance in the Array of type Playlist!!!\n")
}

// Test decodeSongs with invalid input stream
func TestDecodeSongsStreamInvalidInput(t *testing.T) {
	// Load the Sample Data
	filename := "./test/json_object_empty.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 0

	var actualSongs []Song
	c := decodeSongs(dec)
	for s := range c {
		actualSongs = append(actualSongs, *s)
	}
	actualLength := len(actualSongs)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}
// Test decodeSongs with empty input stream
func TestDecodeSongsStreamEmpty(t *testing.T) {
	// Load the Sample Data
	filename := "./test/json_array_empty.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 0

	var actualSongs []Song
	c := decodeSongs(dec)
	for s := range c {
		actualSongs = append(actualSongs, *s)
	}
	actualLength := len(actualSongs)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test decodeSongs with vaild input stream (one song)
func TestDecodeSongsStreamSingleInstance(t *testing.T) {
	// Load the Sample Data
	filename := "./test/songs_array_single.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 1
	expectedSongId := "16"
	expectedArtist := "G-Eazy"
	expectedTitle := "Him & I"

	var actualSongs []Song
	c := decodeSongs(dec)
	for s := range c {
		actualSongs = append(actualSongs, *s)
	}
	actualSong := actualSongs[0]
	actualLength := len(actualSongs)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
	if expectedSongId != actualSong.Id || expectedArtist != actualSong.Artist || expectedTitle != actualSong.Title {
		t.Fatalf("Expected Song Id: %s, Artist: %s, Title: %s, Got Song Id: %s, Artist: %s, Title: %s",
			expectedSongId, expectedArtist, expectedTitle, actualSong.Id, actualSong.Artist, actualSong.Title)
	}
}

// Test decodeSongs with vaild input stream (multiple songs)
func TestDecodeSongssStreamMultipleInstances(t *testing.T) {
	// Load the Sample Data
	filename := "./test/songs_array.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	expectedLength := 40
	expectedSongId := "16"
	expectedArtist := "G-Eazy"
	expectedTitle := "Him & I"

	var actualSongs []Song
	c := decodeSongs(dec)
	for s := range c {
		actualSongs = append(actualSongs, *s)
	}
	actualLength := len(actualSongs)
	var expectedSongs []Song

	// Seek back to the begiing of the file
	ret, err := f.Seek(0, 0)
	if err != nil {
		t.Fatalf("Couldn't seek to beginning of file! %s\nERROR -->\n%#v\n", filename, err)
	}
	fmt.Println("Seeking back to :", ret)
	err = dec.Decode(&expectedSongs)
	if err != nil {
		t.Fatalf("Error decoding JSON to Array of type Playlist! \nERROR -->\n%#v\n", err)
	}

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}

	if !reflect.DeepEqual(expectedSongs, actualSongs) {
		t.Fatalf("Expected Songs (%v) IS NOT EQUVALENT to Actual Songs (%v)!", expectedSongs, actualSongs)
	}

	// Probably unnecessary with the use of DeepEqual
	for _, s := range actualSongs {
		if expectedSongId != s.Id && expectedArtist != s.Artist && expectedTitle != s.Title {
			return
		}
	}
	t.Fatalf("Couldn't find the expected Song instance in the Array of type Song!!!\n")
}

// Test decodeMixtape with vaild input stream (multiple songs)
func TestDecodeMixtape(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_fragmented_playlists.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	var expectedMixtape Mixtape
	var expectedMixtapeIndex MixtapeIndex
	expectedUserLength := 1
	expectedPlaylistsLength := 2
	expectedSongsLength := 3

	actualMixtapeIndex := decodeMixtape(f)
	actualUsersLength := len(actualMixtapeIndex.Users)
	actualPlaylistsLength := len(actualMixtapeIndex.Playlists)
	actualSongsLength := len(actualMixtapeIndex.Songs)

	expectedSongId := "2"
	expectedArtist := "Zedd"
	expectedTitle := "The Middle"

	// Seek back to the begiing of the file
	ret, err := f.Seek(0, 0)
	if err != nil {
		t.Fatalf("Couldn't seek to beginning of file! %s\nERROR -->\n%#v\n", filename, err)
	}
	fmt.Println("Seeking back to :", ret)
	err = dec.Decode(&expectedMixtape)
	if err != nil {
		t.Fatalf("Error decoding JSON to MixtapeIndex! \nERROR -->\n%#v\n", err)
	}
	expectedMixtapeIndex = expectedMixtape.buildIndex()

	if expectedPlaylistsLength != actualPlaylistsLength {
		t.Fatalf("Expected Playlists length to be: %d, Got: %d",
			expectedPlaylistsLength, actualPlaylistsLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Songs length to be: %d, Got: %d",
			expectedSongsLength, actualSongsLength)
	}
	if expectedUserLength != actualUsersLength {
		t.Fatalf("Expected Users length to be: %d, Got: %d",
			expectedUserLength, actualUsersLength)
	}
	if !reflect.DeepEqual(expectedMixtapeIndex, actualMixtapeIndex) {
		t.Fatalf("Expected Songs (%v) IS NOT EQUVALENT to Actual Songs (%v)!", expectedMixtapeIndex, actualMixtapeIndex)
	}

	// Checking to make sure a specific item exists
	for _, s := range actualMixtapeIndex.Songs {
		if expectedSongId != s.Id && expectedArtist != s.Artist && expectedTitle != s.Title {
			return
		}
	}
	t.Fatalf("Couldn't find the expected Song instance in the Array of type Song!!!\n")
}


// 	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~
// 		Tests related decoding Changes
//	~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~~-=-~

// Test decodeChange with empty change
func TestDecodeChangeEmptyInput(t *testing.T) {
	// Load the Sample Data
	filename := "./test/changes_empty"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	change, err := decodeChange(dec)

	if err == nil {
		t.Fatalf("Expected an error! Got the following change: %v", change)
	}
}

// Test decodeChange with invalid NewPlaylistChange
func TestDecodeChangeInvalidNewPlaylistChange(t *testing.T) {
	// Load the Sample Data
	filename := "./test/changes_invalid_newplaylistchange"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	change, err := decodeChange(dec)

	if err == nil {
		t.Fatalf("Expected an error! Got the following change: %v", change)
	}
}

// Test decodeChange with valid NewPlaylistChange
func TestDecodeChangeNewPlaylistChange(t *testing.T) {
	// Load the Sample Data
	filename := "./test/changes_newplaylistchange"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)
	change, err := decodeChange(dec)


	expectedPlaylistId := ""
	expectedUserId := "1"
	expectedSongIds := []string{"6", "8", "11"}

	if actualChange, ok := change.(*NewPlaylistChange); ok {
		if expectedPlaylistId != actualChange.Playlist.Id || expectedUserId != actualChange.Playlist.UserId ||
			!reflect.DeepEqual(expectedSongIds, actualChange.Playlist.SongIds) {
			t.Fatalf("Expected Playlist Id: %s, User Id: %s, Songs: %v\n Got Playlist Id: %s, User Id: %s, Songs: %v\n",
				expectedPlaylistId, expectedUserId, expectedSongIds, actualChange.Playlist.Id,
				actualChange.Playlist.UserId, actualChange.Playlist.SongIds)
		}
	} else {
		t.Fatalf("Could not cast the Change to type NewPlaylistChange!\n")
	}
}

// Test decodeChange with invalid RemovePlaylistChange
func TestDecodeChangeInvalidRemovePlaylistChange(t *testing.T) {
	// Load the Sample Data
	filename := "./test/changes_invalid_removeplaylistchange"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	change, err := decodeChange(dec)

	if err == nil {
		t.Fatalf("Expected an error! Got the following change: %v", change)
	}
}


// Test decodeChange with valid RemovePlaylistChange
func TestDecodeChangeRemovePlaylistChange(t *testing.T) {
	// Load the Sample Data
	filename := "./test/changes_removeplaylistchange"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)
	change, err := decodeChange(dec)


	expectedPlaylistId := "2"

	if actualChange, ok := change.(*RemovePlaylistChange); ok {
		if expectedPlaylistId != actualChange.RemovePlaylistId {
			t.Fatalf("Expected Playlist Id: %s\n Got Playlist Id: %s\n",
				expectedPlaylistId, actualChange.RemovePlaylistId)
		}
	} else {
		t.Fatalf("Could not cast the Change to type RemovePlaylistChange!\n")
	}
}

// Test decodeChange with invalid UpdatePlaylistChange
func TestDecodeChangeInvalidUpdatePlaylistChange(t *testing.T) {
	// Load the Sample Data
	filename := "./test/changes_invalid_updateplaylistchange"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)

	change, err := decodeChange(dec)

	if err == nil {
		t.Fatalf("Expected an error! Got the following change: %v", change)
	}
}

// decodeChange with valid UpdatePlaylistChange
func TestDecodeChangeUpdatePlaylistChange(t *testing.T) {
	// Load the Sample Data
	filename := "./test/changes_updateplaylistchange"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	dec := json.NewDecoder(f)
	change, err := decodeChange(dec)

	expectedPlaylistId := "1"
	expectedSongId := "2"

	if actualChange, ok := change.(*UpdatePlaylistChange); ok {
		if expectedPlaylistId != actualChange.PlaylistId || expectedSongId != actualChange.SongId {
			t.Fatalf("Expected Playlist Id: %s, Song Id: %v\n Got Playlist Id: %s, Song Id: %v\n",
				expectedPlaylistId, expectedSongId, actualChange.PlaylistId, actualChange.SongId)
		}
	} else {
		t.Fatalf("Could not cast the Change to type UpdatePlaylistChange!\n")
	}
}

