// src/components/RecentTransactions.tsx

import { useEffect, useState } from 'react';

import displaytransaction from './displaytransaction';
import Pagination from '@/components/Pagination';
import { getAllTransactions , getIncomes , getExpenses} from '@/services/transactionfetch';



const RecentTransactions = () => {
  const [filter, setFilter] = useState<'all' | 'income' | 'expense'>('all');
  const [currentPage, setCurrentPage] = useState(0);
  const[alltransaction , setalltransaction] = useState([])
  const[allincomes , setallincomes] = useState([])
  const[allexpenses , setallexpenses] = useState([])
  // Filter transactions based on the selected tab
  // const filteredTransactions = transactionsData.filter(transaction => {
  //   if (filter === 'all') return true;
  //   return transaction.type === filter;
  // });
  const[totalPages , settotalpages] = useState(0)

  
 const ITEMS_PER_PAGE = 5

  // Pagination calculations
  
  const startIndex = (currentPage - 1) * ITEMS_PER_PAGE;
  
  // const currentTransactions = filteredTransactions.slice(startIndex, startIndex + ITEMS_PER_PAGE);
  const[incoming , setincoming] = useState(false)
  const[expense , setexpense] = useState(false)
  const[all , setall] = useState(false)
  useEffect (() => {
    const fetch = async () =>{
     try{
      const[response , response2 , response3] = await Promise.all([
        getAllTransactions( currentPage , ITEMS_PER_PAGE),
        getIncomes(currentPage , ITEMS_PER_PAGE),
        getExpenses(currentPage , ITEMS_PER_PAGE ),

      ])
      
      setalltransaction(response.data.content || [])
      settotalpages(response.data.totalPages)
       
      setallincomes(response2.data.content || [])
      settotalpages(response2.data.totalPages)
      
      setallexpenses(response3.data.content || [])
      settotalpages(response3.data.totalPages)
      
     } 
   catch(error){
      console.error('Error fetching transactions:', error);
   }

    }
  
   fetch();

  } , [currentPage])


  useEffect(() => {
    setall(true)
  } , [])


  

  
  return (
    <div className="p-4 md:ml-64">
      {/* Tabs */}
      <div className="flex flex-wrap mb-4">
        <button
          onClick={() => {setFilter('all') , setall(true) , setincoming(false) , setexpense(false)}

        }
          className={`font-bold px-4 py-2 rounded-t-lg ${filter === 'all' ? 'border-b-2 border-blue-500' : 'text-gray-600'}`}
        >
          All Transactions
        </button>
        <button
          onClick={() => {setFilter('income'), setincoming(true) , setall(false) , setexpense(false)}
          }
          className={`font-bold px-4 py-2 rounded-t-lg ${filter === 'income' ? 'border-b-2 border-blue-500' : 'text-gray-600'}`}
        
          >
          Income
        </button>
        <button
          onClick={() =>{ setFilter('expense') ,setexpense(true) , setall(false) , setincoming(false)}}
          className={`font-bold px-4 py-2 rounded-t-lg ${filter === 'expense' ? 'border-b-2 border-blue-500' : 'text-gray-600'}`}
        >
          Expenses
        </button>
        
      </div>
      
      {all && displaytransaction(alltransaction ,"all")}
      {incoming && displaytransaction(allincomes , "income")}
      {expense && displaytransaction(allexpenses , "expense") }
      
        <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={setCurrentPage}
      />
      </div>

      
  );
};

export default RecentTransactions;
