<template>
  <div>
    <div class="row" v-for="devices in result">
      <div class="columns medium-3 medium-6" v-for="device in devices">
        <div class="card">
          <div class="card-divider">
            {{ device.name }}
          </div>
          <div class="card-section">
            {{ device.status }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

const URL_BASE = `${process.env.API_HOST}/api/irkit/`;

export default {
  data() {
    return {
      result: [],
    };
  },
  async created() {
    const url = 'list';
    try {
      const resp = await axios.get(URL_BASE + url);
      this.result = resp.data;
    } catch (e) {
      // eslint-disable-next-line
      console.error(e);
    }
  },
};
</script>
