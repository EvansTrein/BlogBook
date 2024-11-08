<template>
  <div class="conteinerApp">
    <div class="header">
      <div>
        <h1>BlogBook</h1>
      </div>
    </div>
    <div class="bar">
        <h4 v-if="appStore.activeUserName">icon profile</h4>
        <h3>{{ appStore.activeUserName }}</h3>
        <h3>{{ appStore.activeUserEmail }}</h3>
      <div v-if="!appStore.activeUserName">
        <form>
          <input type="text" v-model="name" placeholder="Name" /> <br />
          <input type="text" v-model="email" placeholder="email" /> <br />
          <input
            type="text"
            v-model="password"
            placeholder="password (min 8 symbols)"
          />
          <br />
        </form>
        <button @click="signUp">sign up</button>
        <button @click="signIn">sign in</button>
    </div>
    <button v-if="appStore.activeUserName" @click="logout">logout</button> <br>
    <button v-if="appStore.activeUserName === 'admin'" @click="getAllUsers">view all users</button>
    </div>
    <div class="main">
        <div v-for="item in appStore.sessionDataAllUsers" class="allUsers">
            ID: {{ item.ID }} <br> 
            Email: {{ item.Email }} <br>
            Name: {{ item.Name }} <br>
            <button>Delete</button>
        </div>
      <div v-show="false">
        <h4>User:</h4>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti esse
          ratione aliquam tempora ex maiores dicta. Aliquid vitae a commodi
          minima, officiis vero aut, distinctio laborum perferendis explicabo
          voluptates quas repellendus soluta veritatis magni. Quo sit sunt
          numquam, porro ipsam similique inventore? Itaque perspiciatis
          veritatis officiis magni porro dicta ea nostrum eveniet sit est fugit
          nulla voluptatum, sed reiciendis. Sequi dolore aut iste praesentium.
          Velit vero nemo, sequi totam repudiandae rerum in libero facilis dicta
          maxime recusandae corporis provident, fuga iure, distinctio voluptas
          quibusdam ex modi rem consequuntur ad quisquam animi voluptatibus
          unde! Porro culpa officiis eligendi dolorem error illum?
        </p>
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
import { useAppStore } from "./stores/AppStore"
import { signUp, signIn, logout, getAllUsers } from "./JSFuncsServer";

const appStore = useAppStore();

let name = "";
let email = "";
let password = "";

</script>

<style scoped></style>
