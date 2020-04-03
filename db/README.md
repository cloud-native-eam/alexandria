to get the password out of the secret to login to the DB
`k get secret alexandria-db-root-password -o jsonpath="{.data.password}" | base64 --decode`
