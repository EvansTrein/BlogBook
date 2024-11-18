import { ref } from 'vue'

if (!localStorage.getItem("sessionTab")) {
    localStorage.setItem("sessionTab", "default")
  }

export async function signUp() {
  const responce = await fetch("http://localhost:7000/user/registration", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      name: this.name,
      email: this.email,
      password: this.password,
    }),
  });

  const data = await responce.json();

  if (responce.ok) {
    alert("User is registered, you can now log in")
  } else {
    alert(data.error)
  }
}

export async function signIn() {
  const responce = await fetch("http://localhost:7000/user/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      name: this.name,
      email: this.email,
      password: this.password,
    }),
  });

  const data = await responce.json();
  console.log(data);

  if (responce.ok) {
    localStorage.setItem("activeUser", JSON.stringify(data));
    if (data.user.name === 'admin') {
        localStorage.setItem("sessionTab", "adminSession")
    } else {
        localStorage.setItem("sessionTab", "userSession")
    }
    window.location.href = window.location.origin + "/user/" + data.user.name;
    // window.location.reload();
  } else {
    alert(data.error);
  }
}

export async function logout() {
  localStorage.clear();
  sessionStorage.clear();
  window.location.href = window.location.origin;
  //   window.location.reload();
}

export function useGetAllUsers() {
    const dataAllUsers = ref([])

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

    return {
        dataAllUsers,
        getAllUsers
    }
}