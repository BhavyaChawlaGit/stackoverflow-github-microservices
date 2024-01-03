//Bhavya Chawla A20516957 (bchawla@hawk.iit.edu)
package main

import (
	// Import necessary packages
	// ...

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// ...

	"github.com/google/go-github/github"
	// ...

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// Define a struct for StackOverflow posts
type QAPost struct {
	PostID    int                   `json:"post_id"`
	PostTitle string                `json:"title"`
	PostBody  string                `json:"body"`
	Answers   []QAPostAnswer        `json:"answers"`
}

var (
	githubAPIUsage = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "github_api_usage_per_second",
			Help: "Rate of API calls made to GitHub per second",
		},
		[]string{"endpoint"},
	)
	stackOverflowAPIUsage = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "stackoverflow_api_usage_per_second",
			Help: "Rate of API calls made to StackOverflow per second",
		},
		[]string{"endpoint"},
	)
	dataProcessedPerSecond = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "data_processed_per_second",
			Help: "Amount of data processed per second",
		},
		[]string{"source"},
	)
	totalGitHubAPIUsage = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_github_api_usage",
			Help: "Total number of API calls made to GitHub",
		},
	)
	totalStackOverflowAPIUsage = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_stackoverflow_api_usage",
			Help: "Total number of API calls made to StackOverflow",
		},
	)
	githubAPIUsage2Days = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "github_api_usage_2_days_total",
			Help: "Total number of API calls made to GitHub in the past 2 days",
		},
	)
	githubAPIUsage7Days = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "github_api_usage_7_days_total",
			Help: "Total number of API calls made to GitHub in the past 7 days",
		},
	)
	githubAPIUsage45Days = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "github_api_usage_45_days_total",
			Help: "Total number of API calls made to GitHub in the past 45 days",
		},
	)
	// StackOverflow API usage counters for different durations
	stackOverflowAPIUsage2Days = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "stackoverflow_api_usage_2_days_total",
			Help: "Total number of API calls made to StackOverflow in the past 2 days",
		},
	)
	stackOverflowAPIUsage7Days = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "stackoverflow_api_usage_7_days_total",
			Help: "Total number of API calls made to StackOverflow in the past 7 days",
		},
	)
	stackOverflowAPIUsage45Days = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "stackoverflow_api_usage_45_days_total",
			Help: "Total number of API calls made to StackOverflow in the past 45 days",
		},
	)
)

// Initialize Prometheus metrics
func initializeMetrics() {
	prometheus.MustRegister(githubAPIUsage, stackOverflowAPIUsage, dataProcessedPerSecond)
	prometheus.MustRegister(totalGitHubAPIUsage, totalStackOverflowAPIUsage)
	prometheus.MustRegister(
		githubAPIUsage2Days, githubAPIUsage7Days, githubAPIUsage45Days,
		stackOverflowAPIUsage2Days, stackOverflowAPIUsage7Days, stackOverflowAPIUsage45Days,
	)
}

// Define a struct for StackOverflow answers
type QAPostAnswer struct {
	AnswerID   int    `json:"answer_id"`
	AnswerBody string `json:"body"`
}

// Define a struct for GitHub posts
type GitHubPost struct {
	Type    string // "Question" or "Answer"
	Content string // Body of the post
}

// Function to insert StackOverflow posts and their answers into the database
func insertStackOverflowData(db *sql.DB, posts []QAPost, tagName string) error {
	// Create table for this tag
	err := createStackOverflowTable(db, tagName)
	if err != nil {
		return fmt.Errorf("create table: %v", err)
	}

	questionInsertQuery := fmt.Sprintf(`INSERT INTO so_%s_posts (post_id, title, body, link) VALUES ($1, $2, $3, $4) ON CONFLICT (post_id) DO NOTHING;`, tagName)
	answerInsertQuery := fmt.Sprintf(`INSERT INTO so_%s_answers (answer_id, post_id, body) VALUES ($1, $2, $3) ON CONFLICT (answer_id) DO NOTHING;`, tagName)

	for _, post := range posts {
		// Insert the question
		_, err := db.Exec(questionInsertQuery, post.PostID, post.PostTitle, post.PostBody, "https://stackoverflow.com/q/"+strconv.Itoa(post.PostID))
		if err != nil {
			return fmt.Errorf("insert question: %v", err)
		}

		// Insert each answer for the current question
		for _, answer := range post.Answers {
			_, err := db.Exec(answerInsertQuery, answer.AnswerID, post.PostID, answer.AnswerBody)
			if err != nil {
				return fmt.Errorf("insert answer for question %d: %v", post.PostID, err)
			}
		}
	}
	return nil
}

// Function to create a table dynamically for GitHub data
func createGitHubTable(db *sql.DB, repoName string) error {
	createTableQuery := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS github_%s (
			id SERIAL PRIMARY KEY,
			title TEXT,
			body TEXT,
			labels TEXT[]
		);
	`, repoName)
	_, err := db.Exec(createTableQuery)
	return err
}

// Function to create a table dynamically for StackOverflow data
func createStackOverflowTable(db *sql.DB, tagName string) error {
	// Create posts table
	createPostsTableQuery := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS so_%s_posts (
			post_id INTEGER PRIMARY KEY,
			title TEXT,
			body TEXT,
			link TEXT
		);
	`, tagName)
	_, err := db.Exec(createPostsTableQuery)
	if err != nil {
		return err
	}

	// Create answers table
	createAnswersTableQuery := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS so_%s_answers (
			answer_id INTEGER PRIMARY KEY,
			post_id INTEGER,
			body TEXT,
			FOREIGN KEY (post_id) REFERENCES so_%s_posts(post_id)
		);
	
