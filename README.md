# cli-for-managing-task

It is a CLI to manage your daily tasks.It is just like a todo app except that it is a CLI. 

## Installation
- For linux/unix systems
	- Download the latest binary from [releases](https://github.com/Jatin6256/TaskManager/releases)
	- run the following command to configure to store your tasks

  ``` 
  echo "datafile : <filename with path>" > $HOME/<user>/.TaskManager.yaml
  ```

  - eg of filename with Path  : 
  ```
  $HOME/jatin/.tasks.json
  ```

  - or alternatively set the enviornment variable : 
  ``` 
  TM_DATAFILE=<file name with path>
  ```

- For Other OS Clone this repo and
	Build using:
	```
	$ go build
	```
	- add the compiled binary to path

## Usage

``` 
Task Manager will help you to manage your daily task in simple and easy way

Usage:
  TaskManager [command]

Available Commands:
  add         Add a new task
  completion  generate the autocompletion script for the specified shell
  done        Mark the task done
  edit        edit the task
  help        Help about any command
  list        List your all Task

Flags:
      --config string     config file (default is $HOME/.TaskManager.yaml)
      --dataFile string   data file to store tasks       (default "/home/jatin/.tasks.json")
  -h, --help              help for TaskManager
  -t, --toggle            Help message for toggle

Use "TaskManager [command] --help" for more information about a command.

```




  
  
