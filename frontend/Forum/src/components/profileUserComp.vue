<template>
  <div v-if="activeUser" class="modalWindow">
    <h3>Profile user</h3>
    <div>
      <form>
        <input type="text" v-model="newName" placeholder="Name" /> <br />
        <input type="text" v-model="newEmail" placeholder="email" /> <br />
        <input type="text" v-model="newPassword" placeholder="password (min 8 symbols)" /> <br />
      </form>
      <button @click="updateUser(activeUser.id)">send</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const props = defineProps({
    activeUser: {
      type: Object,
      required: false
    }
})


const newName = ref(props.activeUser?.name || '')
const newEmail = ref(props.activeUser?.email || '')
let newPassword = ""


async function updateUser(id) {
  const userResponce = JSON.parse(localStorage.getItem("activeUser"));
  const accessToken = userResponce.tokens.accessToken;
  const refreshToken = userResponce.tokens.refreshToken;

  const responce = await fetch(`http://localhost:7000/user/${id}/update`, {
    method: "PUT",
    credentials: "omit",
    headers: {
      Authorization: `Bearer ${accessToken}`,
      "Content-Type": "application/json",
      "Refresh-Token": refreshToken,
    },
    body: JSON.stringify({
      name: this.newName,
      email: this.newEmail,
      password: this.newPassword,
    })
  });

  const responceData = await responce.json();
	
  if (!responce.ok) {
    alert(responceData.error);
    return;
  } else {
    localStorage.setItem("activeUser", JSON.stringify(responceData))
    alert("data updated successfully")
    window.location.href = window.location.origin + "/user/" + responceData.name;
  }
}

</script>

<style scoped>
.modalWindow {
  position: absolute;
  z-index: 10;
  top: 35%;
  left: 50%;
  width: 20%;
  height: 20%;
  transform: translate(-50%, -50%);
  background-color: rgb(209, 208, 208);
  text-align: center;
  border: 3px solid black;
  box-shadow: 0 0 30px rgb(5, 5, 5);
}

.modalWindow * button {
  width: 25%;
  margin-top: 5px;
  box-sizing: border-box;
  border: 2px solid black;
  border-radius: 5px;
  background-color: white;
  transition: box-shadow 0.2s ease-in-out;
}

.modalWindow * button:hover {
  background-color: rgb(196, 192, 192);
  box-shadow: 0 0 15px rgb(25, 128, 16);
}

.modalWindow * input {
  margin: 5px;
  width: 80%;
  box-sizing: border-box;
  border-radius: 5px;
}

</style>