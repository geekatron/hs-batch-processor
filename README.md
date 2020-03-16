# Assignment: Batch Processor
## Requirements
1. Ingest an input JSON file
    1. 'mixtape.json' is provided
    2. Input file contains three types
        1. User
            1. id string
            2. name string
        2. Playlist
            1. id string
            2. user_id string -> reference to User
            3. song_ids [] string -> Reference to Song
        3. Songs
            1. id string
            2. artist string
            3. title string
2. Ingest a changes file
    1. Created by the developer
    2. Developer chooses the format for the changes file -> text, YAML, CVS, JSON, etc.
    3. Application will apply the changes from the change file to the mixtape.json data and produce output.json
    4. The _changes_ file should include multiple changes in a single file
        1. Add a new playlist' the playlist should contain at least one song
        2. Remove a playlist
        3. Add an existing song to an existing playlist
3. Produce output.json
    3. Must have the same structure as the 'mixtape.json' input file
    4. Add a README describing how you would scale this application to handle very large input files and/or very large change files
    
## Evaluation Criteria
1. Are all the requirements met?
2. Does the application run successfully on a Mac or Linux?
3. Does your README explain how to run your application? 

# About the batch application
## Notes
I understand one could write a much simpler implementation of this batch processor by:
1. utilizing `ioutil.ReadAll` --> `json.Unmarshall` in conjunction with a `Mixtape struct`
1. building the `MixtapeIndex` from an instance of the `Mixtape struct`
1. utilizing `ioutil.ReadAll` --> `json.Unmarshall` in conjunction with a `Changes struct` to read all of the 
structures into memory.
1. enumerating the `Changes struct` and applying all the changes.
1. transform the `MixtapeIndex` maps into the `Mixtape struct`
1. utilizing `json.Marshal` in conjunction with the updated `Mixtape struct` and write this array of bytes to a file.

However, the above approach would not be very efficient or scalable since the `ioutil.ReadAll` functions have to load 
the entire contents of the file into memory. 
 
In the interest of efficiency and scalability I believe it was appropriate to write this batch processor to work in 
a streaming approach. This does not require loading the entire contents of the input or change files into memory, it 
deals with one object at a time. Further performance could be achieved with the use of Go routines.    
 
## Assumptions
1. Playlist ID is an auto-incrementing integer that has to be unique.
    1. The type based on the input file is a string, so in an ideal case we could use something like a UUID.
1. User ID is an auto-incrementing integer that has to be unique.
1. It's fine to have multiple playlists with the same songs.
    1. e.g. User is making a copy of a playlist to modify
1. Application will not add a Playlist that references an invalid song
    1. Could easily be changed to accommodate invalid song ids to be compatible with an eventual consistency environment.
1. Application should terminate if there is bad data provided via the Input source
1. Application should continue in the presence of a bad change but will not apply the change
1. If there is an issue writing the output file, the application will clean up.
 

## Requirements
1. Ensure the latest version of Go is installed and configured correctly
    1. [Getting Started with Go Lang](https://golang.org/doc/install)
1. Clone the repository to your local machine from [GitHub - Highspot Test](https://github.com/geekatron/hs-batch-processor)
    1. `git clone git@github.com:geekatron/hs-batch-processor.git`
1. Change into the root of the repository
1. Follow the steps below in `How to run the application`

## How to run the application

### How to test the application
1. Open your terminal 
1. Ensure you are in the root directory of the repository
1. Run `go test`

### How to build the application
1. Open your terminal 
1. Ensure you are in the root directory of the repository
1. Run `go build`

### How run the application
1. Execute using `go run`
    1. Open your terminal 
    1. Ensure you are in the root directory of the repository
    1. Run `go run main.go changes.go decode.go encode.go mixtape.go utils.go -i mixtape-data.json -c changes`
1. Executing the binary (hs-batch-processor)
    1. Ensure that you followed the steps in `How to build the application`
    2. Run `./hs-batch-processor -i mixtape-data.json -c changes`
        1. Instructions above are for Linux or OS X. For windows you'll be running an executable (`*.exe`)
        
### Flags for the Application
* `-i <filename>` - Used to specify the input filename. e.g. `-i mixtape-data.json`
    * If you do not specify the `-i` flag, the application will look for the file `mixtape.json`
* `-c <filename>` - Used to specify the changes filename to process e.g. `-c changes`
    * If you do not specify the `-c` flag, the application will look for the file `changes`
* `-p` - Used to specify if you would like to see performance metrics logged to `performance.log`. Off by default.

### How to use the `changes` file
The changes file for this application is meant to represent a stream of JSON data in the format:
* `[object,object,object][object,object]`

The Change objects must be separated by a comma (,). The changes file can accept multiple change sets, which are
separated by arrays ([]).

#### How to `Add New Playlist`?
##### Schema
`{ "playlist": {"user_id": string, "song_ids": [Array of type string] }}`
##### Example
`{ "playlist": {"user_id": "1", "song_ids": [ "6", "8", "11" ] }}`
##### Notes
* The User Id (user_id) and Song Id must exist in order for the change to be applied.

#### How to `Add Song to an Existing Playlist`?
##### Schema
`{"playlist_id": string, "song_id": string}`
##### Example
`{"playlist_id": "1", "song_id": "2"}`
##### Notes
* The Playlist Id (playlist_id) and Song Id (song_id) must exist in order for the change to be applied.

#### How to `Remove a Playlist`?
##### Schema
`{"remove_playlist_id": string }`
##### Example
`{"remove_playlist_id": "2" }`
##### Notes
* The Playlist Id (remove_playlist_id) must exist in order for the change to be applied.

# How to scale the application
## How to accommodate very large input files and/or very large change files?
It depends - please see below for some potential options:
1. As the input file size increases we could split the user, playlist and song data into three separate input files. 
Alternatively, we could write some pre-processing to scan the file in-order to find those three elements and process 
these elements concurrently. With either approach above, we could concurrently load User and Song objects as Playlist 
objects are dependant on User and Song ids.
2. The current implementation streams the data in so we could employ a similar approach as above where we can do some 
pre-processing to split up the changes into smaller change sets and spin up more change processes by using Go Routines. 
This strategy assumes that order of operation is not an issue; the current Change types shouldn't have a problem with 
this since it'll ignore additions to removed playlists & new playlists need Ids before being added to.
3. At one point in time the dataset may become too large for the memory space of a single system or the performance of 
building the index every time may exceed Non-functional requirements. At this point in time I would employ the use of
a database. In this approach, we could populate the database with the initial mixtape and from then on utilize the 
Change processor to stream changes into the database. We could extend the Change types to handle the creation of new
User and Song objects. We would modify decodeMixtape to concurrently Upsert entries into the database. Accordingly, we 
would update the apply functions to make changes to the Database instead of the MixtapeIndex. Results could be then 
queried through the database query language and optionally we could still dump an output of the entire database 
periodically, if necessary (probably backup purposes).
    1. This strategy could be further enhanced with vertical and horizontal partitioning the data based on usage.
    1. Additional caching could also be implemented in-order to improver the performance of read operations. 
    