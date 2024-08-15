import React from 'react'
import BankServiceList from './BankServiceList'
import LifeInsurance from './LifeInsurance'

const page = () => {
  return (
    <div className='flex-col'>
      <div>
      <BankServiceList/>
      </div>
      <div>
        <LifeInsurance/>
      </div>
        
        

        
    </div>
    
  )
}

export default page