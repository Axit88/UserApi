
# How To Run Code

## Locally

### prerequisites
Create database using given SQL script located at - mysql_database/database.sql  
Update your database credential at - src/env/local.env

### Run Rest Server 
Open terminal in project root directory and run command   
<pre> go run src/server/restServer/restServer.go </pre>


### Run GRPC Server 
Open terminal in project root directory and run command   
<pre> go run src/server/grpcServer/grpcServer.go </pre>

## Docker Container

### Prerequisites
Change DATABASE_HOST_URL="db:3306" at src/env/local.env    
Add id_rsa file in project root folder which contain github SSH Key.  
### Run command 
<pre>
docker compose build   
docker compose up 
</pre>

## Get Public Endpoint Using Ngrock 

### prerequisites
Create account at ngrock official website https://ngrok.com/   
You will get auth token.  

### Run Command 
<pre>
docker run -it -e NGROK_AUTHTOKEN=YOUR_AUTH_TOKEN ngrok/ngrok:latest http host.docker.internal:9090  
</pre>
If you want public grpc url then change port 9090 to 8080   

