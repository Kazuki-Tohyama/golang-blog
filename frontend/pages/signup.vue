<template>
  <div class="container">
    <div>
      <logo />
      <h1 class="title">
        frontend
      </h1>
      <h2 class="subtitle">
        Signup
      </h2>
      <form @submit.prevent="signup">
        <p class="error" v-if="error">{{ error }}</p>
        <input
          type="text"
          v-model="username"
          placeholder="username"
          name="username"
        />
        <input
          type="text"
          v-model="password"
          placeholder="password"
          name="password"
        />
        <div><button type="submit" value="signup" /></div>
      </form>
      <nuxt-link to="/">back</nuxt-link>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import Logo from "~/components/Logo.vue";

@Component({
  components: {
    Logo
  }
});

@data({
	return {
		username: '',
		password: '',
		error: null
	}
});

@methods({
	async signup() {
		try {
			await this.$store.dispatch("signup", {
				username: this.username,
				password: this.password
			});
			this.$router.push("/");
		} catch(e) {
			this.error = e.message;
		}
	}
});

export default class Signup extends Vue {}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
