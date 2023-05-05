# Variable collector package    
      simple variable collector respecting those norms : 
     - if the  variable is defined as os env variable it will be collected and it has the 1 priority
     - variable value could be collected from env file by specifying it's path 
  
```go
	//example using variableCollector
	package main

import (
	"fmt"
	"github.com/adamBoualleiguie/variablecollector"
)
func main(){
	variablecollector.NewVariableListConstructor("chef")
	mymap := variablecollector.ExtractVariableValues("/path/to/env/file")
	fmt.Println(mymap["chef"])
	
}

```