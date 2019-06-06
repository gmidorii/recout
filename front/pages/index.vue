<template>
  <section>
    <Header/>
    <App>
      <v-layout column justify-center class="title">
        <h2>{{ title }}</h2>
        <div>{{ subTitle }}</div>
      </v-layout>
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
export default class extends Vue {
  title = "Record + Output = Recout";
  subTitle = "This is a personal output visualization tool.";
}
</script>

<style lang="scss" scoped>
.title {
  margin-top: 10%;
  h2 {
    margin: 0 auto;
  }
  div {
    margin: 10% auto;
  }
}
</style>
