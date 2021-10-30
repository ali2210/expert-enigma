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
        <!-- <vs-sidebar-item id="shop">
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
        </vs-sidebar-item> -->
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
                  <a href="https://m.facebook.com/wizenigma0xGACT/">
                    <i class='bx bxl-facebook-circle' ></i>
                  </a>
              </vs-button>
              <vs-button class='twitter'>
                <a href='https://twitter.com/DwarfWiz'>
                    <i class='bx bxl-twitter' ></i>
                </a>
              </vs-button>
              <vs-button class='instagram'>
                <a href="https://www.instagram.com/wisdomenigma/?fbclid=IwAR0st2X8Oah6gtlQBOHbvBe6DK4S7p-aqUHG7iuF7vpYRT7WCoTRONSmxBs">
                <i class='bx bxl-instagram' ></i>
                </a>
              </vs-button>
              <vs-button class='linkedin'>
                <a href='https://www.linkedin.com/company/wisdom-enigma/?viewAsMember=true'>
                  <i class='bx bxl-linkedin' ></i>
                </a>
              </vs-button>
            </div>
          </div>
          <div class="codebase">
            <span class="commit"> <i class='bx bx-git-commit' > 769 </i></span>
            <div class="fork">
            <vs-button class="forked">
              <a href="https://github.com/ali2210/WizDwarf">
                        <i class='bx bx-git-repo-forked'></i>
              </a>
            </vs-button>
            </div>
          </div>
          <div class="sign">
            <vs-button v-on:click="onGithubLogin($event)" class="github">
              <i class='bx bxl-github'></i>
            </vs-button>
            <vs-button v-on:click="onGoogleLogin($event)" class="google">
              <i class='bx bxl-google'></i>
            </vs-button>
            <vs-button v-on:click="onFacebookLogin($event)" class="facebook-sq">
              <i class='bx bxl-facebook'></i>
            </vs-button>
          </div>
          <div class="payment">
            <vs-button class="payments" dark v-on:click="acceptedBtc">
                <i class='bx bxl-bitcoin'></i>
            </vs-button>
          </div>
        </template>
    </div>
    <div class="parallex">
      <template>
        <img class="abstract" src="../static/louvre.jpg"/>
        <div class="abstractText">
        </div>
        <div class="contentread">
          <img class="dance" src="../static/dance.jpg" alt="dance"/>
          <span class="learn">
            We design abstract computing algorithms and also conjecture problems. These Hard problems serve as primary key to explore limits of computing and the brain.
            We are unware of many conjecture '&' fractalism patterns. These dance are so beautiful and it attract everyone and we love these dance. We can acheive through the Quantum Computing Interface because classical have some limitations.
          </span>
        </div>
        <div class="nft-card">
          <div class="nft-link">
            <vs-card type="3">
               <template #title>
                 <h3> Wizdwarfs </h3>
               </template>
               <template #link>
                  <a href="https://opensea.io/assets/matic/0xe7ea2e2be12c257d376400cb231d8ee51e972bd6/24547768641536978472723590075963550080531792743048580483079221528577123894908/">
                    <i class='bx bxs-crown'> Gemstones </i>
                  </a>
              </template>
               <template #text>
                  <p>
                    We offer our artwork NFTs as royalty. Buy royalty if you like to.
                  </p>
              </template>
            </vs-card>
            <div class="nft-2">
              <vs-card type="5">
                <template #title>
                 <h3> Opal Artwork </h3>
                </template>
                <template #image>
                  <!-- <img src="../static/illusion.jpeg" alt=""/> -->
                </template>
                <template #link>
                  <a href="https://opensea.io/assets/matic/0xe7ea2e2be12c257d376400cb231d8ee51e972bd6/41506717256369801583211002721974187832999321087580976075953145296336844711889/">
                    <i class='bx bxs-shopping-bag'> Opal </i>
                  </a>
                </template>
                <template #text>
                 <p>
                   Software is an illusion image of a mind 
                </p>
                </template>
            </vs-card>
            </div>
          </div>
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
  methods: {
    onGithubLogin : function(event){
      alert("Github login functionality is not enabled yet");
    },
    onFacebookLogin : function(event){
      alert("Fb login functionality is not enabled yet");
    },
    onGoogleLogin : function(event){
      alert("Google login functionality is not enabled yet");
    },
    acceptedBtc : function(){
      const action_btc = document.getElementsByClassName("board")[0];
      const divContent = document.createElement("div");
      divContent.className = "accepted";
      divContent.style.position = "relative";
      divContent.style.top = "2px";
      divContent.style.left = "-96px";
      divContent.style.fontFamily = "system-ui";
      divContent.style.fontSize = "smaller";
      divContent.style.color = "blueviolet";
      const spanText = document.createElement("span");
      spanText.innerHTML = "1NKwtLpNmhyyYCPxTYZXuqNbAxmouB7RkB";
      divContent.appendChild(spanText);
      action_btc.appendChild(divContent);
    }
  },
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
  left:198px;
  top: 73px;
  width:1200px;
}

