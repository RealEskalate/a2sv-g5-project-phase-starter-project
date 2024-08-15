"use client";

import { BarChart, Bar, ResponsiveContainer, LabelList } from "recharts";

const data = [
  { name: "Mon", uv: 2000, pv: 2400, amt: 2400 },
  { name: "Tus", uv: 3000, pv: 2398, amt: 2210 },
  { name: "Wed", uv: 2000, pv: 2800, amt: 2290 },
  { name: "Thu", uv: 2780, pv: 3908, amt: 2000 },
  { name: "Fri", uv: 2890, pv: 2800, amt: 2181 },
  { name: "Sat", uv: 3490, pv: 2300, amt: 2100 },
  { name: "Sun", uv: 3490, pv: 2300, amt: 2100 },
];

const BarChartRechart = () => {
  return (
    <div className="w-full md:w-8/12 h-[350px] bg-white p-6 rounded-3xl">
      <div>
        <h3 className="blue-steel">
          <span className="text-deepNavy">$7,654</span> Debited &{" "}
          <span className="text-deepNavy">$5,420</span> Credited in this week
        </h3>
      </div>
      <ResponsiveContainer width="100%" height="100%">
        <BarChart
          data={data}
          barGap={5}
          margin={{ top: 10, right: 0, left: 0, bottom: 30 }}
        >
          <Bar dataKey="uv" fill="#1A16F3" barSize={30} radius={8}>
            <LabelList dataKey="name" position="bottom" fill="#718EBF" />
          </Bar>
          <Bar dataKey="pv" fill="#FCAA0B" barSize={30} radius={8} />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default BarChartRechart;
