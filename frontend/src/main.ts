import { createApp } from "vue";
import { createPinia } from "pinia";
import VueApexCharts from "vue3-apexcharts";
import App from "./App.vue";
import router from "./router";

import "./style.scss";

// import "./assets/main.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(VueApexCharts);

document.onkeydown = function (evt: any) {
  evt = evt || window.event;
  let isTrigger = false;
  if ("key" in evt) isTrigger = evt.key === "Escape" || evt.key === "Esc";
  else isTrigger = evt.keyCode === 27;

  if (isTrigger) router.push("/");
};

app.mount("#app");
