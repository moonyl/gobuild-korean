gobuild
[![Build Status](https://travis-ci.org/caixw/gobuild.svg?branch=master)](https://travis-ci.org/caixw/gobuild)
[![Build status](https://ci.appveyor.com/api/projects/status/9ux7lbx30k6n5xlr?svg=true)](https://ci.appveyor.com/project/caixw/gobuild)
[![Go version](https://img.shields.io/badge/Go-1.11-brightgreen.svg?style=flat)](https://golang.org) 
======

gobuild 간단한 go 언어 핫 컴파일 도구입니다.
지정한 디렉토리에서 파일 변경(이름 바꾸기, 삭제, 생성, 추가)을 실시간으로 모니터링해서 프로그램을 컴파일하고 실행합니다.


#### 사용 방법:
```shell
gobuild [options] [dependents]

options:
 -h    현재 도움말 정보 표시；
 -v    gobuild와 Go 프로그램의 버전 정보 표시；
 -r    하위 디렉토리를 검색할지 여부. 기본값 true；
 -i    IGNORE로 표시된 로그 내용을 표시할 것인지 여부. 기본값 false. 즉, 표시되지 않음；
 -o    컴파일한 실행 파일 이름.；
 -x    컴파일러에 전달한 매개 변수；
 -ext  모니터링할 확장 프로그램. 기본값은 "go"이며 대/소문자를 구분, 각 확장의 앞뒤 공백을 제거.
       모든 유형의 파일을 모니터링할 경우 *를 사용. 널(NULL) 값을 전달하면 파일을 모니링하지 않음.；
 -main 컴파일할 파일 지정. 기본값 "".

dependents:
 指定其它依赖的目录，只能出现在命令的尾部。
```


#### 사용 예:
```shell
 // 현재 디렉토리의 파일을 모니터링하고 변경 사항이 있는 경우 트리거. go build -main="*.go"
 gobuild

 // ~/Go/src/github.com/issue9/term 디렉토리의 현재 디렉토리와 파일을 모니터링 하고,
 // 변경이 발생하면 트리거. go build -main="main.go"
 gobuild -main=main.go ~/Go/src/github.com/issue9/term
```


#### 지원 플랫폼:

플랫폼 지원 의존성 [colors](https://github.com/issue9/term) 와 [fsnotify](https://gopkg.in/fsnotify.v1) 두 개의 패키지，
현재 지원 플랫폼：Windows, Linux, OSX, BSD。


### 설치

```shell
go get github.com/caixw/gobuild
```


### 라이센스

이 프로젝트는 [MIT](https://opensource.org/licenses/MIT) 오픈 라이센스，자세한 내용은 [LICENSE](LICENSE) 파일에 있습니다.
