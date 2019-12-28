cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main list
You have the following tasks to do: 

1. Check USA dollars1 
2. Check USA dollars2 
3. Check USA dollars2 3333333 
4. Check USA dollars3 
5. Check USA dollars4 
6. Check USA dollars5 
7. Check USA dollars6 
8. Check USA dollars7 
9. Check USA dollars8 
10. Check USA dollars9 
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main Remove
 Remove all task... 
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main list
You have no task to list....

cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main add "Check USA dollars9"
Added:  "Check USA dollars9"  to your task list.
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main list
You have the following tasks to do: 

1. Check USA dollars9 
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main add "Check USA dollars9"
Already find task of:  Check USA dollars9
Added:  "Check USA dollars9"  to your task list.
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main list
You have the following tasks to do: 

1. Check USA dollars9 
cyrilk@cyrilloop-GL553VD:~/go/exe6$ go build main.go 
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main add "Check USA dollars9"
Already find task of:  Check USA dollars9
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main add "Check USA dollars9"
Already find task of:  Check USA dollars9
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main Remove
 Remove all task... 
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main add "Check USA dollars9"
Added:  "Check USA dollars9"  to your task list.
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main add "Check USA dollars9"
Already find task of:  Check USA dollars9
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main list
You have the following tasks to do: 

1. Check USA dollars9 
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main do 1
Marked " 1 " as completed.
cyrilk@cyrilloop-GL553VD:~/go/exe6$ ./main list
You have no task to list....

