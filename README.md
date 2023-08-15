# WallBoard - Ephemeral - No Signup - GPLv3 - Message Board

Generic boiler plate code for use as an anonymous message board, or for building 
other types of social media applications. Built using a component based 
architecture written in Go ([bolt](https://github.com/hartsfield/bolt)), and 
provided under the GNU General Public License version 3 (GPLv3).

# https://walBoard.xyz/

# Requirements
## For Development
  - Go programming environment
  - Redis (v7+)
  - Only tested on Linux
## To Run
  - Redis (v7+)
  - Only tested on Linux
  - No binaries provided (yet)

# Instructions

Clone the repo, and run the following (assuming redis is running on the default port):

    go mod init example.com/m/v2
    go mod tidy
    chmod +x autoload.sh
    ./autoload.sh WallBoard 4534

Now visit `http://localhost:4534` and add some posts:

[example.webm](https://github.com/hartsfield/WallBoard/assets/30379836/326f0e8f-607c-468d-a657-3b294094a340)

# Road map
I plan on forking this to add tags, search, and user accounts, I'll post the fork here when I do. It'll be a re-write of https://tagmachine.xyz, one of my other projects. 

# Development
 - `autoload.sh` re-compiles and restarts the program. Run this file on save in your editor. For vim, I use `:au BufWritePost * silent! execute "!./autoload.sh wallboard 4534" | call timer_start(200, { tid -> execute(':redraw!')})`
