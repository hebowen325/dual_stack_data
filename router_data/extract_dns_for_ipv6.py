import os
from IPy import IP
files=os.listdir("/home/sgl/bwhe/Router/data/dns_ipv6")
writefile=open("/home/sgl/bwhe/Router/data/dnstmp1.txt","w")
for file in files:
    with open("/home/sgl/bwhe/Router/data/dns_ipv6/"+file,"r") as f1:
        lines=f1.readlines()
        for line in lines:
            ip=line.split("\t")
            try:
                if(len(ip[2][:-1])>3 and ip[2][0:4]!="FAIL"):
                    ans=ip[1]+","+ip[2][:-1]
                    writefile.write(ans[::-1]+"\n")
            except IndexError:
                continue
writefile.close()

