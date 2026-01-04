import time
from threading import Lock, Thread

# A global lock stops all transfers.
# A per-user lock stops only the users involved.
# mutex = Lock()


class User:
    def recieve_money(self, amount):
        self.balance += amount

    def send_money(self, amount):
        with self.mutex:
            if self.balance < amount:
                print("can't transfer money, insufficient balance")
                return False
            time.sleep(0.001)
            self.balance -= amount
            return True

    # instance parameters
    def __init__(self, balance, user_no):
        self.balance = balance
        self.user_no = user_no
        self.mutex = Lock()


users = list()
for i in range(100):
    user = User(1000, i)
    users.append(user)

total = sum(user.balance for user in users)
print(total)


def transfer(sender, reciever, amount):
    if sender.send_money(amount):
        reciever.recieve_money(amount)


# for i in range(50):
#     sender = random.choice(users)
#     receiver = random.choice(users)

#     if sender.user_no == receiver.user_no:
#         continue

#     transfer(sender, receiver, 100)

# total = sum(user.balance for user in users)
# print("Total balance:", total)

threads = []
sender = users[0]
for i in range(50):
    # sender = random.choice(users)
    # receiver = random.choice(users)
    receiver = users[i + 1]

    if sender.user_no == receiver.user_no:
        continue

    t = Thread(target=transfer, args=(sender, receiver, 100))
    threads.append(t)

for t in threads:
    t.start()

for t in threads:
    t.join()

total = sum(user.balance for user in users)
print("Total balance:", total)
print("Sender balance:", users[0].balance)

# handle deadlock
