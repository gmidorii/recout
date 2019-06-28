<template>
  <section>
    <App>
      <Login></Login>
    </App>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";
import Login from "~/components/ui/Login.vue";
import Header from "~/components/layouts/Header.vue";
import App from "~/components/layouts/App.vue";

@Component({
  components: {
    App,
    Login
  }
})
export default class extends Vue {}
</script>

