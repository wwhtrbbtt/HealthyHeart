<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { ref } from "vue";
import { GetThreshold, UpdateThreshold } from "../../wailsjs/go/main/App";

const threshold = ref(550);
GetThreshold().then((r) => (threshold.value = r));

const updateThreshold = () => UpdateThreshold(threshold.value);
</script>

<template>
  <div class="mx-16">
    <div class="flex justify-between items-center">
      <h1 class="text-5xl font-bold pt-12">Deine Einstellungen.</h1>
      <router-link to="/" class="px-4 shadow cursor-pointer pm-5 mt-12"
        >Zurück</router-link
      >
    </div>
    <div class="border border-bottom w-96 mt-16 border-c-dark-gray">
      <p>Herzschlaggrenzwert</p>
      <input
        type="number"
        v-model="threshold"
        @input="updateThreshold"
        class="p-3 bg-c-dark-gray rounded shadow-xl"
      />
    </div>

    <div class="circle circle-red"></div>
    <div class="circle circle-blue"></div>
  </div>
</template>

<style lang="scss">
.about {
  .title {
    margin: 30px auto 10px;
    font-size: 38px;
    color: #a150b5;
    text-align: center;
  }
  .content {
    position: relative;
    margin: 36px 20px;
    .comeon {
      position: absolute;
      left: 26px;
      top: 38px;
      max-width: 66%;
      img {
        width: 220px;
        height: 180px;
      }
    }
    .info {
      margin: 0 0 0 33%;
      font-size: 24px;
      text-align: left;
      .info-item {
        margin-bottom: 10px;
        .name {
          line-height: 40px;
          font-size: 28px;
          color: #6d6363;
        }
        .link {
          line-height: 30px;
          font-size: 20px;
          color: #5f6c86;
        }
      }
    }
  }

  .thank {
    height: 68px;
    line-height: 68px;
    margin: 36px auto;
    text-align: center;
    font-size: 40px;
  }
}
</style>
