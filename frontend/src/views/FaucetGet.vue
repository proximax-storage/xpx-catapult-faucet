<template>
  <div class="fauce-container">
    <div class="container-fluid">
      <br>
      <h2 class="text-color-title text-size-30 text-center mt-1rem">
        <b>Sirius-chain testnet</b>
      </h2>
      <p
        class="text-color-black text-size-20 text-center"
      >Top-up your account balance to a maximum of {{xpxMax}} test-XPX every 24 hours.</p>
      <hr>
      <form id="fauceApp" @submit="sendFaucet" novalidate="true">
        <div class="container mt-1rem">
          <div class="input-icon-wrap">
            <span class="input-icon">
              <span class="fa fa-user">
                <img
                  class="icon-wallet"
                  src="@/assets/img/icon-wallet-name-red-16h-proximax-sirius-faucet.svg"
                >
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
            >
          </div>

          <div class="div-alert-war">
            <div v-show="showValidate" :class="alertclass">
              <div :class="div1">{{msjValidate}}</div>
              <div :class="div2">
                <b :class="loaderclass"></b>
              </div>
            </div>
          </div>
          <div class="button-class mt-1rem">
            <button :disabled="isDisabled" type="submit" class="btn">SEND</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>
<script>
import Utils from "@/services/Utils.js";
export default {
  data() {
    return {
      xpxMax: 0,
      cont: 0,
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
  created: function() {
    this.getMaxXpx();
  },
  methods: {
    getMaxXpx: function() {
      console.log(this.amountFormatterSimple(100000000));
      this.cont = this.cont + 1;
      console.log(this.cont);
      this.$apiService
        .get(`faucet/config`)
        .then(response => {
          this.xpxMax = this.amountFormatterSimple(response.data);
        })
        .catch(error => {
          this.xpxMax = 0;
          if (this.cont < 5) {
            this.getMaxXpx();
          }
        });
    },
    amountFormatterSimple: function(amount) {
      const amountDivisibility = Number(amount) / Math.pow(10, 6);
      return amountDivisibility.toLocaleString("en-us", {
        minimumFractionDigits: 0
      });
    },   
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
            this.div1 = "div-afauce-cardlert-text";
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
//########## containers class##########
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.fauce-container {
  display: flex;
  overflow: auto;
  flex-flow: row wrap;
  flex-direction: column;
  justify-content: center;
  margin: 15px 2px 50px 2px;
  align-items: center;
}
.container-fluid {
  padding-top: 2rem;
  margin: 1rem 1rem 1rem 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 90%;
}

//########## inputs class##########
.input-icon-wrap {
  width: 50%;
  border-radius: 20px;
  border: 1px solid silver;
  display: flex;
  flex-flow: row wrap;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  padding: 13px;
}
.input-with-icon {
  // font-size: 18px;
  border: none;
  text-align: center;
  flex: 1;
}
.input-icon,
.input-with-icon {
  padding: 2px;
  border-radius: 20px;
}
.icon-wallet {
  width: 1.5rem;
}

//########## text class##########
.div-alert-text {
  width: 100%;
  text-align: center;
}
.div-loader-text {
  width: 80%;
  text-align: center;
}
.text-color-title {
  //
    font-weight: bold;
  color: #df4c48;
}
.text-color-black {
   font-weight: bold;
  color: #000000;
}
.text-center {
  text-align: center;
}
.text-size-30{
 font-size: 30px;
}
.text-size-20{
 font-size: 20px;
}
.text-size-18{
 font-size: 18px;
}
//########## tasks class##########
input:focus,
textarea:focus,
select:focus {
  outline-offset: none !important;
  outline: none !important;
}

h2 {
  margin-bottom: 0.4rem;
}
form {
  width: 100%;
}
hr {
  border: 1px Solid silver;
  width: 100%;
}

//########## margin and padding class##########
.mt-1rem {
  margin-top: 1rem;
}
.mt-2rem {
  margin-top: 2rem;
}

//##########  alerts class##########
.div-alert-war {
  width: 50%;
  margin: 0.5rem 0rem 0rem 0rem;
  height: 2rem;
   font-size: 18px;
  flex-direction: column;
   display: block;
}
.alert-war {
  padding: 2px 2px;
   display: block;
  width: 100%;
  border-radius: 10px;
  border: 1px solid #f4a400;
  color: #000000;
  
}
.alert-success {
  padding: 2px 2px;
   display: block;
  width: 100%;
  border-radius: 10px;
  border: 1px solid #1eb3aa;
  color: #000000;
}
.alert-error {
  padding: 2px 2px;
   display: block;
  width: 100%;
  border-radius: 10px;
  border: 1px solid #ee6723;
  color: #000000;
}

//##########  button class##########
button:disabled,
button[disabled] {
  border: 1px solid #999999;
  background-color: #cccccc;
  color: #666666;
}
button {
  outline: none;
}
button::-moz-focus-inner {
  border: 0;
}
.button-class {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.btn:hover {
  opacity: 1;
  -webkit-transform: scale(1, 1);
  transform: scale(1, 1);
  background-color: rgb(228, 90, 85);
  box-shadow: 0 1px 16px 0 rgba(0, 0, 0, 0.4);
}
.btn:after {
  content: "";
  background: #df4c48;
  display: block;
  position: absolute;
  padding-top: 300%;
  padding-left: 350%;
  margin-left: -20px !important;
  margin-top: -120%;
  opacity: 0;
  transition: all 0.8s;
}
.btn:active:after {
  padding: 0;
  margin: 0;
  opacity: 1;
  transition: 0s;
}
.btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 22px;
  position: relative;
  background-color: #df4c48;
  border: none;
  margin-top: 15px;
  font-size: 18px;
  color: #ffffff;
  padding: 12px;
  width: 180px;
  text-align: center;
  -webkit-transition-duration: 0.4s; /* Safari */
  transition-duration: 0.4s;
  text-decoration: none;
  font-family: sans-serif;
  overflow: hidden;
  cursor: pointer;
  outline: none;
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
.div-loader-animate {
  width: 15%;
  text-align: right;
}
</style>


