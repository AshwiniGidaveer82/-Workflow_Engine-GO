# Workflow Automation Engine

A lightweight Go workflow automation service that accepts workflow definitions over HTTP and kicks off a (currently stubbed) workflow execution pipeline.

## Project Structure

- `cmd/main.go` - HTTP server entrypoint.
- `configs/config.go` - Application configuration (e.g., server port).
- `internal/api/workflow_handler.go` - HTTP handler to create/start workflows.
- `internal/models/workflow.go` - Data models (`Workflow`, `Step`).
- `internal/engine/workflow_engine.go` - Workflow engine entrypoint (stub).
- `internal/scheduler/scheduler.go` - Scheduler starter (stub).
- `internal/queue/task_queue.go` - Task queue operations (stub).
- `internal/worker/worker.go` - Worker starter (stub).
- `internal/executor/task_executor.go` - Task execution entrypoint (stub).
- `internal/ai/ai_service.go` - AI analysis starter (stub).

## Requirements

- Go `1.26+`

## Build

```bash
go build ./...
```

## Run

```bash
go run ./cmd
```

The server listens on the port configured in `configs/config.go` (default: `:8080`).

## API

### Home

`GET /`

Returns a simple status message.

### Create Workflow

`POST /workflow`

Starts a workflow using the workflow name provided in the request body.

#### Request Body

```json
{
  "id": 1,
  "name": "example-workflow",
  "status": "pending",
  "steps": [
    {
      "id": 1,
      "name": "step-1",
      "type": "task",
      "status": "pending"
    }
  ]
}
```

> Note: The handler currently overwrites `status` to `"pending"`.

#### Response

```json
{
  "message": "Workflow created successfully",
  "workflow": {
    "id": 1,
    "name": "example-workflow",
    "status": "pending",
    "steps": [
      {
        "id": 1,
        "name": "step-1",
        "type": "task",
        "status": "pending"
      }
    ]
  }
}
```

## Notes

Several components are currently stubs (they print to stdout) and are intended to be extended into a full workflow execution system.

- `internal/engine.StartWorkflow(name)` - currently logs the workflow name.
- `internal/scheduler.StartScheduler()` - currently logs scheduler startup.
- Worker/queue/executor/AI are present but not wired to a real execution flow yet.

## License

Add your license information here.
