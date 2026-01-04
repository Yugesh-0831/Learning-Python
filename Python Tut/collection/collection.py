# collections: Counter, namedTuple, OrderedDict, defaultdict, deque
from collections import deque

d = deque()
d.append(1)
d.append(2)
d.appendleft(3)

print(d.popleft())

# d.clear()
d.extendleft([1, 2, 3])
print(d)

# one place to the right
d.rotate(1)
print(len(d))
print(d[len(d) - 1])
