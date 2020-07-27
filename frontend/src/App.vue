<template>
  <v-app>
    <v-app-bar color="white" app flat dense height="70" clipped-left class="d-flex align-baseline">
      <div>
        <!-- Logo -->
        <img class="mr-2" :src="require('@/assets/img/logo.svg')" alt="logo" height="50" />
        <span class="caption font-weight-regular">{{version}}</span>
      </div>
    </v-app-bar>

    <!-- Router -->
    <v-main>
      <router-view :key="$route.path"></router-view>
    </v-main>

    <!-- Footer -->
    <my-footer/>
  </v-app>
</template>

<script>
import Footer from '@/components/Footer.vue'

export default {
  name: 'App',
  data: () => ({
    version: 'v0.3.2'
  }),
  components: {
    'my-footer': Footer
  }
}
</script>
