<template>
  <section>
    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
    <link rel="manifest" href="/site.webmanifest">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5">
    <meta name="msapplication-TileColor" content="#da532c">
    <meta name="theme-color" content="#ffffff">

    <v-app dark>
      <v-navigation-drawer v-model="drawer" cliped fixed app></v-navigation-drawer>
      <v-toolbar app fixed cliped-left>
        <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        <v-toolbar-title>recout</v-toolbar-title>
      </v-toolbar>
      <v-content>
        <v-container class="content" fluid>
          <v-layout column justify-center>
            <a :href="graphDetailUrl" target="_blank">
              <div class="graph">
                <v-img :src="graphLineUrl"></v-img>
              </div>
              <div class="graph">
                <v-img :src="graphUrl"></v-img>
              </div>
            </a>
          </v-layout>
          <v-layout column justify-center>
            <v-form class="output">
              <v-textarea
                class="output-text"
                v-model="output"
                :placeholder="hint"
                rows="2"
                required
              ></v-textarea>
              <v-btn color="success" v-on:click="submit" :loading="loading">submit</v-btn>
              <v-snackbar :value="succeed" color="success" timeout="3000" top>success</v-snackbar>
            </v-form>
            <div>
              <div>{{ record }}</div>
            </div>
          </v-layout>
        </v-container>
      </v-content>
    </v-app>
  </section>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator";
import { State } from "vuex-class";
import axios from "axios";
import { basename } from "path";

@Component({
  components: {}
})
export default class extends Vue {
  graphUrl: string = `${process.env.pixelaUrl}/${process.env.graph}`;
  graphLineUrl: string = `${process.env.pixelaUrl}/${
    process.env.graph
  }?mode=line`;
  graphDetailUrl: string = `${process.env.pixelaUrl}/${process.env.graph}.html`;
  hint: string = "write today your output....";
  recoutUrl: string = `${process.env.recoutUrl}`;

  drawer: boolean = false;
  output: string = "";
  loading: boolean = false;
  succeed: boolean = false;
  record: string = "";

  public async submit() {
    this.loading = true;

    const instance = axios.create({
      baseURL: this.graphUrl,
      timeout: 5000,
      headers: {
        "X-USER-TOKEN": process.env.token
      }
    });

    try {
      const res = await instance.put("/increment");
      await axios.post(`${this.recoutUrl}/recout`, {
        message: this.output
      });
    } catch (error) {
      console.log(error);
    }

    this.loading = false;
    this.record = this.output;
    this.output = "";
    this.succeed = true;
  }
}
</script>

<style scoped lang="scss">
.header {
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  font-size: 32px;
  margin-left: 10px;
}

.content {
  margin: 0 auto;
  width: 80%;

  .graph div {
    margin: 10px auto;
    width: 100%;
  }
}
</style>
