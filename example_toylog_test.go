package toylog_test

import (
  "github.com/aliafshar/toylog"
)

func Example() {
  // We only call this for the example so we can record the output as a test.
  toylog.Stdout()

  toylog.Infoln("I am an info line", "with", "multiple", "elements")
  toylog.Infof("I am a format %s %q", "string", "with the 'normal' features")
  toylog.Debugln("I won't be printed until the log level is changed.")
  toylog.Verbose() // Enables debug logging.
  toylog.Debugln("I will be printed since the log level changed.")

  // Output:
  // I: I am an info line with multiple elements
  // I: I am a format string "with the 'normal' features"
  // D: I will be printed since the log level changed.
}
