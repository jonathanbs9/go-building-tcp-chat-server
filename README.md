# go-building-tcp-chat-server

## :information_source: Description
We build a TCP chat server using Go, which enables clients to communicate with each other. 
We are working with Go's “net” package which very well supports TCP, as well we'll be using channels and goroutines.

## :information_source: Test

:heavy_check_mark: In one terminal

```
go build .
```

```
./go-building-tcp-chat-server.exe
```

:heavy_check_mark: In second terminal

```
telnet localhost 8888
```

## :information_source: Commands

:arrow_forward: '/nickname <name>' - Get a name, otherwise user will stay anonymous.
  
:arrow_forward: '/join <name>'     - Join a room, if room doesn't exist, the new room will be created. User can be only in one room at the same time.
  
:arrow_forward: '/rooms'           - Show list of available rooms to join.

:arrow_forward: '/msg <msg>'       - Broadcast message to everyone in a room.
  
:arrow_forward: '/quit'            - Disconnects from the chat server.
