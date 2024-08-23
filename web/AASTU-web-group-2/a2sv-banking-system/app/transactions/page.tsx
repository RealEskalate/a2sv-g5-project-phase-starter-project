// "use client";

// import React, { useState, useEffect } from 'react';
// import { getCards } from '@/lib/api/cardController';
// import { getSession } from "next-auth/react";
// import Card from '../components/Page2/Card';
// import Tabs from '../components/Tabs';
// import BarChart from '../components/Page2/BarChart';
// import TransactionsList from '../components/Page2/TransactionsList';
// import { Card as CardType } from '@/types/cardController.Interface';
// import { TransactionData, GetTransactionsResponse, PaginatedTransactionsResponse } from '@/types/transactionController.interface';
// import { getTransactions, getTransactionIncomes, getTransactionsExpenses } from '@/lib/api/transactionController';
// import { useRouter } from 'next/navigation';
// import Refresh from '@/app/api/auth/[...nextauth]/token/RefreshToken'; 
// import WhiteCard from '../components/Page2/WhiteCard';

// type Data = {
//   access_token: string;
//   data: string;
//   refresh_token: string;
// };

// type SessionDataType = {
//   user: Data;
// };

// const formatDate = (date: string): string => {
//   const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: '2-digit' };
//   return new Date(date).toLocaleDateString('en-US', options);
// };

// const isPaginatedTransactionsResponse = (response: GetTransactionsResponse | PaginatedTransactionsResponse): response is PaginatedTransactionsResponse => {
//   return (response as PaginatedTransactionsResponse).data !== undefined;
// };

// const isGetTransactionsResponse = (response: GetTransactionsResponse | PaginatedTransactionsResponse): response is GetTransactionsResponse => {
//   return (response as GetTransactionsResponse).transactions !== undefined;
// };

// const Page = () => {
//   const [activeTab, setActiveTab] = useState('All Transactions');
//   const [cards, setCards] = useState<CardType[]>([]);
//   const [transactions, setTransactions] = useState<TransactionData[]>([]);
//   const [page, setPage] = useState(0);
//   const [size] = useState(3);
//   const [loading, setLoading] = useState(false);
//   const [hasMore, setHasMore] = useState(true);
//   const [session, setSession] = useState<SessionDataType | null>(null);

//   const router = useRouter();

//   // useEffect(() => {
//   //   const fetchSessionAndRefreshToken = async () => {
//   //     setLoading(true);
//   //     try {
//   //       const accessToken = await Refresh(); 
//   //       const sessionData = (await getSession()) as SessionDataType | null;
//   //       if (sessionData && sessionData.user) {
//   //         setSession({ user: { ...sessionData.user, access_token: accessToken } });
//   //       } else {
//   //         router.push(/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")});
//   //       }
//   //     } catch (error) {
//   //       console.error("Error fetching session or refreshing token:", error);
//   //       router.push(/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")});
//   //     } finally {
//   //       setLoading(false);
//   //     }
//   //   };

//   //   fetchSessionAndRefreshToken();
//   // }, [router]);

//   useEffect(() => {
//     const loadCards = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           const cardData = await getCards(session.user.access_token, page, size);
//           if (cardData.content.length > 0) {
//             setCards((prevCards) => [
//               ...prevCards.filter(card => !cardData.content.some(newCard => newCard.id === card.id)),
//               ...cardData.content
//             ]);
//             setPage((prevPage) => prevPage + 1);
//             if (cardData.content.length < size) {
//               setHasMore(false);
//             }
//           } else {
//             setHasMore(false);
//           }
//         } catch (error) {
//           console.error('Error fetching cards:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       if (page === 0 && cards.length === 0) {
//         loadCards();
//       }
//     }
//   }, [session, page, size, cards.length]);

//   useEffect(() => {
//     const loadTransactions = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           let response: GetTransactionsResponse | PaginatedTransactionsResponse;
//           switch (activeTab) {
//             case 'Income':
//               response = await getTransactionIncomes(0, 100, session.user.access_token);
//               break;
//             case 'Expense':
//               response = await getTransactionsExpenses(0, 100, session.user.access_token);
//               break;
//             default:
//               response = await getTransactions(0, 100, session.user.access_token);
//           }
//           if (isPaginatedTransactionsResponse(response)) {
//             setTransactions(response.data.content);
//           } else if (isGetTransactionsResponse(response)) {
//             const allTransactions = response.transactions.flatMap(transactionResponse => transactionResponse.data.content);
//             setTransactions(allTransactions);
//           } else {
//             console.error('Unknown response type:', response);
//           }
//         } catch (error) {
//           console.error('Error fetching transactions:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       loadTransactions();
//     }
//   }, [session, activeTab]);

//   const handleTabChange = (tab: string) => {
//     setActiveTab(tab);
//   };

//   return (
//     <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
//       {loading ? (
//         <div className="animate-pulse">
//           {/* Shimmer for Cards Section */}
//           <div className="flex flex-col md:flex-row md:space-x-8 mb-4">
//             <div className="w-full md:w-1/3 lg:w-3/5">
//               <div className="pt-4 flex items-center justify-between">
//                 <div className="h-8 bg-gray-200 rounded-lg w-1/2"></div>
//                 <div className="h-8 bg-gray-200 rounded-lg w-24"></div>
//               </div>
//               <div className="flex overflow-x-auto space-x-6 scrollbar-hide gap-16 mt-4">
//                 <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
//                 <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
//                 <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
//               </div>
//               <div className="mt-4 flex justify-center">
//                 <div className="h-8 bg-gray-200 rounded-lg w-48"></div>
//               </div>
//             </div>

