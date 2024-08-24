// "use client"
// import React, { useState } from "react";
// import Card from "../components/Page2/Card";
// import Tabs from "../components/Tabs";
// import BarChart from "../components/Page2/BarChart";
// import TransactionsList from "../components/Page2/TransactionsList";
// import AddCardForm from "@/app/creditCards/AddCardForm";
// import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
// import { Button } from "@/components/ui/button";

// const dummyCards = [
//   { id: 1, balance: "1000", cardHolder: "John Doe", expiryDate: "2024-10" },
//   { id: 2, balance: "2501", cardHolder: "Jane Smith", expiryDate: "2023-12" },
//   { id: 3, balance: "1002", cardHolder: "Alice Johnson", expiryDate: "2025-07" },
//   { id: 4, balance: "2503", cardHolder: "Michael Brown", expiryDate: "2025-01" },
//   { id: 5, balance: "1004", cardHolder: "Chris Evans", expiryDate: "2024-08" },
//   { id: 6, balance: "2505", cardHolder: "Scarlett Johansson", expiryDate: "2023-09" },
// ];

// const dummyTransactions = [
//   {
//     description: "Grocery Shopping",
//     transactionId: "T123",
//     type: "Expense",
//     receiverUserName: "Store",
//     date: "2024-07-21",
//     amount: "-50",
//     receipt: "/receipt1.pdf",
//   },
//   // (same as before)
// ];

// const formatDate = (date: string): string => {
//   const options: Intl.DateTimeFormatOptions = {
//     year: "numeric",
//     month: "2-digit",
//   };
//   return new Date(date).toLocaleDateString("en-US", options);
// };

// const Page = () => {
//   const [activeTab, setActiveTab] = useState("All Transactions");
//   const [cards, setCards] = useState(dummyCards.slice(0, 2));
//   const [showAllCards, setShowAllCards] = useState(false);
//   const [startIndex, setStartIndex] = useState(0);
//   const [isDialogOpen, setIsDialogOpen] = useState(false);

//   const handleTabChange = (tab: string) => {
//     setActiveTab(tab);
//   };

//   const handleNextCards = () => {
//     if (startIndex + 2 < dummyCards.length) {
//       setStartIndex(startIndex + 2);
//       setCards(dummyCards.slice(startIndex + 2, startIndex + 4));
//     }
//   };

//   const handlePreviousCards = () => {
//     if (startIndex - 2 >= 0) {
//       setStartIndex(startIndex - 2);
//       setCards(dummyCards.slice(startIndex - 2, startIndex));
//     }
//   };

//   const handleCardAddition = (newCard: any) => {
//     setCards([...cards, newCard]);
//   };

//   return (
//     <div className="bg-[#f5f7fa] dark:bg-[#020817] py-4 px-8 max-w-full [&::-webkit-scrollbar]:hidden overflow-x-hidden">
//       {/* Single Div containing header, cards, and bar chart */}
//       <div className="mb-4">
//         {/* Header */}
//         <div className="flex flex-col md:flex-row space-x-4 overflow-hidden">
//           <div className="flex flex-col md:w-2/3">
//             <div className="flex items-center justify-between mb-4">
//               <h2 className="text-xl font-bold text-[#343C6A] dark:text-[#9faaeb]">My Cards</h2>
              
//               {/* Add Card Button with Dialog Trigger */}
//               <DialogTrigger onClick={() => setIsDialogOpen(true)}>
//                 <Button className="border-none">+ Add Card</Button>
//               </DialogTrigger>
//             </div>
            
//             <div className="relative">
//               <div className="flex items-center space-x-2 overflow-x-auto">
//                 {/* Backward Icon */}
//                 {startIndex > 0 && (
//                   <button
//                     onClick={handlePreviousCards}
//                     className="text-[#343C6A] dark:text-[#9faaeb] focus:outline-none"
//                   >
//                     <img src="/back.svg" alt="Show Previous" className="h-6 w-6" />
//                   </button>
//                 )}
//                 <div className="flex overflow-x-auto [&::-webkit-scrollbar]:hidden space-x-2">
//                   {cards.map((item, index) => (
//                     <div key={item.id} className="flex-shrink-0 min-w-[200px]">
//                       <Card
//                         balance={`$${item.balance}`}
//                         cardHolder={item.cardHolder}
//                         validThru={formatDate(item.expiryDate)}
//                         cardNumber="3778 **** **** 1234"
//                         filterClass={index % 2 === 0 ? "" : "filter-black"}
//                         bgColor={
//                           index % 2 === 0
//                             ? "from-[#4C49ED] to-[#0A06F4]"
//                             : "from-white to-gray-200"
//                         }
//                         textColor={index % 2 === 0 ? "text-white" : "text-black"}
//                         iconBgColor="bg-opacity-10"
//                         showIcon={true}
//                       />
//                     </div>
//                   ))}
//                 </div>
//                 {/* Forward Icon */}
//                 {startIndex + 2 < dummyCards.length && (
//                   <button
//                     onClick={handleNextCards}
//                     className="text-[#343C6A] dark:text-[#9faaeb] focus:outline-none"
//                   >
//                     <img src="/forward.svg" alt="Show Next" className="h-6 w-6" />
//                   </button>
//                 )}
//               </div>
//             </div>
//           </div>

