<template>
  <div>
    <div id="firebaseui-auth-container"></div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";
import firebase from "firebase";

@Component({})
export default class extends Vue {
  mounted() {
    const firebaseConfig = {
      apiKey: "AIzaSyAHDcSOiHp6oQ_NVy3WXSwWX3dix3oWLTQ",
      authDomain: "gmidorii-webapp.firebaseapp.com",
      databaseURL: "https://gmidorii-webapp.firebaseio.com",
      projectId: "gmidorii-webapp",
      storageBucket: "gmidorii-webapp.appspot.com",
      messagingSenderId: "773134415366",
      appId: "1:773134415366:web:5b0d2bb88d972113"
    };
    firebase.initializeApp(firebaseConfig);

    if (process.client) {
      const firebaseui = require("firebaseui");
      const config = {
        signInSuccessUrl: "/recout",
        signInOptions: [
          firebase.auth.EmailAuthProvider.PROVIDER_ID,
          firebase.auth.GoogleAuthProvider.PROVIDER_ID
        ]
      };
      const ui = new firebaseui.auth.AuthUI(firebase.auth());
      ui.start("#firebaseui-auth-container", config);
    }
  }
}
</script>

