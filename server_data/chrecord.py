#read txt method three
f2 = open("/home/sgl/bwhe/DNSGrep/fdns_a/record.txt","r")
lines = f2.readlines()
print("00000000000000000000000000000000")
for i in range(len(lines)):
    if(i%3==0):
        s1=";"+lines[i][4:-5]
    if(i%3==1):
        print(lines[i][:-1]+s1)

