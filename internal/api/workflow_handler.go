package api

import (
	"encoding/json"
	"net/http"

	"workflow-engine/internal/engine"
	"workflow-engine/internal/models"
)

func CreateWorkflowHandler(w http.ResponseWriter, r *http.Request) {

	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request body
	var workflow models.Workflow

	err := json.NewDecoder(r.Body).Decode(&workflow)

	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Set default workflow status
	workflow.Status = "pending"

	// Start workflow engine
	engine.StartWorkflow(workflow.Name)

	// Response header
	w.Header().Set("Content-Type", "application/json")

	// Send response
	response := map[string]interface{}{
		"message":  "Workflow created successfully",
		"workflow": workflow,
	}

	json.NewEncoder(w).Encode(response)
}
