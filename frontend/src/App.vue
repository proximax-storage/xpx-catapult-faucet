<template>
  <div id="app">
    <!-- <div>
      <img src="@/assets/img/background-proximax.jpg"  style="width: 100%; height: 300px;">
    </div>-->
    <div class="header">

      <div class="logo"> <img class=" responsive" src="@/assets/img/logo-proximax-sirius-faucet_new.svg"> {{version}}</div>
      
    </div>
    <main>
      <router-view/>
      <footer>
        <p class="copyright">© ProximaX 2022. All Rights Reserved. <b> Sirius Faucet {{version}}</b> </p>
      </footer>
    </main>

    <notifications group="foo"/>
  </div>
</template>

<script>
export default {

  data() { return {
    version: "v0.4.1"
    }}
};
</script>

<style lang="scss">
.header {
  overflow: hidden;
  padding: 20px 10px;
}
.logo {
  float: none;
  display: block;
  text-align: left;
  font-size: 15px;

}
.responsive{
  width: 250px;
}

@media screen and ( max-width: 1024px ) {
img.responsive { width: 200px; }
.logo {
   font-size: 10px;
}
}

@media screen and ( max-width: 800px ) {
img.responsive { width: 170px;
 }
 .logo {
   font-size: 10px;
}
}
@media screen and ( max-width: 500px ) {
img.responsive { width: 150px;
 .logo {
   font-size: 10px;
}

 }
}
@media screen and ( min-width: 1025px) {
img.responsive { width: 300px;}
}
img.responsive { height: auto; }
body {
  overflow: auto;
  background-repeat: no-repeat;
  background-size: contain;
  display: block;
  margin: 0px;
  flex-direction: column;
  background-color: #ffffff;
 font-family: 'Lato', sans-serif;
}

html {
  height: 100%;
}

footer {
  position: relative;
}
footer>p>b {
    color: #df4c48;
}
p.copyright {
  position: fixed;
  margin-bottom: 0%;
  width: 100%;
  color: #000000;
  line-height: 40px;
  font-size: 12px;
  text-align: center;
  bottom: 0;
}
</style>
