# blueocean_go
Blueocean restapi golang package

## Usage

see test.go in test directory

```go
package main

import (
	"fmt"

	goblueoc "github.com/xavi06/blueocean_go"
)

func main() {
	api := goblueoc.NewAPI("http://172.16.151.114:8080", "root", "jks2018")
	//res, err := api.ListPipelines()
	//res, err := api.GetPipeline("tomcat-fccs1-maven")
	//res, err := api.ListPipelineRuns("tomcat-fccs2-maven")
	//res, err := api.ListRunNodes("tomcat-fccs2-maven", "178")
	//res, err := api.ListNodeSteps("tomcat-fccs2-maven", "178", "18")
	//res, err := api.GetStepLog("tomcat-fccs2-maven", "178", "18", "26")
	res, err := api.GetRunLog("tomcat-fccs2-maven", "178")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
```
