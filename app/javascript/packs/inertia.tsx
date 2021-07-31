import { App } from "@inertiajs/inertia-react";
import React from "react";
import { render } from "react-dom";
import axios from "axios";
import { InertiaProgress } from "@inertiajs/progress";

document.addEventListener("DOMContentLoaded", () => {
  InertiaProgress.init();
  const el = document.getElementById("app");

  const element: HTMLElement | null = document.querySelector(
    "meta[name=csrf-token]"
  );
  let csrfToken: string =
    element instanceof HTMLMetaElement ? element.content : "";

  axios.defaults.headers.common["X-CSRF-Token"] = csrfToken;

  render(
    <App
      initialPage={JSON.parse(el.dataset.page)}
      resolveComponent={(name) => require(`../Pages/${name}`).default}
    />,
    el
  );
});
