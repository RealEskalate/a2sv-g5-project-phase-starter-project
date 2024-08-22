import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { trendingArray } from "@/constants";
import { useUser } from "@/contexts/UserContext";

export default function TrendingTable() {
  const { isDarkMode } = useUser();

  return (
    <div className={`text-sm md:text-lg ${isDarkMode ? "text-gray-200" : "text-black"}`}>
      <Table className={`${isDarkMode ? "bg-gray-800" : "bg-white"} rounded-2xl`}>
        <TableHeader>
          <TableRow>
            <TableHead className={`text-sm md:text-lg text-center ${isDarkMode ? "text-gray-400" : "text-black"}`}>
              SL No
            </TableHead>
            <TableHead className={`text-sm md:text-lg text-center ${isDarkMode ? "text-gray-400" : "text-black"}`}>
              Name
            </TableHead>
            <TableHead className={`text-sm md:text-lg text-center ${ isDarkMode ? "text-gray-400" : "text-black"}`}>
              Price
            </TableHead>
            <TableHead className={`text-sm md:text-lg text-center ${isDarkMode ? "text-gray-400" : "text-black"}`}>
              Return
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {trendingArray.map((item) => (
            <TableRow key={item.id}>
              <TableCell className={`${isDarkMode ? "text-gray-300" : "text-black"} text-center`}>
                {item.id}
              </TableCell>
              <TableCell className={`${isDarkMode ? "text-gray-300" : "text-black"} text-center`}>
                {item.name}
              </TableCell>
              <TableCell className={`${isDarkMode ? "text-gray-300" : "text-black"} text-center`}>
                {item.price}
              </TableCell>
              <TableCell  className={`${isDarkMode ? "text-gray-300" : "text-black"} text-center`}>
                {item.return}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
