<template>
  <section>
    <Header/>
    <App>
      <v-layout column justify-center class="title">
        <v-alert :value="true" color="warning" outline>Warning: this app is version β</v-alert>
        <h2>{{ title }}</h2>
        <div class="graph">
          <v-img :src="indexGraphUrl"></v-img>
        </div>
        <div>{{ subTitle }}</div>
      </v-layout>

      <Login></Login>
    </App>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";
import Header from "~/components/layouts/Header.vue";
import App from "~/components/layouts/App.vue";
import Login from "~/components/ui/Login.vue";
import firebase from "firebase";

@Component({
  components: {
    Header,
    App,
    Login
  }
})
export default class extends Vue {
  title = "Record + Output = Recout";
  subTitle = "This is a personal output visualization tool.";
  indexGraphUrl = `${process.env.pixelaUrl}/users/gmidorii/graphs/recout`;

  mounted() {}
}
</script>

<style lang="scss" scoped>
.title {
  h2 {
    font-size: 5vw;
    margin: 0 auto;
  }
  div {
    font-size: 4vw;
    margin: 2em auto;
  }
  .graph {
    margin: 0px auto;
    width: 100%;
  }
}
</style>
