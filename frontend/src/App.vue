<template>
  <div id="app">
    <!-- <div>
      <img src="@/assets/img/background-proximax.jpg"  style="width: 100%; height: 300px;">
    </div>-->
    <div class="header">
      <img class="logo responsive" src="@/assets/img/logo-proximax-sirius-faucet_new.svg">
    </div>
    <main>
      <router-view/>
      <footer>
        <p class="copyright">© ProximaX 2019. All Rights Reserved. <b> Sirius Faucet v0.0.2</b> </p>
      </footer>
    </main>

    <notifications group="foo"/>
  </div>
</template>

<script>
export default {};
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

}
.responsive{
  width: 250px;
}

@media screen and ( max-width: 1024px ) {
img.responsive { width: 200px; }
}

@media screen and ( max-width: 800px ) {
img.responsive { width: 170px; }
}
@media screen and ( max-width: 500px ) {
img.responsive { width: 150px; }
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
  font-family: Open Sans;
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
