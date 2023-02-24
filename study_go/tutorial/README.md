# Tutorial Golang

1. golang 설치
2. GOPATH, PATH 선언
    ```
    vi ./zshrc
    GOPATH={프로젝트 생성할 곳의 pwd}
    PATH=$GOPATH/bin:$PATH
    ```
3. bin/ pkg/ src 폴더 생성

4. 각 튜토리얼 별로 폴더 생성
    ```
    mkdir 1_hello
    ```
5. 모듈 생성하기
    ```
    go mod init tutorial/1_hello
    ```
6. 튜토리얼 진행