# stack_github_microservice
** This project aims to collect data from StackOverflow and GitHub related to specific frameworks/libraries
and store the information in two separate databases, StackoverflowDB and GitHubDB, using Postgres as
the database server. The chosen frameworks/libraries are Prometheus, Docker, and Go. **  

## Technologies Used:  
● Programming Language: Go  
● Containerization: Docker  
● Cloud Platform: Google Cloud Platform (GCP)  
● Database: Postgres  
● Monitoring Tool: Prometheus  


## Steps:
### 1. Created Developer Accounts and API Keys 
** GitHub **  
● Generated a GitHub personal access token  
● Created a personal access token on GitHub  
** StackOverflow **
● Obtained the necessary API key.  
### 2. Running the Microservices
Step 1: Clone the Repository  
bash  
Copy code  
git clone https://github.com/your-username/stackoverflow-github-microservices.git  
cd src  
Step 2: Build and Run Docker Containers  
docker-compose up --build   
### 3. Data Collection
The microservices will start collecting data from StackOverflow and GitHub for the specified
frameworks/libraries. The data will be stored in the StackoverflowDB and GitHubDB databases on the
Postgres server.  
### 4. Prometheus Monitoring
Prometheus is configured to monitor API calls to GitHub and StackOverflow. Metrics are shown at the
end of the report  
### 5. Run Experiments
Three experiments were conducted using data collected over the past 2 days, 7 days, and 45 days.
