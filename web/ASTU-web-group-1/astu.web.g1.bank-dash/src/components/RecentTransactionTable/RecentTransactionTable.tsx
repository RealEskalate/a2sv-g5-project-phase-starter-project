import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import AllTransactionTable from './AllTransactionTable';
import Income from './Income';
import Expense from './Expense';

const RecentTransactionTable = () => {
  return (
    <div>
      <h1 className='text-16px md:text-18px xl:text-22px text-[#333B69] font-semibold mb-5'>
        Recent Transaction
      </h1>
      <div className='bg-white px-2 py-5 rounded-3xl'>
        <Tabs defaultValue='all_transaction' className='w-full'>
          <TabsList className='w-full px-2 md:justify-start bg-white'>
            <TabsTrigger
              value='all_transaction'
              className='data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]'
            >
              All Transaction
            </TabsTrigger>
            <TabsTrigger
              value='income'
              className=' data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]'
            >
              Income
            </TabsTrigger>
            <TabsTrigger
              value='expence'
              className=' data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]'
            >
              Expense
            </TabsTrigger>
          </TabsList>

          <TabsContent value='all_transaction'>
            <AllTransactionTable />
          </TabsContent>
          <TabsContent value='income'>
            <Income />
          </TabsContent>
          <TabsContent value='expence'>
            <Expense />
          </TabsContent>
        </Tabs>
      </div>
    </div>
  );
};

export default RecentTransactionTable;