//           {/* Expense Chart Section */}
//           <div className="flex-grow md:w-1/3">
//             <h2 className="text-xl font-bold text-[#343C6A] dark:text-[#9faaeb] pb-10">My Expense</h2>
//             <BarChart token="dummy-token" />
//           </div>
//         </div>
//       </div>

//       {/* Dialog for Adding New Card */}
//       <Dialog
//         isOpen={isDialogOpen}
//         onClose={() => setIsDialogOpen(false)}
//       >
//         <DialogContent className="sm:max-w-[425px]">
//           <DialogHeader>
//             <DialogTitle>Add New Card</DialogTitle>
//             <DialogDescription>
//               Fill out the form below to add a new card.
//             </DialogDescription>
//           </DialogHeader>
//           <AddCardForm
//             access_token="your_access_token_here"
//             handleAddition={handleCardAddition}
//           />
//           <DialogFooter>
//             <Button variant="outline" onClick={() => setIsDialogOpen(false)}>
//               Close
//             </Button>
//           </DialogFooter>
//         </DialogContent>
//       </Dialog>

//       {/* Recent Transactions Section */}
//       <div className="mb-4 w-full mx-auto">
//         <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A] dark:text-[#9faaeb]">
//           Recent Transactions
//         </h2>
//         <Tabs
//           tabs={["All Transactions", "Income", "Expense"]}
//           activeTab={activeTab}
//           onTabChange={handleTabChange}
//         />
//         <TransactionsList
//           transactions={dummyTransactions.map((transaction) => ({
//             ...transaction,
//             amount: transaction.amount.toString(),
//           }))}
//         />
//       </div>
//     </div>
//   );
// };

// export default Page;


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
  const [showAllCards, setShowAllCards] = useState(false);
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
              ...prevCards.filter((card) =>
                !cardData.content.some((newCard) => newCard.id === card.id)
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
      if (access_token) {
        try {
          setLoading(true);
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
        } catch (error) {
          console.error("Error fetching transactions:", error);
        } finally {
          setLoading(false);
        }
      }
    };

    if (access_token) loadTransactions();
  }, [access_token, activeTab]);

  return (
    <div className="bg-[#f5f7fa] dark:bg-[#020817] py-4 px-8 max-w-full  [&::-webkit-scrollbar]:hidden overflow-x-hidden">
      {loading ? (
        <div className="text-center text-[#343C6A] dark:text-[#9faaeb]">Loading...</div>
      ) : (
        <>
          <div className="mb-4">
            <div className="flex flex-col md:flex-row space-x-4  [&::-webkit-scrollbar]:hidden overflow-hidden">
              <div className="flex flex-col md:w-2/3">
                <div className="flex items-center justify-between mb-4">
                  <h2 className="text-xl font-bold text-[#343C6A] dark:text-[#9faaeb]">My Cards</h2>
                  <Button className="border-none">+ Add Card</Button>
                </div>
                <div className="relative">
                  <div className="flex items-center space-x-2 overflow-x-auto">
                    {cards.map((item, index) => (
                      <div key={item.id} className="flex-shrink-0 min-w-[200px]">
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
                    ))}
                  </div>
                </div>
              </div>
              <div className="flex-grow md:w-1/3">
                <h2 className="text-xl font-bold text-[#343C6A] dark:text-[#9faaeb] pb-10">My Expense</h2>
                {access_token && <BarChart token={access_token} />}
              </div>
            </div>
          </div>

          {/* Recent Transactions */}
          <div className="mb-4 w-full mx-auto">
            <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A] dark:text-[#9faaeb]">Recent Transactions</h2>
            <Tabs tabs={["All Transactions", "Income", "Expense"]} activeTab={activeTab} onTabChange={setActiveTab} />
            <TransactionsList transactions={transactions.map((transaction) => ({ ...transaction, amount: transaction.amount.toString() }))} />
          </div>
        </>
      )}
    </div>
  );
};

export default Page;
