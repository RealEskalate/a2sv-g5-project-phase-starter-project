"use client";

import React, { useState, useEffect } from "react";
import axios from "axios";
import { BarChartComponent } from "./components/BarChartComponent";
import { TableComponent } from "./components/TableComponent";
import TableCard from "./components/TableComponentMobile";
import Card from "../components/common/card";
import columns from "./components/columns";
import creditCardColor from "@/app/CreditCards/cardMockData";
import { useSession } from "next-auth/react";

interface ExtendedUser {
  name?: string;
  email?: string;
  image?: string;
  accessToken?: string;
}

const Transactions: React.FC = () => {
  const [error, setError] = useState<string | null>(null);
  const [activeLink, setActiveLink] = useState<string>("recent");
  const [cardData, setCardData] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const { data: session } = useSession();
  const user = session?.user as ExtendedUser;
  const accessToken = user?.accessToken;
  const [data, setData] = useState<any[]>([]);
  const [expenseData, setExpenseData] = useState<any[]>([]);

  const fetchExpenseData = async () => {
    if (!accessToken) {
      setError("No access token available");
      setLoading(false);
      return;
    }

    try {
      const response = await fetch(
        `https://bank-dashboard-rsf1.onrender.com/transactions/expenses?page=0&size=7`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to fetch expenses");
      }
      const data = await response.json();
      setExpenseData(data.data.content);
    } catch (error) {
      setError((error as Error).message);
    } finally {
      setLoading(false);
    }
  };

<<<<<<< HEAD
  const fetchRecentData = async () => {
    if (!accessToken) {
      setError("No access token available");
      setLoading(false);
      return;
    }

    try {
      const response = await axios.get(
        `https://bank-dashboard-rsf1.onrender.com/transactions?page=0&size=5`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );

      const transformedData = response.data.data.content.map((item: any) => ({
        column1: item.description,
        column2: item.transactionId,
        column3: item.type,
        column4: "N/A", // Update this if you have card info
        column5: new Date(item.date).toLocaleDateString(),
        column6: `$${item.amount.toFixed(2)}`, // Format amount as currency
        column7: "N/A", // Update this if you have receipt info
      }));

      setData(transformedData);
    } catch (error) {
      setError("Failed to fetch data");
    } finally {
      setLoading(false);
    }
  };

  const fetchIncomeData = async () => {
    if (!accessToken) {
      setError("No access token available");
      setLoading(false);
      return;
    }

    try {
      const response = await fetch(
        `https://bank-dashboard-rsf1.onrender.com/transactions/incomes?page=0&size=5`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );

      const data = await response.json();
      setData(data.data.content || []);
    } catch (error) {
      setError("Failed to fetch income data");
    } finally {
      setLoading(false);
    }
  };
=======
  useEffect(() => {
    console.log("Updated expenseData:", expenseData);
  }, [expenseData]);

  useEffect(() => {
    fetchExpenseData();
  }, [accessToken]);
>>>>>>> aastu.web.g5.main

  const fetchCardData = async (page: number) => {
    if (!accessToken) {
      setError("No access token available");
      setLoading(false);
      return;
    }

    try {
      const response = await fetch(
        `https://bank-dashboard-rsf1.onrender.com/cards?page=${page}&size=3`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to fetch cards");
      }

      const data = await response.json();
      setCardData(data.content || []);
    } catch (error) {
      setError((error as Error).message);
    } finally {
      setLoading(false);
    }
  };

<<<<<<< HEAD
=======
  const fetchData = async () => {
    try {
      const response = await axios.get(
        `https://bank-dashboard-rsf1.onrender.com/transactions?page=${0}&size${5}`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );

      const transformedData = response.data.data.content.map((item: any) => ({
        column1: item.description,
        column2: item.transactionId,
        column3: item.type,
        column4: "N/A", // Update this if you have card info
        column5: new Date(item.date).toLocaleDateString(),
        column6: `$${item.amount.toFixed(2)}`, // Format amount as currency
        column7: "N/A", // Update this if you have receipt info
      }));

      setData(transformedData);
    } catch (error) {
      console.error("Failed to fetch data:", error);
      setError("Failed to fetch data. Please check the console for more details.");
    }
  };

>>>>>>> aastu.web.g5.main
  useEffect(() => {
    fetchCardData(0);
    fetchExpenseData(); // Always fetch expense data for BarChartComponent

    switch (activeLink) {
      case "recent":
        fetchRecentData();
        break;
      case "income":
        fetchIncomeData();
        break;
      case "expenses":
        // No action needed here, since expenseData is fetched already
        break;
      default:
        break;
    }
  }, [activeLink, accessToken]);

  const handleLinkClick = (linkName: string) => {
    setActiveLink(linkName);
  };
<<<<<<< HEAD
=======

>>>>>>> aastu.web.g5.main

  return (
    <div className="bg-[#F5F7FA] space-y-8 w-[95%] pt-3 overflow-hidden mx-auto">
      {/* First Row - My Cards and My Expenses */}
      <div className="flex flex-col lg:flex-row gap-6">
        <div className="flex-1 lg:w-[49%] overflow-hidden">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-lg font-semibold font-Inter text-[#343C6A]">
              My Cards
            </h2>
            <button className="text-[#343C6A] font-Inter font-medium">
              + Add Card
            </button>
          </div>
          <div className="flex gap-4">
            {cardData.map((card, index) => (
              <Card
                key={index}
                cardData={card}
                cardColor={creditCardColor[index % creditCardColor.length]}
              />
            ))}
          </div>
        </div>
        <div className="flex-1 lg:w-[49%]">
          <h2 className="text-lg font-semibold mb-4 font-Inter text-[#343C6A]">
            My Expenses
          </h2>
          <BarChartComponent data={expenseData} />
        </div>
      </div>

      {/* Second Row - Links and Table */}
      <div className="flex flex-col w-full">
        <div className="flex flex-row justify-start items-center mb-4 overflow-x-auto">
          <a
            href="#"
            className={`text-lg font-normal text-[#343C6A] mx-2 transition-all ${
              activeLink === "recent" ? "font-bold" : ""
            }`}
            onClick={() => handleLinkClick("recent")}
          >
            Recent Transactions
          </a>
          <a
            href="#"
            className={`text-lg font-normal text-[#343C6A] mx-2 transition-all ${
              activeLink === "income" ? "font-bold" : ""
            }`}
            onClick={() => handleLinkClick("income")}
          >
            Income
          </a>
          <a
            href="#"
            className={`text-lg font-normal text-[#343C6A] mx-2 transition-all ${
              activeLink === "expenses" ? "font-bold" : ""
            }`}
            onClick={() => handleLinkClick("expenses")}
          >
            Expenses
          </a>
        </div>

        <div className="hidden lg:flex flex-col w-full">
          {error ? <div>{error}</div> : <TableComponent columns={columns} data={data} />}
        </div>

        <div className="lg:hidden flex flex-col w-full">
          <TableCard data={data} />
        </div>
      </div>
    </div>
  );
};

export default Transactions;
``
