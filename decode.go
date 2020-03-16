package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Takes an input file and decodes the User, Playlist and Song structs
func decodeMixtape(f *os.File) MixtapeIndex {
	dec := json.NewDecoder(f)
	mi := MixtapeIndex{}
	mi.init()

	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%T: %v", t, t)
		if t, ok := t.(string); ok {
			switch strings.ToLower(t) {
			case "users":
				for u := range decodeUsers(dec) {
					fmt.Println("Decoding user!", u)
					mi.Users[u.Id] = u
				}
				fmt.Println("Done decoding user stream!")
			case "playlists":
				for p := range decodePlaylists(dec) {
					fmt.Println("Decoding Playlist!", p)
					mi.Playlists[p.Id] = p

					id, _ := strconv.Atoi(p.Id)
					fmt.Println(id)
					if id > mi.PlaylistIdUpperBound {
						mi.PlaylistIdUpperBound = id
					}
				}
				fmt.Println("Done decoding Playlist stream!")
			case "songs":
				for s := range decodeSongs(dec) {
					fmt.Println("Decoding Songs!", s)
					mi.Songs[s.Id] = s
				}
				fmt.Println("Done decoding Songs stream!")
			default:
				fmt.Println("Unknown Token!!!: ", t)
			}
		}
	}
	return mi
}

// Takes an input json.Decoder expecting it to be at the "users" token
// Returns a Channel of type User, returning the Decoded structure over the channel
func decodeUsers(d *json.Decoder) chan *User {
	c := make(chan *User)

	go func() {
		for {
			t, err := d.Token()
			if err == io.EOF {
				close(c)
				break
			}
			if err != nil {
				fmt.Printf("Error Decoding User!\nError: %#v\n", err)
				os.Exit(1)
			}
			// Beginning of an Object - try to Unmarshal it
			if s, ok := t.(json.Delim); ok {
				if s.String() == "]" {
					close(c)
					break
				} else if s.String() == "[" {
					for d.More() {
						var user User
						err := d.Decode(&user)
						if err != nil {
							fmt.Printf("Error decoding User Object!\nError --> %#v", err)
							break
						}
						c <- &user
					}
				}
			}
		}
	}()
	return c
}

// Takes an input json.Decoder expecting it to be at the "playlists" token
// Returns a Channel of type Playlist, returning the Decoded structure over the channel
func decodePlaylists(d *json.Decoder) chan *Playlist {
	c := make(chan *Playlist)

	go func() {
		for {
			t, err := d.Token()
			if err == io.EOF {
				close(c)
				break
			}
			if err != nil {
				fmt.Printf("Error Decoding User!\nError: %#v\n", err)
				os.Exit(1)
			}
			// Beginning of an Object - try to Unmarshal it
			if s, ok := t.(json.Delim); ok {
				if s.String() == "]" {
					close(c)
					break
				} else if s.String() == "[" {
					for d.More() {
						var playlist Playlist
						err := d.Decode(&playlist)
						if err != nil {
							fmt.Printf("Error decoding Playlist Object!\nError --> %#v", err)
							break
						}
						c <- &playlist
					}
				}
			}
		}
	}()
	return c
}

// Takes an input json.Decoder expecting it to be at the "songs" token
// Returns a Channel of type Song, returning the Decoded structure over the channel
func decodeSongs(d *json.Decoder) chan *Song {
	c := make(chan *Song)

	go func() {
		for {
			t, err := d.Token()
			if err == io.EOF {
				close(c)
				break
			}
			if err != nil {
				fmt.Printf("Error Decoding User!\nError: %#v\n", err)
				os.Exit(1)
			}
			// Beginning of an Object - try to Unmarshal it
			if s, ok := t.(json.Delim); ok {
				if s.String() == "]" {
					close(c)
					break
				} else if s.String() == "[" {
					for d.More() {
						var song Song
						err := d.Decode(&song)
						if err != nil {
							fmt.Printf("Error decoding Song Object!\nError --> %#v", err)
							break
						}
						c <- &song
					}
				}
			}
		}
	}()
	return c
}

