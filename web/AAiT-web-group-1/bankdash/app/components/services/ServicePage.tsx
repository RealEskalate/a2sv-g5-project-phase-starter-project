'use client';
import React from 'react'
import Image from 'next/image'
import ListCard from './ListCard'
import { useGetServicesQuery } from '@/lib/redux/api/servicesApi'
import { useSession } from 'next-auth/react'
import { redirect } from 'next/navigation'
import { SERVICE } from '@/types/serviceType'

// interface NotificationCardsProps {
//     imgpath: string,
//     title: string,
//     description: string,
//     background: string
// }
const data = {
    page : 1,
    size : 6
}

const ServicePage = () => {
    const { data: session, status } = useSession();
    const {data: serviceData, isError: serviceError, isLoading: serviceLoading} = useGetServicesQuery(data);

    const serviceList = serviceData?.data?.content 
    if (status === 'loading') return <h1 className='text-center text-lg mt-96'>Loading ....</h1>;
    if (status === 'unauthenticated') { 
        redirect('/auth/login');
    }
    if (serviceError) return <h1>Ooops, there seems to be an error</h1>;
    
    if (serviceLoading) return <h1 className='text-center text-lg mt-96'>Loading ....</h1>;


  return (
    <div className='pl-5'>
        <div className='overflow-x-scroll sm:overflow-x-auto flex sm:flex-wrap lg:pb-3 mt-2'>
            <div className='flex flex-wrap rounded-3xl bg-white w-fit h-fit p-3 lg:py-5 sm:px-5 lg:px-8 lg:m-4 lg:mr-5 min-w-56 lg:min-w-64 sm:mb-3 mr-4 '>
                <div className = "rounded-full mr-5 p-2 bg-[#e7edff] my-auto max-w-11 sm:max-w-12 lg:max-w-20">
                    <Image alt='insurance icon' src='/images/life-insurance.svg' width={40} height={40}/>
                </div>
                <div className='my-auto w-fit'>
                    <div className='font-semibold text-base lg:text-xl text-[#232323]'>Life Insurance</div>
                    <div className='text-[#718EBF] text-xs lg:text-base inline-block'>Unlimited protection</div>
                </div>
            </div>
            <div className='flex flex-wrap rounded-3xl bg-white w-fit h-fit p-3 lg:py-5 sm:px-5 lg:px-8 lg:m-4 lg:mr-5 min-w-56 lg:min-w-64 sm:mb-3 mr-4'>
                <div className = "rounded-full mr-5 p-2 bg-[#FFF5D9] my-auto max-w-11 sm:max-w-12 lg:max-w-20">
                    <Image alt='insurance icon' src='/images/bag.svg' width={40} height={40}/>
                </div>
                <div className='my-auto'>
                    <div className='font-semibold text-base lg:text-xl text-[#232323]'>Shopping</div>
                    <div className='text-[#718EBF] text-xs lg:text-base inline-block'>Buy, Think, Grow</div>
                </div>
            </div>
            <div className='flex flex-wrap rounded-3xl bg-white w-fit h-fit p-3 lg:py-5 sm:px-5 lg:px-8 lg:m-4 lg:mr-5 min-w-56 lg:min-w-64 sm:mb-3 mr-4'>
                <div className = "rounded-full mr-5 p-2 bg-[#DCFAF8] my-auto max-w-11 sm:max-w-12 lg:max-w-20">
                    <Image alt='insurance icon' src='/images/shield.svg' width={40} height={40}/>
                </div>
                <div className='my-auto'>
                    <div className='font-semibold text-base lg:text-xl text-[#232323]'>Safety</div>
                    <div className='text-[#718EBF] text-xs lg:text-base inline-block'>We are you allies</div>
                </div>
            </div>
        </div>
        <div className='lg:m-4 sm:m-2 mt-4'>
            <h1 className='text-[#343C6A] lg:text-xl text-lg font-bold pb-4 sm:pb-6'>Bank Services List</h1>

            <div>
                {serviceList?.map((service:any, index: number) => (
                    <ListCard key={index} imgpath={service.icon} heading={{title: service.name, description: service.details}} status={service.status} type={service.type} description3={{title:'', description:''}}/>
                ))}
            </div>
            
            {/* <ListCard imgpath={imgpath} heading={{ title, description }} description1={{ title: Lorem, description: Lorem }} description2={{ title: Lorem, description: Lorem }} description3={{ title: Lorem, description: Lorem }}/>

            <ListCard imgpath={imgpath2} heading={{ title:title2, description }} description1={{ title: Lorem, description: Lorem }} description2={{ title: Lorem, description: Lorem }} description3={{ title: Lorem, description: Lorem }}/>

            <ListCard imgpath={imgpath3} heading={{ title:title3, description }} description1={{ title: Lorem, description: Lorem }} description2={{ title: Lorem, description: Lorem }} description3={{ title: Lorem, description: Lorem }}/>
            <ListCard imgpath={imgpath4} heading={{ title:title4, description }} description1={{ title: Lorem, description: Lorem }} description2={{ title: Lorem, description: Lorem }} description3={{ title: Lorem, description: Lorem }}/>
            <ListCard imgpath={imgpath5} heading={{ title:title5, description }} description1={{ title: Lorem, description: Lorem }} description2={{ title: Lorem, description: Lorem }} description3={{ title: Lorem, description: Lorem }}/> */}
            {/* <span className='ml-40 mt-40 text-lg text-gray-700'>There are no bank services available!</span> */}
            <div className='bg-white'>
            </div>
        </div>
    </div>
    
  )
}

const Lorem = 'Lorem Ipsum'
const imgpath = '/images/loans.svg'
const title = 'Business loans'
const description = 'It is a long established'

const imgpath2 = '/images/CheckingAccount.svg'
const title2 = 'Checking Account'

const imgpath3 = '/images/SavingsAccount.svg'
const title3 = 'Savings Account'

const imgpath4 = '/images/DebitAndCredit.svg'
const title4 = 'Debit and Credit cards'

const imgpath5 = '/images/LifeInsurance.svg'
const title5 = 'Life Insurance'

export default ServicePage