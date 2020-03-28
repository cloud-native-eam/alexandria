#!/bin/bash


# set new admin passwd
curl -X PATCH --header 'accept: application/json' --data-binary @- --dump - http://localhost:8529/_api/user/admin <<EOF
{ 
  "passwd" : "secure" 
}
EOF

echo 'MWYyZDFlMmU2N2Rm' | base64 --decode
# create db for alexandria
curl -X POST --header 'accept: application/json' --data-binary @- --dump - http://localhost:8529/_api/database <<EOF
{ 
  "name" : "ax-db", 
  "users" : [ 
    { 
      "username" : "ax-admin", 
      "passwd" : "topsecret123", 
      "active" : true 
    }, 
    { 
      "username" : "k8s-pharos", 
      "passwd" : "ax123", 
      "active" : true 
    } , 
    { 
      "username" : "aws-pharos", 
      "passwd" : "ax123", 
      "active" : true 
    } 
  ] 
}
EOF

