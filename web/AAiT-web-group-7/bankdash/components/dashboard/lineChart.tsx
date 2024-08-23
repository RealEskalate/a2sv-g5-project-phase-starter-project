'use client';
import {
    AreaChart,
    Area,
    ResponsiveContainer,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
  } from "recharts";
  
  const productSales = [
    { name: "Jul", Balance: 100 },
    { name: "Aug", Balance: 230 },
    { name: "Sep", Balance: 430 },
    { name: "Oct", Balance: 790 },
    { name: "Nov", Balance: 210 },
    { name: "Dec", Balance: 450 },
    { name: "Jan", Balance: 220 },
  ];
  
  const AreaChartComponent = () => {
    return (
      // w-full h-[300px] sm:h-[400px] lg:h-[500px] bg-white rounded-3xl px-3 py-6
      <div className="w-full h-[300px] sm:h-[200px] md:h-[250px]">
        <ResponsiveContainer width='100%' height='100%' className='bg-white rounded-3xl py-2 md:px-3 md:py-6' >
          <AreaChart
            width={500}
            height={80}
            data={productSales}
            margin={{ right: 30 }}
          >
            <defs>
              <linearGradient id="colorVisitors" x1="0" y1="0" x2="0" y2="1">
                <stop offset="5%" stopColor="#2D60FF40" stopOpacity={1} />
                <stop offset="95%" stopColor="#2D60FF00" stopOpacity={0} />
              </linearGradient>
            </defs>
            <YAxis />
            <XAxis dataKey="name" />
            <CartesianGrid stroke="#eaf2f2" />
            <Tooltip
              content={<CustomTooltip active={false} payload={[]} label={""} />}
              cursor={{
                stroke: "red",
                strokeWidth: 2,
              }}
              contentStyle={{
                borderRadius: 4,
              }}
            />
            <Area
              type="basis"
              dataKey="Balance"
              stroke="#2563eb"
              fill="url(#colorVisitors)"
              stackId="1"
              strokeWidth={2}
            />
          </AreaChart>
        </ResponsiveContainer>
        </div>
    );
  };
  
  const CustomTooltip = ({
    active,
    payload,
    label,
  }: {
    active: boolean;
    payload: unknown[];
    label: string;
  }) => {
    if (active && payload && payload.length) {
      return (
        <div className="p-1 bg-[#FF5B5B] flex flex-col gap-1 rounded-md">
          <p className="text-medium text-sm text-white">{label}</p>
          <p className="text-sm text-white">
            Balance
            <span className="ml-2">{(payload[0] as { value: string }).value}</span>
          </p>
        </div>
      );
    }
    return null;
  };
  
  export default AreaChartComponent;