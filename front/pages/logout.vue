<template>
  <section>
    <Header/>
    <App>
      <button v-on:click="logout">logout</button>
    </App>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Mutation, Action } from "nuxt-property-decorator";
import Header from "~/components/layouts/Header.vue";
import App from "~/components/layouts/App.vue";

@Component({
  components: {
    Header,
    App
  }
})
export default class extends Vue {
  @Action resetUser;

  public logout() {
    this.resetUser();
  }
}
</script>