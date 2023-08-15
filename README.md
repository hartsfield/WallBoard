# WallBoard - Ephemeral - No Signup - Message Board

A message board built using a component based architecture written in Go.

# Requirements:
## For development:
  - Go programming environment
  - Redis (v7+)
  - Only tested on Linux
## To Run
  - Redis (v7+)
  - Only tested on Linux
  - No binaries provided (yet)

Clone the repo, and run the following (assuming redis is running on the default port):

    go mod init example.com/m/v2
    go mod tidy
    chmod +x autoload.sh
    ./autoload.sh WallBoard 4534

Now visit `http://localhost:4534` and add some posts. 
