to get the password out of the secret to login to the system_db via root
`k get secret alexandria-db-root-password -o jsonpath="{.data.password}" | base64 --decode`


Set password in /init-container/secrets.yaml

echo -n 'SuperSecrerPW' | base64