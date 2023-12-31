### Learnings

- #### Difference between panic and log.Fatal

    In Go, panic and log.Fatal are both used to halt the execution of a program, but they have some key differences:

    - Output Destination: log.Fatal sends the log message to the configured log output, while panic writes directly to stderr.

    - Stack Trace: panic prints a stack trace, which may not always be relevant to the error. On the other hand, log.Fatal does not print a stack trace.

    - Deferred Functions: When a program panics, deferred functions are executed. However, log.Fatal exits immediately, and deferred functions cannot be run.

    - Usage: panic is typically used for programming errors, where the stack trace is useful for understanding the context of the error. log.Fatal is more suitable for errors that occur under normal operation and do not require a detailed stack trace.

    In Go, the decision to use panic or log.Fatal depends on the nature of the error and how you want your program to handle it.

    Use panic when: You should use panic when your program encounters a situation that should never happen, indicating a serious bug in your program. It's common to use recover to catch the panic and decide whether to continue or abort the program. Panic is typically used for programming errors, where the stack trace is useful for understanding the context of the error.

    Use log.Fatal when: log.Fatal should be used when your program encounters an error that it cannot recover from, such as failing to open a required file or losing network connectivity. Unlike panic, log.Fatal does not provide a stack trace, making it suitable for errors that are expected to occur under normal operation.
