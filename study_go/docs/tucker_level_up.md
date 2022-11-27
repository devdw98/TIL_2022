# 18 슬라이스
### 선언
- 배열 개수를 적지 않고 선언함
```go
var arr [10]int // 배열
var sli []int   // 슬라이스
```

### 초기화
- {} 
    ```go
    var slice = []int{1,2,3}    // slice
    var arr = [...]int{1,2,3}   // array (헷갈리지 말기)
    ```
- 내장함수 make()
    ```go
    var slice = make([]int, 3)  // 길이 3인 슬라이스 선언
    ```

### 요소 추가 - append()
- 요소를 추가해 길이를 늘릴 수 있음

### 슬라이싱
- array[시작인덱스:끝인덱스]
-  배열 일부를 가리키는 슬라이스를 반환하는 것. 새로운 배열이 아닌 배열의 일부를 포인터로 가리키는 슬라이스임

### 슬라이스 복제
- 슬라이싱 + append() 
```go
slice1 := []int{1,2,3,4,5}
slice2 := append([]int{}, slice1)
```
- copy()
```go
slice3 := make([]int, len(slice1))
copy(slice3, slice1)
```

# 19 메서드
- 선언
```go
func (r Receiver) info() int {}
```
    - Receiver : 리시버, 해당 구조체 타입에 메서드가 포함됨을 나타냄. 패키지 안에서 type 키워드로 선언된 타입들이 리시버로서 사용할 수 있음
    - info : 메서드 명
    - int : 반환 타입
    - 리시버가 있으면 메서드, 없으면 함수임

# 20 인터페이스
- 선언
```go
type DuckInterface interface {
    Fly()
    Walk(distance int)
}
```
    - DuckInterface : 인터페이스 명
- 유의사항
    - 메서드에는 반드시 메서드 명이 있어야 함
    - 이름이 같은 메서드는 있을 수 없음
    - 메서드 구현을 포함하지 않음
    - ~er 을 붙여서 인터페이스명을 만드는걸 권장함


# 25 채널과 컨텍스트
### 채널
> 고루틴끼리 메시지를 전달할 수 있는 메세지 큐
- 선언
    ```go
    var chanName chan string = make(chan string)
    ```
    - chanName : 채널 인스턴스 변수
    - chan string : 채널 타입
    - make(chan string) : 
