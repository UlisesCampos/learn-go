## Working with Unix pipes
Unix pipes are useful when we are passing the output of one program to the input of another. For example:

```sh 
$ echo "test case" | wc -l 
```
In a Go application, the left-hand side of the pipe can be read in using os.Stdin, which acts like
a file descriptor. 

Will take an input on the left-hand side of a pipe and return a list of words and their number of 
occurrences. These words will be tokenized in white space.
```sh 
$ echo "test case case test example" | go run main.go
```

## Catching and handling signals 
Signals are a useful way for the user or the OS to kill your running application. Sometimes,
it makes sense to handle these signals in a more graceful way than the default behavior.
Go Provides a mechanism to catching and handle signals. In this recipe, we'll explore the handling of 
signals through the use of a signal that handles the Go routine

```sh 
cd signals 
go build 
./signals
```
From a separate Terminal, determine the PID and kill the application
```sh 
ps -ef | grep signals
kill -SIGTERM 30555
```
