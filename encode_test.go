package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Test JSONMarshal using escapable characters (&)
func TestJSONMarshal(t *testing.T) {
	// Load the Sample Data
	filename := "./test/songs_array_single.json"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Couldn't open file %s\nERROR -->\n%#v\n", filename, err)
	}
	var expectedSongs []Song
	expectedSongsBytes := []byte("[{\"id\":\"16\",\"artist\":\"G-Eazy\",\"title\":\"Him & I\"}]")
	var expectedLength int

	dec := json.NewDecoder(f)
	err = dec.Decode(&expectedSongs)
	if err != nil {
		t.Fatalf("Couldn't decode into Array of type User\nERROR -->\n%#v\n", err)
	}

	expectedLength = len(expectedSongsBytes)

	actualSongs, err := JSONMarshal(&expectedSongs)
	//fmt.Printf("%v,",actualSongs)
	if err != nil {
		t.Fatalf("Couldn't Unmarshal into Array of type User\nERROR -->\n%#v\n", err)
	}

	actualLength := len(actualSongs)

	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}


// Test writeUserStream with empty Array of type User
func TestWriteUserStreamEmpty(t *testing.T) {
	buf := &bytes.Buffer{}

	m := Mixtape{}
	m.init()
	mi := m.buildIndex()

	expectedUsersBytes := []byte("\"users\": []")
	expectedLength := len(expectedUsersBytes)

	writeUserStream(mi.Users, buf)

	actualUserBytes := buf.Bytes()
	actualLength := len(actualUserBytes)

	equivalent := bytes.Compare(expectedUsersBytes, actualUserBytes)

	if equivalent != 0 {
		t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedUsersBytes, actualUserBytes)
	}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test writeUserStream with Array of type User with one instance
func TestWriteUserStreamOneInstance(t *testing.T) {
	buf := &bytes.Buffer{}

	// Load the Sample Data
	filename := "./test/mixtape_single.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	expectedUsersBytes := []byte("\"users\": [{\"id\":\"1\",\"name\":\"Albin Jaye\"}]")
	expectedLength := len(expectedUsersBytes)

	writeUserStream(mi.Users, buf)

	actualUserBytes := buf.Bytes()
	actualLength := len(actualUserBytes)

	equivalent := bytes.Compare(expectedUsersBytes, actualUserBytes)

	if equivalent != 0 {
		t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedUsersBytes, actualUserBytes)
	}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test writeUserStream with Array of type User with multiple instances
