# cglog

Leveled execution logs for Go.

It is a modified version of [glog](https://github.com/chmike/glog),
that initializes the package with an option structure instead of
command line arguments.

The parameters are the following. 

```Go
type Options struct {
	// ToStdErr is true to log to stderr instead of files.
	ToStdErr bool `json:"toStdErr,omitempty"`
	// AlsoToStdErr is true to log to stderr and to files.
	AlsoToStdErr bool `json:"alsoToStdErr,omitempty"`
	// Verbosity sets the log level for V logs (e.g. 3).
	Verbosity int `json:"verbosity,omitempty"`
	// StdErrThreshold set the stderr output threshold to "info", "warning", "error" or "fatal".
	StdErrThreshold string `json:"stdErrThreshold,omitempty"`
	// VModule sets the verbose level per file. V is comma-separated list of pattern=N settings for file-filtered logging.
	// pattern may be a file name (without .go) or a file with wildcard (e.g. gtx*=2).
	VModule string `json:"vmodule,omitempty"`
	// TraceLocation sets a backtrace logging when logging hits line file:N.
	TraceLocation string
	// LogDir sets the log output directory (default is /tmp).
	LogDir string `json:"logdir,omitempty"`
	// MaxSize is the maximum byte size of a log file triggiring rotation (default: 10MB).
	MaxSize int
}
```

Call the `Init(*Options)` function to initialize cglog before any
logging calls. The parameters may be provided by command line
arguments or a configuration file.

`Init` returns the `ErrAlreadyInitialized` error if called more than once.
It also returns an error if a parameter is invalid. If `Init` returns an
error, the behavior of cglog is undefined. It is preferable to consider it
a fatal error.


----

This is an efficient pure Go implementation of leveled logs in the
manner of the open source C++ package
	https://github.com/google/glog

By binding methods to booleans it is possible to use the log package
without paying the expense of evaluating the arguments to the log.
Through the -vmodule flag, the package also provides fine-grained
control over logging at the file level.

The comment from glog.go introduces the ideas:

	Package glog implements logging analogous to the Google-internal
	C++ INFO/ERROR/V setup.  It provides functions Info, Warning,
	Error, Fatal, plus formatting variants such as Infof. It
	also provides V-style logging controlled by the -v and
	-vmodule=file=2 flags.
	
	Basic examples:
	
		glog.Info("Prepare to repel boarders")
	
		glog.Fatalf("Initialization failed: %s", err)
	
	See the documentation for the V function for an explanation
	of these examples:
	
		if glog.V(2) {
			glog.Info("Starting transaction...")
		}
	
		glog.V(2).Infoln("Processed", nItems, "elements")


The repository contains an open source version of the log package
used inside Google. The master copy of the source lives inside
Google, not here. The code in this repo is for export only and is not itself
under development. Feature requests will be ignored.

Send bug reports to golang-nuts@googlegroups.com.
