## wisdomenigma 

    [Docker-Engine]
        Docker run app 

        [pull image]
            $ docker pull wizdwarfs/explorer:latest

        [run]
            $ docker run --network=host -it wizdwarfs/explorer:latest

        [listener address]
            https://localhost:3000    

    [Docker-composer]
        
        [IPFS-Link]
            https://ipfs.io/ipfs/QmcvB8skVZX9fbtUTzvBm8eesDo5SaYRKNFhxvUMvk8fCR
        [run]
            $ docker-compose up
        
        [listener address]
            [traefik]
                localhost:8080/dashboard
        
        [Actives service] (traefik, explorer, consul, influxdb)

        Traefik Instructions:
            -> Click "Service" (Explor-button)
            -> Find "explorer-expert-enigma@docker" and click the service
            -> Find Server URL and paste on your browser tab
        
        [Remove] (as root or privilege )
            -> $ docker volume prune  // this will clear your cache 



    [Routes]        [Description]              [Method]            [searchable] 
    /                home page                  get                     ok
    /projects        project                    get                     ok
    /download        download software          get                     ok
    /sdk             any software dev kit       get                     ok
                       , modules
    /api             api gateways               get                     ok
    /sign            login account              get, post               ok