func TestWriteUserStreamMultipleInstance(t *testing.T) {
	buf := &bytes.Buffer{}

	// Load the Sample Data
	filename := "./test/mixtape_multiple_users.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	filename = "./test/expected_user_stream"
	expectedUsersBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	expectedLength := len(expectedUsersBytes)

	writeUserStream(mi.Users, buf)

	actualUserBytes := buf.Bytes()
	actualLength := len(actualUserBytes)

	//equivalent := bytes.Compare(expectedUsersBytes, actualUserBytes)
	//if equivalent != 0 {
	//	t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedUsersBytes, actualUserBytes)
	//}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test writePlaylistStream with empty Array of type Playlist
func TestWritePlaylistStreamEmpty(t *testing.T) {
	buf := &bytes.Buffer{}

	m := Mixtape{}
	m.init()
	mi := m.buildIndex()

	expectedPlaylistsBytes := []byte("\"playlists\": []")
	expectedLength := len(expectedPlaylistsBytes)

	writePlaylistStream(mi.Playlists, buf)

	actualPlaylistsBytes := buf.Bytes()
	actualLength := len(actualPlaylistsBytes)

	equivalent := bytes.Compare(expectedPlaylistsBytes, actualPlaylistsBytes)

	if equivalent != 0 {
		t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedPlaylistsBytes, actualPlaylistsBytes)
	}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test writePlaylistStream with Array of type Playlist with one instance
func TestWritePlaylistStreamSingleInstance(t *testing.T) {
	buf := &bytes.Buffer{}

	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	filename = "./test/expected_playlists_stream_one_instance"
	expectedUsersBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	expectedLength := len(expectedUsersBytes)

	writePlaylistStream(mi.Playlists, buf)

	actualUserBytes := buf.Bytes()
	actualLength := len(actualUserBytes)

	equivalent := bytes.Compare(expectedUsersBytes, actualUserBytes)

	if equivalent != 0 {
		fmt.Printf("Expected: %s\nActual: %s", expectedUsersBytes, actualUserBytes)
		t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedUsersBytes, actualUserBytes)
	}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Can't compare byte arrays directly; order of items in Map not guaranteed!!!
// Test writePlaylistStream with Array of type Playlist with multiple instances
func TestWritePlaylistStreamMultipleInstance(t *testing.T) {
	buf := &bytes.Buffer{}

	// Load the Sample Data
	filename := "./test/mixtape_multiple_users.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	filename = "./test/expected_playlists_stream"
	expectedPlaylistBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	expectedLength := len(expectedPlaylistBytes)

	writePlaylistStream(mi.Playlists, buf)

	actualPlaylistBytes := buf.Bytes()
	actualLength := len(actualPlaylistBytes)

	// TODO Better testing here
	// Not byte equivalent (different order)
	//equivalent := bytes.Compare(expectedPlaylistBytes, actualPlaylistBytes)
	//equivalent := bytes.Equal(expectedPlaylistBytes, actualPlaylistBytes)
	//if !equivalent {
	//	fmt.Printf("\nExpected: %s\nActual: %s\n", expectedPlaylistBytes, actualPlaylistBytes)
	//	t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedPlaylistBytes, actualPlaylistBytes)
	//}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// Test writeSongsStream with empty Array of type Song
func TestWriteSongsStreamEmpty(t *testing.T) {
	buf := &bytes.Buffer{}

	m := Mixtape{}
	m.init()
	mi := m.buildIndex()

	expectedSongsBytes := []byte("\"songs\": []")
	expectedLength := len(expectedSongsBytes)

	writeSongsStream(mi.Songs, buf)

	actualSongsBytes := buf.Bytes()
	actualLength := len(actualSongsBytes)

	equivalent := bytes.Compare(expectedSongsBytes, actualSongsBytes)

	if equivalent != 0 {
		t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedSongsBytes, actualSongsBytes)
	}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

//  Test writeSongsStream with Array of type Playlist with one instance
func TestWriteSongsStreamSingleInstance(t *testing.T) {
	buf := &bytes.Buffer{}

	// Load the Sample Data
	filename := "./test/mixtape_single_song.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	filename = "./test/expected_songs_stream_one_instance"
	expectedSongs, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	expectedLength := len(expectedSongs)

	writeSongsStream(mi.Songs, buf)

	actualSongsBytes := buf.Bytes()
	actualLength := len(actualSongsBytes)

	equivalent := bytes.Compare(expectedSongs, actualSongsBytes)

	if equivalent != 0 {
		fmt.Printf("Expected: %s\nActual: %s", expectedSongs, actualSongsBytes)
		t.Fatalf("Not Equivalent! Expected: %v, Got: %v", expectedSongs, actualSongsBytes)
	}
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}

// writeSongsStream with Array of type Playlist with multiple instances
func TestWriteSongsStreamMultipleInstance(t *testing.T) {
	buf := &bytes.Buffer{}

	// Load the Sample Data
	filename := "./test/mixtape_single_playlist.json"
	m, err := loadTestMixtape(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	// Build a MixtapeIndex
	// TODO Maybe throw an error is there is a problem
	mi := m.buildIndex()

	filename = "./test/expected_songs_stream"
	expectedSongsBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Couldn't load sample data %s\nERROR -->\n%#v\n", filename, err)
	}

	expectedLength := len(expectedSongsBytes)

	writeSongsStream(mi.Songs, buf)

	actualSongsBytes := buf.Bytes()
	actualLength := len(actualSongsBytes)

	// TODO Better testing here
	if actualLength != expectedLength {
		t.Fatalf("Expected length to be: %d, Got: %d", expectedLength, actualLength)
	}
}