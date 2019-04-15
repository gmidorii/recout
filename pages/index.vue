<template>
  <section>
    <h1 class="header">recout</h1>
    <v-app>
      <div class="content">
        <div class="graph">
          <img :src="graphUrl">
        </div>
        <form>
          <v-textarea v-model="output" required></v-textarea>
          <v-btn color="success" v-on:click="submit" :loading="loading">submit</v-btn>
          <v-snackbar :value="succeed" color="success" timeout="3000" top>success</v-snackbar>
        </form>
        <div>
          <div>{{ record }}</div>
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
  output: string = "";
  graphUrl: string = `${process.env.baseUrl}/${process.env.graph}`;
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

<style scoped>
.header {
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
}

.content {
  width: 80%;
  margin: 0 auto;
}
</style>
