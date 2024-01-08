#Simple List Split

This tool is a simple splitter that find the number of times a string is found in a list and creates a new list based on
the number of matches given.

Example usage
```shell
list-split -f myList.txt -n 4
```

This will read the file `myList.txt` in the current directory and find any strings that have `4` matches and create a new
list of those strings.

## Caveats

The matches will be removed from the original list and written to a new file

## Outputs
The found matches will be written to `ListOfMatches.txt` in the current directory this is ran    
The Remaining matches will be written to `ListOfRemaining.txt` in the current directory 