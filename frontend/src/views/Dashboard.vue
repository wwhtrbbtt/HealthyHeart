<script setup lang="ts">
import { ref } from "vue";
import { GetReports, CreateFile } from "../../wailsjs/go/main/App";

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
      speed: 500,
      animateGradually: {
        enabled: true,
        delay: 150,
      },
      dynamicAnimation: {
        enabled: true,
        speed: 350,
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
    name: "last-10",
    data: [90, 84] as any[],
  },
]);

const reports = ref([] as any[]);
const latest = ref({ Timestamp: 0, Heartrate: 0 });
const last30 = ref([] as any[]);
const last7 = ref([] as any[]);

const getAverage = (reps: any[]) => {
  if (reps.length === 0) return "-";
  let sum = 0;
  reps.forEach((r) => (sum += r.Heartrate));
  return Math.round(sum / reps.length);
};

GetReports().then((r: any[]) => {
  reports.value = r;
  r.forEach((r) => {
    if (r.Timestamp > latest.value.Timestamp) latest.value = r;
    if (r.Timestamp > Date.now() / 1000 - 7 * 3600) last7.value.push(r);
    if (r.Timestamp > Date.now() / 1000 - 30 * 3600) last30.value.push(r);
  });
  series.value[0].data = reports.value.slice(0, 10).map((r) => r.Heartrate);
});

const saveCSV = async () => {
  let csv = "zeitpunkt,puls\n";
  reports.value.forEach(
    (r) =>
      (csv += `"${new Date(r.Timestamp * 1000).toLocaleString()}",${
        r.Heartrate
      }\n`)
  );
  await CreateFile(csv);
};
</script>

<template>
  <div class="mx-16">
    <div class="flex justify-between items-center">
      <div class="">
        <h1 class="text-5xl font-bold pt-12">Deine Werte.</h1>
        <button class="text-c-gray" @click="saveCSV">Exportieren</button>
      </div>
      <router-link to="/" class="px-4 shadow cursor-pointer pm-5"
        >Zurück</router-link
      >
    </div>
    <div class="border border-bottom w-96 mt-16 border-c-dark-gray"></div>

    <div class="grid grid-cols-2 gap-6">
      <div class="border border-c-gray">
        <p class="text-c-gray">Zuletzt gemessen</p>
        <p v-if="latest.Heartrate !== 0">
          <span class="text-8xl">{{ latest.Heartrate }}</span
          ><span class="text-xl text-c-gray">bpm</span>
        </p>
        <p v-else class="text-8xl">-</p>
      </div>
      <div class="border border-c-gray">
        <p class="text-c-gray">Historie</p>
        <div class="w-1/2">
          <apexchart
            type="line"
            :options="chartOptions"
            :series="series"
          ></apexchart>
        </div>
      </div>
      <div class="border border-c-gray">
        <p class="text-c-gray">30-Tage ⌀</p>
        <p>
          <span class="text-8xl">{{ getAverage(last30) }}</span
          ><span class="text-xl text-c-gray">bpm</span>
        </p>
      </div>
      <div class="border border-c-gray">
        <p class="text-c-gray">7-Tage ⌀</p>
        <p>
          <span class="text-8xl">{{ getAverage(last7) }}</span
          ><span class="text-xl text-c-gray">bpm</span>
        </p>
      </div>
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
