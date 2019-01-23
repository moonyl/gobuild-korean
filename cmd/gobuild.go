// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// 一个简单的 Go 语言热编译工具。
//
// 监视指定目录(可同时监视多个目录)下文件的变化，触发`go build`指令，
// 实时编译指定的 Go 代码，并在编译成功时运行该程序。
// 具体命令格式可使用`gobuild -h`来查看。
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/caixw/gobuild"
)

const mainVersion = "0.7.1"

// 与版号相关的变量
var (
	buildDate  string // 由链接器提供此值。
	commitHash string // 由链接器提供此值。
	version    string
)

func init() {
	version = mainVersion
	if len(buildDate) > 0 {
		version += ("+" + buildDate)
	}
}

func main() {
	var showHelp, showVersion, recursive, showIgnore bool
	var mainFiles, outputName, extString, appArgs string

	flag.BoolVar(&showHelp, "h", false, "显示帮助信息；")
	flag.BoolVar(&showVersion, "v", false, "显示版本号；")
	flag.BoolVar(&recursive, "r", true, "是否查找子目录；")
	flag.BoolVar(&showIgnore, "i", false, "是否显示被标记为 IGNORE 的日志内容；")
	flag.StringVar(&outputName, "o", "", "指定输出名称，程序的工作目录随之改变；")
	flag.StringVar(&appArgs, "x", "", "传递给编译程序的参数；")
	flag.StringVar(&extString, "ext", "go", "指定监视的文件扩展，区分大小写。* 表示监视所有类型文件，空值代表不监视任何文件；")
	flag.StringVar(&mainFiles, "main", "", "指定需要编译的文件；")
	flag.Usage = usage
	flag.Parse()

	switch {
	case showHelp:
		flag.Usage()
		return
	case showVersion:
		fmt.Fprintln(os.Stdout, "gobuild", version, "build with", runtime.Version(), runtime.GOOS+"/"+runtime.GOARCH)

		if len(commitHash) > 0 {
			fmt.Fprintln(os.Stdout, "commitHash:", commitHash)

		}
		return
	}

	logs := gobuild.NewConsoleLogs(showIgnore)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := append([]string{wd}, flag.Args()...)

	err = gobuild.Build(logs.Logs, mainFiles, outputName, extString, recursive, appArgs, dirs...)
	if err != nil {
		panic(err)
	}
	logs.Stop()
}

func usage() {
	fmt.Fprintln(os.Stdout, `Gobuild는 파일 변경을 모니터링하여 프로그램을 컴파일하고 실행하는 Go의 핫 컴파일 도구입니다.

사용 방법：
 gobuild [options] [dependents]

 options:`)

	flag.CommandLine.SetOutput(os.Stdout)
	flag.PrintDefaults()

	fmt.Fprintln(os.Stdout, `
 dependents:
  다른 종속성의 디렉토리를 지정하는 것은 명령의 끝에서만 나타날 수 있습니다.


사용 예:

 gobuild
   현재 디렉터리를 모니터링하고 변경 사항이 있는 경우 현재 디렉터리의 *.go 파일을 다시 컴파일합니다.；

 gobuild -main=main.go
   현재 디렉토리를 모니터링하고 변경 사항이 있는 경우 현재 디렉토리의 main.go 파일을 다시 컴파일합니다.；

 gobuild -main="main.go" dir1 dir2
   현재 디렉토리를 모니터링하고 dir1과 dir2가 변경된 경우 main.go 파일을 현재 디렉토리에 다시 컴파일합니다；


NOTE: 숨긴 파일 및 숨긴 디렉터리의 파일은 모니터링하지 않습니다.

소스 코드는 MIT 오픈 소스 라이선스에 따라 사용이 허가 https://github.com/caixw/gobuild`)
}
