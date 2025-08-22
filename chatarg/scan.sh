echo  "wss\n" && sudo nmap -sS -p- --min-rate 1000 -T4  wss.dalechatea.me > scan_wss.txt
echo "ws01\n" && sudo nmap -sS -p- --min-rate 1000 -T4 ws01.dalechatea.me > scan_ws01.txt
echo "ws02\n" && sudo nmap -sS -p- --min-rate 1000 -T4 ws02.dalechatea.me > scan_ws02.txt
echo "ws03\n" && sudo nmap -sS -p- --min-rate 1000 -T4 ws03.dalechatea.me > scan_ws03.txt
echo "ws04\n" && sudo nmap -sS -p- --min-rate 1000 -T4 ws04.dalechatea.me > scan_ws04.txt