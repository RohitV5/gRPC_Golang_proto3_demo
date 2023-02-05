This project is to showcase the implementaion of grpc in golang.

Credit to TutorialEdge on YT https://www.youtube.com/watch?v=BdzYdN_Zd9Q&t=4s&ab_channel=TutorialEdge
Note: Lot of packages and commands are deprecated from the video. Maybe this repo will also have deprecated commands. Please do your own research.



gRPC uses protocol buffers as the format for sending and recievind data.

Rest uses json likewise SOAP uses XML.

Initialize a new go project
-go mod init github.com/rohitv5/grpc_golang_basic

To use protobuf with golang we need to install 
-go get -u google.golang.org/protobuf


-go get -u google.golang.org/grpc

Command to generate go files from proto
   



<!-- HOW TO SETUP PROTOC -->
https://developers.google.com/protocol-buffers/docs/gotutorial



1.copy the protoc file win64 version from the link https://github.com/protocolbuffers/protobuf/releases
2.Extract the contents of the zip file.
3.Bin directory contains the protoc.exe
4.Update the environment variable to point to the bin directory that contains the protoc.exe
5.Re-Open the cmd as the environment variable take effect only on new command window
6.run=> protoc --go_out=.\chat --go-grpc_out=.\chat   chat.proto 



Apps installed by go command are present in C:\Users\verma\go\bin. Add this to path.


Some issues resolved through these links
https://github.com/grpc/grpc-go/issues/3794
https://stackoverflow.com/questions/71777702/service-compiling-successfully-but-message-structs-not-generating-grpc-go