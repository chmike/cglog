package cglog

type Cglog struct {
	tag       string
	verbosity int
}

// New returns a Cglog instance that will print the tag in front of the message.
func New(tag string) *Cglog {
	return &Cglog{tag: tag + ": "}
}

// SetV set the verbosity of the Cglog.
func (c *Cglog) SetV(verbosity int) {
	c.verbosity = verbosity
}

// V returns the verbosity.
func (c *Cglog) V() int {
	return c.verbosity
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (c *Cglog) Info(args ...interface{}) {
	logging.print(infoLog, append([]interface{}{c.tag}, args...))
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (c *Cglog) Infoln(args ...interface{}) {
	logging.println(infoLog, append([]interface{}{c.tag}, args...))
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (c *Cglog) Infof(format string, args ...interface{}) {
	logging.printf(infoLog, c.tag+format, args...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (c *Cglog) Warning(args ...interface{}) {
	logging.print(warningLog, append([]interface{}{c.tag}, args...))
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (c *Cglog) Warningln(args ...interface{}) {
	logging.println(warningLog, append([]interface{}{c.tag}, args...))
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (c *Cglog) Warningf(format string, args ...interface{}) {
	logging.printf(warningLog, c.tag+format, args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (c *Cglog) Error(args ...interface{}) {
	logging.print(errorLog, append([]interface{}{c.tag}, args...))
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (c *Cglog) Errorln(args ...interface{}) {
	logging.println(errorLog, append([]interface{}{c.tag}, args...))
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (c *Cglog) Errorf(format string, args ...interface{}) {
	logging.printf(errorLog, c.tag+format, args...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (c *Cglog) Fatal(args ...interface{}) {
	logging.print(fatalLog, append([]interface{}{c.tag}, args...))
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (c *Cglog) Fatalln(args ...interface{}) {
	logging.println(fatalLog, append([]interface{}{c.tag}, args...))
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (c *Cglog) Fatalf(format string, args ...interface{}) {
	logging.printf(fatalLog, c.tag+format, args...)
}
