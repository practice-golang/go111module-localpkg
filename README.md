# helloworld

[여기 게시물](https://enjoytools.net/xe/index.php?mid=board_nfRq49&document_srl=10948) 첨부용

~~처음 만든 [샘플](https://github.com/practice-golang/helloworld.git)에서 로컬패키지를 모듈로 사용할 때의 공부를 위해 약간 수정~~


## 공부

* 빌드, 실행 순서 : `go mod init` > `go mod tidy` > `go build`
* `go mod init` 실행 시 아래 둘 중 하나를 선택
  * `main.go`의 package main 뒤에  ` // import "package-name"` 추가 > `go mod init`
  * `go mod init package-name`
* 로컬에서만 작업하는 거라면, `package-name`에 `도메인 형식`은 취하지 않아도 된다.
* 하위에 내가 만든 패키지는 `// import "package-name"` 주석은 필요 없다.
  * `main.go`에 import 할 때 `메인패키지/하위패키지` 식으로 입력한다.
  * `vscode-go`가 엥간하면 저장 시점에 자동으로 `import`에 추가해준다.
* 의존성 패키지는 `go mod vendor`로  ```go build -mod vendor```를 이용
