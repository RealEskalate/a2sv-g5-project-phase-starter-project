import { LoanTable } from '@/components/loanTable/loanTable'
import Top from '@/components/Top'
import React from 'react'

function Loan() {
return (
    <div className='flex flex-col items-center'>
            <Top topicName='Loan'/>
            <div className='border rounded-lg w-11/12 shadow-lg'>
                    <LoanTable />
            </div>
    </div>
)
}

export default Loan