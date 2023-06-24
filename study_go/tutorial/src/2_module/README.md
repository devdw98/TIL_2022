# 2 Create a module

- 로컬 모듈 사용하기
    ```
    # go mod edit -replace {모듈명}={위치}
    go mod edit -replace tutorial/2_module/greetings=../greetings

    # 적용
    go mod tidy
    ```
