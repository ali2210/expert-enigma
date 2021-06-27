<template>
  <div class="center examplex">
      <vs-navbar center-collapsed v-model="active">
        <template #left>
          <img :src="'../logo.png'" class="logo" alt="">
        </template>
        <vs-navbar-item :active="active == 'home'" id="home">
          Home
        </vs-navbar-item>
        <vs-navbar-item :active="active == 'shop'" id="shop">
          Shop
        </vs-navbar-item>
        <vs-navbar-item :active="active == 'gallery'" id="gallery">
          Gallery
        </vs-navbar-item>
        <vs-navbar-item :active="active == 'market'" id="market">
          Marketplace
        </vs-navbar-item>
        <template #right>
          <vs-button flat >Login</vs-button>
          <vs-button>Get Started</vs-button>
        </template>
      </vs-navbar>
      <template>
      <div class="window-background">
        <div class="container" style="font-size: larger;
  font-style: oblique;
  font-weight: bolder;
  font-variant: petite-caps;position:relative;left:438px;font-weight:bolder;font-style:oblique;">
            <template>
              <!-- sky 4k image -->
            </template>          
          Explore your limits way beyond the stars , and consciousness?
          <template>
            <div class="btns">
              <vs-button
                color="#7d33ff"
                relief
                border @click="active3=!active3">
                Contribute
              </vs-button>
              <vs-dialog width="300px" not-center v-model="active3">
                <template #header> 
                  <h4 class="not-margin">
                    $ docker run --hostname=wizdwarfs -p 5000:5000 -it ali2210/wiz-dwarfs
                  </h4>
                </template>
              </vs-dialog>  
            </div>
            <div class="btns-right">
               <vs-button
                 color="rgb(59,222,200)"
                  gradient
                  :active="active == 6"
                  @click="active = 6"
                >
                Open Projects
               </vs-button>
            </div>  
          </template> 
          <template>
            <div class="philosphyDiv">
              Life is like a codebase. Every instruction documentated in our DNA-REPO. When our body update
              nature codebase. First of all agent download the data from DNA-REPO, then compile on RNA . These instructions
              generate life code in modular form.   
            </div>
            <div class="alienDiv">
              When unknown code inject in our hardware (Body). That unknown code is alien code / virus.
              Around 70 -80% instructions are written in alien code are similar to us. That is
              reason alien code instruction direcly compile on our body. But when our agent detect
              alien code it trigger preventative measure against hardware.   
            </div>
            <span class="questiondiv"> We don't have modular computer. How can you compile alien instructions? </span>
            <span class="answerdiv"> The Answer is simple. @wizdwarfs building modular computing interface. You can try.</span>
          </template>
          <template>
            <div class="goaldiv">
              <template>
                  <vs-button
                      class="goal1"
                      shadow
                      :active="active == 0"
                      @click="active = 0">
                  </vs-button>
                <div class="linediv"></div>
                <vs-button
                      class="goal2"
                      shadow
                      :active="active == 0"
                      @click="active = 0">
                  </vs-button>
                <div class="linediv1"></div>
                <vs-button
                      class="goal3"
                      shadow
                      :active="active == 0"
                      @click="active = 0">
                  </vs-button>
                  <div class="linediv2"></div>
                <vs-button
                      class="goal4"
                      shadow
                      :active="active == 0"
                      @click="active = 0">
                  </vs-button>
              </template>
            </div>
          </template>
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
    active: 'home'
  }),

  data:()=>({
    active3 : false,
    active : ''
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
.window-background{
  position: absolute;
  top: 0px;
  left: 0px;
  background: black;
  color: white;
  width: 100%;
  border-radius: 10px;
}
.btns{
  position: absolute;
  top: 261px;
  left: 168px;
}
.btns-right{
  position: absolute;
  left: 26px;
  top: 261px;
}
.not-margin{
  color: darkgrey;
  position: relative;
  top: 20px;
}
.logo{
  position: relative;
  left: 332px;
  width: 27%;
}
.philosphyDiv{
  position: absolute;
  top: 428px;
  width: 30%;
  left: -325px;
  color: black;
}
.alienDiv{
  position: relative;
  top: 491px;
  width: 70%;
  color: black;
  left: 181px;
}
.questiondiv{
  position: relative;
  top: 165px;
  left: -875px;
}
.answerdiv{
  position: relative;
  top: 162px;
  right: 747px ;
}
.goaldiv{
  position: absolute;
  top: 607px;
  left: -405px;
}
.linediv{
  position: absolute;
  width: 70px;
  height: 4px;
  background: lawngreen;
}
.goal1{
  position: relative;
  top: 15px;
  left: -12px;
  height: 11px;
  width: 12px;
  background: purple;
}
.goal2{
  position: relative;
  top: -1px;
  left: 61px;
  height: 11px;
  width: 12px;
  background: gold;
}

.linediv1{
  position: absolute;
  width: 70px;
  height: 4px;
  background: lawngreen;
  left: 68px;
  top:21px;
}

.goal3{
  position: relative;
  top: -17px;
  left: 124px;
  height: 11px;
  width: 12px;
  background: seagreen;
}

.linediv2{
  position: absolute;
  width: 70px;
  height: 4px;
  background: lawngreen;
  left: 141px;
  top:21px;
}

.goal4{
  position: relative;
  top: -33px;
  left: 201px;
  height: 11px;
  width: 12px;
  background: seagreen;
}


</style>
