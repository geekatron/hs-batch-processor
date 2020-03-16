package main

import (
	"reflect"
	"testing"
)

// Test initializing Mixtape struct
func TestMixtapeInit(t *testing.T) {
	m := Mixtape{}
	m.init()

	if m.Users == nil  {
		t.Fatalf("Expected empty Array of Type Users! Got nil! ")
	}
	if m.Playlists == nil  {
		t.Fatalf("Expected empty Array of Type Playlist! Got nil! ")
	}
	if m.Songs == nil  {
		t.Fatalf("Expected empty Array of Type Songs! Got nil! ")
	}
	if len(m.Users) != 0 && cap(m.Users) != 0  {
		t.Fatalf("Array of type User should be empty! \n" +
			"  Expected Length: 0, Capacity: 0\n  Actual Length: %d, Capacity: %d",
			len(m.Users), cap(m.Users))
	}

	if len(m.Playlists) != 0 && cap(m.Playlists) != 0  {
		t.Fatalf("Array of type Playlist should be empty! \n" +
			"  Expected Length: 0, Capacity: 0\n  Actual Length: %d, Capacity: %d",
			len(m.Playlists), cap(m.Playlists))
	}

	if len(m.Songs) != 0 && cap(m.Songs) != 0  {
		t.Fatalf("Array of type Song should be empty! \n" +
			"  Expected Length: 0, Capacity: 0\n  Actual Length: %d, Capacity: %d",
			len(m.Songs), cap(m.Songs))
	}
}

// Test initializing MixtapeIndex struct
func TestMixtapeIndexInit(t *testing.T) {
	mi := MixtapeIndex{}
	mi.init()

	if mi.Users == nil  {
		t.Fatalf("Expected empty Array of Type Users! Got nil! ")
	}
	if mi.Playlists == nil  {
		t.Fatalf("Expected empty Array of Type Playlist! Got nil! ")
	}
	if mi.Songs == nil  {
		t.Fatalf("Expected empty Array of Type Songs! Got nil! ")
	}
	if len(mi.Users) != 0  {
		t.Fatalf("Map of key type string, value of User should be empty! \n" +
			"  Expected Length: 0 \n  Actual Length: %d",
			len(mi.Users))
	}
	if len(mi.Playlists) != 0  {
		t.Fatalf("Map of key type string, value of Playlist should be empty! \n" +
			"  Expected Length: 0 \n  Actual Length: %d",
			len(mi.Playlists))
	}
	if len(mi.Songs) != 0  {
		t.Fatalf("Map of key type string, value of Song should be empty! \n" +
			"  Expected Length: 0 \n  Actual Length: %d",
			len(mi.Songs))
	}
	if mi.PlaylistIdUpperBound != -1 {
		t.Fatalf("Expected Upper Bound to be -1. Got %d", mi.PlaylistIdUpperBound)
	}
}

// Test generatePlaylistId on uninitialized MixtapeIndex
func TestGeneratePlaylistIdWUninitializedMixtapeIndex(t *testing.T) {
 	mi := MixtapeIndex{}
 	// TODO Should we allow generating a playlist ID if the MixtapeIndex is empty?
	i := mi.generatePlaylistId()
	// We should get 1 since Go defaults variable declarations to a default value; for int 0 is used
	if i != 1 {
		t.Fatalf("Expected generating Playlist Id! Expected 1, Got %d", i)
	}
}

// Test generatePlaylistId
func TestGeneratePlaylistId(t *testing.T) {
	mi := MixtapeIndex{}
	mi.init()
	if mi.PlaylistIdUpperBound != -1 {
		t.Fatalf("Expected PlaylistIdUpperBound -1, Got %d", mi.PlaylistIdUpperBound)
	}
	i := mi.generatePlaylistId()
	if i != 0 {
		t.Fatalf("Expected generating Playlist Id! Expected 0, Got %d", i)
	}
	if mi.PlaylistIdUpperBound != 0 {
		t.Fatalf("Expected PlaylistIdUpperBound 0, Got %d", mi.PlaylistIdUpperBound)
	}
}

