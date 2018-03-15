import Vue from 'vue';

const EventBus = new Vue();

// eslint-disable-next-line
export const Emit = function (event, ...args) {
  EventBus.$emit(event, ...args);
};
