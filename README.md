# helloworld

[여기 게시물](https://enjoytools.net/xe/index.php?mid=board_nfRq49&document_srl=10948) 첨부용  
처음 만든 [샘플](https://github.com/practice-golang/helloworld.git)에서 로컬패키지를 모듈로 사용할 때의 공부를 위해 약간 수정

## 공부
* go.mod는 모든 패키지에 들어가야 한다.
  * 굳이 도메인 형식으로 넣을 필요는 없다.
* go.mod에 require 항목을 작성 또는 수정해야 한다.
* go.mod에 replace로 경로 별명을 추가한다.
  * ~~~경로 지시가 까다롭(게 느껴진)다. 그냥 처음부터 require와 replace를 혼용하는게 좋을 것 같다.~~~
* import 주석은 GOPATH 밖에서는 반드시 입력해야 한다.
 - package main에 뜨는 오류 때문에 import 주석을 오해했었는데 그게 아니라 build가 끝나야 에디터(vscode)에서 제대로 인식 되었다.
