import React from 'react'
import CardSettingLine from './CardSettingLine'
import Link from 'next/link'

type cardSettingType = {
    icon: string;
    title: string;
    description: string;
    background: string;
}

const CardSetting = () => {
    const options:cardSettingType[] = [
        {icon: '/icons/credit-card.png', title: 'Block Card', description: 'Instantly block your card', background: 'bg-[#FFF5D9]'},
        {icon: '/icons/padlock.png', title: 'Change Pin Code', description: 'Choose another pin code', background: 'bg-[#E7EDFF]'},
        {icon: '/icons/google.png', title: 'Add to Google Pay', description: 'Withdraw without any card', background: 'bg-[#FFE0EB]'},
        {icon: '/icons/apple.png', title: 'Add to Apple Pay', description: 'Withdraw without any card', background: 'bg-[#DCFAF8]'},
        {icon: '/icons/apple.png', title: 'Add to Apple Store', description: 'Withdraw without any card', background: 'bg-[#DCFAF8]'},
    ]
  return (
    <div className='md:w-[335px] md:h-[430px] w-[330px] h-[450px] rounded-3xl bg-white'>
        { options.map((option, index) => (
            <Link key={index} href={''}>
                <CardSettingLine key={index} icon={option.icon} title={option.title} description={option.description} background={option.background} />
            </Link>
        ))}
    </div>
  )
}

export default CardSetting
