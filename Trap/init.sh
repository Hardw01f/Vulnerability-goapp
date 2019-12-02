sed -i -e s/localhost/`curl http://ip4.me | grep size=+ | awk -F'>' '{print $4}' | awk -F'<' '{print $1}'`/g ./Trap/DetailCSRF.gtpl
sed -i -e s/localhost/`curl http://ip4.me | grep size=+ | awk -F'>' '{print $4}' | awk -F'<' '{print $1}'`/g ./Trap/PasswdCSRF.gtpl
