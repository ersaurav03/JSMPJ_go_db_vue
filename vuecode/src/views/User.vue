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
          <h2 class="subtitle">{{info}}</h2>
                    <h2 class="subtitle">{{sort}}</h2>
        </div>
     <div class="container">
        <b-field label="Title">
            <b-input placeholder="Title of comment" v-model="title"></b-input>
        </b-field>
       <b-field label="Message"
            >
            <b-input maxlength="200" type="textarea" v-model="message"></b-input>
        </b-field>
        <a class="button is-info is-medium is-fullwidth" @click="postt">Post it</a>
        </div>
      </div>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Change URL</strong> by   debugger
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
import axios from 'axios'
export default {
    data() {
 return {
 email: this.$route.params.email,
 title:"",
 message:"",
 info:[],
 sort:[],
 postadd:{
     createdat:"",
     title:"",
     message:""
 }
 }
 },
 mounted(){
this.getposts()
this.sortposts()
 },
  methods: {
   postt: function (){
  axios.post('http://localhost:8000/user/'+this.email+"/"+this.title+"/"+this.message)
  .then((response) => {})
  .catch((e) => {
    console.error(e)
  })
   },
   getposts:function(){
            axios
      .get('http://localhost:8000/user/'+this.email)
      .then(response => (this.info = response.data))
   },
   sortposts:function(){

    debugger
    let i
    for(i=0;i<this.info.length;i++){
             
    //   this.postadd.createdat=this.info[i].CreatedAt,
    //   console.log(this.postadd.createdat)
    //   this.postadd.title=this.info[i].Title,
    //   this.postadd.message=this.info[i].Message
      this.sort[i].push({'date':'this.info[i].CreatedAt','title':'this.info[i].Title','message':'this.info[i].Message'})
    }
   }
   }
 

}
</script>