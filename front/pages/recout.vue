<template>
  <section>
    <App>
      <v-layout column justify-center class="part-content">
        <h3>Continues</h3>
        <div class="continues">{{recoutContinues}}</div>
      </v-layout>
      <v-layout column justify-center class="part-content">
        <h3>Graph</h3>
        <a :href="getPixelaGraphDetailUrl" target="_blank">
          <div class="graph">
            <v-img :src="getPixelaGraphUrl"></v-img>
          </div>
          <div class="graph">
            <v-img :src="getPixelaGraphLineUrl"></v-img>
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
            <v-btn block color="primary" v-on:click="submit" :loading="loading">submit</v-btn>
            <v-snackbar :value="succeed" color="success" :timeout="timeout" top>Success</v-snackbar>
            <v-snackbar :value="failed" color="error" :timeout="timeout" top>Error {{ errMessage }}</v-snackbar>
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
    </App>
  </section>
</template>

<script lang="ts">
import { Component, Vue, Action } from "nuxt-property-decorator";
import { State, Getter } from "vuex-class";
import App from "~/components/layouts/App.vue";
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
  components: {
    App
  }
})
export default class extends Vue {
  @Getter userId;
  @Getter pixelaGraphUrl;

  hint: string = "Write your output (Ctrl + Enter)";
  timeout: number = 3000;

  loading: boolean = false;
  succeed: boolean = false;
  failed: boolean = false;

  pastOutputs: Recout[] = [];
  output: string = "";
  recoutContinues: number = null;
  errMessage: string = "";

  async created() {
    this.pastOutputs = await this.loadOutput();
    this.recoutContinues = await this.loadContinues();
  }

  get getPixelaGraphUrl() {
    return this.pixelaGraphUrl;
  }

  get getPixelaGraphLineUrl() {
    return `${this.getPixelaGraphUrl}?mode=line`;
  }

  get getPixelaGraphDetailUrl() {
    return `${this.getPixelaGraphUrl}.html`;
  }

  private async loadOutput(): Promise<Recout[]> {
    try {
      const { data } = await RecoutRepository.get(this.userId);
      return data;
    } catch (error) {
      console.log(error);
      return [];
    }
  }

  private async loadContinues(): Promise<number> {
    try {
      const { data } = await RecoutRepository.getContinues(this.userId);
      return data.count;
    } catch (error) {
      return 0;
    }
  }

  public async submit() {
    this.loading = true;
    try {
      await RecoutRepository.post(this.userId, this.output);
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
</style>
