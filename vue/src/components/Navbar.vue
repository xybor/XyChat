<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">XyChat</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link class="nav-link" to="/">Home</router-link>
          </li>
          <li class="nav-item" v-if="!status.loggedIn">
            <router-link class="nav-link" to="/login">Login</router-link>
          </li>
          <li class="nav-item" v-if="!status.loggedIn">
            <router-link class="nav-link" to="/register">Register</router-link>
          </li>
          <li class="nav-item dropdown" v-if="status.loggedIn">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              id="profile-user"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              Hi {{ user.username }}!
            </a>
            <ul class="dropdown-menu" aria-labelledby="profile-user">
              <li>
                <router-link class="dropdown-item" to="/profile"
                  >Your profile</router-link
                >
              </li>
              <div class="dropdown-divider"></div>
              <li><span class="dropdown-item" @click="logout">Logout</span></li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
import { mapState, mapActions } from "vuex";
export default {
  computed: {
    ...mapState("account", ["status", "user"]),
  },
  methods: {
    ...mapActions("account", ["logout"]),
  },
};
</script>
