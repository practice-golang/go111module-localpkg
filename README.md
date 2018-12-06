# helloworld

[여기 게시물](https://enjoytools.net/xe/index.php?mid=board_nfRq49&document_srl=10948) 첨부용  
처음 만든 [샘플](https://github.com/practice-golang/helloworld.git)에서 로컬패키지를 모듈로 사용할 때의 공부를 위해 약간 수정

## 공부
* go.mod는 모든 패키지에 들어가야 한다.
  * 굳이 도메인 형식으로 넣을 필요는 없다.
* go.mod에 require 항목을 작성 또는 수정해야 한다.
* go.mod에 replace로 경로 별명을 추가한다.
  * ~~~번거롭(게 느껴진)다.~~~
  * require와 replace는 직접 입력해줘야 한다. 이것도 번거로움.
  ```sh
  # 아래와 같이 콘솔에서 입력해도 돼지만, 그래도 번거롭다.
  go mod edit -replace mypkg=../mypkg
  ```
* import 주석은 GOPATH 밖에서는 반드시 입력해야 한다.
  * http 주소 대신, // import "패키지이름" 이렇게 넣어도 된다.
  * package main에 뜨는 오류 때문에 import 주석을 오해했었는데 그게 아니라 build가 끝나야 에디터(vscode)에서 제대로 인식 되었다.
  * GOPATH 안에서 GO111MODULE=on 상태에서는 import 주석이 없을 경우 GOPATH/src/package-name 에서 package-name을 대신 따오는 것 같다.
  * vscode-go에서는 module을 지원하는 대신, GO111MODULE=on과 inferGopath를 동시에 쓸 수 없게 되었다. T-T
  * 의존성 관리를 프로젝트별로 나눠서 하려면 ```go mod vendor```와 ```go build -mod vendor```를 이용하면 된다.
