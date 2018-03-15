<template>
  <div>
    <div class="row" v-for="devices in deviceList">
      <div class="columns medium-3 medium-6" v-for="device in devices">
        <div class="card">
          <div class="card-divider">
            {{ device.name }}
          </div>
          <div v-if="device.status === 'false'"> <!-- off -->
            <div class="card-section">
              <a href="#"  v-on:click="dOn(device.name)">OFF</a>
            </div>
          </div>
          <div v-else> <!-- on -->
            <div class="card-section">
              <a href="#"  v-on:click="dOff(device.name)">ON</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.card {
  margin-bottom: 1rem;
  border: 1px solid #e6e6e6;
  border-radius: 0;
  background: #fefefe;
  -webkit-box-shadow: none;
          box-shadow: none;
  overflow: hidden;
  color: #0a0a0a;
}

.card > :last-child {
  margin-bottom: 0;
}

.card-divider {
  padding: 1rem;
  background: #e6e6e6;
  }
.card-divider > :last-child {
  margin-bottom: 0;
}

.card-section {
  padding: 1rem; }
  .card-section > :last-child {
    margin-bottom: 0; }
</style>


<script>
import { API } from '../utils/Api';

export default {
  name: 'DeviceCard',
  data() {
    return {
      devices: {},
    };
  },
  asyncData: {
    devicesDefault: {},
    devices() {
      return API.getDeviceList();
    },
  },
  computed: {
    deviceList() {
      return this.devices;
    },
  },

  methods: {
    dOff(id) {
      API.changeDeviceOff(id);
    },
    dOn(id) {
      API.changeDeviceOn(id);
    },
  },
};
</script>
