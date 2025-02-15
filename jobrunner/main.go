package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"html/template"
	"net/http"
	"net/smtp"
	"time"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello,", g.Name)
}

func quickStart() {
	// jobrunner 需要先Start然后再添加任务
	jobrunner.Start()
	defer jobrunner.Stop()
	err := jobrunner.Schedule("@every 5s", GreetingJob{Name: "lyb"})
	if err != nil {
		return
	}
	time.Sleep(10 * time.Second)
}

type EmailJob struct {
	Name  string
	Email string
}

type User struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (j EmailJob) Run() {
	e := email.NewEmail()
	e.From = "yubolee1217@163.com"
	e.To = []string{j.Email}
	e.Cc = []string{"yubolee1217@163.com"}
	e.Subject = "Welcome to my home"
	e.Text = []byte(fmt.Sprintf(`Hello, %s. Welcome Back!`, j.Name))
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "yubolee1217@163.com", "YXmqCwMHYqUM8HMS", "smtp.163.com"))
	if err != nil {
		fmt.Printf("failed to send email to %s, err:%v", j.Name, err)
	}
}

func login(c *gin.Context) {
	var u User
	if c.ShouldBind(&u) == nil {
		c.String(http.StatusOK, "login success")
		jobrunner.In(5*time.Second, EmailJob{Name: u.Name, Email: u.Email})
	} else {
		c.String(http.StatusNotFound, "login failed")
	}
}

func jobJson(c *gin.Context) {
	c.JSON(http.StatusOK, jobrunner.StatusJson())
}

func jobHtml(c *gin.Context) {
	t, err := template.ParseFiles("Status.html")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = t.Execute(c.Writer, jobrunner.StatusPage())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
func integratedWeb() {
	r := gin.Default()
	//r.GET("/login", login)
	r.GET("jobrunner/json", jobJson)
	r.GET("jobrunner/html", jobHtml)
	jobrunner.Start()
	defer jobrunner.Stop()
	jobrunner.Every(5*time.Second, GreetingJob{Name: "lyb"})
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
func main() {
	// quickStart()
	integratedWeb()
}
