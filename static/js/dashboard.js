$(function () {


  // =====================================
  // Profit
  // =====================================
  var chart = {
    series: [
      { 
        data:[
          {
            x:new Date('2024-01-01').getTime(),
            y:12,
          },
          {
            x:new Date('2024-02-01').getTime(),
            y:14,
          },
          {
            x:new Date('2024-03-01').getTime(),
            y:2,
          },
          {
            x:new Date('2024-04-01').getTime(),
            y:12,
          },
          {
            x:new Date('2024-05-01').getTime(),
            y:5,
          },
          {
            x:new Date('2024-06-01').getTime(),
            y:7,
          },
        ]
      },
    ],
    dataLabels: {
      position: 'top'
    },
    xaxis:{
      type: 'datetime'
    },
    chart: {
      type: "area",
    },
    stroke: {
      show: true,
      width: 3,
      lineCap: "butt",
      curve: 'smooth',
    },
    colors: ["#5D87FF", "#49BEFF"],
    markers: { size: 0 },
    dataLabels: {
      enabled: false,
    },
    legend: {
      show: false,
    },
    grid: {
      borderColor: "rgba(0,0,0,0.1)",
      strokeDashArray: 3,
      xaxis: {
        lines: {
          show: true,
        },
      },
    },
    tooltip: { 
      theme: "dark" 
    },
    responsive: [
      {
        breakpoint: 600,
        options: {
          plotOptions: {
            bar: {
              borderRadius: 3,
            }
          },
        }
      }
    ]


  };

  var chart = new ApexCharts(document.querySelector("#chart"), chart);
  chart.render();


  // =====================================
  // Breakup
  // =====================================
  var breakup = {
    color: "#adb5bd",
    series: [38, 40, 25],
    labels: ["2022", "2021", "2020"],
    chart: {
      width: 180,
      type: "donut",
      fontFamily: "Plus Jakarta Sans', sans-serif",
      foreColor: "#adb0bb",
    },
    plotOptions: {
      pie: {
        startAngle: 0,
        endAngle: 360,
        donut: {
          size: '75%',
        },
      },
    },
    stroke: {
      show: false,
    },

    dataLabels: {
      enabled: false,
    },

    legend: {
      show: false,
    },
    colors: ["#5D87FF", "#ecf2ff", "#F9F9FD"],

    responsive: [
      {
        breakpoint: 991,
        options: {
          chart: {
            width: 150,
          },
        },
      },
    ],
    tooltip: {
      theme: "dark",
      fillSeriesColor: false,
    },
  };

  var chart = new ApexCharts(document.querySelector("#breakup"), breakup);
  chart.render();



  // =====================================
  // Earning
  // =====================================
  var earning = {
    chart: {
      id: "sparkline3",
      type: "area",
      height: 60,
      sparkline: {
        enabled: true,
      },
      group: "sparklines",
      fontFamily: "Plus Jakarta Sans', sans-serif",
      foreColor: "#adb0bb",
    },
    series: [
      {
        name: "Earnings",
        color: "#49BEFF",
        data: [25, 66, 20, 40, 12, 58, 20],
      },
    ],
    stroke: {
      curve: "smooth",
      width: 2,
    },
    fill: {
      colors: ["#f3feff"],
      type: "solid",
      opacity: 0.05,
    },

    markers: {
      size: 0,
    },
    tooltip: {
      theme: "dark",
      fixed: {
        enabled: true,
        position: "right",
      },
      x: {
        show: false,
      },
    },
  };
  new ApexCharts(document.querySelector("#earning"), earning).render();
})