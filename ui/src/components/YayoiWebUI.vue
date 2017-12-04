<template>
  <div id="app">
    <h1>Yayoi Web UI</h1>
    <device-cards></device-cards>
  </div>
</template>

<script>
import DeviceCards from './DeviceCards';

export default {
  name: 'web-ui',
  components: {
    DeviceCards,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
