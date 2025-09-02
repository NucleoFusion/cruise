<script setup>
import DefaultTheme from 'vitepress/theme'

const { Layout } = DefaultTheme
</script>

<template>
  <Layout>
  <template #doc-before>
    <div class="global-banner">
      Enjoying cruise? Consider giving a star on <a href="https://github.com/NucleoFusion/cruise/tree/main">Github</a>! It helps a lot!
    </div>
    </template>
  </Layout>
</template>

<style>
.global-banner {
  background-color: #9399b2; /* subtle highlight */
  color: #1e1e2e;
  text-align: center;
  padding: 0.5rem;
  font-weight: bold;
  border-radius: 8px;
  margin-bottom: 2vh;
}

.global-banner > a {
  color: #11111b;
  text-decoration: underline;
}
</style>

