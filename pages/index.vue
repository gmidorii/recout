<template>
  <section>
    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
    <link rel="manifest" href="/site.webmanifest">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5">
    <meta name="msapplication-TileColor" content="#da532c">
    <meta name="theme-color" content="#ffffff">

    <v-app>
      <div class="content">
        <h1 class="header">recout</h1>
        <div>
          <div class="graph">
            <a :href="graphDetailUrl" target="_blank">
              <img :src="graphLineUrl">
              <img :src="graphUrl">
            </a>
          </div>
          <form class="output">
            <v-textarea class="output-text" v-model="output" :placeholder="hint" rows="2" required></v-textarea>
            <v-btn color="success" v-on:click="submit" :loading="loading">submit</v-btn>
            <v-snackbar :value="succeed" color="success" timeout="3000" top>success</v-snackbar>
          </form>
          <div>
            <div>{{ record }}</div>
          </div>
        </div>
      </div>
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
  graphUrl: string = `${process.env.baseUrl}/${process.env.graph}`;
  graphLineUrl: string = `${process.env.baseUrl}/${
    process.env.graph
  }?mode=line`;
  graphDetailUrl: string = `${process.env.baseUrl}/${process.env.graph}.html`;
  hint: string = "write today your output....";

  output: string = "";
  loading: boolean = false;
  succeed: boolean = false;
  record: string = "";

  public async submit() {
    this.loading = true;

    const instance = axios.create({
      baseURL: `${process.env.baseUrl}/${process.env.graph}`,
      timeout: 5000,
      headers: {
        "X-USER-TOKEN": process.env.token
      }
    });

    try {
      const res = await instance.put("/increment");
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
  width: 80%;
  margin: 0 auto;

  .output {
    width: 80%;
    margin: 10% auto;
  }
}
</style>
