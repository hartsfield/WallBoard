# WallBoard - Ephemeral - No Signup - GPLv3 - Message Board

Generic boiler plate code for use as an anonymous message board, or for building 
other types of social media applications. Built using a 
component based architecture written in Go 
([bolt](https://github.com/hartsfield/bolt)) and provided under the GNU General 
Public License version 3 (GPLv3).

# https://walBoard.xyz/

# Requirements:
## For Development
  - Go programming environment
  - Redis (v7+)
  - Only tested on Linux
## To Run
  - Redis (v7+)
  - Only tested on Linux
  - No binaries provided (yet)

# Instructions:

Clone the repo, and run the following (assuming redis is running on the default port):

    go mod init example.com/m/v2
    go mod tidy
    chmod +x autoload.sh
    ./autoload.sh WallBoard 4534

Now visit `http://localhost:4534` and add some posts:

[example.webm](https://github.com/hartsfield/WallBoard/assets/30379836/326f0e8f-607c-468d-a657-3b294094a340)
