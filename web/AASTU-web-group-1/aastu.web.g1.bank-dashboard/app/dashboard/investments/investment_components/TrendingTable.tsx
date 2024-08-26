'use client'
import React from "react";
import {Table, TableBody, TableCell, TableHead,TableHeader,TableRow,} from "@/components/ui/table";
import { getTrendingData } from "@/lib/investmentapies";
import { useState,useEffect } from "react";
import { useUser } from "@/contexts/UserContext";
export default function TrendingTable() {
  const { isDarkMode } = useUser();
  const [trendingArray, setTrendingArray] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getTrendingData();
        if (data==null){
          throw new Error("failed to get data");
        }
        if (data.length>5){
          data.splice(5);
        }
        setTrendingArray(data);
        setLoading(false);
        
      }catch (error) {
        console.error("An error occurred on Fetching trending data:", error);
      }
    };

    fetchData();

  },[])

  return ( 
    loading? (
      <div className="w-full mx-auto md:mx-0">
        <div className={`h-10 w-full ${isDarkMode ? "bg-gray-950" : "bg-gray-300"} animate-pulse rounded-md mb-2`}></div>
        <div className={`h-10 w-full ${isDarkMode ? "bg-gray-950" : "bg-gray-300"} animate-pulse rounded-md mb-2`}></div>
        <div className={`h-10 w-full ${isDarkMode ? "bg-gray-950" : "bg-gray-300"} animate-pulse rounded-md mb-2`}></div>
        <div className={`h-10 w-full ${isDarkMode ? "bg-gray-950" : "bg-gray-300"} animate-pulse rounded-md mb-2`}></div>
        <div className={`h-10 w-full ${isDarkMode ? "bg-gray-950" : "bg-gray-300"} animate-pulse rounded-md mb-2`}></div>
        <div className={`h-10 w-full ${isDarkMode ? "bg-gray-950" : "bg-gray-300"} animate-pulse rounded-md mb-2`}></div>
      </div>
    ):
    (<div className={`text-sm md:text-lg ${isDarkMode ? "text-gray-200" : "text-black"}`}>
      <Table className={`${isDarkMode ? "bg-gray-800" : "bg-white"} rounded-2xl`}>
        <TableHeader className={`${isDarkMode ? "bg-[#1a1d2288] rounded-xl":"bg-transparent"} text-white`}>
          <TableRow className={`"bg-transparent ${isDarkMode? "border-b-gray-700": ""}`}>
            <TableHead className={`text-sm md:text-lg text-center ${isDarkMode ? "text-white rounded-tl-xl " : "text-black"}`}>
              SL No
            </TableHead>
            <TableHead className={`text-sm md:text-lg text-center ${isDarkMode ? "text-white" : "text-black"}`}>
              Name
            </TableHead>
            <TableHead className={`text-sm md:text-lg text-center ${ isDarkMode ? "text-white " : "text-black"} rounded-tr-xl`}>
              Type
            </TableHead>
          </TableRow>
        </TableHeader>
            <TableBody>
              {
                trendingArray.map((item:any, idx:number) => (
                  <TableRow key={item.id} className={`"bg-transparent ${isDarkMode? "border-b-gray-700": ""}`}>
                    <TableCell className={`${isDarkMode ? "text-white" : "text-black"} text-center`}>
                      {idx+1}
                    </TableCell>
                    <TableCell className={`${isDarkMode ? "text-white" : "text-black"} text-center`}>
                      {item.companyName}
                    </TableCell>
                    <TableCell className={`${isDarkMode ? "text-white" : "text-black"} text-center`}>
                      {item.type}
                    </TableCell>
                  </TableRow> ))
              }
            </TableBody>
      </Table>
    </div>)
    
  );
}
