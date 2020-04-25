import requests
import re
#url_ip="http://data.caida.org/datasets/topology/ark/ipv4/probe-data/team-1/2019/cycle-20190215/"
#resp=requests.get(url_ip)
#files=re.findall("href=\"(.*warts.gz)\"",resp.text)
f=open("/home/sgl/bwhe/Router/data/data_achieve.sh","w")
#for file in files:
#    f.write("wget -P ipv4probe/ "+url_ip+file+"\n")
for j in range(9,10):
    for i in range(1,10):
        url_domain="http://data.caida.org/datasets/topology/ark/ipv4/dns-names/201"+str(j)+"/0"+str(i)+"/"
        resp=requests.get(url_domain)
        files=re.findall("href=\"(.*201"+str(j)+"0"+str(i)+".*)\"",resp.text)
        for file in files:
            f.write("wget -P /home/sgl/bwhe/Router/data/dns_ipv4/ "+url_domain+file+"\n")
    for i in range(10,13):
        url_domain="http://data.caida.org/datasets/topology/ark/ipv4/dns-names/201"+str(j)+"/"+str(i)+"/"
        resp=requests.get(url_domain)
        files=re.findall("href=\"(.*201"+str(j)+str(i)+".*)\"",resp.text)
        for file in files:
            f.write("wget -P /home/sgl/bwhe/Router/data/dns_ipv4/ "+url_domain+file+"\n")
for j in range(9,10):
    for i in range(1,10):
        url_domain="http://data.caida.org/datasets/topology/ark/ipv6/dns-names/201"+str(j)+"/0"+str(i)+"/"
        resp=requests.get(url_domain)
        files=re.findall("href=\"(.*201"+str(j)+"0"+str(i)+".*)\"",resp.text)
        for file in files:
            f.write("wget -P /home/sgl/bwhe/Router/data/dns_ipv6/ "+url_domain+file+"\n")
    for i in range(10,13):
        url_domain="http://data.caida.org/datasets/topology/ark/ipv6/dns-names/201"+str(j)+"/"+str(i)+"/"
        resp=requests.get(url_domain)
        files=re.findall("href=\"(.*201"+str(j)+str(i)+".*)\"",resp.text)
        for file in files:
            f.write("wget -P /home/sgl/bwhe/Router/data/dns_ipv6/ "+url_domain+file+"\n")
f.close()
