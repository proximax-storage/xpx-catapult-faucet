<template>
  <v-container class="mt-12">
    <!-- Title, subtitle -->
    <v-row>
      <v-col cols="12" md="11" class="text-center mx-auto">
        <p class="primary--text text-h4 font-weight-bold" v-html="title"></p>
        <p class="font-weight-bold" v-html="subtitle"></p>
        <v-divider></v-divider>
      </v-col>
    </v-row>

    <div class="fauce-container">
      <div class="container-fluid">
        <form id="fauceApp" v-on:submit.prevent="sendFaucet">
          <v-row>
            <v-col cols="12" sm="10" md="8" class="mx-auto">
              <v-text-field
                v-model="address"
                dense
                counter
                outlined
                rounded
                maxlength="46"
                name="address"
                type="text"
                placeholder="Enter your testnet address here..."
              >
                <template v-slot:prepend-inner>
                  <v-img
                    class="mr-5 mt-2 mb-4"
                    alt="logo"
                    height="30"
                    width="30"
                    src="@/assets/img/icon-wallet-name-red-16h-proximax-sirius-faucet.svg"
                  ></v-img>
                </template>
              </v-text-field>
            </v-col>

            <v-col cols="12" sm="10" md="8" class="mx-auto">
              <div class="d-flex justify-center">
                <v-btn
                  class="text-transform-none"
                  :disabled="sendButton.disabled"
                  :loading="sendButton.loading"
                  color="primary"
                  rounded
                  width="180"
                  height="40"
                >{{sendButton.text}}</v-btn>
              </div>
            </v-col>
          </v-row>

          <!-- <div class="container mt-1rem">
            <div class="input-icon-wrap">
              <span class="input-icon">
                <span class="fa fa-user">
                  <img
                    class="icon-wallet"
                    src="@/assets/img/icon-wallet-name-red-16h-proximax-sirius-faucet.svg"
                  />
                </span>
              </span>
              <input
                class="input-with-icon"
                id="form-name"
                maxlength="46"
                v-bind:class="[classValdiate]"
                style="width: 100%;"
                type="text"
                v-model="address"
                placeholder="Enter your testnet address here..."
                @input="checkForm($event)"
              />
            </div>

            <div class="div-alert-war">
              <div v-show="showValidate" :class="alertclass">
                <div :class="div1">{{msjValidate}}</div>
                <div :class="div2">
                  <div :class="loaderclass"></div>
                </div>
              </div>
            </div>
            <div class="button-class mt-1rem">
              <button :disabled="isDisabled" type="submit" class="btn">SEND</button>
            </div>
          </div>-->
        </form>
      </div>
    </div>
  </v-container>
</template>
<script>
import Utils from '@/services/Utils.js'

export default {
  data () {
    return {
      title: 'Sirius Chain Testnet',
      subtitle: 'Top-up your account balance, a limit every 24 hours',
      button: {
        disabled: false,
        loading: false,
        text: 'SEND'
      },
      roimer: 'aqui voy yo',
      xpxMaximium: 0,
      cont: 0,
      alertclass: '',
      div1: '',
      div2: '',
      loaderclass: 'loader',
      address: null,
      buttonValidate: true,
      classValdiate: '',
      showValidate: false,
      msjValidate: ''
    }
  },
  beforeMount () {
    this.getConfig()
  },
  methods: {
    getConfig () {
      this.cont = this.cont + 1
      this.$apiService
        .getConfig()
        .then(x => {
          if (x.data && x.data.length > 0) {
            console.log('Configs:', x.data)
            this.configs = x.data
          }
        })
        .catch(e => console.log(e))
      // this.$apiService
      //   .get(`api/faucet/config`)
      //   .then(response => {
      //     console.log('server response', response)
      //     // this.xpxMaximium = this.amountFormatterSimple(response.data);
      //   })
      //   .catch(error => {
      //     this.xpxMaximium = 0;
      //     if (this.cont < 5) {
      //       this.getMaxXpx();
      //     }
      //   });
    },
    amountFormatterSimple: function (amount) {
      const amountDivisibility = Number(amount) / Math.pow(10, 6)
      return amountDivisibility.toLocaleString('en-us', {
        minimumFractionDigits: 0
      })
    },
    checkForm: function (e) {
      this.loaderclass = ''
      this.address = Utils.addressForm(this.address)
      if (Utils.isValid(this.address)) {
        this.showValidate = false
        this.msjValidate = ''
        this.classValdiate = 'success'
      } else {
        this.showValidate = true
        this.msjValidate = 'Address must have 40 characters'
        this.isB = true
        this.classValdiate = 'error'
        this.div1 = 'div-alert-text'
        this.div2 = ''
        this.alertclass = 'alert-war'
      }
    },
    sendFaucet () {
      if (Utils.isValid(this.address)) {
        this.showValidate = true
        this.alertclass = 'alert-war'
        this.loaderclass = 'loader'
        this.div1 = 'div-loader-text'
        this.div2 = 'div-loader-animate'
        this.msjValidate = 'sending XPX    '
        this.buttonValidate = false
        this.$apiService
          .get(`faucet/GetXpx/${Utils.clean(this.address)}`)
          .then(response => {
            this.msjValidate = ''
            this.classValdiate = ''
            this.buttonValidate = true
            this.address = ''
            this.loaderclass = ''
            this.div1 = 'div-loader-text'
            this.div2 = ''
            this.alertclass = 'alert-success'
            this.showValidate = true
            this.msjValidate = `${response.data}`
            setTimeout(
              function () {
                this.showValidate = false
              }.bind(this),
              15000
            )
          })
          .catch(error => {
            this.classValdiate = ''
            this.buttonValidate = true
            this.address = ''
            let msj = null
            msj = error.response.data.message
              ? error.response.data.message
              : error.response.data
            this.loaderclass = ''
            this.div1 = 'div-alert-text'
            this.div2 = ''
            this.alertclass = 'alert-error'
            this.showValidate = true
            this.msjValidate = `${msj}`
            setTimeout(
              function () {
                this.showValidate = false
              }.bind(this),
              10000
            )
          })
      }
    }
  },
  computed: {
    isDisabled: function () {
      return !this.buttonValidate
    },
    sendButton () {
      return this.button
    }
  }
}
</script>
