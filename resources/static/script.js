window.onload = function() {
    createChart('chart')
}

export function createChart(target) {
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
                data: [
                    {x: startDate, y: 120},
                    {x: startDate + 3600000, y: 245},
                    {x: startDate + 3600000 * 2, y: 272},
                ]
            },
            {
                type: 'line',
                name: 'Download',
                data: [
                    {x: startDate, y: 118},
                    {x: startDate + 3600000, y: 223},
                    {x: startDate + 3600000 * 2, y: 250},
                ]
            }
        ],
    })
}