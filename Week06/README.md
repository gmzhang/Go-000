## 作业题目
参考 Hystrix 实现一个滑动窗口计数器。

## 滑动窗口计数器规则

1. 将时间划分为多个区间；
2. 在每个区间内每有一次请求就将计数器加一维持一个时间窗口,占据多个区间；
3. 每经过一个区间的时间,则抛弃最老的一个区间,并纳入最新的一个区间；
4. 如果当前窗口内区间的请求计数总和超过了限制数量,则本窗口内所有的请求都被丢弃。
5. 时间区间的精度越高,算法所需的空间容量就越大。

### 滑动窗口结构

```go
type slidingWindowCounter struct {
	incurRequests    int32 // 窗口内收到的总的请求数
	durationRequests chan int32 // 一个区间持续收到的请求
	accuracy         time.Duration // 区间精度
	snippet          time.Duration // 窗口间隔
	currentRequests  int32 // 当前收到的总的请求数
	allowRequests    int32 // 窗口允许的最大请求数
}
```

### 窗口滚动

```go
func sliding(l *slidingWindowCounter) {
	for {
		select {
		case <-time.After(l.accuracy):
			t := atomic.SwapInt32(&l.incurRequests, 0)
			l.durationRequests <- t
		}
	}
}

func calculate(l *slidingWindowCounter) {
	for {
		<-time.After(l.accuracy)
		if len(l.durationRequests) == cap(l.durationRequests) {
			break
		}
	}
	for {
		<-time.After(l.accuracy)
		t := <-l.durationRequests
		if t != 0 {
			atomic.AddInt32(&l.currentRequests, -t)
		}
	}
}
```
### 收到请求

```go
func (l *slidingWindowCounter) Take() error {
	once.Do(func() {
		go sliding(l)
		go calculate(l)
	})
	curRequest := atomic.LoadInt32(&l.currentRequests)
	if curRequest >= l.allowRequests {
		return ErrExceededLimit
	}
	if !atomic.CompareAndSwapInt32(&l.currentRequests, curRequest, curRequest+1) {
		return ErrExceededLimit
	}
	atomic.AddInt32(&l.incurRequests, 1)
	return nil

}
```