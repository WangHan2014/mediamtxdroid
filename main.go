// replace the offical main function don't need it as library
package mediamtxdroid

import (
	"fmt"
	"os"

	"github.com/bluenviron/mediamtx/internal/core"
)

// running core
var s *core.Core
var ok = false

// start the server with yaml conf,and work directory path (can be null)
// Such as /log , and the log will no longer print on console but /$workDir/mediamtx.log
func StartServer(confPath string, workDir string) {
	// 备份原始输出
	originWriter := os.Stdout
	
	// 提示启动
	fmt.Println("MEDIAMTX DROID Started!! Android AAR by wanghan2")
	
	if workDir != "" {
		// 重定向输出到log
		fmt.Println("MEDIAMTX DROID log to: ",workDir+"/mediamtx.log")
		
		file,_:=os.Create(workDir+"/mediamtx.log")
		// fmt.Println("MEDIAMTX DROID log open: ",file," err: ",err)
		os.Stdout=file
	}

	s, ok = core.New([]string{confPath})
	if !ok {
		// os.Exit(1)
		return
	}
	s.Wait()
	os.Stdout = originWriter
}

// stop the server immediate
func StopServer() {
	os.Exit(0)
}

// CheckIsRunning
func IsRunning() bool {
	return ok
}