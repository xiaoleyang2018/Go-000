package main

import (
	"errors"
	"fmt"
	xerrors "github.com/pkg/errors"
)

var NoSqlRow = errors.New("no sql row")

func dao()(err error){
	//查数据库sql
	//if err != nil{
	return xerrors.Wrap(NoSqlRow,"查数据库sql报错")
	//}
	//return
}

func biz()(err error){
	err = dao()
	return
}

func main() {
	err := biz()
	fmt.Printf("error : %+v \n",err)

}

/*输出：
error : no sql row
查数据库sql报错
main.dao
/Users/silvia/Code/go/src/Go-000/Week02/error.go:14
main.biz
/Users/silvia/Code/go/src/Go-000/Week02/error.go:20
main.main
/Users/silvia/Code/go/src/Go-000/Week02/error.go:25
runtime.main
/usr/local/Cellar/go/1.14.1/libexec/src/runtime/proc.go:203
runtime.goexit
/usr/local/Cellar/go/1.14.1/libexec/src/runtime/asm_amd64.s:1373 */
