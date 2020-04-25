from dns import resolver,reversename,exception
from IPy import IP
import os,sys
x="/home/sgl/bwhe/DNSGrep"
r=resolver.Resolver()
Nameservers=["114.114.114.114","8.8.8.8","180.76.76.76","223.5.5.5"]
f2 = open(x+"/UnionResults.txt","r")
lines = f2.readlines()
fo = open(x+"/Final_result.txt","w")
for t in range(len(lines)):
    domain=lines[t].split(";")[0]
    IPv4=lines[t].split(";")[1].split(',')
    IPv6=lines[t].split(";")[2][:-1].split(',')
    try:
        ipv4=r.query(domain,"A",lifetime=0.1).response.answer
        for ans in ipv4:
            if(ans.to_text().split()[-1] not in IPv4):
                IPv4.append(ans.to_text().split()[-1])
        ipv6=r.query(domain,"AAAA",lifetime=0.1).response.answer
        for ans in ipv6:
            if(ans.to_text().split()[-1] not in IPv6):
                IPv6.append(ans.to_text().split()[-1])
        ansstr=domain+";"+IPv4[0]
        for i in range(1,len(IPv4)):
            ansstr=ansstr+','+IPv4[i]
        ansstr=ansstr+";"+IPv6[0]
        for j in range(1,len(IPv6)):
            ansstr=ansstr+','+IPv6[j]
    except(resolver.NXDOMAIN,resolver.NoAnswer,resolver.NoNameservers,exception.Timeout,ValueError):
        ansstr=domain+";"+IPv4[0]
        for i in range(1,len(IPv4)):
            ansstr=ansstr+','+IPv4[i]
        ansstr=ansstr+";"+IPv6[0]
        for j in range(1,len(IPv6)):
            ansstr=ansstr+','+IPv6[j]
    fo.write(ansstr+"\n")
fo.close()
f2.close()
