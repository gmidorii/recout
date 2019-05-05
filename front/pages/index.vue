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
        <v-toolbar-title>Recout</v-toolbar-title>
      </v-toolbar>
      <v-content>
        <v-container class="content" fluid>
          <v-layout column justify-center class="part-content">
            <h3>Graph</h3>
            <a :href="graphDetailUrl" target="_blank">
              <div class="graph">
                <v-img :src="graphUrl"></v-img>
              </div>
              <div class="graph">
                <v-img :src="graphLineUrl"></v-img>
              </div>
            </a>
          </v-layout>
          <v-layout column justify-center class="part-content">
            <h3>Output</h3>
            <v-card>
              <v-form class="output">
                <v-textarea
                  class="output-text"
                  v-model="output"
                  :label="hint"
                  rows="2"
                  required
                  @keyup.ctrl.enter="submit"
                  auto-grow
                ></v-textarea>
                <v-btn block color="success" v-on:click="submit" :loading="loading">submit</v-btn>
                <v-snackbar :value="succeed" color="success" :timeout="timeout" top>Success</v-snackbar>
                <v-snackbar :value="failed" color="error" :timeout="timeout" top>Error</v-snackbar>
              </v-form>
            </v-card>
          </v-layout>
          <v-layout column class="part-content">
            <h3>Record</h3>
            <div v-for="output in pastOutputs" :key="output.created_at">
              <v-card class="past-output">
                <v-card-text>
                  <div class="recout-date">{{ toDateFormat(output.created_at) }}</div>
                  <div style="white-space: pre-line">{{ output.message }}</div>
                </v-card-text>
              </v-card>
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
import { Recout } from "../types/index";
import { parse, format } from "date-fns";

@Component({
  components: {}
})
export default class extends Vue {
  // TODO: fetch graph name and user name from recout api.
  graphUrl: string = `${process.env.pixelaUrl}/${process.env.graph}`;
  graphLineUrl: string = `${process.env.pixelaUrl}/${
    process.env.graph
  }?mode=line`;
  graphDetailUrl: string = `${process.env.pixelaUrl}/${process.env.graph}.html`;
  hint: string = "Write your output (Ctrl + Enter)";
  recoutUrl: string = `${process.env.recoutUrl}`;
  timeout: number = 3000;

  drawer: boolean = false;
  loading: boolean = false;
  succeed: boolean = false;
  failed: boolean = false;

  pastOutputs: Recout[] = [];
  output: string = "";

  async created() {
    this.pastOutputs = await this.loadOutput();
  }

  private async loadOutput(): Promise<Recout[]> {
    try {
      const res = await axios.get(`${this.recoutUrl}/recout`);
      return res.data;
    } catch (error) {
      console.log(error);
      return [];
    }
  }

  public async submit() {
    this.loading = true;
    try {
      await axios.post(`${this.recoutUrl}/recout`, {
        message: this.output
      });
    } catch (error) {
      console.log(error);
      this.failed = true;
      this.loading = false;
      return;
    }

    this.loading = false;
    this.pastOutputs.unshift({
      message: this.output,
      created_at: this.toDateFormat(new Date())
    });
    this.output = "";
    this.succeed = true;
  }

  public toDateFormat(
    date: string | Date,
    layout: string = "M/D HH:mm"
  ): string {
    const d = parse(date);
    return format(d, layout);
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

  .part-content {
    margin-top: 20px;

    h3 {
      margin: 0 auto;
    }
  }

  .output {
    margin: 0 auto;
    width: 80%;
  }

  .graph div {
    margin: 10px auto;
    width: 100%;
  }

  .past-output {
    margin: 0.5em 0 0.5em 0;
  }

  .recout-date {
    font-size: 0.8em;
  }
}
</style>
