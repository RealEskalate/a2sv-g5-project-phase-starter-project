import ActiveLoansOverviewTable from '@/components/ActiveLoansOverviewTable/ActiveLoansOverviewTable';
import Loansitem from '@/components/LoansItems/Loansitem';
import React from 'react';

export default function page() {
  return (
    <div className='flex flex-col gap-5'> 
      <Loansitem />
      <ActiveLoansOverviewTable />
    </div>
  );
}
