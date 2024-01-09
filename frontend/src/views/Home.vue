<script setup lang="ts">
import { ref } from "vue";
import { IsConnected } from "../../wailsjs/go/main/App";
import router from "../router";

let connected = ref(true);

setInterval(() => {
  IsConnected().then((r) => {
    if (!connected.value && r) {
      console.log("NEWLY CONNECTED DEVICE - INITIATING");
    }
    connected.value = r;
  });
}, 1000);

document.onkeydown = function (evt: any) {
  evt = evt || window.event;
  let isTrigger = false;
  if ("key" in evt) isTrigger = evt.key === "Escape" || evt.key === "Esc";
  else isTrigger = evt.keyCode === 27;

  if (isTrigger) router.push("/");
};
</script>

<template>
  <div class="mx-16">
    <div class="pt-24">
      <h1 class="text-5xl font-bold">Taktvoll. Gesund. Smart.</h1>
      <h1 class="text-5xl mt-1">Dein Herz im Blick</h1>
    </div>
    <p class="mt-10 w-96 text-c-gray">
      <span class="font-bold">HealthyHeart</span
      ><span class="font-light">
        hilft dir dabei, deinen Puls im Blick zu behalten, Krankheiten
        frühzeitig zu erkennen und weiterhin sorgenfrei leben zu können.</span
      >
    </p>

    <div class="border border-bottom w-96 mt-16 border-c-dark-gray"></div>

    <div class="flex gap-6 items-center my-4" v-if="connected">
      <router-link
        to="/measure"
        class="bg-button-blue py-2 px-4 shadow cursor-pointer"
        >Meine Werte messen</router-link
      >
      <router-link to="/dashboard" class="py-2 px-4 cursor-pointer"
        >Historie</router-link
      >
      <router-link to="/settings" class="py-2 px-4 text-c-gray cursor-pointer"
        >Einstellungen</router-link
      >
    </div>
    <div class="flex gap-6 items-center my-4" v-else>
      <router-link
        disabled
        to="/"
        class="bg-button-blue py-2 px-4 shadow cursor-pointer"
        >Sensor nicht verbunden!</router-link
      >
      <router-link to="/dashboard" class="py-2 px-4 cursor-pointer"
        >Historie</router-link
      >
      <router-link to="/settings" class="py-2 px-4 text-c-gray cursor-pointer"
        >Einstellungen</router-link
      >
    </div>

    <div class="flex md:flex-row flex-col gap-14 items-center my-6">
      <div class="border border-c-m-gray p-3 rounded">
        <h3 class="text-button-blue">Frühdiagnose</h3>
        <span class="text-c-m-gray-2">Ergebnisse, die weiterhelfen</span>
      </div>
      <div class="border border-c-m-gray p-3 rounded">
        <h3 class="text-button-blue">Präzision</h3>
        <span class="text-c-m-gray-2">Haargenaue Daten</span>
      </div>
      <div class="border border-c-m-gray p-3 rounded">
        <h3 class="text-button-blue">Benutzerfreundlichkeit</h3>
        <span class="text-c-m-gray-2">Kinderleichte Bedienung</span>
      </div>
    </div>
  </div>
  <div class="circle circle-red"></div>
  <div class="circle circle-blue"></div>
</template>

<style>
.circle {
  position: absolute;
  border-radius: 50%;
  width: 60vh;
  height: 80vh;
  opacity: 0.8;
  -webkit-filter: blur(100px);
  filter: blur(100px);
  -webkit-transform: scale(1, 0.2);
  transform: scale(1, 0.2);
}
.circle-red {
  top: -40%;
  background: #ed43624d;
}
.circle-blue {
  bottom: -40%;
  right: 0%;
  background: #4154ff33;
}
</style>
