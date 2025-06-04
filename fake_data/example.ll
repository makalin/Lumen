; Example program demonstrating the fake data library
declare void @println(i8*)
declare void @print_int(i32)
declare void @print_bool(i1)
declare void @init_random()
declare i8* @generate_first_name()
declare i32 @generate_int(i32, i32)
declare i1 @generate_bool()
declare i8* @generate_date()
declare void @free(i8*)

@str_name = constant [13 x i8] c"Random name: \00"
@str_number = constant [15 x i8] c"Random number: \00"
@str_bool = constant [14 x i8] c"Random bool: \00"
@str_date = constant [13 x i8] c"Random date: \00"

define i32 @main() {
entry:
    ; Initialize random number generator
    call void @init_random()
    
    ; Generate and print a random name
    %name_msg = getelementptr [13 x i8], [13 x i8]* @str_name, i32 0, i32 0
    call void @println(i8* %name_msg)
    %name = call i8* @generate_first_name()
    call void @println(i8* %name)
    call void @free(i8* %name)
    
    ; Generate and print a random number
    %number_msg = getelementptr [15 x i8], [15 x i8]* @str_number, i32 0, i32 0
    call void @println(i8* %number_msg)
    %number = call i32 @generate_int(i32 1, i32 100)
    call void @print_int(i32 %number)
    
    ; Generate and print a random boolean
    %bool_msg = getelementptr [14 x i8], [14 x i8]* @str_bool, i32 0, i32 0
    call void @println(i8* %bool_msg)
    %bool_val = call i1 @generate_bool()
    call void @print_bool(i1 %bool_val)
    
    ; Generate and print a random date
    %date_msg = getelementptr [13 x i8], [13 x i8]* @str_date, i32 0, i32 0
    call void @println(i8* %date_msg)
    %date = call i8* @generate_date()
    call void @println(i8* %date)
    call void @free(i8* %date)
    
    ret i32 0
} 