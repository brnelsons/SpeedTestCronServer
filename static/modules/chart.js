export function createChart(target) {
    return Highcharts.chart(target, {
        chart: {
            backgroundColor: "transparent",
        },
        title: {
            text: 'History',
        },
        xAxis: {
            title: {
                text: 'Date Time'
            },
            type: 'datetime'
        },
        yAxis: {
            title: {
                text: 'Speed (Mbps)'
            }
        },
        plotOptions: {
            series: {
                dataLabels: {
                    enabled: true,
                    format: "{y:.2f} Mbps",
                    numDecimalPlaces: 2,
                },
                enableMouseTracking: false,
            }
        },
        series: [
            {
                id: 'upload',
                type: 'line',
                name: 'Upload',
                data: []
            },
            {
                id: 'download',
                type: 'line',
                name: 'Download',
                data: []
            }
        ],
    })
}

export function updateChartData(chart, data) {
    chart.get('upload').setData(data.map(d => ({x: d['time'], y: d['upload']})));
    chart.get('download').setData(data.map(d => ({x: d['time'], y: d['download']})));
}
