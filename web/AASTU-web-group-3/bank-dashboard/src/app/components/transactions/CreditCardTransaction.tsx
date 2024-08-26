'use client'
import React, { useEffect } from 'react'
import CreditCard from '../CreditCard'
import DashboardBarChart from '../Chart/DashboardBarChart'
import { useDispatch, useSelector } from 'react-redux'
import Link from 'next/link'

import { useGetCardsQuery } from '@/lib/redux/api/cardsApi'
import { setCards, setLoading, setError } from '@/lib/redux/slices/cardsSlice'
import { RootState } from '@/lib/redux/store'


const cardStyles = {
  Debit: {
    backgroundImg:
      "bg-[linear-gradient(107.38deg,#2D60FF_2.61%,#539BFF_101.2%)]",
    textColor: "text-white",
  },
  Primary: {
    backgroundImg:
      "bg-[linear-gradient(107.38deg,#4C49ED_2.61%,#0A06F4_101.2%)]",
    textColor: "text-white",
  },
  Visa: {
    backgroundImg: "bg-black",
    textColor: "text-white",
  },
  Secondary: {
    backgroundImg: "bg-gray-200",
    textColor: "text-black",
  },
};

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

          <div className="creditcards flex gap-5 lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar h-56 lg:justify-start lg:px-4">

          {cards.map((card, index) => {
            const style =
              cardStyles[card.cardType as keyof typeof cardStyles] ||
              cardStyles.Primary;
  
            return (
              <div
                key={index}
                className="credit-card min-h-80 w-[360px] max-w-72 md:max-w-96 flex-shrink-0"
              >
                <CreditCard
                  name={card.cardHolder}
                  balance={String(card.balance)}
                  cardNumber={card.semiCardNumber}
                  validDate={card.expiryDate}
                  backgroundImg={style.backgroundImg}
                  textColor={style.textColor}
                />
              </div>
            );
          })}
        </div>

        </div>
    )
}

export default CreditCardTransaction
