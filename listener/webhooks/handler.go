package handler

import (
	"fmt"
	"goshia/listener/configuration"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func runScript() {
	fmt.Println("Running your script with appropriate arguments")
	cmd := exec.Command(configuration.Config.Script.Interpreter, configuration.Config.Script.Location, configuration.Config.Environment)
	cmd.Dir = configuration.Config.Repository.Folder
	stdout, err := cmd.CombinedOutput()
	fmt.Println(string(stdout))
	if err != nil {
		log.Fatal(err)
	}
}

func checkoutRepo(branch *string) {
	fmt.Println("Running checkout")
	cmd := exec.Command("git", "checkout", *branch)
	cmd.Dir = configuration.Config.Repository.Folder
	stdout, err := cmd.CombinedOutput()
	fmt.Println(string(stdout))

	if err != nil {
		stdout, _ = exec.Command("git", "checkout", "master").CombinedOutput()
		fmt.Println(string(stdout))
	}
}

func pullRepo() {
	fmt.Println("Running git pull... just in case")
	cmd := exec.Command("git", "pull", configuration.Config.Repository.Url)
	cmd.Dir = configuration.Config.Repository.Folder
	stdout, _ := cmd.CombinedOutput()
	fmt.Println(string(stdout))
}

func fetchRepo() {
	fmt.Println("Running git fetch to update branch info")
	cmd := exec.Command("git", "fetch")
	cmd.Dir = configuration.Config.Repository.Folder
	stdout, _ := cmd.CombinedOutput()
	fmt.Println(string(stdout))
}

func cloneOrPullRepo(branch *string) {
	fmt.Println("Running git clone")
	stdout, _ := exec.Command("git", "clone", configuration.Config.Repository.Url, configuration.Config.Repository.Folder).CombinedOutput()
	fmt.Println(string(stdout))
	fetchRepo()
	checkoutRepo(branch)
	pullRepo()
}

func deploy(branch string) {
	cloneOrPullRepo(&branch)
	runScript()
}

func deployHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "accepted"})
	fmt.Println("Webhook called! Deploying...")
	go deploy(context.Query("branch"))
}

func statusHandler(context *gin.Context) {
	context.Status(200)
}

func Serve() {
	service := gin.Default()
	service.GET("/status", statusHandler)
	service.GET("/deploy", deployHandler)
	err := service.Run(":" + configuration.Config.Port)
	if err != nil {
		panic("Cannot start Gin Webserver: " + err.Error())
	}
}