.bxl-facebook-circle, .bxl-instagram, .bxl-twitter, .bxl-linkedin, .bx-git-repo-forked{
  color: white
}

.twitter{
  position: relative;
  top:34px;
}

.instagram{
  position: relative;
  left:67px;
  top:-72px;
}

.linkedin{
  position: relative;
  top: -37px;
  left: 67px;
}

.codebase{
  position: relative;
  height: 100px;
  width: 100px;
  background: white;
  top:171px;
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
  top:-71px;
  left: -24px;
}

.forked{
  position: relative;
  border-radius: 42px;
  top:80px;
  left:-9px;
}

.sign{
  position: absolute;
  left:74px;
  top:577px;
}

.google{
  position: absolute;
  left:46px;
}

.facebook-sq{
  position: absolute;
  top:-43px;
  left:44px;
}

.payment{
  width: 100px;
  position: absolute;
  top:227px;
  left:68px;
  height:41px;
  display: contents;
}

.payments{
  position: relative;
  left:171px;
  top:103px;
  animation: drops 2s alternate infinite;
}

.payments:hover{
  background: #35495e;
}

@keyframes drops {
  0%{
    top:103px;
  }
  5%{
    top: 102px;
  }
  10%{
    top: 101px;
  }
  20%{
    top: 90px;
  }
  30%{
    top: 60px;
  }
  50%{
    top: 30px;
  }
  60%{
    top: -30px;
  }
  70%{
    top: 60px;
    border: 1px dotted red;
    transform: scale3d(1, 1, 1);
  }
  80%{
    top: 90px;
    border: 1px solid gold;
    transform: scale3d(0.5, 0.5, 0.5);
  }
  90%{
    top:103px;
    border : 1px dashed red;
    transform: scale3d(1.5, 1.5, 1.5);
  }
  100%{
    border:none;
  }
}

.bxl-bitcoin{
  background: linear-gradient(90deg,red, orange);
  color : yellow;
  transform: scale3d(1.5,1.5,1.5);
  border-radius: 10px;
  border:none;
}

.parallex{
  position: absolute;
  top:759px;
}

.abstractText{
  width: 805px;
  background:whitesmoke;
  height:19px;
  position: absolute;
}

.contentread{
  position: absolute;
  width: 285px;
  top: 567px;
  left: 200px;
  height:98px;
  color:purple;
}

.dance{
  width: 233px;
  height:127px;
  position: relative;
  top: 28px;
}

.learn{
  position: relative;
  top:-33px;
  left:245px;
}
.nft-card{
  position: relative;
  width: 562px;
  height:539px;
  left:797px;
  top: -523px;
  border-radius: 26px;
}
.nft-link{
  position: relative;
  top:90px;
  left:243px;
  font-size: large;
}
.nft-2{
  position: relative;
  top:127px;
  font-size: large;
}
.bxs-crown{
  color:white
}
</style>