//             {/* Shimmer for BarChart Section */}
//             <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//               <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
//               <div className="h-64 bg-gray-200 rounded-lg"></div>
//             </div>
//           </div>

//           {/* Shimmer for Transactions Section */}
//           <div className="mb-4 md:w-4/5 lg:w-10/12">
//             <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
//             <div className="h-8 bg-gray-200 rounded-lg w-1/4"></div>
//             <div className="h-40 bg-gray-200 rounded-lg mt-4"></div>
//           </div>
//         </div>
//       ) : (
//         <>
//           <div className="mb-4">
//             <div className="flex flex-col md:flex-row md:space-x-8">
//               <div className="w-full md:w-1/3 lg:w-3/5">
//                 <div className="pt-4 flex items-center justify-between">
//                   <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
//                   <button className="px-4 py-2 text-sm font-bold text-[#343C6A] border border-none">
//                     + Add Card
//                   </button>
//                 </div>

//                 <div className="flex overflow-x-auto space-x-6 [&::-webkit-scrollbar]:hidden gap-6 mt-4">
//                   {cards.length > 0 ? (
//                     cards.map((item, index) => (
//                       <div key={item.id} className="flex-shrink-0 w-72">
//                         <Card
//                           balance={$${item.balance}}
//                           cardHolder={item.cardHolder}
//                           validThru={formatDate(item.expiryDate)}
//                           cardNumber="3778 **** **** 1234"
//                           filterClass={index % 2 === 0 ? "" : "filter-black"}
//                           bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "from-white to-gray-200"}
//                           textColor={index % 2 === 0 ? "text-white" : "text-black"}
//                           iconBgColor="bg-opacity-10"
//                           showIcon={true}
//                         />
//                       </div>
//                     ))
//                   ) : (
//                     <div>No cards available</div>
//                   )}
//                 </div>
//               </div>

//               <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//                 <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
//                 <div className="w-full max-h-[200px] flex-grow pt-6">
//                   {session?.user?.access_token && <BarChart token={session.user.access_token} />}
//                 </div>
//               </div>
//             </div>
//           </div>

//           <div className="mb-4 md:w-4/5 lg:w-10/12">
//             <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
//             <Tabs tabs={['All Transactions', 'Income', 'Expense']} activeTab={activeTab} onTabChange={handleTabChange} />
//             <TransactionsList transactions={transactions.map((transaction) => ({
//                 ...transaction,
//                 amount: transaction.amount.toString(),
//               }))}
//             />
//           </div>
//         </>
//       )}
//     </div>
//   );
// };

// export default Page;
"use client";

import React, { useState, useEffect } from 'react';
import { getCards } from '@/lib/api/cardController';
import { getSession } from "next-auth/react";
import Card from '../components/Page2/Card';
import Tabs from '../components/Tabs';
import BarChart from '../components/Page2/BarChart';
import TransactionsList from '../components/Page2/TransactionsList';
import { Card as CardType } from '@/types/cardController.Interface';
import { TransactionData, GetTransactionsResponse, PaginatedTransactionsResponse } from '@/types/transactionController.interface';
import { getTransactions, getTransactionIncomes, getTransactionsExpenses } from '@/lib/api/transactionController';
import { useRouter } from 'next/navigation';
import Refresh from '@/app/api/auth/[...nextauth]/token/RefreshToken'; 
import WhiteCard from '../components/Page2/WhiteCard';

