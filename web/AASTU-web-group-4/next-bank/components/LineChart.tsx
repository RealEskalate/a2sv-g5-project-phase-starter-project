"use client"
import { Area, AreaChart, CartesianGrid, Legend, XAxis, YAxis } from "recharts"
import { Card, CardContent } from "@/components/ui/card"
import { ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent } from "@/components/ui/chart"
import { useEffect, useState } from "react"
import { getRandomBalanceHistory } from "@/services/transactionfetch"

// Mapping of numeric month values to month names
const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

const chartConfig = {
  balance: {
    color: "#1814F3",
  },
} satisfies ChartConfig

export default function LineChart() {
  const [lineChartData, setLineChartData] = useState([]);

  useEffect(() => {
    const fetchLineChart = async () => {
      try {
        const response = await getRandomBalanceHistory();
        
        // Check if the response contains the expected data
        if (response && response.data) {
          // Map API data to chart data structure
          const mappedData = response.data.map((item: { time: string, value: number }) => {
            const monthIndex = parseInt(item.time.split("-")[1]) - 1; // Extract month as an index (0-11)
            return {
              month: monthNames[monthIndex],
              balance: item.value,
            };
          });

          // Log the mapped data to the console
          // console.log("Mapped Data: ", mappedData);

          // Set the mapped data to state
          setLineChartData(mappedData);
        } else {
          console.error("Unexpected API response structure: ", response);
        }
      } catch (error) {
        console.error("Error fetching the random balance history: ", error);
      }
    };
    fetchLineChart();
  }, []);

  return (
    <Card className="w-[100%] rounded-2xl">
      <CardContent>
        <ChartContainer 
          config={chartConfig}
          className="w-[100%]"
        >
          <AreaChart
            accessibilityLayer
            data={lineChartData}  // Use fetched data here
          >
            <defs>
              <linearGradient id="colorBalance" x1="0" y1="0" x2="0" y2="1">
                <stop offset="0%" stopColor={chartConfig.balance.color} stopOpacity={0.2} />
                <stop offset="100%" stopColor={chartConfig.balance.color} stopOpacity={0.05} />
              </linearGradient>
            </defs>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={true}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)} // Show only the first 3 characters of the month
            />
            <YAxis dataKey="balance" axisLine={true} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="line" />}
            />
            <Area
              dataKey="balance"
              type="natural"
              fill="url(#colorBalance)"
              stroke={chartConfig.balance.color}
              fillOpacity={1}
              strokeWidth={3}
            />
            <Legend verticalAlign="top" height={36} />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
