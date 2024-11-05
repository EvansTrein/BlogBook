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
  console.log(data)
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
    console.log(data.tokens)
  }
