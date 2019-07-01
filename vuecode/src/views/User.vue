<template>
  <div>
    <nav class="navbar is-dark" role="navigation" aria-label="main navigation">
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

      <div class="navbar-end">
        <div class="navbar-item">
          <h1>Welcome User :- {{email}}</h1>
        </div>
        <div class="navbar-item">
          <b-button type="is-primary" @click="logoutt" outlined>Logout</b-button>
        </div>
      </div>
    </nav>
    <section class="hero is-primary is-bold is-medium">
      <div class="hero-body">
        <div class="tile is-ancestor">
          <div class="tile is-6 is-vertical is-parent">
            <div class="tile is-child box">
              <ul>
                <li v-bind:key="item.ID" v-for="item in userPostInfo">
                  <!-- <template v-if="index < activityCount"> -->
                  <article class="media">
                    <figure class="media-left">
                      <p class="image is-64x64">
                        <img src="https://bulma.io/images/placeholders/128x128.png" />
                      </p>
                    </figure>
                    <div class="media-content">
                      <div class="content">
                        <p>
                          <strong></strong>
                          <small>{{item.Title}}</small>
                          <small class="is-pulled-right">{{item.Email}}</small>
                          <br />
                          <b>{{item.Message}}</b>
                        </p>
                      </div>
                      <nav class="level is-mobile">
                        <div class="level-left">
                          <a class="level-item">
                            <span class="icon is-small">
                              <i class="fas fa-reply"></i>
                            </span>
                          </a>
                          <a class="level-item">
                            <span class="icon is-small">
                              <i class="fas fa-retweet"></i>
                            </span>
                          </a>
                          <a class="level-item">
                            <span class="icon is-small">
                              <i class="fas fa-heart"></i>
                            </span>
                          </a>
                        </div>
                      </nav>
                    </div>
                    <div class="media-right">
                      <button class="delete" @click="removePost(item.ID)"></button>
                    </div>
                  </article>
                </li>
              </ul>
              <br />
              <div class="has-margin">
                <b-button
                  v-if="activityCount < userPostInfo.length"
                  class="is-small"
                  @click="showMoreLessPosts('addItem')"
                >Show more</b-button>
                <b-button
                  v-if="activityCount > countLessMore"
                  class="is-small"
                  @click="showMoreLessPosts('removeItem')"
                >Show Less</b-button>
              </div>
            </div>
          </div>

          <div class="tile is-6 is-vertical is-parent">
            <div class="tile is-child box">
              <b-field label="Title">
                <b-input placeholder="Title of comment" v-model="title"></b-input>
              </b-field>
              <b-field label="Message">
                <b-input maxlength="200" type="textarea" v-model="message"></b-input>
              </b-field>
              <a class="button is-info is-medium is-fullwidth" @click.prevent="postt">Post it</a>
            </div>
          </div>
        </div>
      </div>
    </section>
    <footer class="footer is-dark">
      <div class="content has-text-centered">
        <p>
          <strong>Change URL</strong> by debugger
          <a href="#">Corporation</a>. The source code is licensed
          <a href="#">MIT</a>. The website content
          is licensed
          <a href="#">CC BY NC SA 4.0</a>.
        </p>
      </div>
    </footer>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      email: this.$route.params.email,
      title: "",
      message: "",
      info: [],
      userPostInfo: [],
      timeLineLength: 0,
      activityCount: 5,
      countLessMore: 5,
      orderActivityCount: 5
    };
  },
  mounted() {
    this.getposts();
  },
  updated() {
    // this.getposts();
  },
  methods: {
    // showMoreLessPosts(status) {
    //   if (status == "addItem") {
    //     this.activityCount += this.countLessMore;
    //   } else {
    //     this.activityCount -= this.countLessMore;
    //   }
    // },
    postt: function() {
      axios
        .post(
          "http://localhost:8000/user/" +
            this.email +
            "/" +
            this.title +
            "/" +
            this.message
        )
        .then(response => {
          this.getposts();
          (this.title = ""), (this.message = "");
        })
        .catch(e => {
          console.error(e);
        });
    },
    getposts: function() {
      axios.get("http://localhost:8000/user/" + this.email).then(response => {
        this.info = response.data;

        this.sortUserdata();
      });
    },
    logoutt: function() {
      this.$router.push("/login");
    },
    sortUserdata: function() {
      this.userPostInfo = [];
      let i;
      for (i = 0; i < this.info.length; i++) {
        if (this.info[i].Email == this.email) {
          this.userPostInfo.push(this.info[i]);
        }
      }
    },
    removePost: function(id) {
      axios
        .delete("http://localhost:8000/user/" + id)
        .then(response => {
          this.getposts();
        })
        .catch(error => {});
    }
  }
};
</script>