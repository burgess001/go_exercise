package main
import (
	"os/exec"
	"context"
	"time"
	"fmt"
)

//执行linux命令并获取标准输出和标准错误
func Cmd(command string) (string,error) {
	cmd := exec.Command("bash","-c",command)
	out, err := cmd.CombinedOutput()
	result := string(out)
	return result,err
}

//执行linux命令获取标准输出和标准错误并设置超时时间
func CmdTime(command string,ts time.Duration)(string,error)  {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*ts)
    cmd := exec.CommandContext(ctx, "bash","-c",command)
	out, err := cmd.CombinedOutput()
	result := string(out)
	return result,err
}

func main () {
	re,err := Cmd("echo hello word")
	fmt.Println(re,err)
	re,err = CmdTime("bash /root/workplace/src/go-exec/i.sh",5)
	fmt.Println(re,err)
}