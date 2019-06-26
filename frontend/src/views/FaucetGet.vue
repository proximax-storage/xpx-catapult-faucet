<template>
  <div class="fauce-card">
    <div class="container-column">
      <div class="card">
        <form id="fauceApp" @submit="sendFaucet" novalidate="true">
          <div class="container">
            <img src="@/assets/img/logo-proximax-xpx-faucet.svg" alt="logo" style="width: 80%">
            <h2 class="text-color-title text-size-vm20 mt-2rem">
              <b>SIRIUS BC TESTNET</b>
            </h2>
            <p
              style="margin: 0;"
              class="text-color text-size-vm16"
            >Faucet Will only send if XPX is less than 100</p>
          </div>
          <div class="input-class mt-2rem">
            <input
              maxlength="46"
              v-bind:class="[classValdiate]"
              style="width: 100%;"
              type="text"
              v-model="address"
              placeholder="Enter your XPX address here..."
              @input="checkForm($event)"
            >
          </div>

          <div class="div-alert-war">
            <div v-show="showValidate" :class="alertclass">
              <div :class="div1">{{msjValidate}}</div>
              <div :class="div2">
                <b style="margin-top: 5px;" :class="loaderclass"></b>
              </div>

              <!-- <div>
                {{msjValidate}}
              </div>
               <div >
                 hola
              </div>-->

              <!-- <b
                class="text-color text-size-vm16-margin"
                style="margin-top: 0.51rem;"
              ></b>
              <b style="margin-top: 5px;" :class="loaderclass"></b>-->
            </div>
            <!-- <div :class="alertclass">
              <b style="margin-top: 5px;" :class="loaderclass"></b>
            </div>-->
          </div>

          <div class="button-class">
            <button :disabled="isDisabled" type="submit" class="btn">Send</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<script>
import Utils from "@/services/Utils.js";
export default {
  data() {
    return {
      alertclass: "",
      div1: "",
      div2: "",
      loaderclass: "loader",
      address: null,
      buttonValidate: true,
      classValdiate: "",
      showValidate: false,
      msjValidate: ""
    };
  },
  methods: {
    checkForm: function(e) {
      this.loaderclass = "";
      // address test  VARC5G-OWFIWG-7JK7JV-Y7DXIS-TQYOID-75ON3G-O22H
      this.address = Utils.addressForm(this.address);
      if (Utils.isValid(this.address)) {
        this.showValidate = false;
        this.msjValidate = "";
        this.classValdiate = "success";
      } else {
        this.showValidate = true;
        this.msjValidate = "Address must have 40 characters";
        this.isB = true;
        this.classValdiate = "error";
        this.div1 = "div-alert-text";
        this.div2 = "";
        this.alertclass = "alert-war";
      }
    },
    sendFaucet() {
      if (Utils.isValid(this.address)) {
        this.showValidate = true;
        this.alertclass = "alert-war";
        this.loaderclass = "loader";
        this.div1 = "div-loader-text";
        this.div2 = "div-loader-animate";
        this.msjValidate = "sending XPX    ";
        this.buttonValidate = false;
        this.$apiService
          .get(`faucet/GetXpx/${Utils.clean(this.address)}`)
          .then(response => {
            this.msjValidate = "";
            this.classValdiate = "";
            this.buttonValidate = true;
            this.address = "";
            this.loaderclass = "";
            this.div1 = "div-alert-text";
            this.div2 = "";
            this.alertclass = "alert-success";
            this.showValidate = true;
            this.msjValidate = `${response.data}`;
            setTimeout(
              function() {
                this.showValidate = false;
              }.bind(this),
              15000
            );
          })
          .catch(error => {
            this.classValdiate = "";
            this.buttonValidate = true;
            this.address = "";
            let msj = null;
            msj = error.response.data.message
              ? error.response.data.message
              : error.response.data;
            this.loaderclass = "";
            this.div1 = "div-alert-text";
            this.div2 = "";
            this.alertclass = "alert-error";
            this.showValidate = true;
            this.msjValidate = `${msj}`;
            setTimeout(
              function() {
                this.showValidate = false;
              }.bind(this),
              10000
            );
          });
      }
    }
  },
  computed: {
    isDisabled: function() {
      return !this.buttonValidate;
    }
  }
};
</script>

