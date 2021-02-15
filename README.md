# MP0
MP0 for Distributed System Spring 2021

## How To Run
In two separate terminals
```bash
go run processA.go
```
```bash
go run processB.go
```

## Package Design
### Main
For running processA and processB.

### Email
Contains Email struct and helper functions such as composing and printing emails.

### TCP
Seperates TCP logic from the main process.
Supports building servers, connecting clients, and sending/encoding/decoding messages in between. 
