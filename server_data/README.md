##This script is used to extract dual-stack domain name with IPv4 and IPv6 addresses from Rapid7 fdns database for Server

* Firstly, run Script_generation_for_server.py, it will visit Rapid7's database and generate corresponding script to extract informaition  
* Secondly, run script.sh, it will
    * download DNS query information automatically( A and AAAA )
    * extract domain name and corresponding IP address, then sort those data according to domain name( IPv4 and IPv6 in two files)
    * use domain names in IPv6 dataset to search for IPv4 addresses, and you can extract dual-stack address
    * the remaining data are those domain names which only have IPv4 address or IPv6 address
    * finally run DNS_Search.py to do A and AAAA query for those dual-stack domain names to expand dual-stack data
* Just run IP_discover.sh, those works can be done automatically
