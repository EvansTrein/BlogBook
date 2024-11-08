import { defineStore } from "pinia";
import { ref } from "vue";

const activeUserLocalSt =
  localStorage.length > 0
    ? JSON.parse(localStorage.getItem("activeUser"))
    : null;

const activeSessionData =
  sessionStorage.hasOwnProperty("sessionUserData")
    ? JSON.parse(sessionStorage.getItem("sessionUserData"))
    : null;

export const useAppStore = defineStore("AppStore", () => {
  const activeUserName =
    activeUserLocalSt !== null ? ref(activeUserLocalSt.user.name) : "";
  const activeUserEmail =
    activeUserLocalSt !== null ? ref(activeUserLocalSt.user.email) : "";
  const admin = ref(false);
//   const sessionDataAllUsers =
//     activeUserLocalSt !== null ? ref(activeSessionData.users) : "";

  return { admin, activeUserName, activeUserEmail };
});
