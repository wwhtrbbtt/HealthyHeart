<script setup lang="ts">
import { ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { GetThreshold, ReportResult } from "../../wailsjs/go/main/App";
import router from "../router";

const datapointsAmount = 25;
const misses = 8;
const chartOptions = {
  grid: { show: false },
  tooltip: { enabled: false },
  theme: {
    mode: "light",
    palette: "palette10",
  },
  chart: {
    fontFamily: "JetBrainsMono",
    toolbar: { show: false },
    animations: {
      easing: "easeinout",
      speed: 1,
      animateGradually: {
        enabled: false,
        delay: 1,
      },
      dynamicAnimation: {
        enabled: false,
        speed: 1,
      },
    },
    id: "vuechart-example",
  },
  xaxis: {
    categories: [""],
  },
};

const series = ref([
  {
    name: "puls",
    data: [] as Number[],
  },
  {
    name: "threshold",
    data: [],
  },
]);

GetThreshold().then((t) => {
  series.value[1].data = Array(datapointsAmount).fill(t);
});

const heartbeat = ref(0);

EventsOn("heartrate", (value: number) => {
  if (value < 20 || value < 200) heartbeat.value = 0;
  heartbeat.value = Math.round(value);
});

let count = 0;
EventsOn("sensordata", (value: number) => {
  if (count++ !== misses) return;
  else count = 0;
  series.value[0].data.push(value);
  console.log("SENSOR NUMBER:", value);
  if (series.value[0].data.length > datapointsAmount)
    series.value[0].data.shift();
});

const save = async () => {
  router.push("/dashboard");
  await ReportResult(heartbeat.value);
};

document.onkeydown = async function (evt: any) {
  evt = evt || window.event;
  let isTrigger = false;
  if ("key" in evt) isTrigger = evt.key === "Enter";
  else isTrigger = evt.keyCode === 27;

  if (isTrigger) await save();
};
</script>

<template>
  <div class="mx-16">
    <div class="flex items-center justify-between">
      <div class="pt-24">
        <h1 class="text-5xl font-bold">Messung läuft...</h1>
        <p class="text-c-gray">Bitte lege deinen Ringfinger auf den Sensor</p>
        <button class="my-2" @click="save">Speichern</button>
      </div>
      <div class="border border-c-gray mx-10 mt-12">
        <p class="text-c-gray">Aktueller Herzschlag</p>
        <p>
          <span class="text-8xl">{{ heartbeat }}</span
          ><span class="text-xl text-c-gray">bpm</span>
        </p>
      </div>
    </div>
    <div class="border border-bottom w-96 mt-16 border-c-dark-gray"></div>

    <div class="w-9/12">
      <apexchart
        type="line"
        :options="chartOptions"
        :series="series"
      ></apexchart>
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
