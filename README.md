
# How To Run Code

## Locally

### prerequisites
Create database using given SQL script located at - mysql_database/database.sql  
Update your database credential at - src/env/local.env  
If you want your request go to MockClient then set IsMock=true located in file - src/constants/constants.go

### Run REST Server 
Open terminal in project root directory and run command   
<pre> go run src/server/restServer/restServer.go </pre>
Postman collection located at - postman/UserApi.postman_collection.json

### Run gRPC Server 
Open terminal in project root directory and run command   
<pre> go run src/server/grpcServer/grpcServer.go </pre>
BloomRPC proto file located at - bloomRPC/userApi.proto
## Docker Container

### Prerequisites
Change DATABASE_HOST_URL="db:3306" at src/env/local.env    
Run this command fetch id_rsa file from local machine  
<pre> bash prep.sh </pre> 

### Run Command 
<pre>
docker compose build   
docker compose up 
</pre>

## Get public endpoint using Ngrock 

### Prerequisites
Create account at ngrock official website https://ngrok.com/   
You will get auth token.  

### Run Command 
<pre> docker run -it -e NGROK_AUTHTOKEN=YOUR_AUTH_TOKEN ngrok/ngrok:latest http host.docker.internal:9090 </pre>
If you want public grpc url then change port 9090 to 8080   


## How to run test file
Go to root directory of the project and run given command it will run all your test files.
<pre> go test -v --cover ./...  </pre>