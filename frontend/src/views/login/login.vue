<template>
  <div class="login">
    <el-row>
      <el-col :span="15" :xs="0"></el-col>
      <el-col :span="9" :xs="24">
        <el-form class="login_form">
          <h1 class="login_title">欢迎使用XXX系统！</h1>
          <el-form-item>
            <el-input class="login_input" :prefix-icon='User' v-model="loginFrom.username"></el-input>
          </el-form-item>
          <el-form-item/>
          <el-form-item>
            <el-input class="login_input" :prefix-icon='Lock' type="password" v-model="loginFrom.password" show-password></el-input>
          </el-form-item>
          <el-form-item/>
          <el-form-item>
            <el-button class="login_button" type="primary" size="default" @click="handleLogin()">登陆</el-button>
          </el-form-item>
        </el-form>

      </el-col>
    </el-row>

  </div>
</template>

<script setup lang="ts">
import {User,Lock} from '@element-plus/icons-vue'
import { reactive} from "vue";

let loginFrom = reactive({username:'',password:''})

</script>

<style scoped>

.login{
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100vh;
  background: url('@/../public/login-img.jpg') no-repeat;
  background-size: cover;
}
.login_title{
  color: hotpink;
}
.login_form{
  position: relative;
  width: 70%;
  top: 25vh;
  backdrop-filter: blur(15px);
  padding: 40px;
  border: 6px solid hotpink;
}
.login_button{
  width: 100%;
  height: 130%;
  background-color: hotpink;
  border-color: hotpink;
  color: white; /* 文字颜色设置为白色，以便更好地对比 */
}
.login_input{
  height: 130%;
  font-size: 130%;
}


</style>