<template>
  <div>
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a
          role="button"
          class="navbar-burger burger"
          aria-label="menu"
          aria-expanded="false"
          data-target="navbarBasicExample"
        >
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item">HOME
          </a>
          <a class="navbar-item">
           <router-link to="/About">ABOUT</router-link>
          </a>
          <a class="navbar-item">
            <router-link to="/Contact">Contact</router-link>
          </a>
          <a class="navbar-item">
            <router-link to="/Login">Login</router-link>
          </a>
         <a class="navbar-item">
            <router-link to="/Register">Register</router-link>
          </a>
        </div>
      </div>
    </nav>
    <section class="hero is-success is-fullheight">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">Welcome User :- {{email}}</h1>
          <h2 class="subtitle"></h2>
        </div>
     <div class="container">
        <b-field label="Title">
            <b-input placeholder="Title of comment"></b-input>
        </b-field>
       <b-field label="Message"
            >
            <b-input maxlength="200" type="textarea"></b-input>
        </b-field>
        <a class="button is-info is-medium is-fullwidth">Post it</a>
        </div>
      </div>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Change URL</strong> by
          <a href="#"> Corporation</a>. The source code is licensed
          <a href="#">MIT</a>. The website content
          is licensed
          <a href="#">CC BY NC SA 4.0</a>.
        </p>
      </div>
    </footer>
  </div>
</template>
<script>
export default {
    data() {
 return {
 email: this.$route.params.email
 }
 }
}
</script>

