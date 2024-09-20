import React from 'react'
import RecentTransaction from './Recent_transaction'
import WeeklyActivityBarChart from './WeeklyActivityBarChart'
import ExpenseStatisticsPieChart from './ExpenseStatisticsPieChart'
import Quick_transfer from './Quick_transfer'
import BalanceHistoryLineGraph from './BalanceHistoryLineGraph'
import Card from '../Common/Card'
// import Mycards from '../components/Common/MyCards'


const LineData = [150, 200, 500, 250, 400, 100, 700, 750, 350, 330, 550, 500]
const sectors = ['Service', 'Others','Shopping', 'Transfer']
const bgColors = ['#FC7900', '#1814F3', '#FA00FF', '#343C6A']
const expenses = [11200, 16200, 12400, 6400, 15000, 12000, 12100]
const months = ['Aug', 'Sep', 'Oct', 'Nov', 'Dec', 'Jan']
const weekdays = ['Sat', 'Sun','Mon', 'Tue', 'Wed', 'Thu', 'Fri' ]
const deposits = [680, 300,575, 200, 425, 350, 605 ]
const withdraws = [200, 260, 300, 400, 300, 200, 100]

const Page = () => {
  return (
    <div className='w-full my-10'>
        <div className="flex justify-around mb-8">
            <div className='w-3/5 '>
                <Card/>
            {/* <Mycards/> */}
            </div>
            <div className='w-2/6'>
                <RecentTransaction recents={[]} />
            </div>
        </div>
        <div className='flex justify-evenly mb-10'>
            <div className='w-3/5  mr-2'>
                <WeeklyActivityBarChart weekdays={weekdays} deposits={deposits} withdraws={withdraws}/>
            </div>
            <div className=' my-auto '>
                <ExpenseStatisticsPieChart sectors={sectors} bgColors={bgColors}/>
            </div>
        </div>
        <div className='flex w-full justify-around'>
            <div className='md:w-2/6 mr-7 h-80'>
                <Quick_transfer/>
            </div>
            <div className="w-3/6 mr-7 mb-5 h-80 ">
                <BalanceHistoryLineGraph balanceHistory={LineData}/>
            </div>
        </div>
    </div>
  )
}

export default Page