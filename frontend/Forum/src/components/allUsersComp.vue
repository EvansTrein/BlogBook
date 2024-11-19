<template>
    <div>
      <div v-for="item in dataAllUsers" class="allUsers">
        ID: {{ item.ID }} <br />
        Email: {{ item.Email }} <br />
        Name: {{ item.Name }} <br />
        <button @click="deleteUser(item.ID)">Delete</button>
      </div>
    </div>
</template>

<script setup>
import { defineProps } from "vue";

const emit = defineEmits(["updateDataAllUsers"]);

const props = defineProps({
  dataAllUsers: {
    type: Array,
    required: false,
  },
});

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
    return;
  } else {
    const filteredUsers = props.dataAllUsers.filter((user) => user.ID !== id);
    emit("updateDataAllUsers", filteredUsers);
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
