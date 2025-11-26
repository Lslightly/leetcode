project(test)

add_executable(C C.cc)
target_compile_options(C PRIVATE -g -O0)
add_executable(C_test C.cc)
target_compile_options(C_test PRIVATE -DTEST -g -O0)

add_executable(F F.cc)
target_compile_options(F PRIVATE -g -O0)
add_executable(F_test F.cc)
target_compile_options(F_test PRIVATE -DTEST -g -O0)
