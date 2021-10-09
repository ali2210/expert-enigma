<template>
  <div class="window">
  <template>
      <vs-sidebar
        absolute
        v-model="active"
        open>
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
          Apps & Services
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
      </vs-sidebar>
  </template>
  <template>
    <div class="billboard">
      <vs-card>
        <template #title>
            <h3> Wizdwarfs </h3>
        </template>
        <template #img>
          <img src="../static/background.png" alt="biocomputer"/>
        </template>
        <template #text>
          <div class="heading">
            <a href="#"> Read more ...</a>
          </div>
        </template>
      </vs-card>
        <div class="leaderboard">
          <vs-card>
            <template #title></template>
            <template #img>
              <img src="../static/download.jpeg" alt="DNA-computing"/>
            </template>
            <template #text>
              <div class="board">
                  Paint the World 
              </div>
            </template>
          </vs-card>
        </div>
        <template >
          <div class="social">
            <div class="socialLinks">
              <vs-button class='facebook'>
                  <i class='bx bxl-facebook-circle' ></i>
              </vs-button>
              <vs-button class='twitter'>
                <i class='bx bxl-twitter' ></i>
              </vs-button>
              <vs-button class='instagram'>
                <i class='bx bxl-instagram' ></i>
              </vs-button>
              <vs-button class='linkedin'>
                <i class='bx bxl-linkedin' ></i>
              </vs-button>
            </div>
          </div>
          <div class="codebase">
            <span class="commit"> <i class='bx bx-git-commit' > 769 </i></span>
          </div>
          <div class="fork">
            <vs-button class="forked">
              <i class='bx bx-git-repo-forked'></i>
            </vs-button>
          </div>
        </template>
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
  position:absolute;
  top: 100px;
  left: 318px;
  width: 50rem;
  shape-rendering: geometricprecision;

}
.heading{
  position: absolute;
  top: 133px;
  left:229px;
  width: 50%;
}

.NuxtLink{
  color: rgb(0,255,0);
}

.leaderboard{
  position: absolute;
  top:203px;
  left:6px;
  width:100%;
}
.board{
  position: relative;
  top:7px;
  left:97px;
  font-size: large;
  color:black;
  font-family: system-ui;
  font-stretch: semi-condensed;
}
.social{
  position:relative;
  height:200px;
  background: black;
  top:300px;
  left:-329px;
  border-radius: 42px;
  width: 1313px;
}
.socialLinks{
  position: relative;
  left:100px;
  top: 45px;
  width:1200px;
}
.twitter{
  position: relative;
  top:39px;
}
.instagram{
  position: relative;
  left:67px;
  top:-66px;
}
.linkedin{
  position: relative;
  top: -27px;
  left: 67px;
}
.codebase{
  position: relative;
  height: 100px;
  width: 100px;
  background: white;
  top:143px;
  border-radius: 63px;
  border:2px dotted blue;
}
.commit{
  position: relative;
  left:10px;
  top: 40px;
  color: blueviolet;
  font-size:large;
}
.fork{
  position: relative;
  top:27px;
}
.forked{
  position: relative;
  border-radius: 42px;
  top:80px;
  left:-9px;
}


</style>
