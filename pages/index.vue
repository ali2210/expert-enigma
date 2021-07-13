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
            <NuxtLink to="/">
              <i 
              shadow 
              dark
              gradient 
              class='bx bx-home'></i>
            </NuxtLink>
          </template>
          Home
        </vs-sidebar-item>
        <vs-sidebar-item id="shop">
          <template #icon>
           <NuxtLink to="/shop">
              <i 
              shadow 
              dark 
              gradient
              class='bx bxs-shopping-bag-alt'></i>
            </NuxtLink>
          </template>
          Shop
        </vs-sidebar-item>
        <vs-sidebar-item id="apps">
          <template #icon>
            <NuxtLink to="/apps">
              <i 
              shadow 
              dark 
              gradient
              class='bx bxs-category'></i>
            </NuxtLink>
          </template>
          Apps
        </vs-sidebar-item>
        <vs-sidebar-item id="news">
          <template #icon>
            <NuxtLink to="/news">
              <i 
              shadow 
              dark 
              gradient
              class='bx bx-news'></i>
            </NuxtLink>
          </template>
          News
        </vs-sidebar-item>
        <vs-sidebar-item id="gallery">
          <template #icon>
            <NuxtLink to="/gallery">
              <i 
              shadow 
              dark 
              gradient
              class='bx bx-images'></i>
            </NuxtLink>
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
              <NuxtLink to="/instagram">
                <i 
                  shadow 
                  dark
                  gradient
                  class='bx bxl-instagram'></i>
              </NuxtLink>
            </template>
            Instagram
          </vs-sidebar-item>
          <vs-sidebar-item id="twitter">
            <template #icon>
                <NuxtLink to="/twitter">
                  <i
                  shadow 
                  dark 
                  gradient
                  class='bx bxl-twitter' ></i>
                </NuxtLink>
            </template>
            Twitter
          </vs-sidebar-item>
          <vs-sidebar-item id="Facebook">
            <template #icon>
               <NuxtLink to="/facebook">
                <i 
                shadow 
                dark
                gradient
                class='bx bxl-facebook' ></i>
               </NuxtLink>
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
                <NuxtLink to="/wallet">
                 <i 
                  shadow 
                  dark 
                  gradient
                  class='bx bxs-coin-stack' ></i>
                </NuxtLink>
            </template>
            Debit
          </vs-sidebar-item>
          <vs-sidebar-item id="sign-in">
            <template #icon>
              <NuxtLink to="/signin">
              <i 
                shadow
                dark
                gradient
                class='bx bxs-user' ></i>
              </NuxtLink>
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
          We are the wizards. Our products are magical. Beleive us!. If you don't then buy our products
          & services.    
      </div>
      <div class="storiesdiv">
        Everyone have a story. Stories are short, long, easy & diffcult. Every story speak itself about a hero who will chase dreams, 
        want to paint the world. In the story, we found how can you pain the world & enhance beauty of a world. 
        Look through eyes of artists. You found nothing ... However, you found something that why you care about hero
        story.   
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
.storiesdiv{
  position: absolute;
  left: 125px;
  width: 300px;
  top: 242px;
}

.NuxtLink{
  color: rgb(0,255,0);
}

</style>
