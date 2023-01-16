# cli tool

## usage:
just write `kotogif`. It'll write path to gif or error :)

### flags:

| Flag           | Type   | Description                                    |
|:---------------|--------|-----------------------------------------------:|
| --help         | none\* | shows help; ignores other flags; equals `help` |
| -o, --output   | string | output file; `output.gif` by default           |
| --tmp          | string | temp directory; `./temp` by default            |
| --not-del-temp | none\* | doesn't delete temp file if put                |
| --overwrite    | none\* | overwrites output file if it exists            |
| --url          | string | url of site (idk why)                          |
| --useragent    | string | `User-Agent` header content\*\*                |
| -t, --timeout  | int    | count of seconds to get gifs; 10 by default    |
| --debug        | none\* | turns on debug log                             |
| --version      | none\* | prints version and exits                       |

*\*none is bool - putting it equals true*
*\*\*by default `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:98.0) Gecko/20100101 Firefox/98.0`*

> Note: if you write path to other video format file in --output, ffmpeg
> will convert to this video format, but this utile made for GIFs :)
