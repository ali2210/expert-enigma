<template>
  <div class="center examplex">
    <template>
    <div class="hidden">
      <vs-sidebar
        absolute
        hover-expand
        reduce
        v-model="active"
        open
        >
        <template #logo>
          <img src="../static/logo.png" alt=""/>
        </template>
        <vs-sidebar-item id="home">
          <template #icon>
            <i 
            shadow 
            dark 
            class='bx bx-home'></i>
          </template>
          Home
        </vs-sidebar-item>
        <vs-sidebar-item id="shop">
          <template #icon>
            <i 
            shadow 
            dark 
            class='bx bxs-shopping-bag-alt'></i>
          </template>
          Shop
        </vs-sidebar-item>
        <vs-sidebar-item id="news">
          <template #icon>
            <i 
            shadow 
            dark 
            class='bx bx-news'></i>
          </template>
          News
        </vs-sidebar-item>
        <vs-sidebar-item id="gallery">
          <template #icon>
            <i 
            shadow 
            dark 
            class='bx bx-images'></i>
          </template>
          Gallery
        </vs-sidebar-item>
        <vs-sidebar-group>
          <template #header>
            <vs-sidebar-item arrow>
              <template #icon>
                <i class='bx bx-group'></i>
              </template>
                Connect
            </vs-sidebar-item>
          </template>
          <vs-sidebar-item id="Instagram">
            <template #icon>
              <i class='bx bxl-instagram'></i>
            </template>
            Instagram
          </vs-sidebar-item>
          <vs-sidebar-item id="twitter">
            <template #icon>
              <i class='bx bxl-twitter' ></i>
            </template>
            Twitter
          </vs-sidebar-item>
          <vs-sidebar-item id="Facebook">
            <template #icon>
              <i class='bx bxl-facebook' ></i>
            </template>
            Facebook
          </vs-sidebar-item>
        </vs-sidebar-group>
        <vs-sidebar-group>
          <template #header>
            <vs-sidebar-item arrow>
              <template #icon>
                <i class='bx bxs-user-circle' ></i>
              </template>
                Profile
            </vs-sidebar-item>
          </template>
          <vs-sidebar-item id="wallet">
            <template #icon>
              <i class='bx bxs-coin-stack' ></i>
            </template>
            Debit
          </vs-sidebar-item>
          <vs-sidebar-item id="sign-in">
            <template #icon>
              <i class='bx bxs-user' ></i>
            </template>
            Sign-in
          </vs-sidebar-item>
        </vs-sidebar-group>
      </vs-sidebar>
    </div>    
  </template>
  <template>
    <div class="billboard">
      <div class="heading">
          We are the wizards. We are building products that are unique, cool & bring ease in your life.    
      </div>
    </div>
  </template>
  </div>
</template>
<script>
import Logo from '~/components/Logo.vue'
import VuesaxLogo from '~/components/VuesaxLogo.vue'


export default {
  data:()=>({
    active:'home',
  }),
  
}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family:
    'Quicksand',
    'Source Sans Pro',
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    Roboto,
    'Helvetica Neue',
    Arial,
    sans-serif;
  display: block;
  font-weight: 300;
  font-size: 55px;
  color: #35495e;
  letter-spacing: 1px;
  text-transform: capitalize;
  margin: 25px 0;
}

.subtitle {
  font-weight: 300;
  font-size: 1.1rem;
  color: #526488;
  word-spacing: 2px;
  padding-bottom: 15px;
  max-width: 600px;
}

.subtitle a {
  font-weight: 500;
  color: inherit;
}

.links {
  padding-top: 15px;
  margin-bottom: 20px;
}

.content-logos {
  display: flex;
  align-items: center;
  justify-content: center;
  min-width: 500px;
}

.plus {
  font-size: 2.5rem;
  margin: 15px;
  color: #35495e;
}

.h3 {
  font-family:
    'Quicksand',
    'Source Sans Pro',
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    Roboto,
    'Helvetica Neue',
    Arial,
    sans-serif;
  font-weight: 400;
  margin: 10px;
}
.logo{
  position: relative;
  left: 332px;
  width: 27%;
}
.billboard{
  position: relative;
  left: 0px;
  top: 233px;
  font-family: monospace;
}
.heading{
  position: absolute;
  width: 300px;
  text-align: center;
  left: 123px;
}



</style>
