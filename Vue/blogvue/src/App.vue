<template>
  <div id="app">
    <div id="nav">
     <users />
    </div>
    <router-view/>
  </div>
</template>

<style lang="stylus">
 @import "./style/stylehome.css";//就是在这里引用

#app
  font-family Avenir, Helvetica, Arial, sans-serif
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale
  text-align center
  color #2c3e50
  margin-top 60px
</style>
