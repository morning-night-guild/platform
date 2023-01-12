# backend

# directory

Top Package
```
.
├── cmd         // entrypoint main.go and .ko.yaml
├── e2e         // end to end test
├── internal    // app code (need unit test code)
└── pkg         // common code (no test required)
```
Like Clean Architecture
```
internal
├── driver      // external information
├── adapter     // framework (controller gateway)
├── usecase     // processing flow (use domain model)
└── model       // domain (business logic)
```

# .vscode

settings.json

```json
{
    "go.toolsEnvVars": {
        "GOFLAGS": "-tags=e2e"
    }
}
```
