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

</style>
