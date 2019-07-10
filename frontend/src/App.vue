<template>
  <div id="app">
    <!-- <div>
      <img src="@/assets/img/background-proximax.jpg"  style="width: 100%; height: 300px;">
    </div>-->
    <main>
      <router-view/>
      <footer>
        <p class="copyright">Copyright © 2019. ProximaX. All Rights Reserved.</p>
      </footer>
    </main>

    <notifications group="foo"/>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss">
body {
  // background-image: url("./assets/img/background-proximax.jpg");
  background-repeat: no-repeat;
  background-size: contain;
  display: block;
  margin: 0px;
  background-color: #ffffff;
  font-family: Open Sans; 
}

html {
  height: 100%;
}


p.copyright {
  background-color: #914090;
  margin-bottom: 0%;
  position: absolute;
  width: 100%;
   color: #ffffff;
  line-height: 40px;
  font-size: 1rem;
  text-align: center;
  bottom: 0;
}
</style>
