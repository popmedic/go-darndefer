# go-darndefer
The gosh darn defer makes my functions slower!

In the process of testing [Channels VS Mutexes](https://github.com/popmedic/go-chanVmutex) in Go, I stumbled upon something that kinda made me stop for a second.  Ofen in Go, I will use a Mutex like so:

``` Go
func (v *Struct) criticalSection(m *sync.Mutex) {
    m.Lock()
    defer m.UnLock()
    // Critical Section Code
    // ...
}
```

but coming from languages without the handy-dandy defer (or what I now call the GOD Darn Defer) I started my examples without the defer line more like:

``` Go
func (v *Struct) criticalSection(m *sync.Mutex) {
    m.Lock()
    // Critical Section Code
    // ...
    m.UnLock()
}
```

because I was NOT going to have a return where I did not unlock. I of course benchmarked the results before I cleaned up by making the unlocks defer functions and I noticed how much longer it took with the defer functions!

This little benchmark test in this package will show what I mean.  Here is what I get with my _MacBook Pro, 2.8 GHz Intel Core i7, 16 GB 1600 MHz DDR3_:

| Function | ns/op |
| --- | --- |
| withDefer | 59.6 ns/op |
| without | 25.4 ns/op |
| withSyncFunc | 26.1 ns/op |

One can tell quickly that defer is not that expensive, ~30 ns/op, but it is overhead, and maybe should be avoided if not needed.  

## What is a withSyncFunc?

Probably asking "what is a withSyncFunc."  Not sure what others call it, but I call them sync functions.  I make these functions so I can have the assurance of defer, but with much less overhead (as seen above.)  The function is quite simple:

``` Go
func syncFunc(l sync.Locker, block func()) {
    l.Lock()
    block()
    l.Unlock()
}
```

It can then be used sort of like defer to insure that the `UnLock()` is always called:

``` Go
syncFunc(myMutex, func(){
    // Critical Section
    // ...
})
```

### Works with close as well

If you like the sync functions like I do, you might also like using this:

``` Go
func closeFunc(c io.Closer, block func()) {
    block()
    c.Close()
}
```

so you can do:

``` Go
closeFunc(f, func(){
    // do stuff with f
    // ...
})
```

## To try the benchmark on your machine

In the terminal:

``` bash
go get github.com/popmedic/go-darndefer/...
cd $GOPATH/src/github.com/popmedic/go-darndefer/darndefer
go test -bench=.
```
