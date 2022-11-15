# golang test practice

## 명령어
```bash
cd {test 파일 위치한 곳}

// 전체 테스트
go test

// 일부 테스트
// {테스트명} 키워드 들어간 테스트케이스 골라 실행함
go test -run {테스트명}

// 벤치마크 테스트
// 모든 테스트, 벤치마크 테스트 실행
go test -bench .
```