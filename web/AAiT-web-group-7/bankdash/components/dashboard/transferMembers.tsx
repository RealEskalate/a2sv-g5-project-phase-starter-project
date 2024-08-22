import Image from 'next/image'
import React from 'react'
import img from '@/public/person.jpg'

interface Person {
    name: string;
    role: string;
}

const TransferMembers = ({person, selectedPerson}: {person: Person, selectedPerson:string | null}) => {
  return (
    <div className="flex flex-col items-center">
              <Image
                src={img}
                alt={person.name}
                className="w-16 h-16 rounded-full object-cover mb-2"
                width={64}
                height={64}
              />
              <p className={`font-semibold ${selectedPerson == person.name? "text-[#5f7df4]": ""} text-sm `}>{person.name}</p>
              <p className={`text-sm text-gray-500 ${selectedPerson == person.name? "text-[#5f7df4]": ""}`}>{person.role}</p>
            </div>
  )
}

export default TransferMembers
