Requirement: 
1) Docker (with docker-hub account)
2) Nomad
3) Golang


Nomad instruction: 
intstallation https://learn.hashicorp.com/tutorials/nomad/get-started-install?in=nomad/get-started

After install: (Debian/Ubuntu): 
                                 sudo nomad agent -dev -bind 0.0.0.0 -log-level INFO
                                 
                                 
                                 



Open [http://127.0.0.1:4646/](http://127.0.0.1:4646/ui/jobs/run)


copy/paste data from microservice.nomad

run the job