// Takes a json.Decoder, expecting it to be at the beginning of a JSON Object to decode returning a concrete Change Type
// Takes a MixtapeIndex, which will be updated with the changes from the json.Decoder
// This function deals with a polyglot list of Change objects, decodes them to a concrete type and applies the change
// If a change fails, it will continue execution --> One could put these failed changes into some kind of queue
//for inspection (file, DB, Queue, etc) to understand where bad data is coming from.
func decodeAndApplyChanges(d *json.Decoder, mi *MixtapeIndex) {
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error w/ Change Token!\nError: %#v\n", err)
			os.Exit(1)
		}

		if s, ok := t.(json.Delim); ok {
			fmt.Println(s)
			// Remove the below check if we want to process multiple change streams [][][]
			//if s.String() == "]" {
			//	break
			//} else
			if s.String() == "[" || s.String() == "{" {
				for d.More() {
					fmt.Println(s.String())
					change, err := decodeChange(d)
					if err != nil {
						msg := err.Error()
						if strings.Contains(msg, "expected comma after array element") {
							break
						}
						fmt.Printf("WARNING: Could not decode object from stream!\nERROR -->\n%#v\n", err)
						continue
						//break
					}
					// No errors Decoding to Change Interface -> Apply the change
					change.apply(mi)
				}
			}
		}
	}
}

// Takes a json.Decoder, expecting it to be at the beginning of a JSON Object to decode returning a concrete Change Type
// This is used to decode the polyglot list of Objects the change parser is dealing with
func decodeChange(d *json.Decoder) (Change, error) {
	// Decoding the update into an interface for reuse
	//	Token will disappear (JSON Object) after Decode is called. We need to try and Decode it up to n times since
	//	we're dealing with a polyglot list of updates.
	var i interface{}
	e := d.Decode(&i)
	if e != nil {
		//fmt.Printf("*ERROR* Cannnot Decode change message into interface{}! \nERROR -->\n%#v\n", e)
		return nil, e
	}
	// Convert the Interface into a ByteArray; will be used to initialize new Reader
	ba, _ := json.Marshal(i)
	fmt.Printf("%s\n", ba)

	// Check for an empty object
	// 	For some reason an empty object will decode into the first available Decoder
	fmt.Printf("# of bytes: %d\n", len(ba))
	match, err := regexp.Match("(^[{]\\s*[}]$)", ba)
	if match {
		return nil, errors.New("Empty JSON Object! Discarding!")
	}

	// Try to create a concrete Change type
	change, err := createNewPlaylistChange(ba)
	if err == nil {
		return change, nil
	}
	change, err = createUpdatePlaylistChange(ba)
	if err == nil {
		return change, nil
	}
	change, err = createRemovePlaylistChange(ba)
	if err == nil {
		return change, nil
	}

	// Couldn't decode <- Unrecognized JSON Change object
	fmt.Printf("*WARNING* Change object could not be detected as recognized type!\nERROR -->\n%#v\n", err)
	return nil, err
}

// Create an instance of NewPlaylistChange from the Slice of type Byte
func createNewPlaylistChange(ba []byte) (Change, error) {
	var c NewPlaylistChange
	d := json.NewDecoder(bytes.NewReader(ba))
	d.DisallowUnknownFields()
	err := d.Decode(&c)
	if err != nil {
		//fmt.Printf("Error Decoding Type --> NewPlaylistChange\nError --> %#v\n", err)
		return nil, err
	}
	return &c, nil
}

// Create an instance of UpdatePlaylistChange from the Slice of type Byte
func createUpdatePlaylistChange(ba []byte) (Change, error) {
	var c UpdatePlaylistChange
	d := json.NewDecoder(bytes.NewReader(ba))
	d.DisallowUnknownFields()
	err := d.Decode(&c)
	if err != nil {
		//fmt.Printf("Error Decoding Type --> UpdatePlaylistChange\nError --> %#v\n", err)
		return nil, err
	}
	return &c, nil
}

// Create an instance of RemovePlaylistChange from the Slice of type Byte
func createRemovePlaylistChange(ba []byte) (Change, error) {
	var c RemovePlaylistChange
	d := json.NewDecoder(bytes.NewReader(ba))
	d.DisallowUnknownFields()
	err := d.Decode(&c)
	if err != nil {
		//fmt.Printf("Error Decoding Type --> RemovePlaylistChange\nError --> %#v\n", err)
		return nil, err
	}
	return &c, nil
}
