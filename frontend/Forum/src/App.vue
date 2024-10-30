<template>
  <div class="conteinerApp">
    <div class="header">
      <div>
        <h1>BlogBook</h1>
      </div>
    </div>
    <div class="bar">
      <div>
        <h4>icon profile</h4>
        <h3>{{ statusForm }}</h3>
        <form>
          <input type="text" v-model="name" placeholder="Name" /> <br />
          <input type="text" v-model="email" placeholder="email" /> <br />
          <input type="text" v-model="password" placeholder="password" /> <br />
          <button @click="signUp">sign up</button>
          <button type="submit" onclick="return false">sign in</button>
          <button type="submit" onclick="return false">logout</button>
        </form>
      </div>
    </div>
    <div class="tape">
      <div>
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

<script>
import { signUp } from './assets/JSFuncsServer';

export default {
  data() {
    return {
      name: "",
      email: "",
      password: "",
      statusForm: localStorage.getItem("statusForm") || "",
    };
  },
  methods: {
    signUp
  },
};
</script>

<style scoped></style>
