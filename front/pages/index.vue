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
            <h3>Continues</h3>
            <div class="continues">{{recoutContinues}}</div>
          </v-layout>
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
                <v-snackbar
                  :value="failed"
                  color="error"
                  :timeout="timeout"
                  top
                >Error {{ errMessage }}</v-snackbar>
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
import { RepositoryFactory } from "../repositories/RepositoryFactory";
import { recout } from "../repositories/RecoutRepository";
import { user } from "~/repositories/UserRepository";
const RecoutRepository: recout = RepositoryFactory.getRecout();
const UserRepository: user = RepositoryFactory.getUser();

@Component({
  components: {}
})
export default class extends Vue {
  // TODO: fetch graph name and user name from recout api.
  hint: string = "Write your output (Ctrl + Enter)";
  timeout: number = 3000;

  drawer: boolean = false;
  loading: boolean = false;
  succeed: boolean = false;
  failed: boolean = false;

  graphUrl: string = "";
  graphLineUrl: string = "";
  graphDetailUrl: string = "";
  pastOutputs: Recout[] = [];
  output: string = "";
  recoutContinues: number = null;
  errMessage: string = "";

  async created() {
    this.pastOutputs = await this.loadOutput();
    this.recoutContinues = await this.loadContinues();
    await this.loadGraph();
  }

  private async loadOutput(): Promise<Recout[]> {
    try {
      const { data } = await RecoutRepository.get();
      return data;
    } catch (error) {
      console.log(error);
      return [];
    }
  }

  private async loadContinues(): Promise<number> {
    try {
      const { data } = await RecoutRepository.getContinues();
      return data.count;
    } catch (error) {
      return 0;
    }
  }

  private async loadGraph() {
    try {
      const { data } = await UserRepository.get();
      this.graphUrl = `${process.env.pixelaUrl}/users/${
        data.account_id
      }/graphs/${data.pixela_graph}`;
      this.graphLineUrl = `${this.graphUrl}?mode=line`;
      this.graphDetailUrl = `${this.graphUrl}.html`;
    } catch (error) {
      this.errMessage = "failed load graph.";
      this.failed = true;
    }
  }

  public async submit() {
    this.loading = true;
    try {
      await RecoutRepository.post(this.output);
    } catch (error) {
      console.log(error);
      this.errMessage = "faild post recout.";
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
  width: 90%;

  .continues {
    margin: 0 auto;
    font-size: 3em;
  }

  .part-content {
    margin-top: 20px;

    h3 {
      margin: 0 auto;
    }
  }

  .output {
    margin: 0 auto;
    width: 85%;
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
