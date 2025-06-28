# linked_list_go

this package is an simple implementation of linked list in golang. this can be used as package if you want but it is training  for me  and if there are dome issues with my code , I'll be happy if you correct me 

## ho to use function 

this is the code and how to use this 

```go
package linkedlist

func testing(){
    //creating new linked list
	a:=createLinkedList("A")
    //inserting into the right of linked list 
	a.insertRight("B")
	a.insertRight("C")
    //inserting on the left of our linked list
	a.insertLeft("D")
    //printing our linked list
	a.showLinkedList()
}
```

output

```bash
D->A->B->C
```
