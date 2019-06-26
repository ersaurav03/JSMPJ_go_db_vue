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
          <a class="navbar-item">
            <router-link to="/">Home</router-link>
          </a>
          <a class="navbar-item">
            <router-link to="/About">About</router-link>
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
          <div class="columns">
            <div class="column">
              <h1 class="title">Login page</h1>
              <h2 class="subtitle"></h2>
            </div>
            <div class="column"></div>
            <div class="column"></div>
            <div class="column">
              <div class="field">
                <label class="label">Email Id</label>
                <div class="control">
                  <input class="input" type="email" v-model="email" placeholder="Email">
                </div>
              </div>
              <div class="field">
                <label class="label">Password</label>
                <div class="control">
                  <input class="input" type="password" v-model="password" placeholder="Pasword">
                </div>
              </div>

              <div class="field is-grouped">
                <div class="control">
                  <button class="button is-link" @click="logins">SignIn</button>
                </div>
              </div>
             
            </div>
          </div>
        </div>
      </div>
      <div v-if="red==true">
          <b-message type="is-success" has-icon>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce id fermentum quam. Proin sagittis, nibh id hendrerit imperdiet, elit sapien laoreet elit
        </b-message>
      </div>
    </section>

    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Change URL</strong> by
          <a href="#">JSMPJ Corporation</a>. The source code is licensed
          <a href="#">MIT</a>. The website content
          is licensed
          <a href="#">CC BY NC SA 4.0</a>.
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import axios from 'axios'
import { Toast } from 'buefy/dist/components/toast'
export default {
data () {
    return {
    info:[],
    email:"",
    password:"",
    red:false
    }
  },
  mounted () {
    this.home()
  },
 methods: {
   home: function (){
     axios
      .get('http://localhost:8000/user')
      .then(response => (this.info = response.data))
   },
   logins: function(){
     let i
      for (i=0;i<this.info.length;i++){
          if(this.email == this.info[i].Email && this.password ==this.info[i].Password){
             this.red=true
             this.$router.push({name:'user',
             params: { email: this.info[i].Email }
             })
             break
          }
          else{
            this.red=false
  
          }
      }
      if(this.red==false){
        Toast.open('Wrong Email id or Password')
      }
   }
 }


}
</script>