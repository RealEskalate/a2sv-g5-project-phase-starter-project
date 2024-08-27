"use client";
import dynamic from "next/dynamic";

const Chart = dynamic(() => import("react-apexcharts"), { ssr: false });

const chartConfig = {
  series: [30, 15, 35, 20],
  options: {
    chart: {
      type: "pie" as "pie",
      width: 280,
      height: 280,
      toolbar: {
        show: false,
      },
    },
    title: {
      text: "",
      align: "center" as "center",
      style: {
        fontSize: "16px",
        fontWeight: "bold",
      },
    },
    dataLabels: {
      enabled: true,
      formatter: (val: number, opts: any) => {
        const labels = ["Transfer", "Service", "Others", "Shopping"];
        return `${labels[opts.seriesIndex]} ${val.toPrecision(2)}%`;
      },
      style: {
        fontSize: "12px",
        fontWeight: "bold",
        colors: ["#fff"],
        innerWidth: "10px",
      },
    },
    plotOptions: {
      pie: {
        donut: {
          size: "85%",
        },
        expandOnClick: true,
        dataLabels: {
          offset: -15,
        },
        customScale: 1,
        offsetY: 5,
        offsetX: 0,
      },
    },
    colors: ["#020617", "#ff8f00", "#00897b", "#1e88e5", "#d81b60"],
    legend: {
      show: false,
    },
    states: {
      hover: {
        filter: {
          type: "darken",
          value: 0.9,
        },
      },
    },
    stroke: {
      show: true,
      width: 2,
    },
  },
};

export default function PieChart() {
  return (
    <div className=" max-h-[300px] h-[300px]  overflow-hidden  bg-white rounded-lg shadow-md">
      <Chart
        options={chartConfig.options}
        series={chartConfig.series}
        type="pie"
        width={"100%"}
        height={"100%"}
      />
    </div>
  );
}
