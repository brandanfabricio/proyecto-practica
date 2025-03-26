import { defineStore } from "pinia";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: localStorage.getItem("token") || null,
    user: null,
  }),

  actions: {
    login(name, password, email) {
      fetch("http://localhost:5000/login", {
        method: "POST",
        body: JSON.stringify({
          name,
          password,
          email,
        }),
      })
        .then((res) => res.json())
        .then((response) => {
          console.log(response);
        })
        .catch((err) => {
          console.error(err);
        });
    },
    logout(){
        this.token = null,
        this.user = null
        localStorage.removeItem('token')
    }
  },
});
