mydict = {"name": "John", "age": 30}
mydict2 = {"name": "Jane", "age": 25, "email": "jane@example.com"}

mydict.update(mydict2)
print(mydict)

mytuple = ("Max",)
mydict[mytuple] = 2
print(mydict)
