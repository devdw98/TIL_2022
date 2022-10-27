# socket type

## protocol
- 통신 규약
- 데이터를 주고 받기 위해 정해놓은 약속

```c
#include <sys/socket.h>

int socket(int domain, int type, int protocol);
```
### domain : 프로토콜 체계 (sys/socket.h)
- PF_INET : IPv4 인터넷 프로토콜 체계

### type: 데이터 전송 방식
- SOCK_STREAM: 신뢰성 있는 순차적인 바이트 기반의 연결지향 데이터 전송 방식 소켓 == TCP socket
    - 버퍼 존재함
    - 손실 시 재전송함
    - 다른 연결지향형 소켓과 1:1 로 연결함
- SOCK_DGRAM: 비연결지향형 소켓 == UDP socket
    - 순서에 상관없이 가장 빠른 전송 지향
    - 손실될 가능성 있음

### protocol: 통신에 사용되는 프로토콜 정보
- domain 에서 지정한 프로토콜 체계 범위 내에서 결정되어야 함