const Page = () => {
  const [activeTab, setActiveTab] = useState('All Transactions');
  const [cards, setCards] = useState<CardType[]>([]);
  const [transactions, setTransactions] = useState<TransactionData[]>([]);
  const [page, setPage] = useState(0);
  const [size] = useState(3);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const [access_token, setAccess_token] = useState("");
  
  const router = useRouter();

  // Fetch session and refresh token only once when the component mounts
  useEffect(() => {
    const fetchSessionAndRefreshToken = async () => {
      setLoading(true);
      try {
        const accessToken = await Refresh();
        console.log("Access Token:", accessToken);
        setAccess_token(accessToken);
      } catch (error) {
        console.error("Error fetching session or refreshing token:", error);
        router.push(`/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
      } finally {
        setLoading(false);
      }
    };

    fetchSessionAndRefreshToken(); // Call it only once when component mounts
  }, [router]); // Empty dependency array with `router`

  useEffect(() => {
    const loadCards = async () => {
      if (access_token) {
        try {
          setLoading(true);
          console.log("Fetching cards...");
          const cardData = await getCards(access_token, page, size);
          if (cardData.content.length > 0) {
            setCards((prevCards) => [
              ...prevCards.filter(card => !cardData.content.some(newCard => newCard.id === card.id)),
              ...cardData.content
            ]);
            setPage((prevPage) => prevPage + 1);
            if (cardData.content.length < size) {
              setHasMore(false);
            }
          } else {
            setHasMore(false);
          }
        } catch (error) {
          console.error('Error fetching cards:', error);
        } finally {
          setLoading(false);
        }
      }
    };

    if (access_token) {
      if (page === 0 && cards.length === 0) {
        loadCards();
      }
    }
  }, [access_token, page, size, cards.length]);

  useEffect(() => {
    const loadTransactions = async () => {
      if (access_token) {
        console.log("Fetching transactions with access token:", access_token);
        try {
          setLoading(true);
          let response: GetTransactionsResponse | PaginatedTransactionsResponse;
          switch (activeTab) {
            case 'Income':
              response = await getTransactionIncomes(0, 100, access_token);
              break;
            case 'Expense':
              response = await getTransactionsExpenses(0, 100, access_token);
              break;
            default:
              response = await getTransactions(0, 100, access_token);
          }
          if (isPaginatedTransactionsResponse(response)) {
            setTransactions(response.data.content);
            console.log("Transaction data:", response.data.content);
          } else if (isGetTransactionsResponse(response)) {
            const allTransactions = response.transactions.flatMap(transactionResponse => transactionResponse.data.content);
            setTransactions(allTransactions);
            console.log("All transactions:", allTransactions);
          } else {
            console.error('Unknown response type:', response);
          }
        } catch (error) {
          console.error('Error fetching transactions:', error);
        } finally {
          setLoading(false);
        }
      }
    };

    if (access_token) {
      loadTransactions();
    }
  }, [access_token, activeTab]);

  const handleTabChange = (tab: string) => {
    setActiveTab(tab);
  };

  return (
    <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
      {loading ? (
        <div className="animate-pulse">
          {/* Shimmer for Cards Section */}
          <div className="flex flex-col md:flex-row md:space-x-8 mb-4">
            <div className="w-full md:w-1/3 lg:w-3/5">
              <div className="pt-4 flex items-center justify-between">
                <div className="h-8 bg-gray-200 rounded-lg w-1/2"></div>
                <div className="h-8 bg-gray-200 rounded-lg w-24"></div>
              </div>
              <div className="flex overflow-x-auto space-x-6 scrollbar-hide gap-16 mt-4">
                <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
                <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
                <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
              </div>
              <div className="mt-4 flex justify-center">
                <div className="h-8 bg-gray-200 rounded-lg w-48"></div>
              </div>
            </div>

            {/* Shimmer for BarChart Section */}
            <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
              <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
              <div className="h-64 bg-gray-200 rounded-lg"></div>
            </div>
          </div>

          {/* Shimmer for Transactions Section */}
          <div className="mb-4 md:w-4/5 lg:w-10/12">
            <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
            <div className="h-8 bg-gray-200 rounded-lg w-1/4"></div>
            <div className="h-40 bg-gray-200 rounded-lg mt-4"></div>
          </div>
        </div>
      ) : (
        <>
          <div className="mb-4">
            <div className="flex flex-col md:flex-row md:space-x-8">
              <div className="w-full md:w-1/3 lg:w-3/5">
                <div className="pt-4 flex items-center justify-between">
                  <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
                  <button className="px-4 py-2 text-sm font-bold text-[#343C6A] border border-none">
                    + Add Card
                  </button>
                </div>

                <div className="flex overflow-x-auto space-x-6 [&::-webkit-scrollbar]:hidden gap-6 mt-4">
                  {cards.length > 0 ? (
                    cards.map((item, index) => (
                      <div key={item.id} className="flex-shrink-0 w-72">
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
                    ))
                  ) : (
                    <div>No cards available</div>
                  )}
                </div>
              </div>

              <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
                <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
                <div className="w-full max-h-[200px] flex-grow pt-6">
                  {access_token && <BarChart token={access_token} />}
                </div>
              </div>
            </div>
          </div>

          <div className="mb-4 md:w-4/5 lg:w-10/12">
            <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
            <Tabs tabs={['All Transactions', 'Income', 'Expense']} activeTab={activeTab} onTabChange={handleTabChange} />
            <TransactionsList transactions={transactions.map((transaction) => ({
                ...transaction,
                amount: transaction.amount.toString(),
              }))}
            />
          </div>
        </>
      )}
    </div>
  );
};

export default Page;



// "use client";

// import React, { useState, useEffect } from 'react';
// import { getCards } from '@/lib/api/cardController';
// import { getSession } from "next-auth/react";
// import Card from '../components/Page2/Card';
// import Tabs from '../components/Tabs';
// import BarChart from '../components/Page2/BarChart';
// import TransactionsList from '../components/Page2/TransactionsList';
// import { Card as CardType } from '@/types/cardController.Interface';
// import { TransactionData, GetTransactionsResponse, PaginatedTransactionsResponse } from '@/types/transactionController.interface';
// import { getTransactions, getTransactionIncomes, getTransactionsExpenses } from '@/lib/api/transactionController';
// import { useRouter } from 'next/navigation';
// import Refresh from '@/app/api/auth/[...nextauth]/token/RefreshToken'; 
// import WhiteCard from '../components/Page2/WhiteCard';

// type Data = {
//   access_token: string;
//   data: string;
//   refresh_token: string;
// };

// type SessionDataType = {
//   user: Data;
// };

// const formatDate = (date: string): string => {
//   const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: '2-digit' };
//   return new Date(date).toLocaleDateString('en-US', options);
// };

// const isPaginatedTransactionsResponse = (response: GetTransactionsResponse | PaginatedTransactionsResponse): response is PaginatedTransactionsResponse => {
//   return (response as PaginatedTransactionsResponse).data !== undefined;
// };

// const isGetTransactionsResponse = (response: GetTransactionsResponse | PaginatedTransactionsResponse): response is GetTransactionsResponse => {
//   return (response as GetTransactionsResponse).transactions !== undefined;
// };

// const Page = () => {
//   const [activeTab, setActiveTab] = useState('All Transactions');
//   const [cards, setCards] = useState<CardType[]>([]);
//   const [transactions, setTransactions] = useState<TransactionData[]>([]);
//   const [page, setPage] = useState(0);
//   const [size] = useState(3);
//   const [loading, setLoading] = useState(false);
//   const [hasMore, setHasMore] = useState(true);
//   const [session, setSession] = useState<SessionDataType | null>(null);

//   const router = useRouter();

//   useEffect(() => {
//     const fetchSessionAndRefreshToken = async () => {
//       setLoading(true);
//       try {
//         const accessToken = await Refresh(); 
//         const sessionData = (await getSession()) as SessionDataType | null;
//         if (sessionData && sessionData.user) {
//           setSession({ user: { ...sessionData.user, access_token: accessToken } });
//         } else {
//           router.push(`/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
//         }
//       } catch (error) {
//         console.error("Error fetching session or refreshing token:", error);
//         router.push(`/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
//       } finally {
//         setLoading(false);
//       }
//     };

//     fetchSessionAndRefreshToken();
//   }, [router]);

//   useEffect(() => {
//     const loadCards = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           const cardData = await getCards(session.user.access_token, page, size);
//           if (cardData.content.length > 0) {
//             setCards((prevCards) => [
//               ...prevCards.filter(card => !cardData.content.some(newCard => newCard.id === card.id)),
//               ...cardData.content
//             ]);
//             setPage((prevPage) => prevPage + 1);
//             if (cardData.content.length < size) {
//               setHasMore(false);
//             }
//           } else {
//             setHasMore(false);
//           }
//         } catch (error) {
//           console.error('Error fetching cards:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       if (page === 0 && cards.length === 0) {
//         loadCards();
//       }
//     }
//   }, [session, page, size, cards.length]);

//   useEffect(() => {
//     const loadTransactions = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           let response: GetTransactionsResponse | PaginatedTransactionsResponse;
//           switch (activeTab) {
//             case 'Income':
//               response = await getTransactionIncomes(0, 100, session.user.access_token);
//               break;
//             case 'Expense':
//               response = await getTransactionsExpenses(0, 100, session.user.access_token);
//               break;
//             default:
//               response = await getTransactions(0, 100, session.user.access_token);
//           }
//           if (isPaginatedTransactionsResponse(response)) {
//             setTransactions(response.data.content);
//           } else if (isGetTransactionsResponse(response)) {
//             const allTransactions = response.transactions.flatMap(transactionResponse => transactionResponse.data.content);
//             setTransactions(allTransactions);
//           } else {
//             console.error('Unknown response type:', response);
//           }
//         } catch (error) {
//           console.error('Error fetching transactions:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       loadTransactions();
//     }
//   }, [session, activeTab]);

//   const handleTabChange = (tab: string) => {
//     setActiveTab(tab);
//   };

//   return (
//     <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
//       {loading ? (
//         <div className="animate-pulse">
//           {/* Shimmer for Cards Section */}
//           <div className="flex flex-col md:flex-row md:space-x-8 mb-4">
//             <div className="w-full md:w-1/3 lg:w-3/5">
//               <div className="pt-4 flex items-center justify-between">
//                 <div className="h-8 bg-gray-200 rounded-lg w-1/2"></div>
//                 <div className="h-8 bg-gray-200 rounded-lg w-24"></div>
//               </div>
//               <div className="flex overflow-x-auto space-x-6 scrollbar-hide gap-16 mt-4">
//                 <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
//                 <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
//                 <div className="w-72 h-40 bg-gray-200 rounded-lg"></div>
//               </div>
//               <div className="mt-4 flex justify-center">
//                 <div className="h-8 bg-gray-200 rounded-lg w-48"></div>
//               </div>
//             </div>

//             {/* Shimmer for BarChart Section */}
//             <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//               <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
//               <div className="h-64 bg-gray-200 rounded-lg"></div>
//             </div>
//           </div>

//           {/* Shimmer for Transactions Section */}
//           <div className="mb-4 md:w-4/5 lg:w-10/12">
//             <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
//             <div className="h-8 bg-gray-200 rounded-lg w-1/4"></div>
//             <div className="h-40 bg-gray-200 rounded-lg mt-4"></div>
//           </div>
//         </div>
//       ) : (
//         <>
//           <div className="mb-4">
//             <div className="flex flex-col md:flex-row md:space-x-8">
//               <div className="w-full md:w-1/3 lg:w-3/5">
//                 <div className="pt-4 flex items-center justify-between">
//                   <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
//                   <button className="px-4 py-2 text-sm font-bold text-[#343C6A] border border-none">
//                     + Add Card
//                   </button>
//                 </div>

//                 <div className="flex overflow-x-auto space-x-6 [&::-webkit-scrollbar]:hidden gap-6 mt-4">
//                   <WhiteCard />
//                   <WhiteCard />
//                   <WhiteCard />
//                   {cards.slice(0, 2).map((item, index) => (  // Slicing to show only 2 cards
//                     <div key={item.id} className="flex-shrink-0 w-72">
//                       <Card
//                         balance={`$${item.balance}`}
//                         cardHolder={item.cardHolder}
//                         validThru={formatDate(item.expiryDate)}
//                         cardNumber="3778 **** **** 1234"
//                         filterClass={index % 2 === 0 ? "" : "filter-black"}
//                         bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "from-white to-gray-200"}
//                         textColor={index % 2 === 0 ? "text-white" : "text-black"}
//                         iconBgColor="bg-opacity-10"
//                         showIcon={true}
//                       />
//                     </div>
//                   ))}
//                 </div>
//               </div>

//               <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//                 <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
//                 <div className="w-full max-h-[200px] flex-grow pt-6">
//                   {session?.user?.access_token && <BarChart token={session.user.access_token} />}
//                 </div>
//               </div>
//             </div>
//           </div>

//           <div className="mb-4 md:w-4/5 lg:w-10/12">
//             <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
//             <Tabs tabs={['All Transactions', 'Income', 'Expense']} activeTab={activeTab} onTabChange={handleTabChange} />
//             <TransactionsList transactions={transactions.map((transaction) => ({
//                 ...transaction,
//                 amount: transaction.amount.toString(),
//               }))}
//             />
//           </div>
//         </>
//       )}
//     </div>
//   );
// };

// export default Page;



// "use client";

// import React, { useState, useEffect } from 'react';
// import { getCards } from '@/lib/api/cardController';
// import { getSession } from "next-auth/react";
// import Card from '../components/Page2/Card';
// import Tabs from '../components/Tabs';
// import BarChart from '../components/Page2/BarChart';
// import TransactionsList from '../components/Page2/TransactionsList';
// import { Card as CardType } from '@/types/cardController.Interface';
// import { TransactionData, GetTransactionsResponse, PaginatedTransactionsResponse } from '@/types/transactionController.interface';
// import { getTransactions, getTransactionIncomes, getTransactionsExpenses } from '@/lib/api/transactionController';
// import { useRouter } from 'next/navigation';
// import Refresh from '@/app/api/auth/[...nextauth]/token/RefreshToken'; 
// import WhiteCard from '../components/Page2/WhiteCard';
// import AddCardForm from '../creditCards/AddCardForm';
// import { Button } from "@/components/ui/button";
// import {
//   Dialog,
//   DialogContent,
//   DialogFooter,
//   DialogHeader,
//   DialogTitle,
//    DialogDescription,
//   DialogTrigger,
// } from "@/components/ui/dialog";

// type Data = {
//   access_token: string;
//   data: string;
//   refresh_token: string;
// };

// type SessionDataType = {
//   user: Data;
// };

// const formatDate = (date: string): string => {
//   const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: '2-digit' };
//   return new Date(date).toLocaleDateString('en-US', options);
// };

// const isPaginatedTransactionsResponse = (response: GetTransactionsResponse | PaginatedTransactionsResponse): response is PaginatedTransactionsResponse => {
//   return (response as PaginatedTransactionsResponse).data !== undefined;
// };

// const isGetTransactionsResponse = (response: GetTransactionsResponse | PaginatedTransactionsResponse): response is GetTransactionsResponse => {
//   return (response as GetTransactionsResponse).transactions !== undefined;
// };

// const Page = () => {
//   const [activeTab, setActiveTab] = useState('All Transactions');
//   const [cards, setCards] = useState<CardType[]>([]);
//   const [transactions, setTransactions] = useState<TransactionData[]>([]);
//   const [page, setPage] = useState(0);
//   const [size] = useState(3);
//   const [loading, setLoading] = useState(false);
//   const [hasMore, setHasMore] = useState(true);
//   const [session, setSession] = useState<SessionDataType | null>(null);

//   const [dialogOpen, setDialogOpen] = useState(false);

//   const router = useRouter();

//   // useEffect(() => {
//   //   const fetchSessionAndRefreshToken = async () => {
//   //     setLoading(true);
//   //     try {
//   //       const accessToken = await Refresh(); 
//   //       const sessionData = (await getSession()) as SessionDataType | null;
//   //       if (sessionData && sessionData.user) {
//   //         setSession({ user: { ...sessionData.user, access_token: accessToken } });
//   //       } else {
//   //         router.push(`/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
//   //       }
//   //     } catch (error) {
//   //       console.error("Error fetching session or refreshing token:", error);
//   //       router.push(`/api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
//   //     } finally {
//   //       setLoading(false);
//   //     }
//   //   };

//   //   fetchSessionAndRefreshToken();
//   // }, [router]);

//   useEffect(() => {
//     const loadCards = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           const cardData = await getCards(session.user.access_token, page, size);
//           if (cardData.content.length > 0) {
//             setCards((prevCards) => [
//               ...prevCards.filter(card => !cardData.content.some(newCard => newCard.id === card.id)),
//               ...cardData.content
//             ]);
//             setPage((prevPage) => prevPage + 1);
//             if (cardData.content.length < size) {
//               setHasMore(false);
//             }
//           } else {
//             setHasMore(false);
//           }
//         } catch (error) {
//           console.error('Error fetching cards:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       if (page === 0 && cards.length === 0) {
//         loadCards();
//       }
//     }
//   }, [session, page, size, cards.length]);

//   useEffect(() => {
//     const loadTransactions = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           let response: GetTransactionsResponse | PaginatedTransactionsResponse;
//           switch (activeTab) {
//             case 'Income':
//               response = await getTransactionIncomes(0, 100, session.user.access_token);
//               break;
//             case 'Expense':
//               response = await getTransactionsExpenses(0, 100, session.user.access_token);
//               break;
//             default:
//               response = await getTransactions(0, 100, session.user.access_token);
//           }
//           if (isPaginatedTransactionsResponse(response)) {
//             setTransactions(response.data.content);
//           } else if (isGetTransactionsResponse(response)) {
//             const allTransactions = response.transactions.flatMap(transactionResponse => transactionResponse.data.content);
//             setTransactions(allTransactions);
//           } else {
//             console.error('Unknown response type:', response);
//           }
//         } catch (error) {
//           console.error('Error fetching transactions:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       loadTransactions();
//     }
//   }, [session, activeTab]);

//   const handleTabChange = (tab: string) => {
//     setActiveTab(tab);
//   };

//   const handleLoadMore = () => {
//     if (hasMore) {
//       setPage(prevPage => prevPage + 1);
//     }
//   };

//   const handleAddition = (newCard: CardType) => {
//     setCards(prevCards => [...prevCards, newCard]);
//   };

//   return (
//     <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
//       {loading ? (
//   <div className="animate-pulse">
//     {/* Shimmer for Cards Section */}
//     <div className="flex flex-col md:flex-row md:space-x-8 mb-8">
//       <div className="w-full md:w-1/3 lg:w-3/5">
//         <div className="pt-4 flex items-center justify-between">
//           <div className="h-8 bg-gray-200 rounded-lg w-1/2"></div>
//           <div className="h-8 bg-gray-200 rounded-lg w-24"></div>
//         </div>
//         <div className="flex overflow-x-auto space-x-6 scrollbar-hide gap-4 mt-4">
//           <div className="w-72 h-44 bg-gray-200 rounded-lg"></div>
//           <div className="w-72 h-44 bg-gray-200 rounded-lg"></div>
//           <div className="w-72 h-44 bg-gray-200 rounded-lg"></div>
//         </div>
//         <div className="mt-4 flex justify-center">
//           <div className="h-10 bg-gray-200 rounded-lg w-48"></div>
//         </div>
//       </div>

//       {/* Shimmer for BarChart Section */}
//       <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//         <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
//         <div className="h-64 bg-gray-200 rounded-lg"></div>
//       </div>
//     </div>

//     {/* Shimmer for Transactions Section */}
//     <div className="mb-4 md:w-4/5 lg:w-10/12">
//       <div className="h-8 bg-gray-200 rounded-lg mb-4"></div>
//       <div className="h-8 bg-gray-200 rounded-lg w-1/4 mb-4"></div>
//       <div className="h-44 bg-gray-200 rounded-lg mt-4"></div>
//     </div>
//   </div>) : (
//         <>
//           <div className="mb-4">
//             <div className="flex flex-col md:flex-row md:space-x-8">
//               <div className="w-full md:w-1/3 lg:w-3/5">
//                 <div className="pt-4 flex items-center justify-between">
//                   <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
                 
//       <Dialog isOpen={dialogOpen} onClose={() => setDialogOpen(false)}>
//         <DialogContent className="sm:max-w-[425px]">
//           <DialogHeader>
//             <DialogTitle>Add New Card</DialogTitle>
//             <DialogDescription>
//               Fill in the details to add a new card.
//             </DialogDescription>
//           </DialogHeader>
//           <AddCardForm access_token={session?.user.access_token!} handleAddition={handleAddition} />
//         </DialogContent>
//         <DialogFooter>
//           <Button onClick={() => setDialogOpen(false)}>Close</Button>
//         </DialogFooter>
//       </Dialog>
      
//       <DialogTrigger onClick={() => setDialogOpen(true)}>
//         <Button variant="outline">+ Add Card</Button>
//       </DialogTrigger>
//                 </div>

//                 <div className="flex overflow-x-auto space-x-6 [&::-webkit-scrollbar]:hidden gap-6 mt-4">
//                   <WhiteCard/>
//                   <WhiteCard/>
//                   <WhiteCard/>
//                   {cards.slice(0, page * 2 + 2).map((item, index) => (
//                     <div key={item.id} className="flex-shrink-0 w-72">
//                       <Card
//                         balance={`$${item.balance}`}
//                         cardHolder={item.cardHolder}
//                         validThru={formatDate(item.expiryDate)}
//                         cardNumber="3778 **** **** 1234"
//                         filterClass={index % 2 === 0 ? "" : "filter-black"}
//                         bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "from-white to-gray-200"}
//                         textColor={index % 2 === 0 ? "text-white" : "text-black"}
//                         iconBgColor="bg-opacity-10"
//                         showIcon={true}
//                       />
//                     </div>
//                   ))}
//                 </div>
//                 {hasMore && (
//                   <div className="mt-4 flex justify-center">
//                     <Button variant="outline" onClick={handleLoadMore}>
//                       Load More Cards
//                     </Button>
//                   </div>
//                 )}
//               </div>

//               <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//                 <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
//                 <div className="w-full max-h-[200px] flex-grow pt-6">
//                   {session?.user?.access_token && <BarChart token={session.user.access_token} />}
//                 </div>
//               </div>
//             </div>
//           </div>

//           <div className="mb-4 md:w-4/5 lg:w-10/12">
//             <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
//             <Tabs tabs={['All Transactions', 'Income', 'Expense']} activeTab={activeTab} onTabChange={handleTabChange} />
//             <TransactionsList transactions={transactions.map((transaction) => ({
//                 ...transaction,
//                 amount: transaction.amount.toString(),
//               }))}
//             />
//           </div>
//         </>
//       )}
//     </div>
//   );
// };

// export default Page;



// "use client";

// import React, { useState, useEffect } from 'react';
// import { getCards } from '@/lib/api/cardController';
// import { getSession } from "next-auth/react";
// import Card from '../components/Page2/Card';
// import Tabs from '../components/Tabs';
// import BarChart from '../components/Page2/BarChart';
// import TransactionsList from '../components/Page2/TransactionsList';
// import { Card as CardType } from '@/types/cardController.Interface';
// import { TransactionData, GetTransactionsResponse, PaginatedTransactionsResponse } from '@/types/transactionController.interface';
// import { getTransactions, getTransactionIncomes, getTransactionsExpenses } from '@/lib/api/transactionController';
// import { useRouter } from 'next/navigation';
// import Refresh from '@/app/api/auth/[...nextauth]/token/RefreshToken'; 
// import WhiteCard from '../components/Page2/WhiteCard';
// import AddCardForm from '../creditCards/AddCardForm';
// import { Button } from "@/components/ui/button";
// import {
//   Dialog,
//   DialogContent,
//   DialogFooter,
//   DialogHeader,
//   DialogTitle,
//   DialogDescription,
//   DialogTrigger,
// } from "@/components/ui/dialog";

// type Data = {
//   access_token: string;
//   data: string;
//   refresh_token: string;
// };

// type SessionDataType = {
//   user: Data;
// };

// const formatDate = (date: string): string => {
//   const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: '2-digit' };
//   return new Date(date).toLocaleDateString('en-US', options);
// };

// const Page = () => {
//   const [activeTab, setActiveTab] = useState('All Transactions');
//   const [cards, setCards] = useState<CardType[]>([]);
//   const [transactions, setTransactions] = useState<TransactionData[]>([]);
//   const [page, setPage] = useState(0);
//   const [size] = useState(3);
//   const [loading, setLoading] = useState(false);
//   const [hasMore, setHasMore] = useState(true);
//   const [session, setSession] = useState<SessionDataType | null>(null);

//   const [dialogOpen, setDialogOpen] = useState(false);

//   const router = useRouter();

//   useEffect(() => {
//     const loadCards = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           const cardData = await getCards(session.user.access_token, page, size);
//           if (cardData.content.length > 0) {
//             setCards((prevCards) => [
//               ...prevCards.filter(card => !cardData.content.some(newCard => newCard.id === card.id)),
//               ...cardData.content
//             ]);
//             setPage((prevPage) => prevPage + 1);
//             if (cardData.content.length < size) {
//               setHasMore(false);
//             }
//           } else {
//             setHasMore(false);
//           }
//         } catch (error) {
//           console.error('Error fetching cards:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       if (page === 0 && cards.length === 0) {
//         loadCards();
//       }
//     }
//   }, [session, page, size, cards.length]);

//   useEffect(() => {
//     const loadTransactions = async () => {
//       if (session?.user?.access_token) {
//         try {
//           setLoading(true);
//           let response: GetTransactionsResponse | PaginatedTransactionsResponse;
//           switch (activeTab) {
//             case 'Income':
//               response = await getTransactionIncomes(0, 100, session.user.access_token);
//               break;
//             case 'Expense':
//               response = await getTransactionsExpenses(0, 100, session.user.access_token);
//               break;
//             default:
//               response = await getTransactions(0, 100, session.user.access_token);
//           }
//           if ('data' in response) {
//             setTransactions(response.data.content);
//           } else {
//             const allTransactions = response.transactions.flatMap(transactionResponse => transactionResponse.data.content);
//             setTransactions(allTransactions);
//           }
//         } catch (error) {
//           console.error('Error fetching transactions:', error);
//         } finally {
//           setLoading(false);
//         }
//       }
//     };

//     if (session?.user?.access_token) {
//       loadTransactions();
//     }
//   }, [session, activeTab]);

//   const handleTabChange = (tab: string) => {
//     setActiveTab(tab);
//   };

//   const handleLoadMore = () => {
//     if (hasMore) {
//       setPage(prevPage => prevPage + 1);
//     }
//   };

//   const handleAddition = (newCard: CardType) => {
//     setCards(prevCards => [...prevCards, newCard]);
//   };

//   return (
//     <div className="bg-[#f5f7fa] dark:bg-[#121212] py-4 px-8 max-w-full">
//       {loading ? (
//         <div className="animate-pulse">
//           <div className="flex flex-col md:flex-row md:space-x-8 mb-8">
//             <div className="w-full md:w-1/3 lg:w-3/5">
//               <div className="pt-4 flex items-center justify-between">
//                 <div className="h-8 bg-gray-200 dark:bg-[#333B69] rounded-lg w-1/2"></div>
//                 <div className="h-8 bg-gray-200 dark:bg-[#333B69] rounded-lg w-24"></div>
//               </div>
//               <div className="flex overflow-x-auto space-x-6 scrollbar-hide gap-4 mt-4">
//                 <div className="w-72 h-44 bg-gray-200 dark:bg-[#333B69] rounded-lg"></div>
//                 <div className="w-72 h-44 bg-gray-200 dark:bg-[#333B69] rounded-lg"></div>
//                 <div className="w-72 h-44 bg-gray-200 dark:bg-[#333B69] rounded-lg"></div>
//               </div>
//               <div className="mt-4 flex justify-center">
//                 <div className="h-10 bg-gray-200 dark:bg-[#333B69] rounded-lg w-48"></div>
//               </div>
//             </div>

//             <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//               <div className="h-8 bg-gray-200 dark:bg-[#333B69] rounded-lg mb-4"></div>
//               <div className="h-64 bg-gray-200 dark:bg-[#333B69] rounded-lg"></div>
//             </div>
//           </div>

//           <div className="mb-4 md:w-4/5 lg:w-10/12">
//             <div className="h-8 bg-gray-200 dark:bg-[#333B69] rounded-lg mb-4"></div>
//             <div className="h-8 bg-gray-200 dark:bg-[#333B69] rounded-lg w-1/4 mb-4"></div>
//             <div className="h-44 bg-gray-200 dark:bg-[#333B69] rounded-lg mt-4"></div>
//           </div>
//         </div>
//       ) : (
//         <>
//           <div className="mb-4">
//             <div className="flex flex-col md:flex-row md:space-x-8">
//               <div className="w-full md:w-1/3 lg:w-3/5">
//                 <div className="pt-4 flex items-center justify-between">
//                   <h2 className="text-xl font-bold text-[#343C6A] dark:text-white">My Cards</h2>
                 
//                   <Dialog isOpen={dialogOpen} onClose={() => setDialogOpen(false)}>
//                     <DialogContent className="sm:max-w-[425px]">
//                       <DialogHeader>
//                         <DialogTitle>Add New Card</DialogTitle>
//                         <DialogDescription>
//                           Fill in the details to add a new card.
//                         </DialogDescription>
//                       </DialogHeader>
//                       <AddCardForm access_token={session?.user.access_token!} handleAddition={handleAddition} />
//                     </DialogContent>
//                     <DialogFooter>
//                       <Button onClick={() => setDialogOpen(false)}>Close</Button>
//                     </DialogFooter>
//                   </Dialog>
                  
//                   <DialogTrigger onClick={() => setDialogOpen(true)}>
//                     <Button variant="outline">+ Add Card</Button>
//                   </DialogTrigger>
//                 </div>

//                 <div className="flex overflow-x-auto space-x-6 [&::-webkit-scrollbar]:hidden gap-6 mt-4">
//                   <WhiteCard/>
//                   <WhiteCard/>
//                   <WhiteCard/>
//                   {cards.slice(0, page * 2 + 2).map((item, index) => (
//                     <div key={item.id} className="flex-shrink-0 w-72">
//                       <Card
//                         balance={`$${item.balance}`}
//                         cardHolder={item.cardHolder}
//                         validThru={formatDate(item.expiryDate)}
//                         cardNumber="3778 **** **** 1234"
//                         filterClass={index % 2 === 0 ? "" : "filter-black"}
//                         bgColor={index % 2 === 0 ? "from-[#4C49ED] to-[#0A06F4]" : "from-white to-gray-200 dark:from-[#333B69] dark:to-[#555B85]"}
//                         textColor={index % 2 === 0 ? "text-white" : "text-black dark:text-white"}
//                         iconBgColor="bg-opacity-10"
//                         showIcon={true}
//                       />
//                     </div>
//                   ))}
//                 </div>
//                 {hasMore && (
//                   <div className="mt-4 flex justify-center">
//                     <Button variant="outline" onClick={handleLoadMore}>
//                       Load More Cards
//                     </Button>
//                   </div>
//                 )}
//               </div>

//               <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//                 <h2 className="text-xl font-bold text-[#343C6A] dark:text-white">My Expense</h2>
//                 <div className="w-full max-h-[200px] flex-grow pt-6">
//                   {session?.user?.access_token && <BarChart token={session.user.access_token} />}
//                 </div>
//               </div>
//             </div>
//           </div>

//           <div className="mb-4 md:w-4/5 lg:w-10/12">
//             <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A] dark:text-white">Recent Transactions</h2>
//             <Tabs tabs={['All Transactions', 'Income', 'Expense']} activeTab={activeTab} onTabChange={handleTabChange} />
//             <TransactionsList transactions={transactions.map((transaction) => ({
//                 ...transaction,
//                 amount: transaction.amount.toString(),
//               }))}
//             />
//           </div>
//         </>
//       )}
//     </div>
//   );
// };

// export default Page;
