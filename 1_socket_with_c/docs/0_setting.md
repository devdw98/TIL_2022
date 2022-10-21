## 0. 개발환경세팅 with mac

1. visualstudiocode 설치 
2. clang 설치 확인
    ```cmd
    clang -- version
    ```
2. extension 설치
    - C/C++
    - Code Runner (for run / option)
    - CodeLLDB (for debug / option)
2. [hello.c](/src/hello.c) 작성
3. 오른쪽 위 debug c/c++ file 실행
4. 실행 전 '디버그 구성 선택' 이 나온다면 'C/C++ : clang 활성 파일 빌드 및 디버그' 누르기 > task.json 파일 생성됨
    - 실행 파일과 디버깅 관련 파일을 한 곳에 모아놓고 싶다면 -args 내 -o 의 경로를 원하는 경로로 변경하기
5. break point 잡아보고 잘 멈추는지 확인!



#### 출처
- [Using Clang in Visual Studio Code](https://code.visualstudio.com/docs/cpp/config-clang-mac#_explore-the-debugger)