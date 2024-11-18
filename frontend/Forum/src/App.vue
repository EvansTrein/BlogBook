<template>
  <div class="conteinerApp">
    <div class="header">
      <div>
        <h1>BlogBook</h1>
      </div>
    </div>
    <div class="bar">
      <h4>icon profile</h4>
      <div v-if="!activeUser">
        <form>
          <input type="text" v-model="name" placeholder="Name" /> <br />
          <input type="text" v-model="email" placeholder="email" /> <br />
          <input type="text" v-model="password" placeholder="password (min 8 symbols)"/> <br />
        </form>
        <button @click="signUp">sign up</button>
        <button @click="signIn">sign in</button>
    </div>
    <button v-if="activeUser" @click="logout">logout</button> <br />
    <button v-if="sessionTab === 'adminSession'" @click="getAllUsers">view all users</button>
    <p class="author"><a target="_blank" href="https://github.com/EvansTrein">powered by Evans Trein's</a></p>
    </div>
    <div class="main">
      <tapeComp v-if="!dataAllUsers.length != 0" />
      <allUsersComp v-bind:data-all-users="dataAllUsers" @updateDataAllUsers="onUpdateDataAllUsers" />
    </div>
    <div class="footer">
      <div>
        <p>footer</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import tapeComp from "./components/tapeComp.vue";
import allUsersComp from "./components/allUsersComp.vue"
import { ref, defineProps } from "vue"
import { signUp, signIn, logout, useGetAllUsers } from "./JSFuncsServer";

let name = "";
let email = "";
let password = "";

const sessionTab = ref(localStorage.getItem("sessionTab"))
const activeUser = ref(JSON.parse(localStorage.getItem("activeUser")))

const { getAllUsers, dataAllUsers } = useGetAllUsers()

defineProps({
    dataAllUsers: {
    type: Array,
    required: false
    }
})

defineEmits(['updateDataAllUsers']); // необязательно писать, тут мы указываем какие пользовательские собыития мы слушаем из доч. комп.

function onUpdateDataAllUsers(newData) {
  dataAllUsers.value = newData;
}

</script>

<style scoped>
</style>
