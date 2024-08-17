import React from 'react'
import CardList from "../../components/CardList"
import DoughnutChart from '@/components/DoughnutChart'

const page = () => {
  return (

    <div className="flex flex-col md:flex-row">
    <DoughnutChart />
    <CardList />
  </div>
  
  )
}

export default page