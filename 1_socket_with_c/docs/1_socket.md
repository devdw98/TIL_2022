# SOCKET

### Network Programming
- 네트워크로 연결되어있는 서로 다른 두 컴퓨터가 데이터를 주고받을 수 있도록 하는 것
- 소켓으로 데이터를 주고받을 수 있게 함!
- -> 네트워크 프로그래밍 == 소켓 프로그래밍
- 소켓: 네트워크를 통한 두 컴퓨터의 연결

### server

1. 소켓 생성 : socket
2. 주소 할당 : bind
3. 대기 : listen
4. 요청 수락 : accept
5. 데이터 송수신 : read/write
6. 연결 종료 : close

### client

1. 소켓 생성 : socket
2. 연결 요청 : connect
3. 데이터 송수신 : read/write
4. 연결 종료 : close

### 실습 ch01
1. 파일 빌드
    ```
    # server
    gcc server.c ../lib/err.c -o hserver.exe
    
    # client
    gcc client.c ../lib/err.c -o hclient.exe
    ```
2. 파일 실행
    ```
    # server exec
    ./hserver.exe 9190

    # client exec
    ./hclient.exe 127.0.0.1 9190
    ```
-  서버 실행 후 포트 열려있는 지 확인
    ```
    sudo lsof -i :9190
    ```

3. 실행 결과
- hserver: 클라이언트에게 메세지 전달 후 종료
- hclient: 메세지 전달 받고 연결 종료

### 파일 디스크립터
> 시스템으로부터 할당 받은 파일 또는 소켓에 부여된 정수, 소켓 생성할 때마다 새로운 숫자를 줌

