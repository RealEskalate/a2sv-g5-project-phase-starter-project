// "use client"
// import { Bar, BarChart, YAxis, Legend, XAxis } from "recharts"
// import {
//   Card,
//   CardContent,
//   CardHeader,
//   CardTitle,
// } from "@/components/ui/card"
// import {
//   ChartConfig,
//   ChartContainer,
//   ChartTooltip,
//   ChartTooltipContent,
// } from "@/components/ui/chart"

// const chartData = [
//   { day: "Sat", debit: 67, credit: 50 },
//   { day: "Sun", debit: 30, credit: 10 },
//   { day: "Mon", debit: 78, credit: 50 },
//   { day: "Tue", debit: 55, credit: 75 },
//   { day: "Wed", debit: 69, credit: 11 },
//   { day: "Thu", debit: 50, credit: 66 },
//   { day: "Fri", debit: 40, credit: 22 },
// ]

// const chartConfig = {
//   debit: {
//     label: "Debit",
//     color: "#1814F3",
//   },
//   credit: {
//     label: "Credit",
//     color: "#FC7900",
//   },
// } satisfies ChartConfig

// export default function Component() {
//   return (
//     <Card className="flex flex-col px-6 w-full lg:max-w-lg lg:max-h-80">
//       <CardHeader>
//         <CardTitle>Weekly Activity</CardTitle>
//       </CardHeader>
//       <CardContent>
//         <ChartContainer config={chartConfig}>
//           <BarChart
//             data={chartData}
//             barCategoryGap="25%"  // Add space between groups of bars (days)
//             barGap={4}            // Add space between bars within a group
            
//           >
//             <XAxis dataKey="day" axisLine={true} tickLine={false} />
//             <ChartTooltip
//               cursor={false}
//               content={<ChartTooltipContent indicator="dashed" />}
//             />
//             <Legend verticalAlign="top" align="right" />
//             <Bar dataKey="debit" fill={chartConfig.debit.color} radius={[20, 20, 0, 0]} />
//             <Bar dataKey="credit" fill={chartConfig.credit.color} radius={[20, 20, 0, 0]} />
//           </BarChart>
//         </ChartContainer>
//       </CardContent>
//     </Card>
//   )
// }
// // export default function Component() {
// //   return (
// //     <Card className="flex flex-col px-6 w-full lg:max-w-lg lg:max-h-80">
// //       <CardHeader>
// //         <CardTitle>Weekly Activity</CardTitle>
// //       </CardHeader>
// //       <CardContent>
// //         <ChartContainer config={chartConfig}>
// //           <BarChart accessibilityLayer data={chartData}>
// //             <XAxis dataKey="day" axisLine={true} tickLine={false} />
// //             <ChartTooltip
// //               cursor={false}
// //               content={<ChartTooltipContent indicator="dashed" />}
// //             />
// //             <Legend verticalAlign="top" align="right" />
// //             <Bar dataKey="debit" fill={chartConfig.debit.color} radius={[20, 20, 0, 0]} />
// //             <Bar dataKey="credit" fill={chartConfig.credit.color} radius={[20, 20, 0, 0]} />
// //           </BarChart>
// //         </ChartContainer>
// //       </CardContent>
// //     </Card>
// //   )
// // }

"use client"
import { Bar, BarChart, YAxis, Legend, XAxis } from "recharts"
import {
  Card,
  CardContent,
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
  { day: "Sat", debit: 50, credit: 20 },
  { day: "Sun", debit: 30, credit: 10 },
  { day: "Mon", debit: 45, credit: 50 },
  { day: "Tue", debit: 33, credit: 15 },
  { day: "Wed", debit: 49, credit: 11 },
  { day: "Thu", debit: 33, credit: 44 },
  { day: "Fri", debit: 40, credit: 22 },
]

const chartConfig = {
  debit: {
    label: "Debit",
    color: "#1814F3",
  },
  credit: {
    label: "Credit",
    color: "#FC7900",
  },
} satisfies ChartConfig

export default function Component() {
  return (
    <Card className="flex flex-col w-full h-55">
      <CardHeader>
        <CardTitle>Weekly Activity</CardTitle>
      </CardHeader>
      <CardContent className="flex-1">
        <ChartContainer config={chartConfig}>
          <BarChart
            data={chartData}
            barCategoryGap="25%"
            barGap={5}
            barSize={30}
            // width="100%"        // Set width to 100% of the container
            // height="100%"       // Set height to 100% of the container
          >
            <XAxis dataKey="day" axisLine={true} tickLine={false} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Legend verticalAlign="top" align="right" />
            <Bar dataKey="debit" fill={chartConfig.debit.color} radius={[20, 20, 0, 0]} />
            <Bar dataKey="credit" fill={chartConfig.credit.color} radius={[20, 20, 0, 0]} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}

// // export default function Component() {
// //   return (
// //     <Card className="flex flex-col w-full">
// //       <CardHeader>
// //         <CardTitle>Weekly Activity</CardTitle>
// //       </CardHeader>
// //       <CardContent>
// //         <ChartContainer config={chartConfig}>
// //           <BarChart accessibilityLayer data={chartData}>
// //             <XAxis dataKey="month" axisLine={true} tickLine={false} />
// //             <ChartTooltip
// //               cursor={false}
// //               content={<ChartTooltipContent indicator="dashed" />}
// //             />
// //             <Legend verticalAlign="top" align="right" />
// //             <Bar dataKey="desktop" fill={chartConfig.desktop.color} radius={[20, 20, 0, 0]} />
// //             <Bar dataKey="credit" fill={chartConfig.mobile.color} radius={[20, 20, 0, 0]} />
// //           </BarChart>
// //         </ChartContainer>
// //       </CardContent>
// //     </Card>
// //   )
// // }
