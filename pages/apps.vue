<template>
    <div class="services">
        We will offer various services for our clients. Customer will choose packages according to your requirements.
        <div class="center">
            <vs-button
            border
            shadow 
            light
            :active="true"
            gradient
            class="business"
            title="services"
            @click="active=!active">
            <i class='bx bx-store-alt'></i>
        </vs-button>
        <vs-button 
            border
            shadow
            light 
            danger
            href="/"
            class="btn-back"
            title="Back"
            :active="true"
            @click="active2=!active2">
            <i class='bx bxs-hand-left'></i>
        </vs-button>
        <vs-dialog overflow-hidden full-screen v-model="active">
        <template #header>
          <h4 class="not-margin">
                Welcome to our Service Box
          </h4>
          <div class="space">
              <i class='bx bx-customize'></i>
          </div>
        </template>
        <template>
            <form class="service-request-form" action="/apps" method="post">
                <vs-button
                  border
                  shadow 
                  light
                  class="docker-btn"
                 @click="active2=!active2">
                    <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAAAXNSR0IArs4c6QAAAVtJREFUSEvl1D1KXGEUxvHfdKKCNgYbJV1aNyC6AxM3EG1TuQNbu0BWEFfgxwZUsFdTWSrESsRYCOkSHnhHrtc7471+FOKBYWbeue/5n/Oc50zPK0fvlfN7X4BJzBZJf7WVtqtEtxjFDH63gXQF7Jak314CsI8FbOIrDsr3MOZxWOlgDsdNHQ3roA/oJ64ClrGFKVziDHtYrUPqgFTyHSf4iIlK5QEslgRH+IArbOCinH8psLtu6oAkSCXVaps6+IkV/MA6rmuVp8AUe2/RcuETpvG3XBgpn6vv+WkcNxjDKdYQGydynlzbdcC/NrYb8sx5SRx5/vSfq0r0XEAGHOnuRRWQlpae2MWd5sMA0TEO6hrRPOZ4dA8ypPg51uwSjdI0zSBnsVYWrA0kQ/08qPJBgJxnwSJXrNYE2ikWfDDQprbb/Nn1tzfy5dUp2gA6JRzmomclGnT57XfwH8+pPxnyQsmqAAAAAElFTkSuQmCC"/>"
                </vs-button>
                <div class="language">
                    <vs-switch v-model="active1" :active1="true" success required> GO </vs-switch>
                </div>
                <div class="modules">
                    <vs-input v-model="value" placeholder="Modules" type="text" :rules="ModulesRules" required/>
                </div>
                <div class="offered con-select">
                    <vs-select
                        filter
                        placeholder="Service"
                        v-model="value1">
                    <vs-option label="Docker" value="1">
                         Docker-Container
                    </vs-option>
                    <vs-option label="Docker-Compose" value="2">
                         Docker-Compose
                    </vs-option>
                    <vs-option label="Docker with Traefik proxy" value="3">
                         Docker with Traefik proxy
                    </vs-option>
                    <vs-option label="Docker Swarm" value="4">
                         Docker Swarm
                    </vs-option>
                    <vs-option label="None" value="5">
                         None
                    </vs-option>
                    </vs-select>
                </div>
                <div class="additional">
                    <vs-select
                        filter
                        placeholder="File-System"
                        v-model="value2">
                    <vs-option label="Disk Mount" value="1">
                         Disk-Mount (Volume)
                    </vs-option>
                    <vs-option label="None" value="2">
                         None
                    </vs-option>
                    </vs-select>
                </div>
                <div class="networks">
                    <vs-select
                        filter
                        placeholder="Networks"
                        v-model="value3">
                    <vs-option label="Configure-Network" value="1">
                         Configure Network
                    </vs-option>
                    <vs-option label="None" value="2">
                         None
                    </vs-option>
                    </vs-select>
                </div>
                <div class="name">
                    <vs-input type="text" v-model="name" placeholder=" Your Name"  :rules="nameRules"/>
                </div>
                <div class="email">
                    <vs-input type="text" v-model="email" placeholder=" Company email address" :rules="emailRules" required/>
                </div>
                <div class="action">
                    <div class="username">
                       <span> Name: </span> {{this.name}}
                    </div>
                    <div class="useremail">
                       <span> Email: </span> {{this.email}}
                    </div>
                    <div class="modules-projects">
                      <span> Project-Modules: </span> {{this.value}} 
                      <span> Modules Prices : </span> {{this.modules_prices}}
                    </div>
                    <div class="project-services">
                        <span> Service: </span> {{this.value1}}
                        <span> Volume-Option: </span> {{this.value2}}
                        <span> Network-Option: </span> {{this.value3}}
                        <span> Project-Services-Charges </span> {{this.charges}}
                    </div>
                    <div class="calculated-charges">
                        <span> {{this.total}} $ </span>
                    </div>
                </div>
                <vs-button border
                    shadow
                    success
                    v-on:click="submit($event)"
                    class="btn-submit">
                    <i class='bx bxs-flag-checkered'></i>
                </vs-button>
            </form>
        </template>
        </vs-dialog>
        </div>
    </div>
