// "use client";
// import React, { useEffect, useState } from 'react';
// import { useSession } from 'next-auth/react';
// import TransactionsList from '../components/Page2/TransactionsList';
// import WhiteCard from '../components/Page2/WhiteCard';
// import CreditCard from '../components/Page2/Card';
// import Tabs from '../components/Tabs';
// import BarChart from '../components/Page2/BarChart';
// import { getTransactions, getIncomes, getExpenses } from '../../transactionController';
// import { getAllCards } from '@/cardController';
// import { TransactionData } from '../../transactionController';

// interface CardData {
//     balance: string;
//     cardHolder: string;
//     validThru: string;
//     cardNumber: string;
//     bgColor: string;
//     textColor: string;
// }

// const Page = () => {
//     const { data: session } = useSession(); // Using NextAuth's useSession hook
//     const [transactions, setTransactions] = useState<TransactionData[]>([]);
//     const [cards, setCards] = useState<CardData[]>([]);
//     const [activeTab, setActiveTab] = useState('All Transactions');
//     const [loading, setLoading] = useState(false);

//     const handleTabChange = (tab: string) => {
//         setActiveTab(tab);
//     };

//     useEffect(() => {
//         const fetchData = async () => {
//             if (!session?.access_token) {
//                 console.error('No access token found');
//                 return;
//             }

//             setLoading(true);
//             try {
//                 const config = {
//                     headers: {
//                         Authorization: `Bearer ${session.access_token}`,  // Using NextAuth's token
//                     },
//                 };

//                 let response;
//                 if (activeTab === 'All Transactions') {
//                     response = await getTransactions(config);
//                 } else if (activeTab === 'Income') {
//                     response = await getIncomes(config);
//                 } else if (activeTab === 'Expense') {
//                     response = await getExpenses(config);
//                 }

//                 if (response && response.success) {
//                     setTransactions(response.data);
//                 } else {
//                     setTransactions([]);
//                 }
//             } catch (error) {
//                 console.error('Error fetching data', error);
//                 setTransactions([]);
//             } finally {
//                 setLoading(false);
//             }
//         };

//         fetchData();
//     }, [activeTab, session?.access_token]);

//     useEffect(() => {
//         const fetchCardData = async () => {
//             if (!session?.access_token) {
//                 console.error('No access token found');
//                 return;
//             }

//             setLoading(true);
//             try {
//                 const config = {
//                     headers: {
//                         Authorization: `Bearer ${session.access_token}`,  // Using NextAuth's token
//                     },
//                 };

//                 const response = await getAllCards(config);

//                 if (response && response.success) {
//                     setCards(response.data);
//                 } else {
//                     setCards([]);
//                 }
//             } catch (error) {
//                 console.error('Error fetching card data', error);
//                 setCards([]);
//             } finally {
//                 setLoading(false);
//             }
//         };

//         fetchCardData();
//     }, [session?.access_token]);

//     return (
//         <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
//             <div className="mb-4">
//                 <div className='flex flex-col md:flex-row md:space-x-8 '>
//                     <div className="w-full md:w-1/3 lg:w-3/5">
//                         <div className="pt-4 flex items-center justify-between">
//                             <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
//                             <button className="px-4 py-2 text-sm font-bold text-[#343C6A] border border-none">
//                                 + Add Card
//                             </button>
//                         </div>

//                         <div className='flex overflow-x-auto space-x-6 scrollbar-hide gap-16 [&::-webkit-scrollbar]:hidden'>
//                             {cards.map((card, index) => (
//                                 <div key={index} className="flex-shrink-0 w-72">
//                                     <CreditCard
//                                         balance={card.balance}
//                                         cardHolder={card.cardHolder}
//                                         validThru={card.validThru}
//                                         cardNumber={card.cardNumber}
//                                         filterClass=""
//                                         bgColor={card.bgColor}
//                                         textColor={card.textColor}
//                                         iconBgColor="bg-opacity-10"
//                                         showIcon={true}
//                                     />
//                                 </div>
//                             ))}
//                             <div className="flex-shrink-0 w-72">
//                                 <WhiteCard />
//                             </div>
//                         </div>
//                     </div>
//                     <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//                         <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
//                         <div className="w-full max-h-[200px] flex-grow pt-6">
//                             <BarChart />
//                         </div>
//                     </div>
//                 </div>
//             </div>
//             <div className="mb-4 md:w-4/5 lg:w-10/12">
//                 <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
//                 <Tabs
//                     tabs={['All Transactions', 'Income', 'Expense']}
//                     activeTab={activeTab}
//                     onTabChange={handleTabChange}
//                 />
//                 {loading ? (
//                     <div>Loading...</div>
//                 ) : (
//                     <TransactionsList transactions={transactions} />
//                 )}
//             </div>
//         </div>
//     );
// };

