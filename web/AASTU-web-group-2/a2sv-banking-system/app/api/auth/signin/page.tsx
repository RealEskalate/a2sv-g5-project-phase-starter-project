import React from 'react'
import { getServerSession } from 'next-auth'
import { options } from '../[...nextauth]/options'
import { redirect } from 'next/navigation' 
import Forms from './components/Forms'
const page = async () => {
    const session = await getServerSession(options)
    if(session){
        redirect("/")
    }
  return (
    <Forms></Forms>
  )
}

export default page