// Test generatePlaylistId on MixtapeIndex with fragmented playlists (e.g. 1, 3, 5...)
// 	i.e. Playlists have been removed
func TestGeneratePlaylistIdWFragmentedPlaylists(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_fragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	if mi.PlaylistIdUpperBound != 5 {
		t.Fatalf("Expected PlaylistIdUpperBound 5, Got %d", mi.PlaylistIdUpperBound)
	}
	i := mi.generatePlaylistId()
	if i != 6 {
		t.Fatalf("Expected generating Playlist Id! Expected 6, Got %d", i)
	}
}

// Test generatePlaylistId on MixtapeIndex with non-fragmented Playlists (e.g. 1, 2, 3...)
//	i.e. Playlists have not been removed
func TestGeneratePlaylistIdWNonFragmentedPlaylists(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	if mi.PlaylistIdUpperBound != 2 {
		t.Fatalf("Expected PlaylistIdUpperBound 2, Got %d", mi.PlaylistIdUpperBound)
	}
	i := mi.generatePlaylistId()
	if i != 3 {
		t.Fatalf("Expected generating Playlist Id! Expected 3, Got %d", i)
	}
}

// Test userExists with valid user on uninitialized MixtapeIndex
func TestUserExistsWUninitializedMixtapeIndex(t *testing.T) {
	mi := MixtapeIndex{}

	ue := mi.userExists("0")
	if ue == true {
		t.Fatalf("Expected userExists false, Got %t", ue)
	}
}

// Test userExists with valid user
func TestUserExistsWValidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	id := "1"
	exists := mi.userExists(id)
	if exists == false {
		t.Fatalf("Expected userExists true, Got %t for Id %s", exists, id)
	}
}

// Test userExists with invalid user
func TestUserExistsWVInvalidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	id := "2"
	exists := mi.userExists(id)
	if exists == true {
		t.Fatalf("Expected userExists false, Got %t for Id %s", exists, id)
	}
}

// Test playlistExists with valid id on uninitialized MixtapeIndex
func TestPlaylistExistsWUninitializedMixtapeIndex(t *testing.T) {
	mi := MixtapeIndex{}

	exists := mi.playlistExists("0")
	if exists == true {
		t.Fatalf("Expected playlistExists false, Got %t", exists)
	}
}
// Test playlistExists with valid id
func TestPlaylistExistsWValidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	id := "1"
	exists := mi.playlistExists(id)
	if exists == false {
		t.Fatalf("Expected playlistExists true, Got %t for Id %s", exists, id)
	}
}

// Test playlistExists with invalid id
func TestPlaylistExistsWVInvalidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	id := "9"
	exists := mi.playlistExists(id)
	if exists == true {
		t.Fatalf("Expected playlistExists false, Got %t for Id %s", exists, id)
	}
}

// Test songsExist with valid id on uninitialized MixtapeIndex
func TestSongsExistWUninitializedMixtapeIndex(t *testing.T) {
	mi := MixtapeIndex{}
	id := "0"
	exists := mi.songsExist([]string{id})
	if exists == true {
		t.Fatalf("Expected songsExist false, Got %t", exists)
	}
}

// Test songsExist with valid id
func TestSongsExistWValidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	id := "1"
	exists := mi.playlistExists(id)
	if exists == false {
		t.Fatalf("Expected songsExist true, Got %t for Id %s", exists, id)
	}
}

// Test songsExist with invalid id
func TestSongsExistWVInvalidId(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	id := "9"
	exists := mi.songsExist([]string{id})
	if exists == true {
		t.Fatalf("Expected songsExist false, Got %t for Id %s", exists, id)
	}
}

// Test songsExist with multiple ids
func TestSongsExistsWMultipleValidIds(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	ids := []string{"1", "2"}
	exists := mi.songsExist(ids)
	if exists == false {
		t.Fatalf("Expected songsExist false, Got %t for Ids %s", exists, ids)
	}
}

