<template>
  <section>
    <v-app dark>
      <Header/>
      <v-content>
        <v-container class="content" fluid>
          <slot></slot>
        </v-container>
      </v-content>
      <Footer/>
    </v-app>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Getter } from "nuxt-property-decorator";
import Header from "~/components/layouts/Header.vue";
import Footer from "~/components/layouts/Footer.vue";

@Component({
  components: {
    Header,
    Footer
  }
})
export default class extends Vue {}
</script>

<style lang="scss" scoped>
.content {
  margin: 0 auto;
  width: 85%;
  max-width: 800px;
}
</style>
