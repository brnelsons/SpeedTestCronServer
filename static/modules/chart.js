export function createChart(target, data) {
    const startDate = new Date(2024, 1, 1, 2, 0).getTime();
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
            line: {
                dataLabels: {
                    enabled: true,
                    format: "{y} Mbps",
                },
                enableMouseTracking: false,
            }
        },
        series: [
            {
                type: 'line',
                name: 'Upload',
                data: data.map(d => ({x: d['x'], y: d['upload']}))
            },
            {
                type: 'line',
                name: 'Download',
                data: data.map(d => ({x: d['x'], y: d['download']}))
            }
        ],
    })
}