import React, { useState } from 'react'

const RecentTransactions = () => {
  const [activeTab, setActiveTab] = useState('all transactions') 
  const renderContent = ()=>{
    switch(activeTab){
        case 'all transactions':
            return
        case 'expenses':
            return 
        case 'income':
            return
    }
  }
  
  
  return (
    <div className='w-11/12 mt-4 ml-2'>
        <div className='border-b flex justify-start gap-4'>

        </div>
      
    </div>
  )
}

export default RecentTransactions
