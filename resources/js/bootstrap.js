import "htmx.org";
import Alpine from "alpinejs";
import axios from "axios";

window.Alpine = Alpine;
window.axios = axios;

window.axios.defaults.headers.common["X-Requested-With"] = "XMLHttpRequest";
Alpine.start();
