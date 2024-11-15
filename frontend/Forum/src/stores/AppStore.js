import { defineStore } from "pinia";
import { ref, reactive } from "vue";

localStorage.setItem("sessionTab", "default")

export const useAppStore = defineStore("AppStore", () => {
    const sessionTab = ref(localStorage.getItem("sessionTab"))
    const activeUser = reactive(JSON.parse(localStorage.getItem("activeUser")))
    
  return { sessionTab, activeUser };
});


