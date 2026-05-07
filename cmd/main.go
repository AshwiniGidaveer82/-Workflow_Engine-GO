package main

import (
	"fmt"
	"net/http"

	"workflow-engine/configs"
	"workflow-engine/internal/api"
	"workflow-engine/internal/scheduler"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Workflow Automation Engine Running 🚀")
}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/workflow", api.CreateWorkflowHandler)

	scheduler.StartScheduler()

	fmt.Println("Server running on port", configs.ServerPort)

	err := http.ListenAndServe(configs.ServerPort, nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
