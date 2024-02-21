<template>
  <div class="app">
    <RouterView/>
  </div>
</template>


<script lang="ts">
export default {
  name:'App'
}
</script>
<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<style scoped>
.app{
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100vh;
}
</style>
