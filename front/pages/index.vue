<template>
  <section>
    <App>
      <v-layout column justify-center class="title">
        <v-alert :value="true" color="warning">{{warningMessage}}</v-alert>
        <h2>{{ title }}</h2>
        <div class="subtitle">{{ subTitle }}</div>
      </v-layout>

      <Login></Login>

      <v-layout class="content" column justify-center>
        <div class="section">
          <div class="label">
            {{recordGraphMessage}}
            (with
            <a href="https://pixe.la/" target="_blank">pixela</a>)
          </div>
          <div class="graph">
            <v-img :src="indexGraphUrl"></v-img>
          </div>
        </div>
      </v-layout>
    </App>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "nuxt-property-decorator";
import App from "~/components/layouts/App.vue";
import Login from "~/components/ui/Login.vue";
import firebase from "firebase";

@Component({
  components: {
    App,
    Login
  }
})
export default class extends Vue {
  title = "Record + Output = Recout";
  subTitle = "This is a personal output visualization tool.";
  indexGraphUrl = `${process.env.pixelaUrl}/users/gmidorii/graphs/recout`;
  warningMessage = "Warning: this app is version β";
  recordGraphMessage = "Record Graph";

  mounted() {}
}
</script>

<style lang="scss" scoped>
.title {
  h2 {
    font-size: 36px;
    margin: 0 auto;
    text-align: center;
  }
  div {
    margin: 2em auto;
  }
  .subtitle {
    font-size: 16px;
    color: #afaaaa;
    text-align: center;
  }
}

.content {
  margin: 4em 0 4em 0;
  .section {
    .label {
      margin: 1em 0 0.5em 0;
      font-size: 20px;
      text-align: center;
    }
  }
}

.graph {
  margin: 0px auto;
  width: 100%;
}
</style>
