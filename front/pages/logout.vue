<template>
  <section>
    <App>
      <v-layout justify-center>
        <v-btn @click="dialog = true">logout</v-btn>
      </v-layout>
      <v-dialog v-model="dialog" max-width="300">
        <v-card>
          <v-card-title>{{dialogTitle}}</v-card-title>
          <v-card-actions>
            <v-btn @click="disagree">no</v-btn>
            <v-btn @click="agree">yes</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </App>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Mutation, Action } from "nuxt-property-decorator";
import App from "~/components/layouts/App.vue";

@Component({
  components: {
    App
  }
})
export default class extends Vue {
  dialogTitle = "Do you want to logout?";

  @Action resetUser;
  dialog: boolean = false;

  public disagree() {
    this.dialog = false;
  }

  public async agree() {
    await this.resetUser();
    this.dialog = false;
    this.$router.push("/");
  }
}
</script>