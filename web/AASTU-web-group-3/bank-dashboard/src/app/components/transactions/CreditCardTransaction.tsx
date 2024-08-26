'use client'
import React, { useEffect } from 'react'
import CreditCard from '../CreditCard'
import DashboardBarChart from '../Chart/DashboardBarChart'
import { useDispatch, useSelector } from 'react-redux'
import Link from 'next/link'

import { useGetCardsQuery } from '@/lib/redux/api/cardsApi'
import { setCards, setLoading, setError } from '@/lib/redux/slices/cardsSlice'
import { RootState } from '@/lib/redux/store'


const CreditCardTransaction = () => {
    const dispatch = useDispatch()
    const {
        cards, 
        loading,
        error
    } = useSelector((state:RootState)=>state.cards )

    const {
        data:cardsData,
        isLoading:cardsLoading,
        isError:errorCard
    } = useGetCardsQuery({size:5, page:0})


    useEffect(()=>{
        dispatch(setLoading(cardsLoading))
        if (cardsData){
            dispatch(setCards(cardsData.content))
        }
        if (errorCard){
            dispatch(setError("Error on fetching data"))
        }
    }, [cardsData, errorCard, cardsLoading, dispatch])

    if(loading) return <div>Loading...</div>
    if(error) return <div>{error}</div>
  
    return (
        <div className="lg:w-[65%] xl:w-[68%] rounded-xl bg-[#F5F7FA] dark:bg-[#0f1a2b]">
          <div className="credit-card-info flex justify-between h-16 items-center ">
            <h1 className="font-semibold text-[#343C6A] dark:text-white">My cards</h1>
            <Link href="/creditcardpage#add-new-card">
              <h1 className="text-[#2D60FF] dark:text-darkPrimary">+ Add Card</h1>
            </Link>  
          </div>
          <div className="cards flex gap-5 lg:gap-1 lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar h-56 lg:justify-between xl:gap-10">
            {
                cards.map((card)=>(
                    <div key={card.id} className="flex-shrink-0 min-w-60 w-73 lg:w-60 lg:h-48 xl:w-96 xl:h-56 items-center">
                      <CreditCard
                        name={card.cardHolder}
                        balance={card.balance.toString()} // Convert the balance to a string
                        cardNumber={card.semiCardNumber}
                        validDate={card.expiryDate}
                        backgroundImg="bg-blue-500"
                        textColor="text-white"
                      />
                    </div>
                ))
            }
          </div>
        </div>
    )
}

export default CreditCardTransaction
