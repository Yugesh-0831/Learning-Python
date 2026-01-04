from timeit import default_timer as timer

my_list = ["a"] * 1000000

start = timer()
new_str = ""
for i in my_list:
    new_str += i
stop = timer()
print(stop - start)

start = timer()
new_str = "".join(my_list)
stop = timer()
print(stop - start)
