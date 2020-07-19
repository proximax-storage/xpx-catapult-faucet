<template>
  <v-container class="mt-12">
    <!-- Title, subtitle -->
    <v-row>
      <v-col cols="12" md="11" class="text-center mx-auto">
        <div>
          <span class="primary--text text-h4 font-weight-bold" v-html="title"></span>
        </div>
        <div class="mt-2">
          <span class="font-weight-bold" v-html="subtitle"></span>
        </div>
        <v-divider></v-divider>
      </v-col>
    </v-row>

    <div class="fauce-container">
      <div class="container-fluid">
        <v-form v-model="isValidForm" ref="form">
          <v-row>
            <v-col cols="12" sm="7" md="8" class="mx-auto mt-5" v-if="mosaicsList.length > 0">
              <div class="d-flex flex-wrap justify-center">
                <v-card v-ripple v-for="(item, i) in listMosaic" :key="i" outlined tile>
                  <div
                    @click="mosaicSelected = item.name"
                    v-ripple
                    class="pa-4 cursor-p text-center"
                    :class="mosaicSelected === item.name ? 'primary white--text': ''"
                  >
                    <span class="caption font-weight-medium text-body-2" v-html="item.name"></span>
                    <br />
                    <span class="caption font-weight-normal text-caption">Maximium <b>{{item.max}}</b></span>
                  </div>
                </v-card>
              </div>
            </v-col>

            <v-col cols="12" sm="10" md="8" class="mx-auto mt-5 pb-0">
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
                :rules="[configForm.rules.required, configForm.rules.min, configForm.rules.max, validateAddress]"
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

            <v-col cols="11" sm="9" md="7" class="mx-auto pt-0" v-if="showValidate">
              <v-alert
                dark
                rounded
                dense
                outlined
                class="mb-0"
                border="left"
                :type="typeMessage"
              >{{msjValidate}}</v-alert>
              <!-- <v-text-field class="pt-0 mt-0" color="success" loading disabled></v-text-field> -->
            </v-col>

            <v-col cols="12" sm="10" md="8" class="mx-auto">
              <div class="d-flex justify-center">
                <v-btn
                  :disabled="sendButton.disabled"
                  :loading="sendButton.loading"
                  class="text-transform-none"
                  color="primary"
                  rounded
                  width="180"
                  height="40"
                  @click="sendFaucet"
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
        </v-form>
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
      address: '',
      configForm: {},
      count: 0,
      button: {
        disabled: false,
        loading: false,
        text: 'SEND'
      },
      mosaicSelected: 'prx.xpx',
      isValidForm: false,
      sendingForm: false,
      showValidate: false,
      typeMessage: '',
      msjValidate: '',
      mosaicsList: []
    }
  },
  beforeMount () {
    this.configForm = {
      rules: {
        required: v => !!v || 'Address field is required',
        min: v =>
          (v && v.length >= 46) || 'The address field must be 46 characters.',
        max: v =>
          (v && v.length <= 46) || 'The address field must be 46 characters.'
      }
    }
    this.getConfig()
  },
  methods: {
    getConfig () {
      this.count = this.count + 1
      this.$apiService
        .getConfig()
        .then(x => {
          if (x.data && x.data.length > 0) {
            console.log('Configs:', x.data)
            this.mosaicsList = x.data || []
          }
        })
        .catch(e => {
          if (this.count < 5) this.getConfig()
        })
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
    validateAddress: function (e) {
      this.address = Utils.addressForm(this.address)
      // if (Utils.isValid(this.address)) {
      //   this.showValidate = false
      //   this.msjValidate = ''
      //   this.classValdiate = 'success'
      // } else {
      //   this.showValidate = true
      //   this.msjValidate = 'Address must have 40 characters'
      //   this.isB = true
      //   this.classValdiate = 'error'
      //   this.div1 = 'div-alert-text'
      //   this.div2 = ''
      //   this.alertclass = 'alert-war'
      // }
      return Utils.isValid(this.address)
    },
    sendFaucet () {
      if (Utils.isValid(this.address)) {
        this.showValidate = true
        this.typeMessage = 'warning'
        this.msjValidate = 'Sending....'
        // this.buttonValidate = false
        // this.alertclass = 'alert-war'
        // this.loaderclass = 'loader'
        // this.div1 = 'div-loader-text'
        // this.div2 = 'div-loader-animate'
        this.sendingForm = true
        this.$apiService
          .getMosaic(Utils.clean(this.address), this.mosaicSelected)
          .then(response => {
            this.sendingForm = false
            this.typeMessage = 'success'
            this.msjValidate = `${response.data}` || 'Success'
            if (this.$refs.form) this.$refs.form.reset()
            setTimeout(() => {
              this.showValidate = false
              this.msjValidate = ''
            }, 15000)
            console.log('response', response)
            /* this.msjValidate = ''
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
            ) */
          })
          .catch(error => {
            console.log('catchhhh error', error)
            if (this.$refs.form) this.$refs.form.reset()
            this.sendingForm = false
            this.typeMessage = 'error'
            // this.classValdiate = ''
            // this.buttonValidate = true
            // this.address = ''
            // let msj = null
            this.msjValidate = error.response.data.message
              ? error.response.data.message
              : error.response.data
            setTimeout(() => {
              this.showValidate = false
              this.msjValidate = ''
            }, 10000)
            // this.loaderclass = ''
            // this.div1 = 'div-alert-text'
            // this.div2 = ''
            // this.alertclass = 'alert-error'
            // this.showValidate = true
            // this.msjValidate = `${msj}`
            // setTimeout(
            //   function () {
            //     this.showValidate = false
            //   }.bind(this),
            //   10000
            // )
          })
      }
    }
  },
  computed: {
    // isDisabled: function () {
    //   return !this.buttonValidate
    // },
    sendButton () {
      const b = this.button
      b.disabled = !this.isValidForm || this.sendingForm
      b.loading = this.sendingForm
      return this.button
    },
    listMosaic () {
      const m = this.mosaicsList
      m.forEach(element => {
        element.max = this.amountFormatterSimple(element.maxQuantity)
      })
      console.log(m)
      return m
    }
  }
}
</script>

<style>
.v-messages__message {
  font-size: 16px !important;
  margin-top: 6px !important;
  margin-bottom: 6px !important;
  font-weight: 500 !important;
}

.cursor-p {
  cursor: pointer;
}
</style>
