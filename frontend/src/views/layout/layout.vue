<template>

    <el-container >
      <el-aside width="200px">
        <el-scrollbar>
          <p v-for="item in 200" :key="item" class="scrollbar-demo-item">{{ item }}</p>
        </el-scrollbar>

      </el-aside>

      <el-container>

        <el-header>
          Header
        </el-header>
        <el-main>
          Main
        </el-main>

      </el-container>
    </el-container>

</template>

<script setup lang="ts">

</script>


<style scoped>
.el-aside {
  height: 100vh;
  color: white;
  background: purple;
}
.el-header {
  position: relative;
  background-color: deeppink;
  color: white;
}
.el-main{
  position: relative;
  background-color: white;
  color: black;
}

</style>