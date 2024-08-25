"use client";
import { useEffect, useState } from "react";
import displaytransaction from "./displaytransaction";
import Pagination from "@/components/Pagination";
import {
  getAllTransactions,
  getIncomes,
  getExpenses,
} from "@/services/transactionfetch";
import { TbFileSad } from "react-icons/tb";
import { colors } from "@/constants/index";

const RecentTransactions = () => {
  const [filter, setFilter] = useState<"all" | "income" | "expense">("all");
  const [currentPage, setCurrentPage] = useState(1);
  const [alltransaction, setAllTransaction] = useState([]);
  const [allincomes, setAllIncomes] = useState([]);
  const [allexpenses, setAllExpenses] = useState([]);
  const [totalPages, setTotalPages] = useState(0);
  const [isLoading, setIsLoading] = useState<"loading" | "success" | "error">(
    "loading"
  );
  const [activeTab, setActiveTab] = useState<"all" | "income" | "expense">(
    "all"
  );

  const ITEMS_PER_PAGE = 5;

  useEffect(() => {
    const fetchData = async () => {
      setIsLoading("loading");
      try {
        const [response, response2, response3] = await Promise.all([
          getAllTransactions(currentPage, ITEMS_PER_PAGE),
          getIncomes(currentPage, ITEMS_PER_PAGE),
          getExpenses(currentPage, ITEMS_PER_PAGE),
        ]);

        setAllTransaction(response.data.content || []);
        setTotalPages(response.data.totalPages);

        setAllIncomes(response2.data.content || []);
        setTotalPages(response2.data.totalPages);

        setAllExpenses(response3.data.content || []);
        setTotalPages(response3.data.totalPages);

        setIsLoading("success");
      } catch (error) {
        console.error("Error fetching transactions:", error);
        setIsLoading("error");
      }
    };

    fetchData();
  }, [currentPage]);

  if (isLoading === "loading") {
    return (
      <div className="flex flex-wrap lg:ml-72 px-10 flex-col gap-4">
        <div className="flex flex-wrap mb-4">
          <button className="font-bold px-4 py-2 rounded-t-lg text-gray-600">
            All Transactions
          </button>
          <button className="font-bold px-4 py-2 rounded-t-lg text-gray-600">
            Income
          </button>
          <button className="font-bold px-4 py-2 rounded-t-lg text-gray-600">
            Expenses
          </button>
        </div>
        <div className="w-full h-12 bg-gray-300 rounded-t-lg animate-pulse"></div>
        <div className="w-full h-12 bg-gray-300 rounded-t-lg animate-pulse"></div>
        <div className="w-full h-12 bg-gray-300 rounded-t-lg animate-pulse"></div>
      </div>
    );
  }

  if (isLoading === "error") {
    return (
      <div className="flex flex-wrap lg:ml-72 px-10 flex-col gap-4">
        <div className="flex flex-wrap mb-4">
          <button className="font-bold px-4 py-2 rounded-t-lg text-gray-600">
            All Transactions
          </button>
          <button className="font-bold px-4 py-2 rounded-t-lg text-gray-600">
            Income
          </button>
          <button className="font-bold px-4 py-2 rounded-t-lg text-gray-600">
            Expenses
          </button>
        </div>
        <div
          className="p-4 text-[#993d4b] flex flex-col gap-4 h-200px justify-center items-center rounded relative"
          role="alert"
        >
          <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
          <div>
            <span className="block sm:inline text-[18px]">
              {" "}
              Couldnt get your data bro...
            </span>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="p-4 md:ml-64">
      <div className="flex flex-wrap mb-4 dark:text-blue-500">
        <button
          onClick={() => {
            setFilter("all");
            setActiveTab("all");
          }}
          className={`font-bold px-4 py-2 rounded-t-lg ${
            filter === "all" ? "border-b-2 border-blue-500" : "text-gray-600"
          }`}
        >
          All Transactions
        </button>
        <button
          onClick={() => {
            setFilter("income");
            setActiveTab("income");
          }}
          className={`font-bold px-4 py-2 rounded-t-lg ${
            filter === "income" ? "border-b-2 border-blue-500" : "text-gray-600"
          }`}
        >
          Income
        </button>
        <button
          onClick={() => {
            setFilter("expense");
            setActiveTab("expense");
          }}
          className={`font-bold px-4 py-2 rounded-t-lg ${
            filter === "expense"
              ? "border-b-2 border-blue-500"
              : "text-gray-600"
          }`}
        >
          Expenses
        </button>
      </div>

      {activeTab === "all" && displaytransaction(alltransaction, "all")}
      {activeTab === "income" && displaytransaction(allincomes, "income")}
      {activeTab === "expense" && displaytransaction(allexpenses, "expense")}

      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={setCurrentPage}
      />
    </div>
  );
};

export default RecentTransactions;
