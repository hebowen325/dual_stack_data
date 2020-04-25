##This script is used to extract dual-stack domain name with IPv4 and IPv6 addresses from CAIDA database for Router

* Firstly, run data_achieve.py, it will visit CAIDA's database and generate corresponding script to extract informaition  
* Secondly, run data_acheive.sh, it will download DNS information automatically( A and AAAA )
* Thirdly, run extract_dns_for_ipv4.py and extract_dns_for_ipv6.py to get data, then sort those data according to domain name( IPv4 and IPv6 in two files)
* Then use domain names in IPv6 dataset to search for IPv4 addresses, you can extract dual-stack address
* the remaining data are those domain names which only have IPv4 address or IPv6 address
* Just run CollectIP.sh, those works can be done automatically
