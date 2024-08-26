"use client";

import React, { useState, useEffect } from "react";
import { getCards } from "@/lib/api/cardController";
import { getSession } from "next-auth/react";
import Card from "../components/Page2/Card";
import Tabs from "../components/Tabs";
import BarChart from "../components/Page2/BarChart";
import TransactionsList from "../components/Page2/TransactionsList";
import { Card as CardType } from "@/types/cardController.Interface";
import {
  TransactionData,
  GetTransactionsResponse,
  PaginatedTransactionsResponse,
} from "@/types/transactionController.interface";
import {
  getTransactions,
  getTransactionIncomes,
  getTransactionsExpenses,
} from "@/lib/api/transactionController";
import { useRouter } from "next/navigation";
import Refresh from "@/app/api/auth/[...nextauth]/token/RefreshToken";
import {Dialog} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button";

// Utility to format dates
const formatDate = (date: string): string => {
  const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "2-digit",
  };
  return new Date(date).toLocaleDateString("en-US", options);
};

const Page = () => {
  const [activeTab, setActiveTab] = useState("All Transactions");
  const [cards, setCards] = useState<CardType[]>([]);
  const [transactions, setTransactions] = useState<TransactionData[]>([]);
  const [page, setPage] = useState(0);
  const [size] = useState(3);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const [access_token, setAccess_token] = useState("");
  const [startIndex, setStartIndex] = useState(0);
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const router = useRouter();

  // Load session and refresh token
  useEffect(() => {
    const fetchSessionAndRefreshToken = async () => {
      setLoading(true);
      try {
        const accessToken = await Refresh();
        setAccess_token(accessToken);
      } catch (error) {
        console.error("Error fetching session or refreshing token:", error);
        router.push(`/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
        router.push(`/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
      } finally {
        setLoading(false);
      }
    };

    fetchSessionAndRefreshToken();
  }, [router]);

  // Fetch cards
  useEffect(() => {
    const loadCards = async () => {
      if (access_token) {
        try {
          setLoading(true);
          const cardData = await getCards(access_token, page, size);
          if (cardData.content.length > 0) {
            setCards((prevCards) => [
              ...prevCards.filter(
                (card) => !cardData.content.some((newCard) => newCard.id === card.id)
              ),
              ...cardData.content,
            ]);
            setPage((prevPage) => prevPage + 1);
            if (cardData.content.length < size) setHasMore(false);
          } else setHasMore(false);
        } catch (error) {
          console.error("Error fetching cards:", error);
        } finally {
          setLoading(false);
        }
      }
    };

    if (access_token && page === 0 && cards.length === 0) {
      loadCards();
    }
  }, [access_token, page, size, cards.length]);

  // Fetch transactions
  useEffect(() => {
    const loadTransactions = async () => {
      setLoading(true);
      try {
        if (access_token) {
          let response: GetTransactionsResponse | PaginatedTransactionsResponse;
          switch (activeTab) {
            case "Income":
              response = await getTransactionIncomes(0, 100, access_token);
              break;
            case "Expense":
              response = await getTransactionsExpenses(0, 100, access_token);
              break;
            default:
              response = await getTransactions(0, 100, access_token);
          }

          if ("data" in response) {
            setTransactions(response.data.content);
          } else if ("transactions" in response) {
            const allTransactions = response.transactions.flatMap(
              (transactionResponse) => transactionResponse.data.content
            );
            setTransactions(allTransactions);
          } else {
            console.error("Unknown response type:", response);
          }
        }
      } catch (error) {
        console.error("Error fetching transactions:", error);
      } finally {
        setLoading(false);
      }
    };

    if (access_token) loadTransactions();
  }, [access_token, activeTab]);

  const handleTabChange = (tab: string) => {
    setActiveTab(tab);
  };

  const cardsPerPage = window.innerWidth >= 1200 ? 2 : 1;

  const handleNextCards = () => {
    if (startIndex + cardsPerPage < cards.length) {
      setStartIndex((prevIndex) => prevIndex + cardsPerPage);
    }
  };

  const handlePreviousCards = () => {
    if (startIndex - cardsPerPage >= 0) {
      setStartIndex((prevIndex) => prevIndex - cardsPerPage);
    }
  };

  const handleCardAddition = (newCard: any) => {
    setCards([...cards, newCard]);
  };

  return (
    <div className="bg-[#f5f7fa] dark:bg-[#090b0e] py-4 px-4 md:px-8 max-w-full overflow-x-hidden">
      <div className="mb-4">
        <div className="flex flex-col md:flex-row space-x-4">
          {/* Card Section */}
          <div className="flex flex-col w-full md:w-1/2 lg:w-3/5">
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-xl font-black text-[#343C6A] dark:text-[#9faaeb]">My Cards</h2>
              <DialogTrigger onClick={() => setIsDialogOpen(true)}>
                <Button className="border-none">+ Add Card</Button>
              </DialogTrigger>
            </div>
            <div className="relative">
              <div className="flex flex-col md:flex-row overflow-x-auto [&::-webkit-scrollbar]:hidden">
                <div className="flex items-center space-x-1">
                  {startIndex > 0 && (
                    <button
                      onClick={handlePreviousCards}
                      className="text-[#343C6A] dark:text-[#9faaeb] focus:outline-none "
                    >
                      <img src="/back.svg" alt="Show Previous" className="h-6 w-6" />
                    </button>
                  )}
                  {cards.length > 0 ? ( 
                    window.innerWidth < 768 ?(
                      cards
                      .map((item, index) => (
                    <div key={item.id} className="flex-shrink-0 overflow-x-auto min-w-[350px] [&::-webkit-scrollbar]:hidden">
                      <Card
                        balance={`$${item.balance}`}
                        cardHolder={item.cardHolder}
                        validThru={formatDate(item.expiryDate)}
                        cardNumber="3778 **** **** 1234"
                        filterClass={index % 2 === 0 ? "" : "filter-black"}
                        bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "from-white to-gray-200"}
                        textColor={index % 2 === 0 ? "text-white" : "text-black"}
                        iconBgColor="bg-opacity-10"
                        showIcon={true}
                      />
                    </div>))):(
                    cards
                      .slice(startIndex, startIndex + (window.innerWidth < 900 ? 1 : 2))
                      .map((item, index) => (
                        <div key={item.id} className="flex-shrink-0 min-w-[300px] md:min-w-[300px] lg:min-w-[350px]">
                          <Card
                            balance={`$${item.balance}`}
                            cardHolder={item.cardHolder}
                            validThru={formatDate(item.expiryDate)}
                            cardNumber="3778 **** **** 1234"
                            filterClass={index % 2 === 0 ? "" : "filter-black"}
                            bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "from-white to-gray-200"}
                            textColor={index % 2 === 0 ? "text-white" : "text-black"}
                            iconBgColor="bg-opacity-10"
                            showIcon={true}
                          />
                        </div>
                      )))
                  ) : (
                    [...Array(2)].map((_, i) => (
                      <div key={i} className="border dark:border-[#333B69] rounded-3xl my-4 mx-2 animate-pulse">
                        <div className="relative w-full bg-gradient-to-b from-gray-200 dark:from-[#333B69] to-gray-300 dark:to-[#555B85] text-transparent rounded-3xl shadow-md h-[230px] min-w-[350px]">
                          <div className="flex justify-between items-start px-6 pt-6">
                            <div>
                              <p className="text-xs font-semibold bg-gray-300 dark:bg-[#555B85] rounded w-16 h-4 mb-2"></p>
                              <p className="text-xl font-medium bg-gray-300 dark:bg-[#555B85] rounded w-24 h-6"></p>
                            </div>
                            <div className="w-8 h-8 bg-gray-300 dark:bg-[#555B85] rounded-full"></div>
                          </div>
                          <div className="flex justify-between gap-12 mt-4 px-6">
                            <div>
                              <p className="text-xs font-medium bg-gray-300 dark:bg-[#555B85] rounded w-16 h-4 mb-2"></p>
                              <p className="font-medium text-base bg-gray-300 dark:bg-[#555B85] rounded w-24 h-6"></p>
                            </div>
                            <div className="pr-8">
                              <p className="text-xs font-medium bg-gray-300 dark:bg-[#555B85] rounded w-16 h-4 mb-2"></p>
                              <p className="font-medium text-base md:text-lg bg-gray-300 dark:bg-[#555B85] rounded w-24 h-6"></p>
                            </div>
                          </div>
                          <div className="relative mt-8 flex justify-between py-4 items-center">
                            <div className="absolute inset-0 w-full h-full bg-gradient-to-b from-white/20 dark:from-gray-700/20 to-transparent z-0"></div>
                            <div className="ml-4 relative z-10 text-base font-medium px-6 bg-gray-300 dark:bg-[#555B85] rounded w-40 h-6"></div>
                            <div className="flex justify-end relative z-10 px-6">
                              <div className="w-10 h-10 bg-gray-300 dark:bg-[#555B85] rounded-full"></div>
                            </div>
                          </div>
                        </div>
                      </div>
                    ))
                  )}
                  {startIndex +  (window.innerWidth < 900 ? 1 : 2) < cards.length && (
                    <button
                      onClick={handleNextCards}
                     className="text-[#343C6A] dark:text-[#9faaeb] focus:outline-none"
                    >
                      <img src="/forward.svg" alt="Show Next" className="h-6 w-6" />
                    </button>
                  )}
                </div>
              </div>
            </div>
          </div>

          {/* Bar Chart Section */}
          <div className="w-full md:w-1/2 lg:w-1/5 flex-shrink-0  min-w-[300px] lg:min-w-[350px] pr-2 pt-2">
            <h2 className="text-xl font-black text-[#343C6A] dark:text-[#9faaeb] pb-8 ">My Expense</h2>
            {access_token && <BarChart token={access_token} />}
          </div>
        </div>
      </div>

      <Dialog isOpen={isDialogOpen} onClose={() => setIsDialogOpen(false)}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Add New Card</DialogTitle>
            <DialogDescription>
              Fill out the form below to add a new card.
            </DialogDescription>
          </DialogHeader>
          <AddCardForm
            access_token={access_token}
            handleAddition={handleCardAddition}
          />
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsDialogOpen(false)}
            >
              Close
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Recent Transactions */}
      <div className="mb-4 w-full mx-auto">
        <h2 className="text-xl font-black mb-4 pt-6 text-[#343C6A] dark:text-[#9faaeb]">
          Recent Transactions
        </h2>
        <Tabs
          tabs={["All Transactions", "Income", "Expense"]}
          activeTab={activeTab}
          onTabChange={handleTabChange}
        />
        <TransactionsList
          transactions={transactions.map((transaction) => ({
            ...transaction,
            amount: transaction.amount.toString(),
          }))}
          loading={loading}
          activeTab={activeTab}
        />
      </div>
    </div>
  );
};

export default Page;
