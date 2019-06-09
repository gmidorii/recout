<template>
  <section>
    <Header/>
    <App>
      <Login></Login>
    </App>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";
import Header from "~/components/layouts/Header.vue";
import App from "~/components/layouts/App.vue";

@Component({
  components: {
    Header,
    App
  }
})
export default class extends Vue {}
</script>