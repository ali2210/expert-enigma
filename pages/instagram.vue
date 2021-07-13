<template>
   <div>
      <vs-sidebar
            right
            absolute
             v-model="active0"
            open>
        <template #logo>
            <img src="../static/logo.png" alt=""/>
        </template>
        <template #header>
        </template>
        <vs-sidebar-item id="insta">
             <div>
                 <div class="center">
                    <vs-tooltip>
                     <vs-button
                        border
                        shadow 
                        dark
                        :active="true"
                     @click="active=!active">
                         <i class='bx bxl-instagram'></i>
                    </vs-button>
                        Connect
                    <template #tooltip>
                         Now you will connect instagram with your account
                    </template>
                    </vs-tooltip>
                </div>
                <div class="center">
                    <vs-tooltip>
                      <vs-button
                        border
                        shadow 
                        primary
                        gradient
                        :active1="true"
                        @click="active1=!active1">
                         <i class='bx bxl-instagram-alt'></i>
                     </vs-button>
                        Insta us
                     <template #tooltip>
                         Follow our instagram-page
                     </template>
                    </vs-tooltip>
                </div>
             </div>
        </vs-sidebar-item>
      </vs-sidebar>
       <vs-dialog width="550px" not-center v-model="active">
        <template #header>
          <h4 class="center">
            Welcome back
          </h4>
        </template>
       </vs-dialog>
   </div>
</template>

<script>
    export default{
        data:() =>({
            active: false,
            active1:false,
            active0: '',
            insta:'',
        })
    }
</script>

<style>
</style>