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
const chartData = [
  { day: "Saturday", debit: 186, credit: 80 },
  { day: "Sunday", debit: 305, credit: 200 },
  { day: "Monday", debit: 237, credit: 120 },
  { day: "Tuesday", debit: 73, credit: 190 },
  { day: "Wednesday", debit: 209, credit: 130 },
  { day: "Thursday", debit: 214, credit: 140 },
  { day: "Friday", debit: 214, credit: 140 }
]

const chartConfig = {
  debit: {
    label: "Debit",
    color: "hsl(var(--chart-1))",
  },
  mobile: {
    label: "Credit",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig

export function DebitCreditOver() {
  return (
    <Card className="rounded-3xl  ">
      <CardHeader>
        <div className="flex justify-between ">
        <CardTitle className="text-base font-normal font-inter text-[#718EBF]"><span className="text-black">$7,560</span> Debited & <span className="text-black">$5,420</span> Credited in this Week</CardTitle>
        <div className="flex gap-5">
        <div className="flex items-center">
          <div className="border border-[#4C78FF] w-[15px] h-[15px] rounded-sm bg-[#4C78FF]">
          </div>
          <p className="font-inter font-normal text-base text-[#718EBF]">Debit</p>
        </div>
        <div className="flex items-center">
          <div className="border border-[#FCAA0B] w-[15px] h-[15px] rounded-sm bg-[#FCAA0B]">
          </div>
          <p className="font-inter font-normal text-base text-[#718EBF]">Credit</p>
        </div>
        </div>
        </div>
        {/* <CardDescription>January - June 2024</CardDescription> */}
      </CardHeader>
      <CardContent  >
        <ChartContainer config={chartConfig} className="h-[350px] w-[100%]">
          <BarChart accessibilityLayer data={chartData} >
            
            <CartesianGrid vertical={false} className="h-[70%]" />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="debit" fill="#1A16F3" radius={10} />
            <Bar dataKey="credit" fill="#FCAA0B" radius={10} />
          
          </BarChart>
        </ChartContainer>
      </CardContent>
      {/* <CardFooter className="flex-col items-start gap-2 text-sm">
        <div className="flex gap-2 font-medium leading-none">
          Trending up by 5.2% this month <TrendingUp className="h-4 w-4" />
        </div>
        <div className="leading-none text-muted-foreground">
          Showing total visitors for the last 6 months
        </div>
      </CardFooter> */}
    </Card>
  )
}
