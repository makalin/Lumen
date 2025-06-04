; Ray tracer example in Lumen/LLVM IR
; Demonstrates structs, arrays, and basic math operations

%struct.Vector3 = type { double, double, double }
%struct.Sphere = type { %struct.Vector3, double }
%struct.Ray = type { %struct.Vector3, %struct.Vector3 }

; Vector3 operations
declare double @add(double, double)
declare double @mul(double, double)
declare double @div(double, double)
declare double @sqrt(double)
declare void @println(i8*)

; Constants
@str_result = constant [15 x i8] c"Ray trace: %f\0A\00"
@str_done = constant [12 x i8] c"Done!\0A\00"

; Vector3 dot product
define double @dot(%struct.Vector3* %a, %struct.Vector3* %b) {
entry:
  %a_x = getelementptr %struct.Vector3, %struct.Vector3* %a, i32 0, i32 0
  %a_y = getelementptr %struct.Vector3, %struct.Vector3* %a, i32 0, i32 1
  %a_z = getelementptr %struct.Vector3, %struct.Vector3* %a, i32 0, i32 2
  %b_x = getelementptr %struct.Vector3, %struct.Vector3* %b, i32 0, i32 0
  %b_y = getelementptr %struct.Vector3, %struct.Vector3* %b, i32 0, i32 1
  %b_z = getelementptr %struct.Vector3, %struct.Vector3* %b, i32 0, i32 2
  
  %x1 = load double, double* %a_x
  %x2 = load double, double* %b_x
  %y1 = load double, double* %a_y
  %y2 = load double, double* %b_y
  %z1 = load double, double* %a_z
  %z2 = load double, double* %b_z
  
  %xx = call double @mul(double %x1, double %x2)
  %yy = call double @mul(double %y1, double %y2)
  %zz = call double @mul(double %z1, double %z2)
  
  %xy = call double @add(double %xx, double %yy)
  %result = call double @add(double %xy, double %zz)
  
  ret double %result
}

; Ray-sphere intersection
define double @intersect(%struct.Ray* %ray, %struct.Sphere* %sphere) {
entry:
  %ray_origin = getelementptr %struct.Ray, %struct.Ray* %ray, i32 0, i32 0
  %ray_dir = getelementptr %struct.Ray, %struct.Ray* %ray, i32 0, i32 1
  %sphere_center = getelementptr %struct.Sphere, %struct.Sphere* %sphere, i32 0, i32 0
  %sphere_radius = getelementptr %struct.Sphere, %struct.Sphere* %sphere, i32 0, i32 1
  
  ; Calculate discriminant
  %oc = alloca %struct.Vector3
  %oc_x = getelementptr %struct.Vector3, %struct.Vector3* %oc, i32 0, i32 0
  %oc_y = getelementptr %struct.Vector3, %struct.Vector3* %oc, i32 0, i32 1
  %oc_z = getelementptr %struct.Vector3, %struct.Vector3* %oc, i32 0, i32 2
  
  %a = call double @dot(%struct.Vector3* %ray_dir, %struct.Vector3* %ray_dir)
  %b = call double @mul(double 2.0, double %a)
  %c = call double @dot(%struct.Vector3* %oc, %struct.Vector3* %oc)
  %c1 = call double @mul(double %c, double -1.0)
  %radius_squared = call double @mul(double %sphere_radius, double %sphere_radius)
  %c2 = call double @add(double %c1, double %radius_squared)
  
  %discriminant = call double @mul(double %b, double %b)
  %discriminant1 = call double @mul(double 4.0, double %a)
  %discriminant2 = call double @mul(double %discriminant1, double %c2)
  %discriminant3 = call double @add(double %discriminant, double %discriminant2)
  
  %cmp = fcmp olt double %discriminant3, 0.0
  br i1 %cmp, label %no_intersection, label %has_intersection

no_intersection:
  ret double -1.0

has_intersection:
  %t1 = call double @mul(double %b, double -1.0)
  %t2 = call double @sqrt(double %discriminant3)
  %t3 = call double @add(double %t1, double %t2)
  %t4 = call double @mul(double 2.0, double %a)
  %t = call double @div(double %t3, double %t4)
  ret double %t
}

define i32 @main() {
entry:
  ; Create a ray
  %ray = alloca %struct.Ray
  %ray_origin = getelementptr %struct.Ray, %struct.Ray* %ray, i32 0, i32 0
  %ray_dir = getelementptr %struct.Ray, %struct.Ray* %ray, i32 0, i32 1
  
  ; Create a sphere
  %sphere = alloca %struct.Sphere
  %sphere_center = getelementptr %struct.Sphere, %struct.Sphere* %sphere, i32 0, i32 0
  %sphere_radius = getelementptr %struct.Sphere, %struct.Sphere* %sphere, i32 0, i32 1
  store double 5.0, double* %sphere_radius
  
  ; Calculate intersection
  %t = call double @intersect(%struct.Ray* %ray, %struct.Sphere* %sphere)
  
  ; Print result
  %msg_ptr = getelementptr [15 x i8], [15 x i8]* @str_result, i32 0, i32 0
  call void @println(i8* %msg_ptr)
  
  %done_ptr = getelementptr [12 x i8], [12 x i8]* @str_done, i32 0, i32 0
  call void @println(i8* %done_ptr)
  
  ret i32 0
} 