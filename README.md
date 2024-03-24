# GOTEX: A simple tcp chat app write in Go #

### Features v1 ###

1. A user can connect to server by using a either a specific client or any other net tool such as netcat or telnet
2. First version of the app will only contain a chat room where all the users will be able to write and see what others write as well.
3. When a user either connects or disconnect from room, a message will be shown
4. On the server a log will be displayed with notifications about user connections. In a future version this log could be saved in a file.
5. It is 100% self-learning purpose.

### Components v1 ###
1. Server
2. Client - first iterations will use third party net tools -
3. Messages struct
4. Manage client connections by using routines
5. Read and Write messages by using channels. First iterations will broadcast the messages to other users (chat room)


   