</template>

<script>
export default{
    data:() =>({
        active : false,
        active1: false,
        active2: false,
        value: '',
        value1: '',
        value2: '',
        value3: '',
        name: '',
        email: '',
        modules_prices : 0,
        charges: 0,
        total: 0,
        message: '',
        ModulesRules : [
            v => !!v || 'Modules must require',
            v => (v && v.max > 200) || 'Modules must be at least 200',
        ],
        nameRules : [
            u => !!u || 'Empty Field',
            u => (u && u.length <= 10) || 'Name too long',
        ],
        emailRules : [
            w => !!w || 'Email must be required',
            w => (w && w.length <= 50) || 'Email must be at least 50 characters',
        ]
    }),

    methods: {
        submit : function(event){
            if(this.value != ""){
               this.total += this.value * 10 
               switch(this.value1){
                   case "1":
                       this.total += 500;
                       this.message = "Docker-image: [1]. Application want cloud access ";
                       break;
                    case "2":
                        this.total += 700;
                        this.message = "Docker-container: [1]. Already application containerized";
                        break;
                    case "3":
                        this.total += 1200;
                        this.message = "Docker-container [1], additional features such as reverse proxy & load balancing: [1]. ";
                        break;
                    case "4":
                        this.total += 1350;
                        this.message = "Docker-container [3] as swarm mode";
                        break;
                    case "5":
                        this.total += 0;
                        this.message = "No service choose";
                        break;
                    default:
                        this.total += 0;
               }
            }
            switch(this.value2){
                case "1":
                    this.total += 200;
                    this.message = "Volume-Mount";
                break;
                case "2":
                    this.total += 0;
                    this.message = "No Volume-Mounting"
                default:
                    this.total += 0;
            }
            switch(this.value3){
                case "1":
                    this.total += 500;
                    this.message = "Networking configuration";
                break;
                case "2":
                    this.total += 0;
                    this.message = "No Network"
                default:
                    this.total += 0;
            }
        }
    }

    

    
}

</script>

<style>

.services{
    position: absolute;
    top: 200px;
    left:200px;
}
.center{
    position: absolute;
    top: 50px;
    left:350px;
}   

.business{
 width: 50px;
 height: 50px;
}

.space{
    position: absolute;
    left: 503px;
}
.not-margin{
    margin: 0px;
    font-weight: normal;
    padding: 10px;
}
.language{
    position: absolute;
    top: 56px;
}
.docker-btn{
    position: absolute;
    left:401px;
    top: 0px;
}
.modules{
    position: absolute;
    top: 53px;
    left: 81px;
}
.offered{
    position: absolute;
    top: 124px;
    left: 78px;
}
.additional{
    position: absolute;
    left:78px;
    top: 200px;
}
.networks{
    position: absolute;
    left:78px;
    top: 275px;
}
.btn-submit{
    position: absolute;
    top: 337px;
    left:161px;
    width: 50px;
    height:50px;
}
.name{
    position: absolute;
    left:313px;
    top: 53px;
}
.email{
    position: absolute;
    left: 315px;
    top: 124px;
}
.action{
    position: absolute;
    left:324px;
    top: 169px;
}

.useremail{
    position: absolute;
    top:40px;
}
.modules-projects{
    position: absolute;
    top: 93px;
}
.project-services{
    position: relative;
    top: 129px;
}
.calculated-charges{
    position: absolute;
    top:206px;
    left:133px;
}
.btn-back{
    position: absolute;
    top: 0px;
    left: -70px;
    width: 50px;
    height:50px;
}
</style>
