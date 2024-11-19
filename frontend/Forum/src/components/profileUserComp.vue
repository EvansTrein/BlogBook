<template>
 <div v-if="activeUser" class="pofrileUser">
  <hr>
  <h3>Profile user</h3>
  {{ activeUser.user }}
  <hr>
  <button @click="showChangeForm = !showChangeForm">change</button>
  <div v-if="showChangeForm">
    <form>
      <input type="text" v-model="newName" placeholder="Name" /> <br />
      <input type="text" v-model="newEmail" placeholder="email" /> <br />
      <input type="text" v-model="newPassword" placeholder="password (min 8 symbols)"/> <br />
    </form>
    <button @click="updateUser(activeUser.user.id)">send</button>
  </div>
 </div>
</template>

<script setup>
import { ref } from "vue";

const emit = defineEmits(["updateUserData"])

const props = defineProps({
    activeUser: {
      type: Object,
      required: false
    }
})

const showChangeForm = ref(false)
let newName = ""
let newEmail = ""
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
  console.log(responceData)
  console.log(responce)
  // if (!responce.ok) {
  //   alert(responceData.error);
  //   return;
  // } else {
  //   // тут нужно получить новые данные для Vue
  //   emit('updateUserData', newName, newEmail)
  // }
}

</script>

<style scoped></style>