import {createChart} from "./modules/chart.js";
import {Service} from "./modules/service.js";

const service = new Service()
var historyChart;
window.onload = function() {
    setInterval(() => {
        service.getHistory()
            .then(data => {
                historyChart = createChart('chart', data)
            })
    }, 5000)
}