# StackOverflow_Github_Microservice
This project aims to collect data from StackOverflow and GitHub related to specific frameworks/libraries
and store the information in two separate databases, StackoverflowDB and GitHubDB, using Postgres as
the database server. The chosen frameworks/libraries are Prometheus, Docker, and Go. 

## Technologies Used:  
● Programming Language: Go  
● Containerization: Docker  
● Cloud Platform: Google Cloud Platform (GCP)  
● Database: Postgres  
● Monitoring Tool: Prometheus  


## Steps:
### 1. Created Developer Accounts and API Keys 
**GitHub**  
● Generated a GitHub personal access token  
● Created a personal access token on GitHub  
  
**StackOverflow**  
● Obtained the necessary API key.  
### 2. Running the Microservices
**Step 1:** Clone the Repository  
bash  
Copy code  
git clone https://github.com/your-username/stackoverflow-github-microservices.git  
cd src  
  
**Step 2:** Build and Run Docker Containers  
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

## Metrics

![image](https://github.com/BhavyaChawlaGit/stackoverflow-github-microservices/assets/112718303/1f1ba778-8218-4ad4-959a-5ec8e5747a48)

![image](https://github.com/BhavyaChawlaGit/stackoverflow-github-microservices/assets/112718303/d3d1bccf-1596-4f8d-a2f4-30230c045dcb)



# Specification:
1. Used Go languages, Docker, GCP (Google Cloud Platform), Postgres to create microservices to collect data of the Posts (Question/body and Answer/body) from StackOverflow and Issues from GitHub and stored the data collected in two databases: StackoverflowDB and GitHubDB using Postgres database server.  
  
2. Collected the data from StackOverflow and GitHub for the following frameworks/libraries:  
a. Prometheus  
i. https://stackoverflow.com/search?q=Prometheus  
ii. https://github.com/prometheus/prometheus  
b. Selenium  
c. OpenAl  
d. Docker  
e. Milvus  
f. Go  
i. https://stackoverflow.com/search?q=golang ii. https://github.com/golang/go  
4. You will need to create developer accounts and API KEYS on GitHub and StackExchange  
https://api.github.com  
https://docs.github.com/en  
Here is the URL to guide you on how to generate your GITHUB_TOKEN  
https://help.github.com/articles/creating-an-access-token-for-command-line-use/  
Create your GitHub personal access tokens from  
https://github.com/settings/tokens  
https://api.stackexchange.com/  
https://api.stackexchange.com/docs https://stackapps.com/apps/oauth/register  
5. Use Prometheus to show how many API calls to GitHub and StackOverflow were made per second, and how much data was collected per second.  
6. Run 3 experiments using data collected for the past 2 days, 7 days, and 45 days.  



