<template>
  <div v-if="sessionDataAllUsers" >
    <!-- <div v-for="(item, index) in sessionDataAllUsers.users" :key="item.ID" class="allUsers">
    ID: {{ item.ID }} <br />
    Email: {{ item.Email }} <br />
    Name: {{ item.Name }} <br />
    <button>Delete</button>
    </div> -->
  </div>
</template>

<script setup>
import { ref } from "vue"

const sessionDataAllUsers = ref(JSON.parse(sessionStorage.getItem("sessionAllUsersData")))
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
