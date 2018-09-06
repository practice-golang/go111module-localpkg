# helloworld

처음 만든 [샘플](https://github.com/practice-golang/helloworld.git)에서 로컬패키지를 모듈로 사용할 때의 공부를 위해 약간 수정

## 공부
* go.mod는 모든 패키지에 들어가야 한다.
  * 굳이 도메인 형식으로 넣을 필요는 없다.
* go.mod에 require 항목을 작성 또는 수정해야 한다.
* go.mod에 replace로 경로 별명을 추가한다.
  * 경로 지시가 까다롭(게 느껴진)다. 그냥 처음부터 require와 replace를 혼용하는게 좋을 것 같다.
