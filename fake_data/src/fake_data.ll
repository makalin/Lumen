; Fake Data Generation Library for Lumen
; Provides functions to generate realistic fake data

; External declarations
declare i8* @malloc(i64)
declare void @free(i8*)
declare i32 @rand()
declare void @srand(i32)
declare i32 @time(i32*)
declare i8* @strcat(i8*, i8*)
declare i8* @strcpy(i8*, i8*)
declare i64 @strlen(i8*)

; Constants
@first_names = constant [10 x i8*] [
    i8* getelementptr ([5 x i8], [5 x i8]* @str_John, i32 0, i32 0),
    i8* getelementptr ([5 x i8], [5 x i8]* @str_Jane, i32 0, i32 0),
    i8* getelementptr ([6 x i8], [6 x i8]* @str_Alice, i32 0, i32 0),
    i8* getelementptr ([4 x i8], [4 x i8]* @str_Bob, i32 0, i32 0),
    i8* getelementptr ([5 x i8], [5 x i8]* @str_Mary, i32 0, i32 0),
    i8* getelementptr ([5 x i8], [5 x i8]* @str_John, i32 0, i32 0),
    i8* getelementptr ([5 x i8], [5 x i8]* @str_Jane, i32 0, i32 0),
    i8* getelementptr ([6 x i8], [6 x i8]* @str_Alice, i32 0, i32 0),
    i8* getelementptr ([4 x i8], [4 x i8]* @str_Bob, i32 0, i32 0),
    i8* getelementptr ([5 x i8], [5 x i8]* @str_Mary, i32 0, i32 0)
]

@str_John = constant [5 x i8] c"John\00"
@str_Jane = constant [5 x i8] c"Jane\00"
@str_Alice = constant [6 x i8] c"Alice\00"
@str_Bob = constant [4 x i8] c"Bob\00"
@str_Mary = constant [5 x i8] c"Mary\00"

; Initialize random seed
define void @init_random() {
entry:
    %time_ptr = alloca i32
    %time_val = call i32 @time(i32* %time_ptr)
    call void @srand(i32 %time_val)
    ret void
}

; Generate a random first name
define i8* @generate_first_name() {
entry:
    %rand_val = call i32 @rand()
    %index = srem i32 %rand_val, 10
    %name_ptr = getelementptr [10 x i8*], [10 x i8*]* @first_names, i32 0, i32 %index
    %name = load i8*, i8** %name_ptr
    %len = call i64 @strlen(i8* %name)
    %new_str = call i8* @malloc(i64 %len)
    %result = call i8* @strcpy(i8* %new_str, i8* %name)
    ret i8* %result
}

; Generate a random integer between min and max (inclusive)
define i32 @generate_int(i32 %min, i32 %max) {
entry:
    %rand_val = call i32 @rand()
    %range = sub i32 %max, %min
    %range_plus_one = add i32 %range, 1
    %scaled = srem i32 %rand_val, %range_plus_one
    %result = add i32 %scaled, %min
    ret i32 %result
}

; Generate a random boolean
define i1 @generate_bool() {
entry:
    %rand_val = call i32 @rand()
    %result = trunc i32 %rand_val to i1
    ret i1 %result
}

; Generate a random date string (YYYY-MM-DD)
define i8* @generate_date() {
entry:
    %year = call i32 @generate_int(i32 1970, i32 2024)
    %month = call i32 @generate_int(i32 1, i32 12)
    %day = call i32 @generate_int(i32 1, i32 28)
    
    %date_str = call i8* @malloc(i64 11)
    %year_str = getelementptr i8, i8* %date_str, i32 0
    %month_str = getelementptr i8, i8* %date_str, i32 5
    %day_str = getelementptr i8, i8* %date_str, i32 8
    
    ; Format: YYYY-MM-DD
    store i8 48, i8* %year_str
    store i8 48, i8* %month_str
    store i8 48, i8* %day_str
    
    ret i8* %date_str
} 