package main

import (
	"reflect"
	"testing"
)

// Test NewPlaylistChange with invalid User
func TestApplyNewPlaylistChangeWInvalidUser(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_simple_no_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := NewPlaylistChange{Playlist:Playlist{
		Id:      "",
		UserId:  "2",
		SongIds: []string{"1", "2"},
	}}

	change.apply(&mi)

	// If there was no bad data, a new playlist with the id of 6 should have been created
	// Considered a success if the pointer to mi.Playlists["6"] is nil
	if mi.Playlists["6"] != nil{
		t.Fatalf("Expected playlist to be nil! Got: %v", mi.Playlists["6"])
	}

	expectedLength := 0
	actualLength := len(mi.Playlists)
	if expectedLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedLength, actualLength)
	}
}

// Test NewPlaylistChange with invalid Song
func TestApplyNewPlaylistChangeWInvalidSong(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_simple_no_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := NewPlaylistChange{Playlist:Playlist{
		Id:      "",
		UserId:  "1",
		SongIds: []string{"1", "5"},
	}}

	change.apply(&mi)

	expectedLength := 0
	actualLength := len(mi.Playlists)
	if expectedLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedLength, actualLength)
	}
}

// Test creating an empty playlist
func TestApplyNewPlaylistChangeWNoSongs(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_simple_no_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := NewPlaylistChange{Playlist:Playlist{
		Id:      "",
		UserId:  "1",
		SongIds: []string{},
	}}

	change.apply(&mi)

	expectedLength := 1
	actualLength := len(mi.Playlists)
	if expectedLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedLength, actualLength)
	}

	expectedPlaylistId := "0"
	expectedPlaylistUserId := "1"
	expectedSongs := []string{}
	actualPlaylist := *mi.Playlists[expectedPlaylistId]
	if actualPlaylist.Id != expectedPlaylistId || actualPlaylist.UserId != expectedPlaylistUserId {
		t.Fatalf("Expected Playlist with Id %s, User Id %s --> Got Id: %s, User Id: %s ",
			expectedPlaylistId, expectedPlaylistUserId, actualPlaylist.Id, actualPlaylist.UserId)
	}
	if !reflect.DeepEqual(expectedSongs, actualPlaylist.SongIds) {
		t.Fatalf("Expected Playlist with Songs %s, --> Got Songs: %s",
			expectedSongs, actualPlaylist.SongIds)
	}
}

// Test NewPlaylistChange with one Song
func TestApplyNewPlaylistChangeWValidSong(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_simple_no_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := NewPlaylistChange{Playlist:Playlist{
		Id:      "",
		UserId:  "1",
		SongIds: []string{"1"},
	}}

	change.apply(&mi)

	expectedLength := 1
	actualLength := len(mi.Playlists)
	if expectedLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedLength, actualLength)
	}

	expectedPlaylistId := "0"
	expectedPlaylistUserId := "1"
	expectedSongs := []string{"1"}
	actualPlaylist := *mi.Playlists[expectedPlaylistId]
	if actualPlaylist.Id != expectedPlaylistId || actualPlaylist.UserId != expectedPlaylistUserId {
		t.Fatalf("Expected Playlist with Id %s, User Id %s --> Got Id: %s, User Id: %s ",
			expectedPlaylistId, expectedPlaylistUserId, actualPlaylist.Id, actualPlaylist.UserId)
	}
	if !reflect.DeepEqual(expectedSongs, actualPlaylist.SongIds) {
		t.Fatalf("Expected Playlist with Songs %s, --> Got Songs: %s",
			expectedSongs, actualPlaylist.SongIds)
	}
}

