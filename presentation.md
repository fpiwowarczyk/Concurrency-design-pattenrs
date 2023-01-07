---
author: Filip Piwowarczyk
date: dd MM, YYYY
paging: Slide %d / %d
extensions:
  - render
---
Agenda:

1. Przypomnienie czym jest concurrency (szybko)
2. Przyblizenie czym sa channele i jak duzo to zmienia w porownaniu do zwyklych jezykw (szybko)
3. Generator
4. Fan-out, fan-in 
5. Pipeline
6. Timeout
7. Generator

---
# Main idea
Main idea, you can all of these thing in different languages but in GO it is easy to create because of its design.

---
## Concurrency 


---
## Goroutines 

``` go 
go f() 
go g(1,2)
```

---
## Channels

```
c := make(chan int)
go func() { c <- 3}()
n := <- c
```
---

## Cool example 

~~~xargs cat 
./ping-pong.go
~~~


~~~xargs cat 
./ping-pong.md
~~~


~~~xargs cat 
./ping-pong.md
~~~

~~~xargs cat 
Conccurency-design-patterns/ping-pong.md
~~~


~~~xargs cat 
/ping-pong.md
~~~

---
## Przykład z ostatniego talka z uzyciem channeli 
 Przypomnil share state by communication not communicate by sharing state 
---
## FanOut

Here is description how to make fan out 


---
# Funnelin


```
~~~graph-easy --as=boxart
[ A ] - to -> [ B ]
~~~
```