// Test songsExist with multiple ids, invalid
func TestSongsExistsWMultipleValidAndInvalidIds(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_nonfragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()
	id := []string{"1", "2", "9"}
	exists := mi.songsExist(id)
	if exists == true {
		t.Fatalf("Expected songsExist true, Got %t for Id %s", exists, id)
	}
}


// Test buildIndex on uninitialized Mixtape
func TestBuildIndexWUninitializedMixtape(t *testing.T) {
	m := Mixtape{}
	mi := m.buildIndex()
	id := -1
	if mi.PlaylistIdUpperBound != id {
		t.Fatalf("Expected -1, Got %d ", mi.PlaylistIdUpperBound)
	}
	if len(mi.Users) != 0 {
		t.Fatalf("Expected Users length of 1, Got %d ", len(mi.Users))
	}
	if len(mi.Playlists) != 0 {
		t.Fatalf("Expected Playlists length of 1, Got %d ", len(mi.Playlists))
	}
	if len(mi.Songs) != 0 {
		t.Fatalf("Expected Songs length of 1, Got %d ", len(mi.Songs))
	}
}

// Test buildIndex on empty Mixtape
func TestBuildIndexWEmptyMixtape(t *testing.T) {
	m := Mixtape{}
	m.init()
	mi := m.buildIndex()
	id := -1
	if mi.PlaylistIdUpperBound != id {
		t.Fatalf("Expected -1, Got %d ", mi.PlaylistIdUpperBound)
	}
	if len(mi.Users) != 0 {
		t.Fatalf("Expected Users length of 1, Got %d ", len(mi.Users))
	}
	if len(mi.Playlists) != 0 {
		t.Fatalf("Expected Playlists length of 1, Got %d ", len(mi.Playlists))
	}
	if len(mi.Songs) != 0 {
		t.Fatalf("Expected Songs length of 1, Got %d ", len(mi.Songs))
	}
}


// Test buildIndex on Mixtape with fragmented Playlist Ids
func TestBuildIndexWMixtape(t *testing.T) {
	// Load the Sample Data
	filename := "./test/mixtape_fragmented_playlists.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	expectedUpperBound := 5
	expectedUserLength := 1
	expectedPlaylistLength := 2
	expectedSongsLength := 3
	if mi.PlaylistIdUpperBound != expectedUpperBound {
		t.Fatalf("Expected %d, Got %d ", expectedUpperBound, mi.PlaylistIdUpperBound)
	}
	if len(mi.Users) != expectedUserLength {
		t.Fatalf("Expected Users length of %d, Got %d ", expectedUserLength, len(mi.Users))
	}
	if len(mi.Playlists) != expectedPlaylistLength {
		t.Fatalf("Expected Playlists length of %d, Got %d ", expectedPlaylistLength, len(mi.Playlists))
	}
	if len(mi.Songs) != expectedSongsLength {
		t.Fatalf("Expected Songs length of %d, Got %d ", expectedSongsLength, len(mi.Songs))
	}

	expectedUserId := "1"
	expectedUserName := "Albin Jaye"
	actualUser := *mi.Users[expectedUserId]
	if actualUser.Id != expectedUserId || actualUser.Name != expectedUserName {
		t.Fatalf("Expected User with Id %s, Name %s --> Got Id: %s, Name: %s ",
			expectedUserId, expectedUserName, actualUser.Id, actualUser.Name)
	}

	expectedPlaylistId := "2"
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

	expectedSongId := "3"
	expectedArtist := "The Weeknd"
	expectedTitle := "Pray For Me"
	actualSong := *mi.Songs[expectedSongId]
	if actualSong.Id != expectedSongId || actualSong.Artist != expectedArtist ||
		actualSong.Title != expectedTitle {
		t.Fatalf("Expected Song with Id %s, Artist %s, Title %s --> Got Id: %s, Artist: %s, Title: %s ",
			expectedSongId, expectedArtist, expectedTitle, actualSong.Id, actualSong.Artist, actualSong.Title)
	}

}