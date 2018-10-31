package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// underlyingError 会返回已知的操作系统相关错误的潜在错误值。
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

// Errno 代表某种错误的类型。
type Errno int

func (e Errno) Error() string {
	return "errno " + strconv.Itoa(int(e))
}

func main() {
	var err error
	// 示例1。
	_, err = exec.LookPath(os.DevNull)
	fmt.Printf("error1: %s\n", err)
	if execErr, ok := err.(*exec.Error); ok {
		//重构err
		execErr.Name = os.TempDir()
		execErr.Err = os.ErrNotExist
	}
	fmt.Printf("error2: %s\n", err)
	fmt.Println()

	// 示例2。
	err = os.ErrPermission
	if os.IsPermission(err) {
		fmt.Printf("error(permission): %s\n", err)
	} else {
		fmt.Printf("error(other): %s\n", err)
	}
	os.ErrPermission = os.ErrExist
	if os.IsPermission(err) {
		fmt.Printf("error(permission): %s\n", err)
	} else {
		fmt.Printf("error(other): %s\n", err)
	}
	fmt.Println()

	// 示例3。
	const (
		ERR0 = Errno(0)
		ERR1 = Errno(1)
		ERR2 = Errno(2)
	)
	var myErr error = Errno(0)
	switch myErr {
	case ERR0:
		fmt.Println("ERR0")
	case ERR1:
		fmt.Println("ERR1")
	case ERR2:
		fmt.Println("ERR2")
	}

	//打印结果
	//error1: exec: "NUL": executable file not found in %PATH%
	//error2: exec: "C:\\Users\\yq155\\AppData\\Local\\Temp": file does not exist
	//
	//error(permission): permission denied
	//error(other): permission denied
	//
	//ERR0
}
