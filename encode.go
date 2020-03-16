package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func writeUserStream(m map[string]*User, w io.Writer) {
	buf := &bytes.Buffer{}
	index := 0
	size := len(m)
	buf.WriteString("\"users\": [")

	for i, v := range m {
		fmt.Println("Writing interface{] with index: ", i)
		//b, err := json.Marshal(v)
		b, err := JSONMarshal(v)
		if err != nil {
			fmt.Println("Error converting interface{} to []byte! ", err)
			os.Exit(1)
		}
		buf.Write(b)
		if index < size - 1 {
			buf.WriteString(",")
		}
		n, err := buf.WriteTo(w)
		fmt.Printf("Wrote %d bytes to Writer from buffer!", n)
		if err != nil {
			fmt.Printf("Error writing Byte Buffer to Writer!\nERROR -->\n%#v\n", err)
			os.Exit(1)
		}
		index++
	}

	buf.WriteString("]")
	buf.WriteTo(w)
}

func writePlaylistStream(m map[string]*Playlist, w io.Writer) {
	var buf bytes.Buffer
	index := 0
	size := len(m)
	buf.WriteString("\"playlists\": [")

	for i, v := range m {
		fmt.Println("Writing Playlist with index: ", i)
		//b, err := json.Marshal(v)
		b, err := JSONMarshal(v)
		if err != nil {
			fmt.Println("Error converting Playlist to []byte! ", err)
			os.Exit(1)
		}
		buf.Write(b)
		if index < size - 1 {
			buf.WriteString(",")
		}
		n, err := buf.WriteTo(w)
		fmt.Printf("Wrote %d bytes to Writer from buffer!", n)
		if err != nil {
			fmt.Printf("Error writing Byte Buffer to Writer!\nERROR -->\n%#v\n", err)
			os.Exit(1)
		}
		index++
	}
	buf.WriteString("]")
	buf.WriteTo(w)
}

// Enumerate over the map of Pointers to type Song and write them to the provided Writer.
// Function will create the songs json attribute and an array as the value:
// 	`"songs": [ ]`
func writeSongsStream(m map[string]*Song, w io.Writer) {
	var buf bytes.Buffer
	index := 0
	size := len(m)
	buf.WriteString("\"songs\": [")

	for i, v := range m {
		fmt.Println("Writing Songs with index: ", i)
		//b, err := json.Marshal(v)
		b, err := JSONMarshal(v)
		if err != nil {
			fmt.Println("Error converting Song to []byte! ", err)
			os.Exit(1)
		}
		buf.Write(b)
		if index < size - 1 {
			buf.WriteString(",")
		}
		n, err := buf.WriteTo(w)
		fmt.Printf("Wrote %d bytes to Writer from buffer!", n)
		if err != nil {
			fmt.Printf("Error writing Byte Buffer to Writer!\nERROR -->\n%#v\n", err)
			os.Exit(1)
		}
		index++
	}
	buf.WriteString("]")
	buf.WriteTo(w)
}


// Create a custom JSONMarshal function in-order to work around the default of SetEscapeHtml(true).
// e.g. We don't want `&` to become `\u0026`
func JSONMarshal(t interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(t)
	// Make sure to trim the last character in the buffer as json.Encode adds a new line character
	buf.Truncate(buf.Len() - 1)
	return buf.Bytes(), err
}
