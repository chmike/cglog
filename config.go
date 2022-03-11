package cglog

import (
	"errors"
	"strconv"
	"sync/atomic"
)

type Options struct {
	// ToStdErr is true to log to stderr instead of files.
	ToStdErr bool `json:"toStdErr,omitempty"`
	// AlsoToStdErr is true to log to stderr and to files.
	AlsoToStdErr bool `json:"alsoToStdErr,omitempty"`
	// Verbosity sets the log level for V logs (e.g. 3).
	Verbosity int `json:"verbosity,omitempty"`
	// StdErrThreshold set the stderr output threshold to "info", "warning", "error" or "fatal" (default: error).
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

// An Option is a function setting logginT options.
type Option func(*loggingT) error

// initOnce is non-zero when initialized.
var initOnce uint32

// ErrAlreadyInitialized is returned by Init() if glog is already initialized.
var ErrAlreadyInitialized = errors.New("already initialized")

// Init initializes glog with the given options. Don’t call any glog functions before
// calling the glog.Init() function. Returns ErrAlreadyInitialized if called more than
// once.
func Init(options Options) error {
	if !atomic.CompareAndSwapUint32(&initOnce, 0, 1) {
		return ErrAlreadyInitialized
	}
	logging.toStderr = options.ToStdErr
	logging.alsoToStderr = options.AlsoToStdErr
	logDir = options.LogDir
	if err := logging.verbosity.Set(strconv.Itoa(options.Verbosity)); err != nil {
		return err
	}

	if options.StdErrThreshold != "" {
		if err := logging.stderrThreshold.Set(options.StdErrThreshold); err != nil {
			return err
		}
	} else {
		logging.stderrThreshold.Set("error")
	}
	if err := logging.vmodule.Set(options.VModule); err != nil {
		return err
	}
	if options.TraceLocation != "" {
		if err := logging.traceLocation.Set(options.TraceLocation); err != nil {
			return err
		}
	}
	if options.MaxSize != 0 {
		MaxSize = uint64(options.MaxSize)
	}
	logging.setVState(logging.verbosity, nil, false)
	CopyStandardLogTo("INFO")
	go logging.flushDaemon()
	return nil
}
