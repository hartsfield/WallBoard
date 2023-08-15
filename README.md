# WallBoard - Ephemeral - No Signup - Message Board

A message board built using a component based architecture written in Go.

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

Now visit `http://localhost:4534` and add some posts. 

[Screencast from 2023-08-15 16-47-45.webm](https://github.com/hartsfield/WallBoard/assets/30379836/d548af58-397f-4a53-af69-9935842b770a)
