# gpmdp-np
### Google Play Music Desktop Player - Now Playing

A tool to export the currently playing song (and relevant metadata) to a file

### Config

A configuration file, `config.toml`, must be present to run gpmdp-np. A sample configuration file is included.

**Configuration Keys**
| Key           | Usage         |
| ------------- | ------------- |
| `UpdateTextOnPause` | If true, the output text will be set to `''` while no song is playing.
| `DownloadTrackCover` | If true, the track's cover art will be downloaded to `cover.png`.

**Output Files**

At least one array entry of an `Output` must be specified. Feel free to include as many as you wish.

**Output Keys**
| Key           | Usage         |
| ------------- | ------------- |
| `Path`        | The relative path to the file. (Note: this should not be prefixed with `./` to indicate the current directory)
| `Format`      | A string representing the output this file should contain. See the `Format Types` section on what keys can be placed in this string.

**Format Types**

| Key           | Becomes       |
| ------------- | ------------- |
| `{{track}}`   | The title of the track |
| `{{artist}}`  | The artist of the track |
| `{{album}}`   | The name of the album |