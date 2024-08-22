"use client"

import { Bar, BarChart } from "recharts"

import { ChartConfig, ChartContainer } from "@/components/ui/chart"

const chartData = [
  { month: "January", desktop: 186},
  { month: "February", desktop: 305},
  { month: "March", desktop: 237},
  { month: "April", desktop: 73},
  { month: "May", desktop: 209},
  { month: "June", desktop: 214},
]

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#2563eb",
  },
} satisfies ChartConfig

export function BarChartt() {
  return (
    <ChartContainer config={chartConfig} className="min-h-[200px] w-full">
      <BarChart accessibilityLayer data={chartData}>
        <Bar dataKey="desktop" fill="var(--color-desktop)" radius={4} />
      </BarChart>
    </ChartContainer>
  )
}
