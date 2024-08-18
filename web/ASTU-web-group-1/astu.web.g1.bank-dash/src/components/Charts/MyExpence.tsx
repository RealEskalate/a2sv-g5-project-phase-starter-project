"use client";

import { BarChart, Bar, ResponsiveContainer, LabelList, Cell } from "recharts";

interface ExpenceChart {
  name: string;
  uv: number;
}

const data: ExpenceChart[] = [
  { name: "Jan", uv: 4000 },
  { name: "Feb", uv: 3000 },
  { name: "Mar", uv: 2000 },
  { name: "Apr", uv: 2780 },
  { name: "Jun", uv: 1890 },
  { name: "Jul", uv: 3490 },
];

const MyExpence = () => {
  return (
    <div className="w-full mt-5 lg:mt-0">
      <p className="text-[#333B69] pb-2 font-semibold">My Expense</p>
      <div className="bg-white p-6 rounded-3xl h-[190px] md:h-[250px] lg:h-[190px] w-full">
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={data}
            margin={{ top: 10, right: 0, left: 0, bottom: 30 }}
            width={150}
          >
            <Bar
              dataKey="uv"
              fill={data[0].name === "Apr" ? "#16DBCC" : "#EDF0F7"}
              barSize={25}
            >
              {data.map((entry, index) => (
                <Cell
                  cursor="pointer"
                  fill={entry["name"] === "Jun" ? "#16DBCC" : "#EDF0F7"}
                  key={`cell-${index}`}
                  radius={8}
                />
              ))}

              <LabelList dataKey="name" position="bottom" fill="#718EBF" />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default MyExpence;