// export default Page;


"use client";
import React, { useEffect, useState } from 'react';
import TransactionsList from '../components/Page2/TransactionsList';
import WhiteCard from '../components/Page2/WhiteCard';
import CreditCard from '../components/Page2/Card';
import Tabs from '../components/Tabs';
import BarChart from '../components/Page2/BarChart';
import { getTransactions, getIncomes, getExpenses } from '../../transactionController';
import { getAllCards } from '@/cardController';
import { TransactionData } from '../../transactionController';

interface CardData {
    balance: string;
    cardHolder: string;
    validThru: string;
    cardNumber: string;
    bgColor: string;
    textColor: string;
}

const Page = () => {
    const [transactions, setTransactions] = useState<TransactionData[]>([]);
    const [cards, setCards] = useState<CardData[]>([]);
    const [activeTab, setActiveTab] = useState('All Transactions');
    const [loading, setLoading] = useState(false);

    const handleTabChange = (tab: string) => {
        setActiveTab(tab);
    };

    useEffect(() => {
        const fetchData = async () => {
            const accessToken = 'your-hardcoded-access-token';  // Replace with your token

            setLoading(true);
            try {
                const config = {
                    headers: {
                        Authorization: `Bearer ${accessToken}`,  // Hardcoded token
                    },
                };

                let response;
                if (activeTab === 'All Transactions') {
                    response = await getTransactions(config);
                } else if (activeTab === 'Income') {
                    response = await getIncomes(config);
                } else if (activeTab === 'Expense') {
                    response = await getExpenses(config);
                }

                if (response && response.success) {
                    setTransactions(response.data);
                } else {
                    setTransactions([]);
                }
            } catch (error) {
                console.error('Error fetching data', error);
                setTransactions([]);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [activeTab]);

    useEffect(() => {
        const fetchCardData = async () => {
            const accessToken = 'your-hardcoded-access-token';  // Replace with your token

            setLoading(true);
            try {
                const config = {
                    headers: {
                        Authorization: `Bearer ${accessToken}`,  // Hardcoded token
                    },
                };

                const response = await getAllCards(config);

                if (response && response.success) {
                    setCards(response.data);
                } else {
                    setCards([]);
                }
            } catch (error) {
                console.error('Error fetching card data', error);
                setCards([]);
            } finally {
                setLoading(false);
            }
        };

        fetchCardData();
    }, []);

    return (
        <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
            <div className="mb-4">
                <div className='flex flex-col md:flex-row md:space-x-8 '>
                    <div className="w-full md:w-1/3 lg:w-3/5">
                        <div className="pt-4 flex items-center justify-between">
                            <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
                            <button className="px-4 py-2 text-sm font-bold text-[#343C6A] border border-none">
                                + Add Card
                            </button>
                        </div>

                        <div className='flex overflow-x-auto space-x-6 scrollbar-hide gap-16 [&::-webkit-scrollbar]:hidden'>
                            {cards.map((card, index) => (
                                <div key={index} className="flex-shrink-0 w-72">
                                    <CreditCard
                                        balance={card.balance}
                                        cardHolder={card.cardHolder}
                                        validThru={card.validThru}
                                        cardNumber={card.cardNumber}
                                        filterClass=""
                                        bgColor={card.bgColor}
                                        textColor={card.textColor}
                                        iconBgColor="bg-opacity-10"
                                        showIcon={true}
                                    />
                                </div>
                            ))}
                            <div className="flex-shrink-0 w-72">
                                <WhiteCard />
                            </div>
                        </div>
                    </div>
                    <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
                        <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
                        <div className="w-full max-h-[200px] flex-grow pt-6">
                            <BarChart />
                        </div>
                    </div>
                </div>
            </div>
            <div className="mb-4 md:w-4/5 lg:w-10/12">
                <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
                <Tabs
                    tabs={['All Transactions', 'Income', 'Expense']}
                    activeTab={activeTab}
                    onTabChange={handleTabChange}
                />
                {loading ? (
                    <div>Loading...</div>
                ) : (
                    <TransactionsList transactions={transactions} />
                )}
            </div>
        </div>
    );
};

export default Page;




// eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJtaWhlcmV0IiwiaWF0IjoxNzI0MTUxNjU3LCJleHAiOjE3MjQyMzgwNTd9.vLdFTv3vuvVASXepVATvawmuxuFEwXXjD-lWKnl-qhBem-gd6-u5TvNPg2xf7KbY




// "use client";
// import React, { useEffect, useState } from 'react';
// import { useRouter } from 'next/router'; 
// import TransactionsList from '../components/Page2/TransactionsList';
// import WhiteCard from '../components/Page2/WhiteCard';
// import CreditCard from '../components/Page2/Card';  
// import Tabs from '../components/Tabs';
// import BarChart from '../components/Page2/BarChart';
// import { getTransactions, getIncomes, getExpenses } from '../../transactionController';
// import { TransactionData } from '../../transactionController';
// import { getSession } from '../../authController'; 
// const Page = () => {
//   const [transactions, setTransactions] = useState<TransactionData[]>([]);
//   const [activeTab, setActiveTab] = useState('All Transactions');
//   const [loading, setLoading] = useState(true); 
//   const [session, setSession] = useState(false); 
//   const route = useRouter(); 

//   useEffect(() => {
//     const fetchSession = async () => {
//       const sessionData = await getSession();
//       if (sessionData?.user) {
//         setSession(true);
//       } else {
//         route.push(`/api/auth/signin?callbackUrl=${encodeURIComponent('/accounts')}`);
//       }
//       setLoading(false); 
//     };

//     fetchSession();
//   }, [route]);

//   useEffect(() => {
//     const fetchData = async () => {
//       if (!session) return; 
      
//       setLoading(true);
//       try {
//         let response;
//         if (activeTab === 'All Transactions') {
//           response = await getTransactions();
//         } else if (activeTab === 'Income') {
//           response = await getIncomes();
//         } else if (activeTab === 'Expense') {
//           response = await getExpenses();
//         }

//         if (response && response.success) {
//           setTransactions(response.data); 
//         } else {
//           setTransactions([]); 
//         }
//       } catch (error) {
//         console.error('Error fetching data', error);
//         setTransactions([]); 
//       } finally {
//         setLoading(false);
//       }
//     };

//     fetchData();
//   }, [activeTab, session]); 

//   if (loading) return <div>Loading...</div>;

//   return (
//     <div className="bg-[#f5f7fa] py-4 px-8 max-w-full">
//       <div className="mb-4">
//         <div className='flex flex-col md:flex-row md:space-x-8 '>
//           <div className="w-full md:w-1/3 lg:w-3/5">
//             <div className="pt-4 flex items-center justify-between">
//               <h2 className="text-xl font-bold text-[#343C6A]">My Cards</h2>
//               <button className="px-4 py-2 text-sm font-bold text-[#343C6A] border border-none">
//                 + Add Card
//               </button>
//             </div>

//             <div className='flex overflow-x-auto space-x-6 scrollbar-hide gap-16 [&::-webkit-scrollbar]:hidden'>
//               <div className="flex-shrink-0 w-72">
//                 <CreditCard
//                   balance="$5,756"
//                   cardHolder="Eddy Cusuma"
//                   validThru="12/22"
//                   cardNumber="3778 **** **** 1234"
//                   filterClass=""
//                   bgColor="from-[#4C49ED] to-[#0A06F4]"
//                   textColor="text-white"
//                   iconBgColor="bg-opacity-10"
//                   showIcon={true}
//                 />
//               </div>
//               <div className="flex-shrink-0 w-72">
//                 <WhiteCard />
//               </div>
//               <div className="flex-shrink-0 w-72">
//                 <CreditCard
//                   balance="$7,250"
//                   cardHolder="John Doe"
//                   validThru="11/23"
//                   cardNumber="1234 **** **** 5678"
//                   filterClass=""
//                   bgColor="from-[#F49E0A] to-[#F06A24]"
//                   textColor="text-white"
//                   iconBgColor="bg-opacity-10"
//                   showIcon={true}
//                 />
//               </div>
//             </div>
//           </div>
//           <div className="w-full md:w-1/3 lg:w-1/5 mt-8 md:mt-0 pt-4 pb-8">
//             <h2 className="text-xl font-bold text-[#343C6A]">My Expense</h2>
//             <div className="w-full max-h-[200px] flex-grow pt-6">
//               <BarChart />
//             </div>
//           </div>
//         </div>
//       </div>
//       <div className="mb-4 md:w-4/5 lg:w-10/12">
//         <h2 className="text-xl font-bold mb-4 pt-6 text-[#343C6A]">Recent Transactions</h2>
//         <Tabs
//           tabs={['All Transactions', 'Income', 'Expense']}
//           activeTab={activeTab}
//           onTabChange={setActiveTab}
//         />
//         {loading ? (
//           <div>Loading...</div>
//         ) : (
//           <TransactionsList transactions={transactions} />
//         )}
//       </div>
//     </div>
//   );
// };

// export default Page;
