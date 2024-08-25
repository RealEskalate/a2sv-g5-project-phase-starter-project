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
              Type
            </TableHead>
          </TableRow>
        </TableHeader>
        
          {loading ? (
            <TableBody className="animate-pulse gap-2" >
              {
                Array.from({length:5}).map(() =>(
                  <TableRow >
                    <TableCell className={ `${isDarkMode ? "bg-gray-700":"bg-gray-200"} h-[25px]  rounded mx-auto my-2`} colSpan={3}></TableCell>
                    
                  </TableRow> ))
              }
            </TableBody>
          ):(
            <TableBody>
              {
                trendingArray.map((item:any, idx:number) => (
                  <TableRow key={item.id}>
                    <TableCell className={`${isDarkMode ? "text-gray-300" : "text-black"} text-center`}>
                      {idx+1}
                    </TableCell>
                    <TableCell className={`${isDarkMode ? "text-gray-300" : "text-black"} text-center`}>
                      {item.companyName}
                    </TableCell>
                    <TableCell className={`${isDarkMode ? "text-gray-300" : "text-black"} text-center`}>
                      {item.type}
                    </TableCell>
                  </TableRow> ))
              }
            </TableBody>
        )}
      </Table>
    </div>
  );
}