<style lang="scss">
h2 {
  margin-bottom: 0.4rem;
}
.mt-1rem {
  margin-top: 1rem;
}
.mt-2rem {
  margin-top: 2rem;
}
.text-size-vm26 {
  font-size: 26px;
}
.text-size-vm40 {
  font-size: 40px;
}
.text-color-title {
  // font-size: 30px;
  color: #0a8e9b;
}
.fauce-card {
  display: flex;
  overflow: auto;
  flex-flow: row wrap;
  justify-content: center;
  position: fixed;
  align-items: center;
  width: 100%;
  height: 100%;
}
.input-class {
  padding: 1px 43px 18px 0px;
  background-color: #ffffff;
  height: 3rem;
  border-radius: 6px; /* 6px rounded corners */
}
.card {
  margin: 1rem 1rem 1rem 1rem;
  height: 27rem;
  width: 519px;
}
.container {
  display: flex;
  flex-direction: column;
  padding: 2px 16px;
  align-items: center;
}
.div-alert-war {
  margin: 0.5rem 0rem 0rem 0rem;
  height: 2rem;
  flex-direction: row;
  display: flex;
}
.alert-war {
  padding: 2px 2px;
  display: flex;
  width: 100%;
  background-color: #f4a400;
  border-radius: 6px; /* 6px rounded corners */
  color: #ffffff;
}

.alert-success {
  padding: 2px 2px;
  display: flex;
  width: 100%;
  background-color: #1eb3aa;
  border-radius: 6px; /* 6px rounded corners */
  color: #ffffff;
}
.alert-error {
  padding: 2px 2px;
  display: flex;
  width: 100%;
  background-color: #ee6723;
  border-radius: 6px; /* 6px rounded corners */
  color: #ffffff;
}
.container-column {
  flex-direction: column;
}
.button-class {
  display: flex;
  flex-direction: column;

  align-items: center;
}
button:disabled,
button[disabled] {
  border: 1px solid #999999;
  background-color: #cccccc;
  color: #666666;
}

.btn:hover {
  opacity: 1;
  -webkit-transform: scale(1, 1);
  transform: scale(1, 1);
  background-color: #1eb3aa;
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.4);
}
.btn:after {
  content: "";
  background: #0a8e9b;
  display: block;
  position: absolute;
  padding-top: 300%;
  padding-left: 350%;
  margin-left: -20px !important;
  margin-top: -120%;
  opacity: 0;
  transition: all 0.8s;
}
button {
  outline: none;
}
.btn:active:after {
  padding: 0;
  margin: 0;
  opacity: 1;
  transition: 0s;
}
.btn {
  border-radius: 22px;
  position: relative;
  background-color: #0a8e9b;
  border: none;
  margin-top: 15px;
  font-size: 25px;
  color: #ffffff;
  padding: 12px;
  width: 280px;
  text-align: center;
  -webkit-transition-duration: 0.4s; /* Safari */
  transition-duration: 0.4s;
  text-decoration: none;
  font-family: sans-serif;
  overflow: hidden;
  cursor: pointer;
  outline: none;
}
button::-moz-focus-inner {
  border: 0;
}
mb-2rem {
  margin-bottom: 2rem;
}
input {
  margin: 25px 25px;
  width: 200px;
  display: block;
  text-align: center;
  border: none;
  color: #000000;
  font-size: 15px;
  border-bottom: 1px Solid #999999;
  background: -webkit-linear-gradient(
    top,
    rgba(255, 255, 255, 0) 100%,
    white 4%
  );
  font-weight: bold;
}

input:focus,
input:valid {
  box-shadow: none;
  outline: none;
  background-position: 0 0;
}
input:focus {
  font-size: 15px;
  visibility: visible !important;
}

.loader {
  border: 4px solid #f3f3f3;
  border-radius: 50%;
  border-top: 4px solid #f4a400;
  width: 8px;
  height: 8px;
  position: fixed;
  -webkit-animation: spin 2s linear infinite; /* Safari chrome */
  -moz-animation: spin 2s linear infinite; /* firefox */
  -o-animation: spin 2s linear infinite; /* opera */
  animation: spin 2s linear infinite; /* Safari */
}
/* Safari */
@-webkit-keyframes spin {
  0% {
    -webkit-transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
  }
}
@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.div-alert-text {
  width: 100%;
  text-align: center;
}

.div-loader-text {
  width: 80%;
  text-align: center;
}
.div-loader-animate {
  width: 15%;
  text-align: right;
}
</style>


