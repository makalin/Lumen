; Declare external functions from lumen_std
declare i32 @add(i32, i32)
declare void @println(i8*)

@str_result = constant [8 x i8] c"Result:\0A\00"

define i32 @main() {
entry:
  %sum = call i32 @add(i32 3, i32 4)
  %msg_ptr = getelementptr [8 x i8], [8 x i8]* @str_result, i32 0, i32 0
  call void @println(i8* %msg_ptr)
  ret i32 0
}