// Test NewPlaylistChange with multiple Songs
func TestApplyNewPlaylistChangeWValidSongs(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_simple_no_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := NewPlaylistChange{Playlist:Playlist{
		Id:      "",
		UserId:  "1",
		SongIds: []string{"1", "2", "3"},
	}}

	change.apply(&mi)

	expectedLength := 1
	actualLength := len(mi.Playlists)
	if expectedLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedLength, actualLength)
	}

	expectedPlaylistId := "0"
	expectedPlaylistUserId := "1"
	expectedSongs := []string{"1", "2", "3"}
	actualPlaylist := *mi.Playlists[expectedPlaylistId]
	if actualPlaylist.Id != expectedPlaylistId || actualPlaylist.UserId != expectedPlaylistUserId {
		t.Fatalf("Expected Playlist with Id %s, User Id %s --> Got Id: %s, User Id: %s ",
			expectedPlaylistId, expectedPlaylistUserId, actualPlaylist.Id, actualPlaylist.UserId)
	}
	if !reflect.DeepEqual(expectedSongs, actualPlaylist.SongIds) {
		t.Fatalf("Expected Playlist with Songs %s, --> Got Songs: %s",
			expectedSongs, actualPlaylist.SongIds)
	}
}

// Test RemovePlaylistChange with invalid Playlist Id
func TestApplyRemovePlaylistChangeWInvalidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := RemovePlaylistChange{RemovePlaylistId: "0"}

	change.apply(&mi)

	expectedPlaylistLength := 1
	actualLength := len(mi.Playlists)
	if expectedPlaylistLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedPlaylistLength, actualLength)
	}
}

// Test RemovePlaylistChange with valid Playlist Id
func TestApplyRemovePlaylistChangeWValidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := RemovePlaylistChange{RemovePlaylistId: "1"}

	change.apply(&mi)

	expectedPlaylistLength := 0
	actualLength := len(mi.Playlists)
	if expectedPlaylistLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedPlaylistLength, actualLength)
	}
}

// Test UpdatePlaylistChange with invalid Playlist Id
func TestApplyUpdatePlaylistChangeWInvalidPlaylistId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := UpdatePlaylistChange{
		PlaylistId: "0",
		SongId:     "1",
	}

	change.apply(&mi)

	expectedPlaylistLength := 1
	actualLength := len(mi.Playlists)
	if expectedPlaylistLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedPlaylistLength, actualLength)
	}
}

// Test UpdatePlaylistChange with invalid Song Id
func TestApplyUpdatePlaylistChangeWInvalidSongId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := UpdatePlaylistChange{
		PlaylistId: "1",
		SongId:     "5",
	}

	change.apply(&mi)

	expectedPlaylistLength := 1
	actualLength := len(mi.Playlists)
	if expectedPlaylistLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedPlaylistLength, actualLength)
	}

	expectedPlaylistId := "1"
	expectedPlaylistSongsLength := 1
	actualPlaylistSongsLength := len(mi.Playlists[expectedPlaylistId].SongIds)
	if expectedPlaylistSongsLength != actualPlaylistSongsLength {
		t.Fatalf("Expected playlist with %d songs, Got: %v", expectedPlaylistSongsLength, actualPlaylistSongsLength)
	}
}

// Test UpdatePlaylistChange with valid data
func TestApplyUpdatePlaylistChange(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := UpdatePlaylistChange{
		PlaylistId: "1",
		SongId:     "2",
	}

	change.apply(&mi)

	expectedPlaylistLength := 1
	actualLength := len(mi.Playlists)
	if expectedPlaylistLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedPlaylistLength, actualLength)
	}

	expectedPlaylistId := "1"
	expectedPlaylistSongsLength := 2
	actualPlaylistSongsLength := len(mi.Playlists[expectedPlaylistId].SongIds)
	if expectedPlaylistSongsLength != actualPlaylistSongsLength {
		t.Fatalf("Expected playlist with %d songs, Got: %v", expectedPlaylistSongsLength, actualPlaylistSongsLength)
	}
}

