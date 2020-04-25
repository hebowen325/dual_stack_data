ulimit -n 10240
rm /home/sgl/bwhe/Router/data/dns_ipv4/*
rm /home/sgl/bwhe/Router/data/dns_ipv6/*
rm /home/sgl/bwhe/Router/data/data_ipv4/*
rm /home/sgl/bwhe/Router/data/data_ipv6/*
python3 /home/sgl/bwhe/Router/data/data_achieve.py
chmod 755 /home/sgl/bwhe/Router/data/data_achieve.sh
/home/sgl/bwhe/Router/data/data_achieve.sh
:<<!
for f in /home/sgl/bwhe/Router/data/ipv4probe/*;do gzip -d  $f ;done
for f in /home/sgl/bwhe/Router/data/ipv4probe/*;do sc_warts2json $f |jq ".hops"|grep "addr" >>/home/sgl/bwhe/Router/data/iptmp1.txt ;done
LC_COLLATE=C sort /home/sgl/bwhe/Router/data/iptmp1.txt | uniq > /home/sgl/bwhe/Router/data/iptmp2.txt 
python3 /home/sgl/bwhe/Router/data/extract_ipv4_addr.py
rm /home/sgl/bwhe/Router/data/iptmp*
!
for f in /home/sgl/bwhe/Router/data/dns_ipv4/*;do gzip -d  $f ;done
python3 /home/sgl/bwhe/Router/data/extract_dns_for_ipv4.py
LC_COLLATE=C sort -k 1,1 -t "," /home/sgl/bwhe/Router/data/dnstmp1.txt| uniq >/home/sgl/bwhe/Router/data/data_ipv4/domain_name.txt
rm /home/sgl/bwhe/Router/data/dnstmp1.txt
split -l 30000000 /home/sgl/bwhe/Router/data/data_ipv4/domain_name.txt /home/sgl/bwhe/Router/data/data_ipv4/fileChunk
for f in /home/sgl/bwhe/Router/data/dns_ipv6/*;do gzip -d  $f ;done
python3 /home/sgl/bwhe/Router/data/extract_dns_for_ipv6.py
LC_COLLATE=C sort -k 1,1 -t "," /home/sgl/bwhe/Router/data/dnstmp1.txt | uniq >/home/sgl/bwhe/Router/data/data_ipv6/domain_name.txt
rm /home/sgl/bwhe/Router/data/dnstmp1.txt
go run /home/sgl/bwhe/Router/data/dnsgrep.go >/home/sgl/bwhe/Router/data/Final_result1.txt
go run /home/sgl/bwhe/Router/data/dnsgrep2.go >>/home/sgl/bwhe/Router/data/Final_result1.txt
C_ALL=C sort -k 2 -t ';' -u /home/sgl/bwhe/Router/data/Final_result1.txt |uniq >/home/sgl/bwhe/Router/data/Final_result.txt
rm /home/sgl/bwhe/Router/data/Final_result1.txt
