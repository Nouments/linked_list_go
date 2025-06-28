package linkedlist


func main(){
	a:=createLinkedList("A")
	a.insertRight("B")
	a.insertRight("C")
	a=a.insertLeft("D")
	a.showLinkedList()
}

