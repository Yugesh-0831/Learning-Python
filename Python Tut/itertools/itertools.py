# itertools -> product, permutations, combinations, accumulate, groupby, and infinite iterators
import operator
from itertools import (
    accumulate,
    combinations,
    combinations_with_replacement,
    groupby,
    permutations,
    product,
    repeat,
)

a = [1, 2, 3, 4]
b = [3, 4]

prod = product(a, b, repeat=2)

perm = permutations(a, 2)

comb = combinations(a, 2)

comb_with_replacement = combinations_with_replacement(a, 2)

# prefix iterations
acc = accumulate(a)

acc = accumulate(a, func=operator.mul)

acc = accumulate(a, func=max)
# print(list(acc))


# grouping elements for a logic (key is return value, and values are list of values)
def smaller_than_3(x):
    return x < 3


grp = groupby(a, key=lambda x: x < 3)
# print(grp)
# for key, value in grp:
#     print(key, list(value))

for i in repeat(10, 3):
    print(i)
