import React from 'react'
import { Color } from 'chart.js'
import { colors  , logo} from '@/constants'
import DesktopCreditCart from '@/components/DesktopCreditCard'
import RecentTransaction from '@/components/Recent Transaction'
import { icons } from 'lucide-react'
import { text } from 'stream/consumers'
import BarChart from '@/components/BarChart'

const page = () => {
  return (
    <div className='px-6 '>
      <div className='flex flex-col  justify-between md:flex-row'>
        <div>
          <div className={`${colors.navbartext} flex justify-between`}>
            <h1 className='font-bold text-2xl'>My Cards</h1>
            <p className='my-2'> See All</p>
          </div>
          <div className='flex gap-3 overflow-x-auto md:overflow-x-hidden max-w-lg'>
            <div className='flex py-3'> <DesktopCreditCart/></div>
            <div className='flex py-3'> <DesktopCreditCart/></div>
          </div>
        </div>
        <div>
          <div className={`${colors.navbartext} flex justify-between py-4`}>
            <h1 className='font-bold text-2xl'>Recent Transaction</h1>
          </div>
          <div className=''>
            <RecentTransaction icon={logo.RT1} color={colors.lightorange} colortext={colors.textred} name="deposit from my" date = "23 january 2023" amount='-850'/>
            <RecentTransaction icon={logo.RT2} color={colors.lightblue} colortext={colors.textgreen} name="deposit from my" date = "23 january 2023" amount='+2500'/>
            <RecentTransaction icon={logo.RT3} color={colors.lightgreen} colortext={colors.textgreen} name="deposit from my" date = "23 january 2023" amount='+5400'/>
          </div>
        </div>
      </div>
      <div className={`${colors.navbartext} flex justify-between py-4`}>
        <h1 className='font-bold text-2xl'>Weekly Activity</h1>
      </div>
      <div><BarChart/></div>
    </div>
  )
}

export default page