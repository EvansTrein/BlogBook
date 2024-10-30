export async function signUp() {
    try {
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
      if (responce.ok) {
        const data = await responce.json();
        console.log(data);
      } else {
        console.error("Ошибка json с сервера:", responce.status);
      }
    } catch (error) {
      console.error("Ошибка:", error);
    }
}
