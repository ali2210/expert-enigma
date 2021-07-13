<template>
<div>
      <vs-sidebar
            right
            absolute
             v-model="active"
            open>
        <template #logo>
            <img src="../static/logo.png" alt=""/>
        </template>
        <template #header>
        </template>
        <vs-sidebar-item id="signup">
            <div class="center">
                    <vs-tooltip>
                      <vs-button
                        border
                        shadow 
                        success
                        gradient
                        :active0="true"
                        @click="active0=!active0">
                         <i class='bx bxs-user-plus'></i>
                     </vs-button>
                        Sign-up
                     <template #tooltip>
                         New Member
                     </template>
                    </vs-tooltip>
                </div>
                <div class="center">
                    <NuxtLink to="/signin">
                    <vs-tooltip>
                      <vs-button
                        border
                        shadow 
                        primary
                        gradient
                        :active1="true"
                        @click="active1=!active1">
                         <i class='bx bxs-user-badge'></i>
                     </vs-button>
                        Sign-in
                     <template #tooltip>
                         Already Member
                     </template>
                    </vs-tooltip>
                    </NuxtLink>
                </div>
                </vs-sidebar-item>
                </vs-sidebar>
                <vs-dialog overflow-hidden full-screen not-center v-model="active0">
                     <template #header>
                        <h4 class="center">
                            User-Profile
                        </h4>
                     </template>
                     <template>
                         <div class="con-form">
                                <div class="row">
                                    <div class="col sm-4">
                                        <vs-input type="text" v-model="input1" placeholder="Name">
                                    <template #icon>
                                        <i class='bx bxs-user-circle'></i>
                                    </template>
                                    </vs-input>
                                    </div>
                                </div>
                                
                         </div>
                     </template>
                </vs-dialog>
            </div>
</template>


<script>
export default{
    data:()=>({
        active : '',
        active0: false,
        active1: false,
        input1: ''
    })
}
</script>


<style></style>