while [[ 1 ]]; do
	read -p "Enter new dns: " dns
	dig +subnet=$dns @ns1.google.com www.google.com
done