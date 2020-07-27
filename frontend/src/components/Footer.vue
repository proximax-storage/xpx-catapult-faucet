<template>
  <v-footer app absolute>
    <v-row>
      <v-col cols="12" class="text-center">
        <div class="text-caption font-weight-normal">
          <span>© ProximaX 2020. All Rights Reserved.</span>
          <span class="font-weight-bold primary--text">&nbsp; Sirius Faucet {{version}}</span>
        </div>
      </v-col>
    </v-row>
  </v-footer>
</template>

<script>
export default {
  data: () => ({
    version: 'v0.3.2'
  })
}
</script>
