package errs

type Handling struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	ERR_BIND_OBJECT = "Error to Bind Object"
	ERR_SERVICE     = "Service Error Execution"
	ERR_REPO        = "Repository Error Execution"
	ERR_DB_CONN     = "database.Conn Error to retrive db:"
	ERR_EXEC_QUERY  = "A error ocurred to Exec Query:"
	ERR_READ_FILES  = "A error ocorred to Read Files"
)
