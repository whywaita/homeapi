<template>
  <div>
    <div class="row" v-for="devices in deviceList">
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
};
</script>
