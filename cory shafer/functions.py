def hello_function(greeting, name='You'):
    print(f"{greeting}, {name}!")

# hello_function('Hi', name = 'Yugesh')

def student_info(*args, **kwargs):
    """ Tells what it is supposed to do """
    # tuple unpacking
    print(args)
    # dictionary unpacking
    print(kwargs)

courses = ('Math', 'Science')
info = {'name': 'Yugesh', 'age': 20}

# student_info('Math', 'Science', 'English', name='Yugesh', age=20)
student_info(*courses, **info)
