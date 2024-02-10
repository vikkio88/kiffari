import './app.css';

import App from './App.svelte';
import TimeAgo from "javascript-time-ago";
import en from "javascript-time-ago/locale/en";
TimeAgo.addDefaultLocale(en);

const app = new App({
  target: document.getElementById('app'),
});

export default app;
