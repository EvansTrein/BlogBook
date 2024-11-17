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
      <tapeComp v-if="!dataAllUsers"></tapeComp>
      <!-- <allUsersComp></allUsersComp> -->
      <div v-if="dataAllUsers">
        <div>
          <div v-for="item in dataAllUsers" class="allUsers">
            ID: {{ item.ID }} <br />
            Email: {{ item.Email }} <br />
            Name: {{ item.Name }} <br />
            <button @click="deleteUser(item.ID)">Delete</button>
          </div>
        </div>
      </div>
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
// import allUsersComp from "./components/allUsersComp.vue"
import { ref } from "vue"
import { signUp, signIn, logout } from "./JSFuncsServer";

let name = "";
let email = "";
let password = "";

const dataAllUsers = ref(null)

const sessionTab = ref(localStorage.getItem("sessionTab"))
const activeUser = ref(JSON.parse(localStorage.getItem("activeUser")))

async function getAllUsers() {
  const activeUser = JSON.parse(localStorage.getItem("activeUser"));
  const accessToken = activeUser.tokens.accessToken;
  // const refreshToken = activeUser.tokens.refreshToken;

  const responce = await fetch("http://localhost:7000/users", {
    method: "GET",
    credentials: "omit",
    headers: {
      Authorization: `Bearer ${accessToken}`,
      // "Refresh-Token": refreshToken,
    },
  });

  const responceData = await responce.json();

  if (!responce.ok) {
    alert("You need to log into your profile");
    logout();
    return
  };

  dataAllUsers.value = responceData.users
}

async function deleteUser(id) {
  const activeUser = JSON.parse(localStorage.getItem("activeUser"));
  const accessToken = activeUser.tokens.accessToken;

  const responce = await fetch("http://localhost:7000/user/delete", {
    method: "DELETE",
    credentials: "omit",
    headers: {
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({ userId: id }),
  });

  const responceData = await responce.json();
  if (!responce.ok) {
    alert(responceData.error);
    return
  } else {
    dataAllUsers.value = dataAllUsers.value.filter(user => user.ID !== id);
  }

}

</script>

<style scoped>
.allUsers {
    padding: 10px;
    margin: 10px;
    border: 2px solid black;
    border-radius: 15px;
    width: 290px;
    text-align: center;
    display: inline-block;
}
</style>
