<template>
  <section>
    <h1 class="header">recout</h1>
    <v-app>
      <div class="content">
        <div class="graph">
          <img src="https://pixe.la/v1/users/gmidorii/graphs/dev-recout">
        </div>
        <form>
          <v-text-field v-model="output" required></v-text-field>
          <v-btn color="success" v-on:click="submit">submit</v-btn>
        </form>
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

  public async submit() {
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
  }
}
</script>

<style scoped>
.header {
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
}

.cards {
  display: flex;
  flex-wrap: wrap;
}

.content {
  width: 80%;
  margin: 0 auto;
}
</style>
