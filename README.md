# GO-Web-Practice-Projects

## Building Go File

```
pwd
/Users/akshay/GO-Projects/01-HTTP-Simple-Server
akshay@localhost 01-HTTP-Simple-Server % ls -lrt
total 8
-rw-r--r--  1 akshay  akshay  398 12 Sep 22:15 main.go

akshay@localhost 01-HTTP-Simple-Server % go build -o ./bin/ main.go
akshay@localhost 01-HTTP-Simple-Server % ls -lrt
total 8
-rw-r--r--  1 akshay  akshay  398 12 Sep 22:15 main.go
drwxr-xr-x  3 akshay  akshay   96 13 Sep 14:51 bin

akshay@localhost 01-HTTP-Simple-Server % cd bin
akshay@localhost bin % ls -lrt
total 12032
-rwxr-xr-x  1 akshay  akshay  6157344 13 Sep 14:51 main

akshay@localhost bin % ./main 
server started in port : 8080
```
# GO-Fundamental-Learning
