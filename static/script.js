import {createChart, updateChartData} from "./modules/chart.js";
import {Service} from "./modules/service.js";

const service = new Service()
var historyChart;

function updateChart() {
    service.getHistory()
        .then(data => {
            updateChartData(historyChart, data)
        })
}

window.onload = function() {
    historyChart = createChart('chart');
    updateChart()
    setInterval(updateChart, 5000)
}