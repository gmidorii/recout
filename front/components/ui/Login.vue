<template>
  <div>
    <div id="firebaseui-auth-container"></div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";
import firebase from "firebase";
import "firebase/auth";
import { State, Mutation } from "vuex-class";
import { providers, auth } from "~/plugins/firebaseinit";

@Component({})
export default class extends Vue {
  mounted() {
    if (process.client) {
      const firebaseui = require("firebaseui");
      const config = {
        signInSuccessUrl: "/recout",
        signInOptions: providers,
        signInFlow: "popup",
        callbacks: {
          signInSuccessWithAuthResult: function(authResult, redirectUrl) {
            return true;
          }
        }
      };
      const ui =
        firebaseui.auth.AuthUI.getInstance() ||
        new firebaseui.auth.AuthUI(auth);
      ui.start("#firebaseui-auth-container", config);
    }
  }
}
</script>