// Test UpdatePlaylistChange with valid data
func TestApplyUpdatePlaylistChangeDuplicateSong(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	change := UpdatePlaylistChange{
		PlaylistId: "1",
		SongId:     "1",
	}

	change.apply(&mi)

	expectedPlaylistLength := 1
	actualLength := len(mi.Playlists)
	if expectedPlaylistLength != actualLength {
		t.Fatalf("Expected playlist length: %d, Got: %v", expectedPlaylistLength, actualLength)
	}

	expectedPlaylistId := "1"
	expectedPlaylistSongsLength := 2
	expectedSongIds := []string{"1", "1"}
	actualPlaylist := *mi.Playlists[expectedPlaylistId]
	actualPlaylistSongsLength := len(mi.Playlists[expectedPlaylistId].SongIds)
	if expectedPlaylistSongsLength != actualPlaylistSongsLength {
		t.Fatalf("Expected playlist with %d songs, Got: %v", expectedPlaylistSongsLength, actualPlaylistSongsLength)
	}
	if !reflect.DeepEqual(expectedSongIds, actualPlaylist.SongIds) {
		t.Fatalf("Expected Playlist with Songs %s, --> Got Songs: %s",
			expectedSongIds, actualPlaylist.SongIds)
	}
}

// THESE TESTS ARE FOR THE DECODE AND APPLY CHANGES FUNCTION
// Test invalid change to make sure it wasn't applied (use simple mixtape)
func TestApplyChangeFileInvald(t *testing.T) {
	// Load the Sample Data
	mixtapeFilename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(mixtapeFilename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", mixtapeFilename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	changesFilename := "./test/changes_multiple_invalid"
	applyChangeFile(changesFilename, mi)

	expectedPlaylistLength := 1
	expectedSongsLength := 3
	expectedUsersLength := 1

	actualPlaylistLength := len(mi.Playlists)
	actualSongsLength := len(mi.Songs)
	actualUsersLength := len(mi.Users)
	if expectedPlaylistLength != actualPlaylistLength {
		t.Fatalf("Expected Playlist Length: %d, Got: %d", expectedPlaylistLength, actualPlaylistLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Songs Length: %d, Got: %d", expectedSongsLength, actualSongsLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Users Length: %d, Got: %d", expectedUsersLength, actualUsersLength)
	}
}

// Test NewPlaylistChange,RemovePlaylistChange,UpdatePlaylistChange with multiple valid changes
func TestApplyChangeFileValid(t *testing.T) {
	// Load the Sample Data
	mixtapeFilename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(mixtapeFilename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", mixtapeFilename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	changesFilename := "./test/changes_multiple_valid"
	applyChangeFile(changesFilename, mi)

	expectedPlaylist := Playlist{
		Id:      "2",
		UserId:  "1",
		SongIds: []string{"1", "2", "3"},
	}
	expectedPlaylistLength := 1
	expectedSongsLength := 3
	expectedUsersLength := 1

	actualPlaylistLength := len(mi.Playlists)
	actualSongsLength := len(mi.Songs)
	actualUsersLength := len(mi.Users)
	actualPlaylist := mi.Playlists["2"]

	if expectedPlaylistLength != actualPlaylistLength {
		t.Fatalf("Expected Playlist Length: %d, Got: %d", expectedPlaylistLength, actualPlaylistLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Songs Length: %d, Got: %d", expectedSongsLength, actualSongsLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Users Length: %d, Got: %d", expectedUsersLength, actualUsersLength)
	}
	if expectedPlaylist.Id != actualPlaylist.Id {
		t.Fatalf("Expected Playlist Id Length: %s, Got: %s", expectedPlaylist.Id, actualPlaylist.Id)
	}
	if expectedPlaylist.UserId != actualPlaylist.UserId {
		t.Fatalf("Expected Playlist Id Length: %s, Got: %s", expectedPlaylist.Id, actualPlaylist.Id)
	}
	if !reflect.DeepEqual(expectedPlaylist.SongIds, actualPlaylist.SongIds) {
		t.Fatalf("Expected Playlist Songs: %v, Got: %v", expectedPlaylist.SongIds, actualPlaylist.SongIds)
	}
}

// Test NewPlaylistChange,RemovePlaylistChange,UpdatePlaylistChange with valid & invalid changes
func TestApplyChangeFileValidAndInvalid(t *testing.T) {
	// Load the Sample Data
	mixtapeFilename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(mixtapeFilename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", mixtapeFilename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	changesFilename := "./test/changes_multiple_valid_and_invalid"
	applyChangeFile(changesFilename, mi)

	expectedPlaylist := Playlist{
		Id:      "2",
		UserId:  "1",
		SongIds: []string{"1", "2", "3"},
	}
	expectedPlaylistLength := 1
	expectedSongsLength := 3
	expectedUsersLength := 1

	actualPlaylistLength := len(mi.Playlists)
	actualSongsLength := len(mi.Songs)
	actualUsersLength := len(mi.Users)
	actualPlaylist := mi.Playlists["2"]

	if expectedPlaylistLength != actualPlaylistLength {
		t.Fatalf("Expected Playlist Length: %d, Got: %d", expectedPlaylistLength, actualPlaylistLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Songs Length: %d, Got: %d", expectedSongsLength, actualSongsLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Users Length: %d, Got: %d", expectedUsersLength, actualUsersLength)
	}
	if expectedPlaylist.Id != actualPlaylist.Id {
		t.Fatalf("Expected Playlist Id Length: %s, Got: %s", expectedPlaylist.Id, actualPlaylist.Id)
	}
	if expectedPlaylist.UserId != actualPlaylist.UserId {
		t.Fatalf("Expected Playlist Id Length: %s, Got: %s", expectedPlaylist.Id, actualPlaylist.Id)
	}
	if !reflect.DeepEqual(expectedPlaylist.SongIds, actualPlaylist.SongIds) {
		t.Fatalf("Expected Playlist Songs: %v, Got: %v", expectedPlaylist.SongIds, actualPlaylist.SongIds)
	}
}

// Test NewPlaylistChange,RemovePlaylistChange,UpdatePlaylistChange with multiple valid change sets (e.g. [][])
func TestApplyChangeFileValidMultipleChangeSets(t *testing.T) {
	// Load the Sample Data
	mixtapeFilename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(mixtapeFilename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", mixtapeFilename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	changesFilename := "./test/changes_multiple_valid_streams"
	applyChangeFile(changesFilename, mi)

	expectedPlaylists := []Playlist{
		Playlist{
			Id:      "2",
			UserId:  "1",
			SongIds: []string{"1", "2", "3"},
		},
		Playlist{
			Id:     "3",
			UserId: "1",
			SongIds: []string{"2", "3"},
			},
		}

	expectedPlaylistLength := 2
	expectedSongsLength := 3
	expectedUsersLength := 1

	actualPlaylistLength := len(mi.Playlists)
	actualSongsLength := len(mi.Songs)
	actualUsersLength := len(mi.Users)

	if expectedPlaylistLength != actualPlaylistLength {
		t.Fatalf("Expected Playlist Length: %d, Got: %d", expectedPlaylistLength, actualPlaylistLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Songs Length: %d, Got: %d", expectedSongsLength, actualSongsLength)
	}
	if expectedSongsLength != actualSongsLength {
		t.Fatalf("Expected Users Length: %d, Got: %d", expectedUsersLength, actualUsersLength)
	}

	for _, expectedPlaylist := range expectedPlaylists{
		actualPlaylist := mi.Playlists[expectedPlaylist.Id]
		if expectedPlaylist.Id != actualPlaylist.Id {
			t.Fatalf("Expected Playlist Id Length: %s, Got: %s", expectedPlaylist.Id, actualPlaylist.Id)
		}
		if expectedPlaylist.UserId != actualPlaylist.UserId {
			t.Fatalf("Expected Playlist Id Length: %s, Got: %s", expectedPlaylist.Id, actualPlaylist.Id)
		}
		if len(expectedPlaylist.SongIds) != len(actualPlaylist.SongIds) {
			t.Fatalf("Expected Playlist Songs: %v, Got: %v", len(expectedPlaylist.SongIds), len(actualPlaylist.SongIds))
		}
	}
}