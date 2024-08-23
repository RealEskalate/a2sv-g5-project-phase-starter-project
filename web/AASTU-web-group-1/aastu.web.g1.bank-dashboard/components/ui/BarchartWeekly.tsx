"use client"

import { TrendingUp } from "lucide-react"
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts"

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart"
import { useUser } from "@/contexts/UserContext"
const chartData = [
  { days: "Sunday", debit: 186, credit: 80 },
  { days: "Monday", debit: 305, credit: 200 },
  { days: "Tuesday", debit: 237, credit: 120 },
  { days: "Wednesday", debit: 73, credit: 190 },
  { days: "Thursday", debit: 209, credit: 130 },
  { days: "Friday", debit: 214, credit: 140 },
  { days: "Saturday", debit: 214, credit: 140 },
]

const chartConfig = {
  debit: {
    label: "Debit",
    color: "#1A16F3",
  },
  credit: {
    label: "Credit",
    color: "#FCAA0B",
  },
} satisfies ChartConfig

export function ChartWeekly() {
  const { isDarkMode } = useUser();
  return (
    <Card
      className={` ${
        isDarkMode ? "bg-gray-800 border-none " : "bg-white"
      } `}
    >
      <CardHeader>
        <CardDescription className="hidden md:block">
          {" "}
          $7,560 <span className="text-[#718EBF]">Debited</span> & $5,420{" "}
          <span className="text-[#718EBF]">Credited</span> in this Week{" "}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer
          config={chartConfig}
          className="lg:w-[800px] sm:w-[600px] w-[300px] h-60 "
        >
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="days"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="debit" fill="#1A16F3" radius={4} />
            <Bar dataKey="credit" fill="#FCAA0B" radius={4} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
